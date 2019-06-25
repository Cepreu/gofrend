package preparer

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	ivr "github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/xmlparser"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// PrepareFile creates necessary DialogFlow intents and uploads XML to google storage
func PrepareFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	script, err := xmlparser.NewIVRScript(bufio.NewReader(f))
	if err != nil {
		return err
	}
	prepareIntents(script)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return uploadData(data)
}

func prepareIntents(script *ivr.IVRScript) error {
	intents, err := intentsGenerator(script)
	if err != nil {
		return err
	}

	projectID := "f9-test-agent"
	ctx := context.Background()
	client, err := dialogflow.NewIntentsClient(ctx, option.WithCredentialsFile("F9-Test-Agent-0925974a682a.json"))
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
