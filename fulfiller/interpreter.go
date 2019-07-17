package fulfiller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// Interpreter can process modules and populate its WebhookResponse as necessary
type Interpreter struct {
	SessionID       string
	Session         *Session
	Script          *ivr.IVRScript
	ScriptHash      string
	QueryResult     *dialogflowpb.QueryResult
	WebhookResponse *dialogflowpb.WebhookResponse
}

// Interpret - main interpreter loop
func Interpret(wr dialogflowpb.WebhookRequest, script *ivr.IVRScript, scriptHash string) (*dialogflowpb.WebhookResponse, error) {
	interpreter := &Interpreter{
		SessionID:   wr.Session,
		Script:      script,
		ScriptHash:  scriptHash,
		QueryResult: wr.QueryResult,
		WebhookResponse: &dialogflowpb.WebhookResponse{
			FulfillmentMessages: []*dialogflowpb.Intent_Message{
				{
					Message: &dialogflowpb.Intent_Message_Text_{
						Text: &dialogflowpb.Intent_Message_Text{
							Text: []string{},
						},
					},
				},
			},
		},
	}

	displayName := wr.QueryResult.Intent.DisplayName
	moduleID := utils.DisplayNameToModuleID(displayName)
	module, err := getModuleByID(script, ivr.ModuleID(moduleID))
	if err != nil {
		return nil, err
	}
	module, err = interpreter.processInitial(module, displayName)
	if err != nil {
		return nil, err
	}
	for module != nil {
		module, err = interpreter.process(module)
		if err != nil {
			return nil, err
		}
	}
	return interpreter.WebhookResponse, interpreter.Session.close()
}

// ProcessInitial process a module and returns the next module to be processed
func (interpreter *Interpreter) processInitial(module ivr.Module, displayName string) (ivr.Module, error) {
	switch v := module.(type) {
	case *ivr.IncomingCallModule:
		return interpreter.processIncomingCall(v)
	case *ivr.InputModule:
		return interpreter.processInputInitial(v)
	case *ivr.MenuModule:
		return interpreter.processMenuInitial(v, displayName)
	default:
		panic("Not implemented")
	}
}

// Process processes a module and returns the next module to be processed
func (interpreter *Interpreter) process(module ivr.Module) (ivr.Module, error) {
	switch v := module.(type) {
	case *ivr.IfElseModule:
		return interpreter.processIfElse(v)
	case *ivr.PlayModule:
		return interpreter.processPlay(v)
	case *ivr.InputModule:
		return interpreter.processInput(v)
	case *ivr.MenuModule:
		return interpreter.processMenu(v)
	case *ivr.QueryModule:
		return interpreter.processQuery(v)
	case *ivr.HangupModule:
		return interpreter.processHangup(v)
	default:
		panic("Not implemented")
	}
}

func (interpreter *Interpreter) initSession() error {
	session, err := initSession(interpreter.SessionID, interpreter.Script)
	if err != nil {
		return err
	}
	interpreter.Session = session
	return nil
}

func (interpreter *Interpreter) loadSession() error {
	session, err := loadSession(interpreter.SessionID, interpreter.Script)
	if err != nil {
		return err
	}
	interpreter.Session = session
	return nil
}

func (interpreter *Interpreter) processIncomingCall(module *ivr.IncomingCallModule) (ivr.Module, error) {
	err := interpreter.initSession()
	if err != nil {
		return nil, err
	}
	return getModuleByID(interpreter.Script, module.GetDescendant())
}

func (interpreter *Interpreter) processInputInitial(module *ivr.InputModule) (ivr.Module, error) {
	err := interpreter.loadSession()
	if err != nil {
		return nil, err
	}
	parameters := interpreter.QueryResult.Parameters.Fields
	for name, value := range parameters {
		interpreter.Session.setParameter(name, value)
	}
	return getModuleByID(interpreter.Script, module.GetDescendant())
}

func (interpreter *Interpreter) processInput(module *ivr.InputModule) (ivr.Module, error) {
	interpreter.addResponseText(module.VoicePromptIDs)
	interpreter.populateWebhookContext(module.GetID())
	return nil, interpreter.Session.save()
}

func (interpreter *Interpreter) processMenuInitial(module *ivr.MenuModule, displayName string) (ivr.Module, error) {
	err := interpreter.loadSession()
	if err != nil {
		return nil, err
	}
	branchName := utils.MenuDisplayNameToBranchName(displayName)
	var branch *ivr.OutputBranch
	for _, b := range module.Branches {
		if b.Name == branchName {
			branch = b
			break
		}
	}
	return getModuleByID(interpreter.Script, branch.Descendant)
}

func (interpreter *Interpreter) processMenu(module *ivr.MenuModule) (ivr.Module, error) {
	interpreter.addResponseText(module.VoicePromptIDs)
	interpreter.populateWebhookContext(module.GetID())
	return nil, interpreter.Session.save()
}

func (interpreter *Interpreter) populateWebhookContext(moduleID ivr.ModuleID) {
	interpreter.WebhookResponse.OutputContexts = []*dialogflowpb.Context{
		&dialogflowpb.Context{
			Name:          utils.MakeInputContextName(utils.MakeInputDisplayName(interpreter.ScriptHash, moduleID)),
			LifespanCount: 1,
		},
	}
}

func (interpreter *Interpreter) addResponseText(VoicePromptIDs ivr.ModulePrompts) {
	promptStrings := VoicePromptIDs.TransformToAI(interpreter.Script.Prompts)
	expression := regexp.MustCompile("@.+@")
	f := func(s string) string {
		varName := s[1 : len(s)-1]
		variable, found := interpreter.Session.getParameter(varName)
		if !found {
			return ""
		}
		return variable.valueAsString()
	}
	for i := range promptStrings {
		promptStrings[i] = expression.ReplaceAllStringFunc(promptStrings[i], f)
	}
	intentMessageText := interpreter.WebhookResponse.FulfillmentMessages[0].GetText()
	intentMessageText.Text = append(intentMessageText.Text, promptStrings...)
}

// func (interpreter *Interpreter) processSkillTransfer(module *ivr.SkillTransferModule) (ivr.Module, error) {
// 	config, err := cloud.GetConfig()
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = createCallback(config["DOMAIN_NAME"], config["CAMPAIGN_NAME"])
// }

func (interpreter *Interpreter) processQuery(module *ivr.QueryModule) (ivr.Module, error) {
	interpreter.addResponseText(module.VoicePromptIDs)
	body, err := utils.CmdUnzip(module.RequestInfo.Base64)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(module.Method, module.URL, bytes.NewReader([]byte(body)))
	for _, h := range module.Headers {
		val := interpreter.Script.Variables[h.Value]
		request.Header.Add(h.Key, val.Value.StringValue)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	for _, responseInfo := range module.ResponseInfos {
		if response.StatusCode >= responseInfo.HTTPCodeFrom && response.StatusCode <= responseInfo.HTTPCodeTo {
			switch responseInfo.ParsingMethod {
			case "REG_EXP":
				expression, err := regexp.Compile(responseInfo.Regexp.RegexpBody)
				if err != nil {
					return nil, err
				}
				matches := expression.FindStringSubmatch(string(contents))
				if matches != nil {
					for i, match := range matches[1:] {
						if responseInfo.TargetVariables[i] != "" {
							err = interpreter.Session.setParameterString(responseInfo.TargetVariables[i], match)
							if err != nil {
								return nil, err
							}
						}
					}
				}
			default:
				return nil, fmt.Errorf("Parsing method not understood: '%s'", responseInfo.ParsingMethod)
			}

		}
	}
	return getModuleByID(interpreter.Script, module.GetDescendant())
}

func (interpreter *Interpreter) processIfElse(module *ivr.IfElseModule) (ivr.Module, error) {
	var conditionsPass bool
	conditions := module.BranchIf.Cond.Conditions
	if module.BranchIf.Cond.CustomCondition == "" { // Parser currently does not populate this field
		module.BranchIf.Cond.CustomCondition = "ALL"
	}
	switch module.BranchIf.Cond.CustomCondition {
	case "ALL":
		conditionsPass = true
		for _, condition := range conditions {
			err := populateCondition(interpreter.Session, condition, interpreter.Script)
			if err != nil {
				return nil, err
			}
			passes, err := conditionPasses(condition, interpreter.Script)
			if err != nil {
				return nil, err
			}
			if !passes {
				conditionsPass = false
			}
		}
	}
	if conditionsPass {
		return getModuleByID(interpreter.Script, module.BranchIf.Descendant)
	}
	return getModuleByID(interpreter.Script, module.BranchElse.Descendant)
}

func populateCondition(session *Session, condition *ivr.Condition, script *ivr.IVRScript) error {
	storageVar, ok := session.getParameter(string(condition.LeftOperand))
	if !ok {
		return fmt.Errorf("Error finding session variable with name: %s", condition.LeftOperand)
	}
	script.Variables[condition.LeftOperand].Value.StringValue = storageVar.value()

	if script.Variables[condition.RightOperand].VarType != ivr.VarConstant {
		storageVar, ok := session.getParameter(string(condition.RightOperand))
		if !ok {
			return fmt.Errorf("Error finding session variable with name: %s", condition.RightOperand)
		}
		script.Variables[condition.RightOperand].Value.StringValue = storageVar.value()
	}
	return nil
}

func conditionPasses(condition *ivr.Condition, script *ivr.IVRScript) (bool, error) {
	var leftVal, rightVal *ivr.Variable

	leftVal = script.Variables[condition.LeftOperand]
	if condition.RightOperand != "" {
		rightVal = script.Variables[condition.RightOperand]
	}
	switch condition.ComparisonType {
	case "MORE_THAN":
		switch leftVal.ValType {
		case ivr.ValInteger, ivr.ValTime:
			left, err1 := strconv.Atoi(leftVal.Value.StringValue)
			right, err2 := strconv.Atoi(rightVal.Value.StringValue)
			if err1 == nil && err2 == nil {
				return left > right, nil
			}
		case ivr.ValNumeric:
			left, err1 := strconv.ParseFloat(leftVal.Value.StringValue, 64)
			right, err2 := strconv.ParseFloat(rightVal.Value.StringValue, 64)
			if err1 == nil && err2 == nil {
				return left > right, nil
			}
		case ivr.ValCurrency, ivr.ValCurrencyPound, ivr.ValCurrencyEuro:
			left, err1 := strconv.ParseFloat(leftVal.Value.StringValue[2:], 64)
			right, err2 := strconv.ParseFloat(rightVal.Value.StringValue[2:], 64)
			if err1 == nil && err2 == nil {
				return left > right, nil
			}
		default:
			return leftVal.Value.StringValue >= rightVal.Value.StringValue, nil
		}
	}
	return false, fmt.Errorf("Incorrect comparison operands: %v >= %v", leftVal, rightVal)
}

func (interpreter *Interpreter) processPlay(module *ivr.PlayModule) (ivr.Module, error) {
	interpreter.addResponseText(module.VoicePromptIDs)
	return getModuleByID(interpreter.Script, module.GetDescendant())
}

func (interpreter *Interpreter) processHangup(module *ivr.HangupModule) (ivr.Module, error) {
	return nil, interpreter.Session.delete()
}

func getModuleByID(script *ivr.IVRScript, moduleID ivr.ModuleID) (ivr.Module, error) {
	module, ok := script.Modules[moduleID]
	if !ok {
		return nil, fmt.Errorf("Module not found with ID: %s", moduleID)
	}
	return module, nil
}

func getIncomingCallModule(script *ivr.IVRScript) ivr.Module {
	for _, module := range script.Modules {
		if _, ok := module.(*ivr.IncomingCallModule); ok {
			return module
		}
	}
	return nil
}
