package fulfiller

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/xmlparser"
	"github.com/Cepreu/gofrend/utils"
	"github.com/golang/protobuf/jsonpb"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// HandleWebhook performs DialogFlow fulfillment for the F9 Agent
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	script := getScript("ivr_scripts/is_large_test.five9ivr")
	utils.PrettyLog(script)

	wr := dialogflowpb.WebhookRequest{}
	err := jsonpb.Unmarshal(r.Body, &wr)
	if err != nil {
		log.Fatalf("Error unmarshalling dialogflow request: %v", err)
	}

	response := Interpret(wr, script)

	marshaler := jsonpb.Marshaler{}
	marshaler.Marshal(w, response)

	//Need to implement storing session
}

func getScript(fname string) *ivr.IVRScript {
	f, err := os.Open(fname)
	if err != nil {
		log.Panic(err)
	}
	script, err := xmlparser.NewIVRScript(bufio.NewReader(f))
	if err != nil {
		log.Panic(err)
	}
	return script
}

func getModuleByID(script *ivr.IVRScript, ID string) (m ivr.Module) { // probably an unnecessary function
	m, ok := script.Modules[ivr.ModuleID(ID)]
	if !ok {
		log.Print("No match.") //Debug
		utils.PrettyLog(script.Modules)
		return nil
	}
	return m
}

func isInputOrMenu(module ivr.Module) bool {
	switch module.(type) {
	case *ivr.MenuModule:
		return true
	case *ivr.InputModule:
		return true
	default:
		return false
	}
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
