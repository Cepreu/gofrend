package dialogflow

import (
	"fmt"
	"log"

	ivr "github.com/Cepreu/gofrend/ivrparser"
	"github.com/Cepreu/gofrend/utils"
	"github.com/davecgh/go-spew/spew"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func (dp *DialogflowProcessor) GetIntent(intentID string) string {
	req := dialogflowpb.GetIntentRequest{
		// Required. The name of the intent.
		// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
		Name: fmt.Sprintf("projects/%s/agent/intents/%s", dp.projectID, intentID),
		// Optional. The language to retrieve training phrases, parameters and rich
		// messages for. If not specified, the agent's default language is used.
		LanguageCode: "en",
		// // Optional. The resource view to apply to the returned intent.
		// IntentView: IntentView{},
	}
	response, err := dp.intentClient.GetIntent(dp.ctx, &req)
	if err != nil {
		log.Fatalf("Error in communication with Dialogflow %s", err.Error())
		return ""
	}
	utils.PrettyPrint(response)
	return spew.Sdump(response)
}

func (dp *DialogflowProcessor) CreateIntent() string {
	req := dialogflowpb.CreateIntentRequest{
		//TODO: fill request struct fields
		Parent: `projects/marysbikeshop-f198d/agent`,
		Intent: &dialogflowpb.Intent{
			DisplayName: "Menu1",
			// Required. Indicates whether webhooks are enabled for the intent.
			WebhookState: (dialogflowpb.Intent_WebhookState(
				dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_UNSPECIFIED"])),
			// Optional. The priority of this intent. Higher numbers represent higher
			// priorities. Zero or negative numbers mean that the intent is disabled.
			Priority: 500000,
			// Optional. Indicates whether this is a fallback intent.
			IsFallback: false,
			// Optional. Indicates whether Machine Learning is disabled for the intent.
			// Note: If `ml_diabled` setting is set to true, then this intent is not
			// taken into account during inference in `ML ONLY` match mode. Also,
			// auto-markup in the UI is turned off.
			MlDisabled: false,
			// Optional. The list of context names required for this intent to be
			// triggered.
			// Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context ID>`.
			InputContextNames: []string{},
			// Optional. The collection of event names that trigger the intent.
			// If the collection of input contexts is not empty, all of the contexts must
			// be present in the active user session for an event to trigger this intent.
			Events: []string{"MENU2"},
			// Optional. The collection of examples/templates that the agent is
			// trained on.
			TrainingPhrases: []*dialogflowpb.Intent_TrainingPhrase{
				&dialogflowpb.Intent_TrainingPhrase{
					Name: utils.GenUUIDv4(),
					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
						&dialogflowpb.Intent_TrainingPhrase_Part{Text: "Book me a sedan tomorrow"},
					},
				},
				&dialogflowpb.Intent_TrainingPhrase{
					Name: utils.GenUUIDv4(),
					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
						&dialogflowpb.Intent_TrainingPhrase_Part{Text: "I need an SUV in Moscow"},
					},
				},
				&dialogflowpb.Intent_TrainingPhrase{
					Name: utils.GenUUIDv4(),
					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
						&dialogflowpb.Intent_TrainingPhrase_Part{Text: "I need a sports car today in London"},
					},
				},
				&dialogflowpb.Intent_TrainingPhrase{
					Name: utils.GenUUIDv4(),
					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
						&dialogflowpb.Intent_TrainingPhrase_Part{Text: "Reserve a sports car"},
					},
				},
				&dialogflowpb.Intent_TrainingPhrase{
					Name: utils.GenUUIDv4(),
					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
						&dialogflowpb.Intent_TrainingPhrase_Part{Text: "Book me a convertible tomorrow"},
					},
				},
				&dialogflowpb.Intent_TrainingPhrase{
					Name: utils.GenUUIDv4(),
					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
						&dialogflowpb.Intent_TrainingPhrase_Part{Text: "Can you book me a car in Berlin?"},
					},
				},
			},
			// Optional. The name of the action associated with the intent.
			Action: "input.welcome",
			// Optional. The collection of contexts that are activated when the intent
			// is matched. Context messages in this collection should not set the
			// parameters field. Setting the `lifespan_count` to 0 will reset the context
			// when the intent is matched.
			// Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context ID>`.
			OutputContexts: []*dialogflowpb.Context{},
			// Optional. Indicates whether to delete all contexts in the current
			// session when this intent is matched.
			ResetContexts: false,
			// Optional. The collection of parameters associated with the intent.
			Parameters: []*dialogflowpb.Intent_Parameter{},
			// Optional. The collection of rich messages corresponding to the
			// `Response` field in the Dialogflow console.
			Messages: []*dialogflowpb.Intent_Message{
				{
					Message: &dialogflowpb.Intent_Message_Text_{
						Text: &dialogflowpb.Intent_Message_Text{
							Text: []string{
								"Welcome. I can tell you the shop hours, or I can set up an appointment. Which would you like?",
								"Welcome. I can tell you the shop hours, or I can make an appointment. What can I do for you?",
								"Hello there. I can tell you the shop hours, or I can schedule an appointment. How may I help you today?",
							},
						},
					},
				},
			},
			// Optional. The list of platforms for which the first response will be
			// taken from among the messages assigned to the DEFAULT_PLATFORM.
			DefaultResponsePlatforms: []dialogflowpb.Intent_Message_Platform{},
			// The unique identifier of the root intent in the chain of followup intents.
			// It identifies the correct followup intents chain for this intent.
			// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
			RootFollowupIntentName: "",
			// The unique identifier of the parent intent in the chain of followup
			// intents.
			// It identifies the parent followup intent.
			// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
			ParentFollowupIntentName: "",
			// Optional. Collection of information about all followup intents that have
			// name of this intent as a root_name.
			FollowupIntentInfo: []*dialogflowpb.Intent_FollowupIntentInfo{},
		},
	}
	response, err := dp.intentClient.CreateIntent(dp.ctx, &req)
	if err != nil {
		log.Fatalf("Error in communication with Dialogflow %s", err.Error())
		return ""
	}
	utils.PrettyPrint(response)
	return spew.Sdump(response)
}

func intentsGenerator(ivrScript *ivr.IVRScript) (intents []*dialogflowpb.Intent, err error) {
	for _, m := range ivrScript.Modules {
		switch v := m.(type) {
		case *ivr.MenuModule:
			return menu2intents(ivrScript, v)
		}
	}
	return nil, nil
}

func menu2intents(ivrScript *ivr.IVRScript, menu *ivr.MenuModule) (intents []*dialogflowpb.Intent, err error) {
	mainIntent := &dialogflowpb.Intent{
		DisplayName: menu.Name,
		WebhookState: (dialogflowpb.Intent_WebhookState(
			dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_UNSPECIFIED"])),
		Priority:                 500000,
		MlDisabled:               false,
		InputContextNames:        []string{},
		Events:                   []string{"MENU2"},
		TrainingPhrases:          []*dialogflowpb.Intent_TrainingPhrase{},
		Action:                   "input.welcome",
		OutputContexts:           []*dialogflowpb.Context{},
		Parameters:               []*dialogflowpb.Intent_Parameter{},
		DefaultResponsePlatforms: []dialogflowpb.Intent_Message_Platform{},
		FollowupIntentInfo:       []*dialogflowpb.Intent_FollowupIntentInfo{},
	}

	im := dialogflowpb.Intent_Message{
		Message: &dialogflowpb.Intent_Message_Text_{
			Text: &dialogflowpb.Intent_Message_Text{
				Text: menu.VoicePromptIDs.TransformToAI(ivrScript.Prompts),
			},
		},
	}
	mainIntent.Messages = append(mainIntent.Messages, &im)
	intents = append(intents, mainIntent)

	/// Add an intent per a branch:
	intntPerBranch := make(map[string][]string)
	for _, i := range menu.Items {
		intntPerBranch[i.Action.Name] = append(intntPerBranch[i.Action.Name], i.Prompt.TransformToAI(ivrScript.Prompts))
	}

	for _, br := range menu.Branches {
		b := br.Name
		tfs := []*dialogflowpb.Intent_TrainingPhrase{}
		for _, t := range intntPerBranch[b] {
			tfs = append(tfs,
				&dialogflowpb.Intent_TrainingPhrase{
					Name: utils.GenUUIDv4(),
					Type: dialogflowpb.Intent_TrainingPhrase_Type(1),
					Parts: []*dialogflowpb.Intent_TrainingPhrase_Part{
						&dialogflowpb.Intent_TrainingPhrase_Part{Text: t},
					},
				})
		}

		branchIntent := &dialogflowpb.Intent{
			DisplayName: menu.Name + " " + b,
			WebhookState: (dialogflowpb.Intent_WebhookState(
				dialogflowpb.Intent_WebhookState_value["WEBHOOK_STATE_UNSPECIFIED"])),
			Priority:                 500000,
			IsFallback:               false,
			MlDisabled:               false,
			InputContextNames:        []string{},
			Events:                   []string{"MENU3"},
			TrainingPhrases:          tfs,
			Action:                   "input." + b,
			OutputContexts:           []*dialogflowpb.Context{},
			ResetContexts:            false,
			Parameters:               []*dialogflowpb.Intent_Parameter{},
			DefaultResponsePlatforms: []dialogflowpb.Intent_Message_Platform{},
			RootFollowupIntentName:   "",
			ParentFollowupIntentName: "",
			FollowupIntentInfo:       []*dialogflowpb.Intent_FollowupIntentInfo{},
		}

		buildIntentBranch(ivrScript, branchIntent, br.Descendant)
		intents = append(intents, branchIntent)
	}
	return
}

func buildIntentBranch(ivrScript *ivr.IVRScript, branchIntent *dialogflowpb.Intent, nextModule ivr.ModuleID) {
	m := ivrScript.Modules[nextModule]
	switch v := m.(type) {
	case *ivr.InputModule:
		var param = &dialogflowpb.Intent_Parameter{
			Name:        utils.GenUUIDv4(),
			DisplayName: v.Name,
			// Optional. The definition of the parameter value. It can be:
			// - a constant string,
			// - a parameter value defined as `$parameter_name`,
			// - an original parameter value defined as `$parameter_name.original`,
			// - a parameter value from some context defined as
			//   `#context_name.parameter_name`.
			Value: "$" + v.Grammar.MRVvariable,
			// Optional. The default value to use when the `value` yields an empty
			// result.
			// Default values can be extracted from contexts by using the following
			// syntax: `#context_name.parameter_name`.
			//DefaultValue string
			// Optional. The name of the entity type, prefixed with `@`, that
			// describes values of the parameter. If the parameter is
			// required, this must be provided.
			EntityTypeDisplayName: "@" + v.Grammar.MRVtype,
			Mandatory:             true,
			// Optional. The collection of prompts that the agent can present to the
			// user in order to collect value for the parameter.
			Prompts: v.VoicePromptIDs.TransformToAI(ivrScript.Prompts),
			IsList:  false,
		}
		branchIntent.Parameters = append(branchIntent.Parameters, param)
		//				if v.Collapsible {
		// Recursively check the next module
		buildIntentBranch(ivrScript, branchIntent, v.Descendant)
		//				}

	case *ivr.PlayModule:
		msg := dialogflowpb.Intent_Message{
			Message: &dialogflowpb.Intent_Message_Text_{
				Text: &dialogflowpb.Intent_Message_Text{
					Text: v.VoicePromptIDs.TransformToAI(ivrScript.Prompts),
				},
			},
		}
		branchIntent.Messages = append(branchIntent.Messages, &msg)
		//				buildIntentBranch(ivrScript, branchIntent, im.GetDescendant())
		fmt.Println("~~~~~~~~~~~~~~~~~~", v.VoicePromptIDs.TransformToAI(ivrScript.Prompts))
	default:
		buildIntentBranch(ivrScript, branchIntent, m.GetDescendant())
	}
}
