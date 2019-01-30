package main

import (
	"fmt"
	"log"

	"github.com/Cepreu/gofrend/utils"
	"github.com/davecgh/go-spew/spew"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func (dp *DialogflowProcessor) ExportAgent() string {
	req := dialogflowpb.ExportAgentRequest{
		Parent: fmt.Sprintf(`projects/%s`, dp.projectID),
		// // Optional. The Google Cloud Storage URI to export the agent to.
		// // Note: The URI must start with
		// // "gs://". If left unspecified, the serialized agent is returned inline.
		// AgentUri             string
	}
	op, err := dp.agentClient.ExportAgent(dp.ctx, &req)
	if err != nil {
		log.Fatalf("Error in communication with Dialogflow %s", err.Error())
		return ""
	}
	response, err := op.Wait(dp.ctx)
	if err != nil {
		log.Fatalf("Error in communication with Dialogflow %s", err.Error())
		return ""
	}

	utils.PrettyPrint(response)
	return spew.Sdump(response)
}
