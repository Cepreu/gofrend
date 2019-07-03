package fulfiller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/vars"
	"github.com/Cepreu/gofrend/utils"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// Interpreter can process modules and populate its WebhookResponse as necessary
type Interpreter struct {
	Session         *Session
	Script          *ivr.IVRScript
	ScriptHash      string
	QueryResult     *dialogflowpb.QueryResult
	WebhookResponse *dialogflowpb.WebhookResponse
}

// Interpret - main interpreter loop
func Interpret(wr dialogflowpb.WebhookRequest, script *ivr.IVRScript, scriptHash string) (*dialogflowpb.WebhookResponse, error) {
	sessionID := wr.Session
	session, err := loadSession(sessionID, script)

	interpreter := &Interpreter{
		Session:     session,
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

	moduleID := utils.DisplayNameToModuleID(wr.QueryResult.Intent.DisplayName)
	module, err := getModuleByID(script, ivr.ModuleID(moduleID))
	if err != nil {
		return nil, err
	}
	if !isInputOrMenu(module) {
		return nil, fmt.Errorf("Expected input or menu module to start processing, instead got: %T", module)
	}
	module, err = interpreter.processInitial(module)
	if err != nil {
		return nil, err
	}
	for module != nil {
		module, err = interpreter.process(module)
		if err != nil {
			return nil, err
		}
	}
	return interpreter.WebhookResponse, nil
}

// ProcessInitial process a module and returns the next module to be processed
func (interpreter *Interpreter) processInitial(module ivr.Module) (ivr.Module, error) {
	switch v := module.(type) {
	case *ivr.InputModule:
		return interpreter.processInputInitial(v)
	}
	return nil, nil
}

// Process processes a module and returns the next module to be processed
func (interpreter *Interpreter) process(module ivr.Module) (ivr.Module, error) {
	utils.PrettyLog(module)
	switch v := module.(type) {
	case *ivr.IfElseModule:
		return interpreter.processIfElse(v)
	case *ivr.PlayModule:
		return interpreter.processPlay(v)
	case *ivr.InputModule:
		return interpreter.processInput(v)
	case *ivr.QueryModule:
		return interpreter.processQuery(v)
	case *ivr.HangupModule:
		return nil, nil
	}
	return nil, nil
}

func (interpreter *Interpreter) processInputInitial(module *ivr.InputModule) (ivr.Module, error) {
	parameters := interpreter.QueryResult.Parameters.Fields
	for name, value := range parameters {
		interpreter.Session.setParameter(name, value)
	}
	return getModuleByID(interpreter.Script, module.GetDescendant())
}

func (interpreter *Interpreter) processInput(module *ivr.InputModule) (ivr.Module, error) {
	interpreter.addResponseText(module.VoicePromptIDs)
	interpreter.WebhookResponse.OutputContexts = []*dialogflowpb.Context{
		&dialogflowpb.Context{
			Name:          utils.MakeContextName(utils.MakeDisplayName(interpreter.ScriptHash, module.GetID())),
			LifespanCount: 1,
		},
	}
	return nil, interpreter.Session.save()
}

func (interpreter *Interpreter) addResponseText(VoicePromptIDs ivr.ModulePrompts) {
	promptStrings := VoicePromptIDs.TransformToAI(interpreter.Script.Prompts)
	intentMessageText := interpreter.WebhookResponse.FulfillmentMessages[0].GetText()
	intentMessageText.Text = append(intentMessageText.Text, promptStrings...)
}

func (interpreter *Interpreter) processQuery(module *ivr.QueryModule) (ivr.Module, error) {
	interpreter.addResponseText(module.VoicePromptIDs)
	body, err := utils.CmdUnzip(module.RequestInfo.Base64)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(module.Method, module.URL, bytes.NewReader([]byte(body)))
	for _, h := range module.Headers {
		request.Header.Add(h.Key, h.Value.Value.String())
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
	if module.BranchIf.Cond.ConditionGrouping == "" { // Parser currently does not populate this field
		module.BranchIf.Cond.ConditionGrouping = "ALL"
	}
	switch module.BranchIf.Cond.ConditionGrouping {
	case "ALL":
		conditionsPass = true
		for _, condition := range conditions {
			populateCondition(interpreter.Session, condition)
			passes, err := conditionPasses(condition)
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

func populateCondition(session *Session, condition *ivr.Condition) {
	varName := condition.LeftOperand.VariableName
	if varName != "" { // varName is empty if comparison is against constant
		storageVar, ok := session.getParameter(varName)
		if !ok {
			log.Fatalf("Error finding session variable with name: %s", varName)
		}
		condition.LeftOperand.Value = storageVar.value()
	}
	varName = condition.RightOperand.VariableName
	if varName != "" {
		storageVar, ok := session.getParameter(varName)
		if !ok {
			log.Fatalf("Error finding session variable with name: %s", varName)
		}
		condition.RightOperand.Value = storageVar.value()
	}
}

func conditionPasses(condition *ivr.Condition) (bool, error) {
	switch condition.ComparisonType {
	case "MORE_THAN":
		var left float64
		switch v := condition.LeftOperand.Value.(type) {
		case *vars.Integer:
			left = float64(v.Value)
		case *vars.Numeric:
			left = v.Value
		default:
			log.Fatalf("Expected int or numeric in more than comparison, instead got: %T", v)
		}
		var right float64
		switch v := condition.RightOperand.Value.(type) {
		case *vars.Integer:
			right = float64(v.Value)
		case *vars.Numeric:
			right = v.Value
		default:
			log.Fatalf("Expected int or numeric in more than comparison, instead got: %T", v)
		}
		log.Printf("Left: %f, Right: %f.", left, right)
		return left > right, nil
	}
	return false, nil
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

func isInputOrMenu(module ivr.Module) bool {
	switch module.(type) {
	case *ivr.MenuModule:
		return true
	case *ivr.InputModule:
		return true
	default:
		return false
	}
}
