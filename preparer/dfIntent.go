package preparer

import (
	ivr "github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func intentsGenerator(ivrScript *ivr.IVRScript) (intents []*dialogflowpb.Intent, err error) {
	for _, m := range ivrScript.Modules {
		switch v := m.(type) {
		case *ivr.InputModule:
			intent, err := input2intent(ivrScript, v)
			return []*dialogflowpb.Intent{intent}, err
		}
	}
	return nil, nil
}

func input2intent(ivrScript *ivr.IVRScript, input *ivr.InputModule) (intent *dialogflowpb.Intent, err error) {
	intent = &dialogflowpb.Intent{
		DisplayName: string(input.GetID()),
		WebhookState: (dialogflowpb.Intent_WebhookState(
			dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_UNSPECIFIED"])),
		Priority:          500000,
		MlDisabled:        false,
		InputContextNames: []string{},
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

// func menu2intents(ivrScript *ivr.IVRScript, menu *ivr.MenuModule) (intents []*dialogflowpb.Intent, err error) {
// 	mainIntent := &dialogflowpb.Intent{
// 		DisplayName: menu.Name,
// 		WebhookState: (dialogflowpb.Intent_WebhookState(
// 			dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_UNSPECIFIED"])),
// 		Priority:                 500000,
// 		MlDisabled:               false,
// 		InputContextNames:        []string{},
// 		Events:                   []string{"MENU2"},
// 		TrainingPhrases:          []*dialogflowpb.Intent_TrainingPhrase{},
// 		Action:                   "input.welcome",
// 		OutputContexts:           []*dialogflowpb.Context{},
// 		Parameters:               []*dialogflowpb.Intent_Parameter{},
// 		DefaultResponsePlatforms: []dialogflowpb.Intent_Message_Platform{},
// 		FollowupIntentInfo:       []*dialogflowpb.Intent_FollowupIntentInfo{},
// 	}

// 	im := dialogflowpb.Intent_Message{
// 		Message: &dialogflowpb.Intent_Message_Text_{
// 			Text: &dialogflowpb.Intent_Message_Text{
// 				Text: menu.VoicePromptIDs.TransformToAI(ivrScript.Prompts),
// 			},
// 		},
// 	}
// 	mainIntent.Messages = append(mainIntent.Messages, &im)
// 	intents = append(intents, mainIntent)

// 	/// Add an intent per a branch:
// 	intntPerBranch := make(map[string][]string)
// 	for _, i := range menu.Items {
// 		intntPerBranch[i.Action.Name] = append(intntPerBranch[i.Action.Name], i.Prompt.TransformToAI(ivrScript.Prompts))
// 	}

// 	for _, br := range menu.Branches {
// 		b := br.Name
// 		tfs := []*dialogflowpb.Intent_TrainingPhrase{}
// 		for _, t := range intntPerBranch[b] {
// 			tfs = append(tfs,
// 				&dialogflowpb.Intent_TrainingPhrase{
// 					Name: utils.GenUUIDv4(),
// 					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
// 					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
// 						&dialogflowpb.Intent_TrainingPhrase_Part{Text: t},
// 					},
// 				})
// 		}

// 		branchIntent := &dialogflowpb.Intent{
// 			DisplayName: menu.Name + " " + b,
// 			WebhookState: (dialogflowpb.Intent_WebhookState(
// 				dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_UNSPECIFIED"])),
// 			Priority:                 500000,
// 			IsFallback:               false,
// 			MlDisabled:               false,
// 			InputContextNames:        []string{},
// 			Events:                   []string{"MENU3"},
// 			TrainingPhrases:          tfs,
// 			Action:                   "input." + b,
// 			OutputContexts:           []*dialogflowpb.Context{},
// 			ResetContexts:            false,
// 			Parameters:               []*dialogflowpb.Intent_Parameter{},
// 			DefaultResponsePlatforms: []dialogflowpb.Intent_Message_Platform{},
// 			RootFollowupIntentName:   "",
// 			ParentFollowupIntentName: "",
// 			FollowupIntentInfo:       []*dialogflowpb.Intent_FollowupIntentInfo{},
// 		}

// 		buildIntentBranch(ivrScript, branchIntent, br.Descendant)
// 		intents = append(intents, branchIntent)
// 	}
// 	return
// }

// func buildIntentBranch(ivrScript *ivr.IVRScript, branchIntent *dialogflowpb.Intent, nextModule ivr.ModuleID) {
// 	m := ivrScript.Modules[nextModule]
// 	switch v := m.(type) {
// 	case *ivr.InputModule:
// 		var param = &dialogflowpb.Intent_Parameter{
// 			Name:        utils.GenUUIDv4(),
// 			DisplayName: v.Name,
// 			// Optional. The definition of the parameter value. It can be:
// 			// - a constant string,
// 			// - a parameter value defined as `$parameter_name`,
// 			// - an original parameter value defined as `$parameter_name.original`,
// 			// - a parameter value from some context defined as
// 			//   `#context_name.parameter_name`.
// 			Value: "$" + v.Grammar.MRVvariable,
// 			// Optional. The default value to use when the `value` yields an empty
// 			// result.
// 			// Default values can be extracted from contexts by using the following
// 			// syntax: `#context_name.parameter_name`.
// 			//DefaultValue string
// 			// Optional. The name of the entity type, prefixed with `@`, that
// 			// describes values of the parameter. If the parameter is
// 			// required, this must be provided.
// 			EntityTypeDisplayName: "@" + v.Grammar.MRVtype,
// 			Mandatory:             true,
// 			// Optional. The collection of prompts that the agent can present to the
// 			// user in order to collect value for the parameter.
// 			Prompts: v.VoicePromptIDs.TransformToAI(ivrScript.Prompts),
// 			IsList:  false,
// 		}
// 		branchIntent.Parameters = append(branchIntent.Parameters, param)
// 		//				if v.Collapsible {
// 		// Recursively check the next module
// 		buildIntentBranch(ivrScript, branchIntent, v.Descendant)
// 		//				}

// 	case *ivr.PlayModule:
// 		msg := dialogflowpb.Intent_Message{
// 			Message: &dialogflowpb.Intent_Message_Text_{
// 				Text: &dialogflowpb.Intent_Message_Text{
// 					Text: v.VoicePromptIDs.TransformToAI(ivrScript.Prompts),
// 				},
// 			},
// 		}
// 		branchIntent.Messages = append(branchIntent.Messages, &msg)
// 		//				buildIntentBranch(ivrScript, branchIntent, im.GetDescendant())
// 		fmt.Println("~~~~~~~~~~~~~~~~~~", v.VoicePromptIDs.TransformToAI(ivrScript.Prompts))
// 	default:
// 		buildIntentBranch(ivrScript, branchIntent, m.GetDescendant())
// 	}
// }
