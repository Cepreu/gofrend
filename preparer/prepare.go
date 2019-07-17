package preparer

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/Cepreu/gofrend/cloud"
	ivr "github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/xmlparser"
	"github.com/Cepreu/gofrend/utils"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// Prepare creates necessary Dialogflow intents and uploads XML to gcp storage
func Prepare(data []byte) error {
	err := cloud.UploadXML(data)
	if err != nil {
		return err
	}

	script, err := xmlparser.NewIVRScript(bytes.NewReader(data))
	utils.PrettyLog(script)
	if err != nil {
		return err
	}

	scriptHash := utils.HashToString(data)

	return prepareIntents(script, scriptHash)
}

// PrepareFile creates necessary DialogFlow intents and uploads XML to gcp storage
func PrepareFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return Prepare(data)
}

func prepareIntents(script *ivr.IVRScript, scriptHash string) error {
	intents, err := generateIntents(script, scriptHash)
	if err != nil {
		return fmt.Errorf("Error generating intents: %v", err)
	}

	ctx := context.Background()
	client, err := dialogflow.NewIntentsClient(ctx, option.WithCredentialsFile(cloud.GcpCredentialsFileName))
	if err != nil {
		return err
	}

	for _, intent := range intents {
		request := &dialogflowpb.CreateIntentRequest{
			Parent: fmt.Sprintf("projects/%s/agent", cloud.GcpProjectID),
			Intent: intent,
		}

		softDeleteIntent(ctx, client, request)
		if err != nil {
			return err
		}

		_, err = client.CreateIntent(ctx, request)
		if err != nil {
			return err
		}
	}
	return nil
}

func softDeleteIntent(ctx context.Context, client *dialogflow.IntentsClient, request *dialogflowpb.CreateIntentRequest) {
	intentsIterator := client.ListIntents(ctx, &dialogflowpb.ListIntentsRequest{Parent: request.Parent})
	intent, err := intentsIterator.Next()
	for err == nil && intent != nil {
		if intent.DisplayName == request.Intent.DisplayName {
			client.DeleteIntent(ctx, &dialogflowpb.DeleteIntentRequest{Name: intent.Name})
			break
		}
		intent, err = intentsIterator.Next()
	}
}
