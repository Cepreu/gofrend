package fulfiller

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"testing"

	ivr "github.com/Cepreu/gofrend/ivrparser"
	"github.com/Cepreu/gofrend/utils"
	"github.com/golang/protobuf/jsonpb"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

var script *ivr.IVRScript

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestScript(t *testing.T) {
	var fname = "test_files/is_large_test.five9ivr"
	f, err := os.Open(fname)
	check(err)
	script, err = ivr.NewIVRScript(bufio.NewReader(f))
	check(err)
	utils.PrettyPrint(script)
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fname := "test_files/is_large_test.five9ivr"
	f, _ := os.Open(fname)
	script, _ = ivr.NewIVRScript(bufio.NewReader(f))

	wr := dialogflowpb.WebhookRequest{}
	err := jsonpb.Unmarshal(r.Body, &wr)
	if err != nil {
		log.Fatalf("Error unmarshalling dialogflow request: %v", err)
	}

	value, ok := wr.QueryResult.Parameters.Fields["value"]
	if !ok {
		log.Fatal("Request does not contain parameter 'value'")
	}
	valueFloat := value.GetNumberValue()
	var messageText string
	if valueFloat > 9 {
		messageText = "Your number is large!"
	} else {
		messageText = "Your number is small!"
	}
	webhookResponse := dialogflowpb.WebhookResponse{
		FulfillmentMessages: []*dialogflowpb.Intent_Message{
			{
				Message: &dialogflowpb.Intent_Message_Text_{
					Text: &dialogflowpb.Intent_Message_Text{
						Text: []string{messageText},
					},
				},
			},
		},
	}
	marshaler := jsonpb.Marshaler{}
	marshaler.Marshal(w, &webhookResponse)
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
