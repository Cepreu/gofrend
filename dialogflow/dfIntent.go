package main

import (
	"fmt"
	"log"

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
