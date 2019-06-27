package preparer

import (
	"fmt"

	ivr "github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func intentsGenerator(ivrScript *ivr.IVRScript, scriptHash string) (intents []*dialogflowpb.Intent, err error) {
	for _, m := range ivrScript.Modules {
		switch v := m.(type) {
		case *ivr.InputModule:
			intent, err := input2intent(ivrScript, v, scriptHash)
			return []*dialogflowpb.Intent{intent}, err
		}
	}
	return nil, nil
}

func input2intent(ivrScript *ivr.IVRScript, input *ivr.InputModule, scriptHash string) (intent *dialogflowpb.Intent, err error) {
	displayName := utils.MakeDisplayName(scriptHash, input.GetID())
	intent = &dialogflowpb.Intent{
		DisplayName: displayName,
		WebhookState: (dialogflowpb.Intent_WebhookState(
			dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_ENABLED"])),
		Priority:          500000,
		MlDisabled:        false,
		InputContextNames: []string{fmt.Sprintf("projects/f9-dialogflow-converter/agent/sessions/-/contexts/%s", displayName)},
		Events:            []string{},
		TrainingPhrases: []*dialogflowpb.Intent_TrainingPhrase{
			&dialogflowpb.Intent_TrainingPhrase{
				Name: utils.GenUUIDv4(),
				Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
				Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
					&dialogflowpb.Intent_TrainingPhrase_Part{
						Text:        "1",
						EntityType:  "@sys.number",
						Alias:       "Value",
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
						Alias:       "Value",
						UserDefined: true,
					},
				},
			},
		},
		Action:         "input.sayvalue",
		OutputContexts: []*dialogflowpb.Context{},
		Parameters: []*dialogflowpb.Intent_Parameter{
			{
				Name:                  utils.GenUUIDv4(),
				DisplayName:           "Value",
				Value:                 "$Value",
				EntityTypeDisplayName: "@sys.number",
				Mandatory:             true,
			},
		},
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
	return
}
