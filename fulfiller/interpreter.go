package fulfiller

import (
	"log"

	ivr "github.com/Cepreu/gofrend/ivrparser"
	"github.com/Cepreu/gofrend/vars"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// Interpreter can process modules and populate its WebhookResponse as necessary
type Interpreter struct {
	Session         *Session
	Script          *ivr.IVRScript
	QueryResult     *dialogflowpb.QueryResult
	WebhookResponse *dialogflowpb.WebhookResponse
}

// ProcessInitial process a module and returns the next module to be processed
func (interpreter *Interpreter) ProcessInitial(module ivr.Module) ivr.Module {
	switch v := module.(type) {
	case *ivr.InputModule:
		return interpreter.processInputInitial(v)
	}
	return nil
}

// Process processes a module and returns the next module to be processed
func (interpreter *Interpreter) Process(module ivr.Module) ivr.Module {
	switch v := module.(type) {
	case *ivr.IfElseModule:
		return interpreter.processIfElse(v)
	case *ivr.PlayModule:
		return interpreter.processPlay(v)
	case *ivr.HangupModule:
		return nil
	}
	return nil
}

func (interpreter *Interpreter) processInputInitial(module *ivr.InputModule) ivr.Module {
	parameters := interpreter.QueryResult.Parameters.Fields
	for name, value := range parameters {
		interpreter.Session.setParameter(name, value)
	}
	return getModuleByID(interpreter.Script, string(module.GetDescendant()))
}

func (interpreter *Interpreter) processIfElse(module *ivr.IfElseModule) (next ivr.Module) {
	var conditionsPass bool
	conditions := module.BranchIf.Cond.Conditions
	if module.BranchIf.Cond.ConditionGrouping == "" { // Parser currently does not populate this field
		module.BranchIf.Cond.ConditionGrouping = "ALL"
	}
	switch module.BranchIf.Cond.ConditionGrouping {
	case "ALL":
		conditionsPass = true
		for _, condition := range conditions {
			populateCondition(condition, interpreter.Session.Variables)
			if !passes(condition) {
				conditionsPass = false
			}
		}
	}
	if conditionsPass {
		next = getModuleByID(interpreter.Script, string(module.BranchIf.Descendant))
	} else {
		next = getModuleByID(interpreter.Script, string(module.BranchElse.Descendant))
	}
	return
}

func populateCondition(condition *ivr.Condition, variables map[string]*vars.Variable) {
	varName := condition.LeftOperand.VariableName
	variable, ok := variables[varName]
	if !ok {
		log.Fatalf("Error finding session variable with name: %s", varName)
	}
	condition.LeftOperand.Value = variable.Value
	varName = condition.RightOperand.VariableName
	variable, ok = variables[varName]
	if !ok {
		log.Fatalf("Error finding session variable with name: %s", varName)
	}
	condition.RightOperand.Value = variable.Value
}

func passes(condition *ivr.Condition) bool {
	log.Printf("Comparison Type: %s", condition.ComparisonType)
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
		switch v := condition.LeftOperand.Value.(type) {
		case *vars.Integer:
			right = float64(v.Value)
		case *vars.Numeric:
			right = v.Value
		default:
			log.Fatalf("Expected int or numeric in more than comparison, instead got: %T", v)
		}
		log.Printf("Left: %f, Right: %f.", left, right)
		return left > right
	}
	return false
}

func (interpreter *Interpreter) processPlay(module *ivr.PlayModule) ivr.Module {
	promptStrings := module.VoicePromptIDs.TransformToAI(interpreter.Script.Prompts)
	intentMessageText := interpreter.WebhookResponse.FulfillmentMessages[0].GetText()
	intentMessageText.Text = append(intentMessageText.Text, promptStrings...)
	return getModuleByID(interpreter.Script, string(module.GetDescendant()))
}
