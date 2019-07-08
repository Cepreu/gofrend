package preparer

import (
	"fmt"

	ivr "github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func generateIntents(ivrScript *ivr.IVRScript, scriptHash string) ([]*dialogflowpb.Intent, error) {
	intents := []*dialogflowpb.Intent{}
	for _, module := range ivrScript.Modules {
		intent, err := createIntent(ivrScript, module, scriptHash)
		if err != nil {
			return nil, err
		}
		intents = append(intents, intent)
	}
	return intents, nil
}

func createIntent(ivrScript *ivr.IVRScript, module ivr.Module, scriptHash string) (*dialogflowpb.Intent, error) {
	displayName := utils.MakeDisplayName(scriptHash, module.GetID())
	trainingPhrases := createTrainingPhrases(module)
	parameters := createParameters(module)
	events := createEvents(module)
	intent := &dialogflowpb.Intent{
		DisplayName: displayName,
		WebhookState: (dialogflowpb.Intent_WebhookState(
			dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_ENABLED"])),
		Priority:                 500000,
		MlDisabled:               false,
		InputContextNames:        []string{utils.MakeContextName(displayName)},
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

func createTrainingPhrases(module ivr.Module) []*dialogflowpb.Intent_TrainingPhrase {
	switch v := module.(type) {
	case *ivr.IncomingCallModule:
		return []*dialogflowpb.Intent_TrainingPhrase{}
	case *ivr.InputModule:
		parameterName := v.Grammar.MRVvariable
		return []*dialogflowpb.Intent_TrainingPhrase{
			&dialogflowpb.Intent_TrainingPhrase{
				Name: utils.GenUUIDv4(),
				Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					&dialogflowpb.Intent_TrainingPhrase_Part{
						Text:        "1",
						EntityType:  "@sys.number",
						Alias:       parameterName,
						UserDefined: true,
					},
				},
			},
			&dialogflowpb.Intent_TrainingPhrase{
				Name: utils.GenUUIDv4(),
				Type: 1,
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					&dialogflowpb.Intent_TrainingPhrase_Part{
						Text: "My number is ",
					},
					&dialogflowpb.Intent_TrainingPhrase_Part{
						Text:        "2",
						EntityType:  "@sys.number",
						Alias:       parameterName,
						UserDefined: true,
					},
				},
			},
		}
	default:
		panic("Not implemented")
	}
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
	default:
		panic("Not implemented")
	}
}
