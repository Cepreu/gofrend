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
	wr := dialogflowpb.WebhookRequest{}
	err := jsonpb.Unmarshal(r.Body, &wr)
	if err != nil {
		log.Panic(err)
	}

	hash := getScriptHash(wr)
	storageScript, err := cloud.DownloadScript(hash)
	if err != nil {
		log.Panic(err)
	}
	script := storageScript.GetScript()

	response, err := Interpret(wr, script, hash)
	if err != nil {
		log.Panic(err)
	}

	marshaler := jsonpb.Marshaler{}
	marshaler.Marshal(w, response)
	log.Print(marshaler.MarshalToString(response))
}

func getScriptHash(webhookRequest dialogflowpb.WebhookRequest) string {
	return utils.DisplayNameToScriptHash(webhookRequest.QueryResult.Intent.DisplayName)
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
