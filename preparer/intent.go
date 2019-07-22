package preparer

import (
	"fmt"
	"strings"

	ivr "github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func generateIntents(ivrScript *ivr.IVRScript, scriptHash string) ([]*dialogflowpb.Intent, error) {
	intents := []*dialogflowpb.Intent{}
	for _, module := range ivrScript.Modules {
		if menu, ok := module.(*ivr.MenuModule); ok {
			branches := menu.Branches
			for _, branch := range branches {
				intent, err := createIntent(ivrScript, menu, scriptHash, branch.Name)
				if err != nil {
					return nil, err
				}
				if intent != nil {
					intents = append(intents, intent)
				}
			}
		} else if requiresIntent(module) {
			intent, err := createIntent(ivrScript, module, scriptHash, "")
			if err != nil {
				return nil, err
			}
			intents = append(intents, intent)
		}
	}
	return intents, nil
}

func createIntent(ivrScript *ivr.IVRScript, module ivr.Module, scriptHash string, branchName string) (*dialogflowpb.Intent, error) {
	displayName := createDisplayName(module, scriptHash, branchName)
	trainingPhrases := createTrainingPhrases(module, ivrScript, branchName)
	if len(trainingPhrases) == 0 {
		return nil, nil
	}
	parameters := createParameters(module)
	events := createEvents(module)
	inputContextNames := createInputContextNames(module, displayName)
	intent := &dialogflowpb.Intent{
		DisplayName: displayName,
		WebhookState: (dialogflowpb.Intent_WebhookState(
			dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_ENABLED"])),
		Priority:                 500000,
		MlDisabled:               false,
		InputContextNames:        inputContextNames,
		Events:                   events,
		TrainingPhrases:          trainingPhrases,
		Action:                   "input.sayvalue",
		OutputContexts:           []*dialogflowpb.Context{},
		Parameters:               parameters,
		DefaultResponsePlatforms: []dialogflowpb.Intent_Message_Platform{},
		FollowupIntentInfo:       []*dialogflowpb.Intent_FollowupIntentInfo{},
		Messages: []*dialogflowpb.Intent_Message{
			{
				Message: &dialogflowpb.Intent_Message_Text_{
					Text: &dialogflowpb.Intent_Message_Text{
						Text: []string{
							"Dummy response.",
						},
					},
				},
			},
		},
	}
	return intent, nil
}

func createDisplayName(module ivr.Module, scriptHash string, branchName string) string {
	switch module.(type) {
	case *ivr.IncomingCallModule:
		return utils.MakeInputDisplayName(scriptHash, module.GetID())
	case *ivr.InputModule:
		return utils.MakeInputDisplayName(scriptHash, module.GetID())
	case *ivr.MenuModule:
		return utils.MakeMenuDisplayName(scriptHash, module.GetID(), branchName)
	case *ivr.SkillTransferModule:
		return utils.MakeInputDisplayName(scriptHash, module.GetID())
	default:
		panic("Not implemented")
	}
}

func createTrainingPhrases(module ivr.Module, script *ivr.IVRScript, branchName string) []*dialogflowpb.Intent_TrainingPhrase {
	switch v := module.(type) {
	case *ivr.IncomingCallModule:
		return []*dialogflowpb.Intent_TrainingPhrase{
			&dialogflowpb.Intent_TrainingPhrase{
				Name: utils.GenUUIDv4(),
				Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					&dialogflowpb.Intent_TrainingPhrase_Part{
						Text: "invoke",
					},
				},
			},
		}
	case *ivr.InputModule:
		parameterName := v.Grammar.MRVvariable
		return []*dialogflowpb.Intent_TrainingPhrase{
			&dialogflowpb.Intent_TrainingPhrase{
				Name: utils.GenUUIDv4(),
				Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					{
						Text:        "1",
						EntityType:  "@sys.number",
						Alias:       parameterName,
						UserDefined: true,
					},
				},
			},
			&dialogflowpb.Intent_TrainingPhrase{
				Name: utils.GenUUIDv4(),
				Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					{
						Text: "My number is ",
					},
					{
						Text:        "2",
						EntityType:  "@sys.number",
						Alias:       parameterName,
						UserDefined: true,
					},
				},
			},
		}
	case *ivr.MenuModule:
		phrases := []*dialogflowpb.Intent_TrainingPhrase{}
		for _, item := range v.Items {
			var phrase *dialogflowpb.Intent_TrainingPhrase
			if item.Action.Name == branchName {
				phrase = &dialogflowpb.Intent_TrainingPhrase{
					Name: utils.GenUUIDv4(),
					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
						{
							Text: getMenuItemText(item, script.Prompts),
						},
					},
				}
				phrases = append(phrases, phrase)
			}
		}
		return phrases
	case *ivr.SkillTransferModule:
		return []*dialogflowpb.Intent_TrainingPhrase{
			{
				Name: utils.GenUUIDv4(),
				Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					{
						Text:        "7818641522",
						EntityType:  "@sys.phone-number",
						Alias:       "callback-number",
						UserDefined: true,
					},
				},
			},
			{
				Name: utils.GenUUIDv4(),
				Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					{
						Text: "My number is ",
					},
					{
						Text:        "781-864-1234",
						EntityType:  "@sys.number",
						Alias:       "callback-number",
						UserDefined: true,
					},
				},
			},
			{
				Name: utils.GenUUIDv4(),
				Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					{
						Text: "My phone number is ",
					},
					{
						Text:        "3392232900",
						EntityType:  "@sys.number",
						Alias:       "callback-number",
						UserDefined: true,
					},
				},
			},
		}
	default:
		panic("Not implemented")
	}
}

func getMenuItemText(item *ivr.MenuItem, scriptPrompts ivr.ScriptPrompts) string {
	promptID := item.Prompt.LangPrArr[0].PrArr[0]
	prompt := scriptPrompts[promptID]
	return prompt.TransformToAI()
}

func createParameters(module ivr.Module) []*dialogflowpb.Intent_Parameter {
	switch v := module.(type) {
	case *ivr.IncomingCallModule:
		return []*dialogflowpb.Intent_Parameter{}
	case *ivr.InputModule:
		parameterName := v.Grammar.MRVvariable
		return []*dialogflowpb.Intent_Parameter{
			{
				Name:                  utils.GenUUIDv4(),
				DisplayName:           parameterName,
				Value:                 fmt.Sprintf("$%s", parameterName),
				EntityTypeDisplayName: "@sys.number",
				Mandatory:             true,
			},
		}
	case *ivr.MenuModule:
		return []*dialogflowpb.Intent_Parameter{}
	case *ivr.SkillTransferModule:
		return []*dialogflowpb.Intent_Parameter{
			{
				Name:                  utils.GenUUIDv4(),
				DisplayName:           "callback-number",
				Value:                 "$callback-number",
				EntityTypeDisplayName: "@sys.phone-number",
				Mandatory:             true,
			},
		}
	default:
		panic("Not implemented")
	}
}

func createEvents(module ivr.Module) []string {
	switch module.(type) {
	case *ivr.IncomingCallModule:
		return []string{"WELCOME"}
	case *ivr.InputModule:
		return []string{}
	case *ivr.MenuModule:
		return []string{}
	case *ivr.SkillTransferModule:
		return []string{}
	default:
		panic("Not implemented")
	}
}

func createInputContextNames(module ivr.Module, displayName string) []string {
	switch module.(type) {
	case *ivr.IncomingCallModule:
		return []string{}
	case *ivr.InputModule:
		return []string{utils.MakeInputContextName(displayName)}
	case *ivr.MenuModule:
		return []string{utils.MakeMenuContextName(displayName)}
	case *ivr.SkillTransferModule:
		return []string{utils.MakeInputContextName(displayName)}
	default:
		panic("Not implemented")
	}
}

func requiresIntent(module ivr.Module) bool {
	switch v := module.(type) {
	case *ivr.IncomingCallModule: // Eventually need parser to stop parsing pickup on hangup module
		if strings.Contains(v.Name, "Hangup") {
			return false
		}
		return true
	case *ivr.InputModule:
		return true
	case *ivr.SkillTransferModule:
		return true
	default:
		return false
	}
}
