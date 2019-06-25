package fulfiller

import (
	"bytes"
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
	script, err := getScript("ivr_scripts/is_large_test.five9ivr")
	if err != nil {
		log.Panic(err)
	}
	utils.PrettyLog(script)

	wr := dialogflowpb.WebhookRequest{}
	err = jsonpb.Unmarshal(r.Body, &wr)
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

func getScript(hash string) (*ivr.IVRScript, error) {
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
