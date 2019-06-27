package fulfiller

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/Cepreu/gofrend/cloud"
	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/xmlparser"
	"github.com/Cepreu/gofrend/utils"
	"github.com/golang/protobuf/jsonpb"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// HandleWebhook performs DialogFlow fulfillment for the F9 Agent
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	wr := dialogflowpb.WebhookRequest{}
	err := jsonpb.Unmarshal(r.Body, &wr)
	if err != nil {
		log.Panic(err)
	}

	script, err := getScript(wr)
	if err != nil {
		log.Panic(err)
	}

	response, err := Interpret(wr, script)
	if err != nil {
		log.Panic(err)
	}

	marshaler := jsonpb.Marshaler{}
	marshaler.Marshal(w, response)

	//Need to implement storing session
}

func getScriptHash(webhookRequest dialogflowpb.WebhookRequest) (string, error) {
	contextNames := webhookRequest.QueryResult.Intent.InputContextNames
	if len(contextNames) != 1 {
		return "", fmt.Errorf("Length of input contexts expected to be 1, in reality: %d", len(contextNames))
	}
	return utils.ContextToHash(contextNames[0])
}

func getScriptFromHash(hash string) (*ivr.IVRScript, error) {
	data, err := cloud.DownloadXML(hash)
	if err != nil {
		return nil, err
	}
	script, err := xmlparser.NewIVRScript(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return script, nil
}

func getScript(webhookRequest dialogflowpb.WebhookRequest) (*ivr.IVRScript, error) {
	hash, err := getScriptHash(webhookRequest)
	if err != nil {
		return nil, err
	}
	return getScriptFromHash(hash)
}

// func HandleWebhook(w http.ResponseWriter, r *http.Request) {
// 	wr := dialogflowpb.WebhookRequest{}
// 	err := jsonpb.Unmarshal(r.Body, &wr)
// 	if err != nil {
// 		log.Fatalf("Error unmarshalling dialogflow request: %v", err)
// 	}
// 	value, ok := wr.QueryResult.Parameters.Fields["value"]
// 	if !ok {
// 		log.Fatal("Request does not contain parameter 'value'")
// 	}
// 	valueFloat := value.GetNumberValue()
// 	var messageText string
// 	if valueFloat > 9 {
// 		messageText = "Your number is large!"
// 	} else {
// 		messageText = "Your number is small!"
// 	}
// 	webhookResponse := dialogflowpb.WebhookResponse{
// 		FulfillmentMessages: []*dialogflowpb.Intent_Message{
// 			{
// 				Message: &dialogflowpb.Intent_Message_Text_{
// 					Text: &dialogflowpb.Intent_Message_Text{
// 						Text: []string{messageText},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	marshaler := jsonpb.Marshaler{}
// 	marshaler.Marshal(w, &webhookResponse)
// }
