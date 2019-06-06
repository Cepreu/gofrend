package dialogflow

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// DialogflowProcessor has all the information for connecting with Dialogflow
type DialogflowProcessor struct {
	projectID        string
	authJSONFilePath string
	lang             string
	timeZone         string
	sessionClient    *dialogflow.SessionsClient
	intentClient     *dialogflow.IntentsClient
	agentClient      *dialogflow.AgentsClient
	ctx              context.Context
}

// NLPResponse is the struct for the response
type NLPResponse struct {
	Intent     string            `json:"intent"`
	Confidence float32           `json:"confidence"`
	Entities   map[string]string `json:"entities"`
}

var dp DialogflowProcessor

func main() {
	dp.init("f9-test-agent", "F9-Test-Agent-0925974a682a", "en", "America/Los_Angeles")
	http.HandleFunc("/", requestHandler)
	fmt.Println("Started listening...")
	http.ListenAndServe(":5000", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//POST method, receives a json to parse
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		type inboundMessage struct {
			Message string
		}
		var m inboundMessage
		err = json.Unmarshal(body, &m)
		if err != nil {
			panic(err)
		}

		// Use NLP
		//	response := dp.processNLP(m.Message, "testUser")
		//	response := dp.GetIntent("6ad8beb8-7382-4267-8070-cbb3b506b45e")
		//response := dp.CreateIntent()
		response := dp.ExportAgent()

		fmt.Printf("%#v", response)
		w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(response)
		w.Write([]byte(response))
	}
}

func (dp *DialogflowProcessor) init(a ...string) (err error) {
	dp.projectID = a[0]
	dp.authJSONFilePath = a[1]
	dp.lang = a[2]
	dp.timeZone = a[3]

	// Auth process: https://dialogflow.com/docs/reference/v2-auth-setup

	dp.ctx = context.Background()
	sessionClient, err := dialogflow.NewSessionsClient(dp.ctx, option.WithCredentialsFile(dp.authJSONFilePath))
	if err != nil {
		log.Fatal("Error in auth with Dialogflow")
	}
	dp.sessionClient = sessionClient

	intentClient, err := dialogflow.NewIntentsClient(dp.ctx, option.WithCredentialsFile(dp.authJSONFilePath))
	if err != nil {
		log.Fatal("Error in auth with Dialogflow")
	}
	dp.intentClient = intentClient

	agentClient, err := dialogflow.NewAgentsClient(dp.ctx, option.WithCredentialsFile(dp.authJSONFilePath))
	if err != nil {
		log.Fatal("Error in auth with Dialogflow")
	}
	dp.agentClient = agentClient

	return
}

func (dp *DialogflowProcessor) processNLP(rawMessage string, username string) (r NLPResponse) {
	sessionID := username
	request := dialogflowpb.DetectIntentRequest{
		Session: fmt.Sprintf("projects/%s/agent/sessions/%s", dp.projectID, sessionID),
		QueryInput: &dialogflowpb.QueryInput{
			Input: &dialogflowpb.QueryInput_Text{
				Text: &dialogflowpb.TextInput{
					Text:         rawMessage,
					LanguageCode: dp.lang,
				},
			},
		},
		QueryParams: &dialogflowpb.QueryParameters{
			TimeZone: dp.timeZone,
		},
	}
	response, err := dp.sessionClient.DetectIntent(dp.ctx, &request)
	if err != nil {
		log.Fatalf("Error in communication with Dialogflow %s", err.Error())
		return
	}
	queryResult := response.GetQueryResult()
	if queryResult.Intent != nil {
		r.Intent = queryResult.Intent.DisplayName
		r.Confidence = float32(queryResult.IntentDetectionConfidence)
	}
	r.Entities = make(map[string]string)
	params := queryResult.Parameters.GetFields()
	if len(params) > 0 {
		for paramName, p := range params {
			fmt.Printf("Param %s: %s (%s)", paramName, p.GetStringValue(), p.String())
			extractedValue := extractDialogflowEntities(p)
			r.Entities[paramName] = extractedValue
		}
	}
	return
}

func extractDialogflowEntities(p *structpb.Value) (extractedEntity string) {
	kind := p.GetKind()
	switch kind.(type) {
	case *structpb.Value_StringValue:
		return p.GetStringValue()
	case *structpb.Value_NumberValue:
		return strconv.FormatFloat(p.GetNumberValue(), 'f', 6, 64)
	case *structpb.Value_BoolValue:
		return strconv.FormatBool(p.GetBoolValue())
	case *structpb.Value_StructValue:
		s := p.GetStructValue()
		fields := s.GetFields()
		extractedEntity = ""
		for key, value := range fields {
			if key == "amount" {
				extractedEntity = fmt.Sprintf("%s%s", extractedEntity, strconv.FormatFloat(value.GetNumberValue(), 'f', 6, 64))
			}
			if key == "unit" {
				extractedEntity = fmt.Sprintf("%s%s", extractedEntity, value.GetStringValue())
			}
			if key == "date_time" {
				extractedEntity = fmt.Sprintf("%s%s", extractedEntity, value.GetStringValue())
			}
			// @TODO: Other entity types can be added here
		}
		return extractedEntity
	case *structpb.Value_ListValue:
		list := p.GetListValue()
		if len(list.GetValues()) > 1 {
			// @TODO: Extract more values
		}
		extractedEntity = extractDialogflowEntities(list.GetValues()[0])
		return extractedEntity
	default:
		return ""
	}
}
