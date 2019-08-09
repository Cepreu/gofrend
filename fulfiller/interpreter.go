package fulfiller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Cepreu/gofrend/cloud"
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
	case *ivr.SkillTransferModule:
		return interpreter.processSkillTransferInitial(v)
	default:
		panic("Not implemented")
	}
}

// Process processes a module and returns the next module to be processed
func (interpreter *Interpreter) process(module ivr.Module) (ivr.Module, error) {
	switch v := module.(type) {
	case *ivr.IfElseModule:
		module, err := interpreter.processIfElse(v)
		if err != nil {
			log.Panicf("Error processing IfElse module: %v", err)
		}
		return module, err
	case *ivr.CaseModule:
		module, err := interpreter.processCase(v)
		if err != nil {
			log.Panicf("Error processing Case module: %v", err)
		}
		return module, err
	case *ivr.PlayModule:
		module, err := interpreter.processPlay(v)
		if err != nil {
			log.Panicf("Error processing Play module: %v", err)
		}
		return module, err
	case *ivr.InputModule:
		module, err := interpreter.processInput(v)
		if err != nil {
			log.Panicf("Error processing Input module: %v", err)
		}
		return module, err
	case *ivr.MenuModule:
		module, err := interpreter.processMenu(v)
		if err != nil {
			log.Panicf("Error processing Menu module: %v", err)
		}
		return module, err
	case *ivr.QueryModule:
		module, err := interpreter.processQuery(v)
		if err != nil {
			log.Panicf("Error processing Query module: %v", err)
		}
		return module, err
	case *ivr.SkillTransferModule:
		module, err := interpreter.processSkillTransfer(v)
		if err != nil {
			log.Panicf("Error processing SkillTransfer module: %v", err)
		}
		return module, err
	case *ivr.SetVariableModule:
		module, err := interpreter.processSetVariable(v)
		if err != nil {
			log.Panicf("Error processing SetVariable module: %v", err)
		}
		return module, err
	case *ivr.HangupModule:
		module, err := interpreter.processHangup(v)
		if err != nil {
			log.Panicf("Error processing Hangup module: %v", err)
		}
		return module, err
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
		interpreter.Session.setParameterFromPBVal(name, value)
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
	/*
	* Todo.
	* Saving the selected menu option value requires more thought than in current implmementation.
	* In IVR, menu options and branches are separate constructs. Branches are expressed in this
	* module's ivr package, whereas menu options simply become training phrases for dialogflow intents.
	* In order to match the behavior of an IVR script, it will be necessary to match the user input
	* to the appropriate menu option, and save this in the desired variable.
	* Right now, a cop-out is taken: we assume the user input from dialogflow will exactly match a menu
	* option, so we can simply save the user input in our variable.
	 */
	if string(module.SaveInputVariable) != "" {
		interpreter.Session.setParameter(string(module.SaveInputVariable), interpreter.QueryResult.QueryText)
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
	expression := regexp.MustCompile("@.+?@")
	f := func(s string) string {
		varName := s[1 : len(s)-1]
		variable := interpreter.Session.getParameter(varName)
		return variable.Value
	}
	for i := range promptStrings {
		promptStrings[i] = expression.ReplaceAllStringFunc(promptStrings[i], f)
	}
	intentMessageText := interpreter.WebhookResponse.FulfillmentMessages[0].GetText()
	intentMessageText.Text = append(intentMessageText.Text, promptStrings...)
}

func (interpreter *Interpreter) addSingleResponseText(text string) {
	intentMessageText := interpreter.WebhookResponse.FulfillmentMessages[0].GetText()
	intentMessageText.Text = append(intentMessageText.Text, text)
}

func (interpreter *Interpreter) processSkillTransferInitial(module *ivr.SkillTransferModule) (ivr.Module, error) {
	err := interpreter.loadSession()
	if err != nil {
		return nil, err
	}
	callbackNumber := interpreter.QueryResult.Parameters.Fields["callback-number"].GetStringValue()
	config, err := cloud.GetConfig()
	if err != nil {
		return nil, err
	}
	err = createCallback(config[cDomainNameConfigKey], config[cCampaignNameConfigKey], callbackNumber, interpreter.makeModuleParams(module))
	if err != nil {
		return nil, err
	}
	interpreter.addSingleResponseText(fmt.Sprintf("Connecting an agent to %s...", callbackNumber))
	return nil, interpreter.Session.delete()
}

func (interpreter *Interpreter) makeModuleParams(module ivr.Module) map[string]string {
	params := map[string]string{cModuleIDParamKey: string(module.GetID())}
	for _, variable := range interpreter.Session.Data.Variables {
		params[string(variable.ID)] = variable.Value
	}
	return params
}

func (interpreter *Interpreter) processSkillTransfer(module *ivr.SkillTransferModule) (ivr.Module, error) {
	interpreter.addSingleResponseText("Please enter your phone number.")
	interpreter.populateWebhookContext(module.GetID())
	return nil, interpreter.Session.save()
}

func (interpreter *Interpreter) processQuery(module *ivr.QueryModule) (ivr.Module, error) {
	interpreter.addResponseText(module.VoicePromptIDs)
	body, err := utils.CmdUnzip(module.RequestInfo.Base64)
	if err != nil {
		return nil, err
	}
	utils.LogWithoutNewlines(body)
	log.Printf("Number of replacements: %d", len(module.RequestInfo.Replacements))
	for _, replacement := range module.RequestInfo.Replacements {
		log.Printf("Replacement location: %d, VariableName: %s", replacement.Position, replacement.VariableName)
		variable := interpreter.Session.getParameter(replacement.VariableName)
		body = body[:replacement.Position] + variable.Value + body[replacement.Position:]
	}
	request, err := http.NewRequest(module.Method, module.URL, bytes.NewReader([]byte(body)))
	for _, h := range module.Headers {
		variable := interpreter.Script.Variables[h.Value]
		request.Header.Add(h.Key, variable.Value)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panicf("Error doing http request: %v", err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicf("Error reading contents of http response: %v", err)
	}
	utils.LogWithoutNewlines(string(contents))
	for _, responseInfo := range module.ResponseInfos {
		if response.StatusCode >= responseInfo.HTTPCodeFrom && response.StatusCode <= responseInfo.HTTPCodeTo {
			switch responseInfo.ParsingMethod {
			case "REG_EXP":
				expression, err := regexp.Compile(responseInfo.Regexp.RegexpBody)
				if err != nil {
					return nil, err
				}
				matches := expression.FindStringSubmatch(string(contents))
				utils.PrettyLog(matches)
				if matches != nil {
					for i, match := range matches[1:] {
						if responseInfo.TargetVariables[i] != "" {
							interpreter.Session.setParameter(responseInfo.TargetVariables[i], match)
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
	conditionsPass, err := interpreter.conditionsPass(&module.BranchIf)
	if err != nil {
		return nil, err
	}
	if conditionsPass {
		return getModuleByID(interpreter.Script, module.BranchIf.Descendant)
	}
	return getModuleByID(interpreter.Script, module.BranchElse.Descendant)
}

func (interpreter *Interpreter) processCase(module *ivr.CaseModule) (ivr.Module, error) {
	for _, branch := range module.Branches[:len(module.Branches)-1] {
		conditionsPass, err := interpreter.conditionsPass(branch)
		if err != nil {
			return nil, err
		}
		if conditionsPass {
			return getModuleByID(interpreter.Script, branch.Descendant)
		}
	}
	return getModuleByID(interpreter.Script, module.Branches[len(module.Branches)-1].Descendant)
}

func (interpreter *Interpreter) conditionsPass(branch *ivr.OutputBranch) (bool, error) {
	conditionsPass := true
	for _, condition := range branch.Cond.Conditions { // Eventually needs to examine CustomCondition field and implement condition logic - currently assumes ALL
		passes, err := interpreter.conditionPasses(condition)
		if err != nil {
			return false, err
		}
		if !passes {
			conditionsPass = false
		}
	}
	return conditionsPass, nil
}

func (interpreter *Interpreter) conditionPasses(condition *ivr.Condition) (bool, error) {
	leftVal := interpreter.Session.getParameter(string(condition.LeftOperand))
	rightVal := interpreter.Session.getParameter(string(condition.RightOperand))
	log.Print(condition.ComparisonType)
	switch condition.ComparisonType {
	case "MORE_THAN", "LESS_THAN":
		log.Print(leftVal.ValType)
		log.Print(rightVal.ValType)
		switch leftVal.ValType {
		case ivr.ValInteger, ivr.ValTime:
			left, err1 := strconv.Atoi(leftVal.Value)
			right, err2 := strconv.Atoi(rightVal.Value)
			if err1 == nil && err2 == nil {
				log.Printf("Left: %d, Right: %d.", left, right)
				if condition.ComparisonType == "MORE_THAN" {
					return left > right, nil
				}
				return left < right, nil
			}
			log.Printf("Err1: %v, err2: %v.", err1, err2)
		case ivr.ValNumeric:
			left, err1 := strconv.ParseFloat(leftVal.Value, 64)
			right, err2 := strconv.ParseFloat(rightVal.Value, 64)
			if err1 == nil && err2 == nil {
				if condition.ComparisonType == "MORE_THAN" {
					return left > right, nil
				}
				return left < right, nil
			}
		case ivr.ValCurrency, ivr.ValCurrencyPound, ivr.ValCurrencyEuro:
			left, err1 := strconv.ParseFloat(leftVal.Value[2:], 64)
			right, err2 := strconv.ParseFloat(rightVal.Value[2:], 64)
			if err1 == nil && err2 == nil {
				if condition.ComparisonType == "MORE_THAN" {
					return left > right, nil
				}
				return left < right, nil
			}
		default:
			if condition.ComparisonType == "MORE_THAN" {
				return leftVal.Value > rightVal.Value, nil
			}
			return leftVal.Value < rightVal.Value, nil
		}
	case "EQUALS":
		return leftVal.Value == rightVal.Value, nil
	}
	return false, fmt.Errorf("incorrect comparison operands with types %v and %v and values '%s' and '%s'", leftVal.ValType, rightVal.ValType, leftVal.Value, rightVal.Value)
}

func (interpreter *Interpreter) processPlay(module *ivr.PlayModule) (ivr.Module, error) {
	interpreter.addResponseText(module.VoicePromptIDs)
	return getModuleByID(interpreter.Script, module.GetDescendant())
}

func (interpreter *Interpreter) processSetVariable(module *ivr.SetVariableModule) (ivr.Module, error) {
	for _, expr := range module.Exprs {
		lval := interpreter.Session.getParameter(expr.Lval)
		switch expr.Rval.FuncDef {
		case "STRING#PUT#KVLIST#STRING#STRING":
			kvlistvar := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			key := interpreter.Session.getParameter(string(expr.Rval.Params[1]))
			val := interpreter.Session.getParameter(string(expr.Rval.Params[2]))
			if lval.ValType != ivr.ValString || kvlistvar.ValType != ivr.ValKVList || key.ValType != ivr.ValString || val.ValType != ivr.ValString {
				log.Panicf("Type mismatch") // TODO make verbose
			}
			kvlist := ivr.StringToKVList(kvlistvar.Value)
			lval.Value = kvlist.Put(key.Value, val.Value)
			kvlistvar.Value = kvlist.ToString()
		case "INTEGER#SIZE#KVLIST":
			kvlistvar := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			if lval.ValType != ivr.ValInteger || kvlistvar.ValType != ivr.ValKVList {
				log.Panicf("Type mismatch")
			}
			kvlist := ivr.StringToKVList(kvlistvar.Value)
			lval.Value = strconv.Itoa(len(*kvlist))
		case "STRING#GET_KEY#KVLIST#INTEGER":
			kvlistvar := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			index := interpreter.Session.getParameter(string(expr.Rval.Params[1]))
			if lval.ValType != ivr.ValString || kvlistvar.ValType != ivr.ValKVList || index.ValType != ivr.ValInteger {
				log.Panicf("Type mismatch")
			}
			kvlist := ivr.StringToKVList(kvlistvar.Value)
			i, err := strconv.Atoi(index.Value)
			if err != nil {
				panic(err)
			}
			lval.Value = kvlist.GetKey(i)
		case "INTEGER#SUM#INTEGER#INTEGER":
			int1 := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			int2 := interpreter.Session.getParameter(string(expr.Rval.Params[1]))
			if lval.ValType != ivr.ValInteger || int1.ValType != ivr.ValInteger || int2.ValType != ivr.ValInteger {
				log.Panicf("Type mismatch")
			}
			i1, err := strconv.Atoi(int1.Value)
			if err != nil {
				panic(err)
			}
			i2, err := strconv.Atoi(int2.Value)
			if err != nil {
				panic(err)
			}
			lval.Value = strconv.Itoa(i1 + i2)
		case "STRING#REMOVE#KVLIST#STRING":
			kvlistvar := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			key := interpreter.Session.getParameter(string(expr.Rval.Params[1]))
			if lval.ValType != ivr.ValString || kvlistvar.ValType != ivr.ValKVList || key.ValType != ivr.ValString {
				log.Panicf("Type mismatch")
			}
			kvlist := ivr.StringToKVList(kvlistvar.Value)
			lval.Value = kvlist.Remove(key.Value)
			kvlistvar.Value = kvlist.ToString()
		case "__COPY__":
			source := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			lval.Value = source.Value
		case "STRING#TOSTRING#INTEGER":
			integer := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			if lval.ValType != ivr.ValString || integer.ValType != ivr.ValInteger {
				log.Panicf("Type mismatch")
			}
			lval.Value = integer.Value
		case "CURRENCY#SUM#CURRENCY#CURRENCY":
			currency1 := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			currency2 := interpreter.Session.getParameter(string(expr.Rval.Params[1]))
			if lval.ValType != ivr.ValCurrency || currency1.ValType != ivr.ValCurrency || currency2.ValType != ivr.ValCurrency {
				log.Panicf("Type mismatch")
			}
			amount1, err := strconv.ParseFloat(currency1.Value[3:], 64)
			if err != nil {
				panic(err)
			}
			amount2, err := strconv.ParseFloat(currency2.Value[3:], 64)
			if err != nil {
				panic(err)
			}
			total, err := ivr.NewUSCurrencyValue(amount1 + amount2)
			if err != nil {
				panic(err)
			}
			lval.Value = total
		case "STRING#CONCAT#STRING#STRING":
			string1 := interpreter.Session.getParameter(string(expr.Rval.Params[0]))
			string2 := interpreter.Session.getParameter(string(expr.Rval.Params[1]))
			if lval.ValType != ivr.ValString || string1.ValType != ivr.ValString || string2.ValType != ivr.ValString {
				log.Panicf("Type mismatch")
			}
			lval.Value = string1.Value + string2.Value
		default:
			panic("Unsupported function ID: " + expr.Rval.FuncDef)
		}
	}
	return getModuleByID(interpreter.Script, module.GetDescendant())
}

func (interpreter *Interpreter) processHangup(module *ivr.HangupModule) (ivr.Module, error) {
	config, err := cloud.GetConfig()
	if err != nil {
		return nil, err
	}
	err = createTermination(config[cDomainNameConfigKey], config[cCampaignNameConfigKey], interpreter.makeModuleParams(module))
	if err != nil {
		return nil, err
	}
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
