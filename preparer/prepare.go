package preparer

import (
	"context"
	"fmt"
	"html"
	"strings"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/Cepreu/gofrend/cloud"
	ivr "github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/xmlparser"
	"github.com/Cepreu/gofrend/utils"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func Prepare(scriptName, campaignName, username, temporaryPassword string) error {
	auth := createAuthString(username, temporaryPassword)
	xmlDefinition, err := getIvrFromF9(auth, scriptName)
	if err != nil {
		return err
	}
	xmlDefinition = html.UnescapeString(xmlDefinition)

	script, err := xmlparser.NewIVRScript(strings.NewReader(xmlDefinition))
	if err != nil {
		return err
	}
	script.Name = scriptName

	scriptHash := utils.HashToString([]byte(xmlDefinition))

	err = PrepareScript(script, scriptHash)
	if err != nil {
		return err
	}

	return configureF9(auth, campaignName, script)
}

func PrepareScript(script *ivr.IVRScript, scriptHash string) error {
	err := cloud.UploadScript(ivr.MakeStorageScript(script), scriptHash)
	if err != nil {
		return err
	}

	err = cloud.UpdateConfig(map[string]string{cloud.GcpConfigDomainNameKeyString: "Product Management DW", cloud.GcpConfigCampaignNameKeyString: "sergei_inbound"})
	if err != nil {
		return err
	}

	return prepareIntents(script, scriptHash)
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
