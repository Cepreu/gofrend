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
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// PrepareFile creates necessary DialogFlow intents and uploads XML to gcp storage
func PrepareFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = cloud.UploadXML(data)
	if err != nil {
		return err
	}

	script, err := xmlparser.NewIVRScript(bytes.NewReader(data))
	if err != nil {
		return err
	}

	return prepareIntents(script)
}

func prepareIntents(script *ivr.IVRScript) error {
	intents, err := intentsGenerator(script)
	if err != nil {
		return err
	}

	projectID := "f9-test-agent"
	ctx := context.Background()
	client, err := dialogflow.NewIntentsClient(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		return err
	}

	for _, intent := range intents {
		request := &dialogflowpb.CreateIntentRequest{
			Parent: fmt.Sprintf("projects/%s/agent", projectID),
			Intent: intent,
		}

		_, err = client.CreateIntent(ctx, request)
		if err != nil {
			return err
		}
	}
	return nil
}
