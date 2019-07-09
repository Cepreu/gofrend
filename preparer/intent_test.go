package preparer

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/Cepreu/gofrend/cloud"
	"github.com/Cepreu/gofrend/ivr/xmlparser"
	"github.com/Cepreu/gofrend/utils"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func checkNil(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntent(t *testing.T) {
	var filename = "test_files/is_large_test.five9ivr"
	data, err := ioutil.ReadFile(filename)
	checkNil(err, t)
	scriptHash := utils.HashToString(data)
	script, err := xmlparser.NewIVRScript(bytes.NewReader(data))
	checkNil(err, t)
	intents, err := generateIntents(script, scriptHash)
	checkNil(err, t)

	request := &dialogflowpb.CreateIntentRequest{
		Parent: fmt.Sprintf("projects/%s/agent", cloud.GcpProjectID),
		Intent: intents[0],
	}

	ctx := context.Background()
	client, err := dialogflow.NewIntentsClient(ctx, option.WithCredentialsFile(cloud.GcpCredentialsFileName))
	checkNil(err, t)

	softDeleteIntent(ctx, client, request)

	_, err = client.CreateIntent(ctx, request)
	checkNil(err, t)
}

func TestComparisonIntent(t *testing.T) {
	var filename = "test_files/comparison_test.five9ivr"
	data, err := ioutil.ReadFile(filename)
	checkNil(err, t)
	scriptHash := utils.HashToString(data)
	script, err := xmlparser.NewIVRScript(bytes.NewReader(data))
	checkNil(err, t)
	intents, err := generateIntents(script, scriptHash)
	checkNil(err, t)
	if len(intents) != 2 {
		t.Fatalf("Len of intents expected to be 2, instead: %d", len(intents))
	}

	ctx := context.Background()
	client, err := dialogflow.NewIntentsClient(ctx, option.WithCredentialsFile(cloud.GcpCredentialsFileName))
	checkNil(err, t)
	for _, intent := range intents {
		request := &dialogflowpb.CreateIntentRequest{
			Parent: fmt.Sprintf("projects/%s/agent", cloud.GcpProjectID),
			Intent: intent,
		}
		softDeleteIntent(ctx, client, request)
		checkNil(err, t)
		_, err = client.CreateIntent(ctx, request)
		checkNil(err, t)
	}
}
