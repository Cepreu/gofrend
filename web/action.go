package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserAction struct {
	Name     string                 `json:"name"`
	ScriptID string                 `json:"scriptId"`
	ModuleID string                 `json:"moduleId"`
	BranchID string                 `json:"branchId"`
	Args     map[string]interface{} `json:"args"`
}

type SessionStateResp struct {
	Count int            `json:"count"`
	Items []SessionState `json:"items"`
	Error *APIError      `json:"error"`
}
type SessionState struct {
	SessionURL          string                 `json:"sessionURL"`
	ModuleID            string                 `json:"moduleId"`
	ScriptID            int                    `json:"scriptId"`
	Stage               int                    `json:"stage"`
	IsFinal             bool                   `json:"isFinal"`
	IsFeedbackRequested bool                   `json:"isFeedbackRequested"`
	IsBackAvailable     bool                   `json:"isBackAvailable"`
	Variables           map[string]interface{} `json:"variables"`
	Module              ModuleDescription      `json:"module"`
}
type ModuleDescription struct {
	ModuleID   string    `json:"moduleId"`
	ModuleType string    `json:"moduleType"`
	ModuleName string    `json:"moduleName"`
	Prompts    []Prompt  `json:"prompts"`
	Captions   []Caption `json:"caption"`
	Languages  []string  `json:"languages"`
}
type Prompt struct {
	PromptType string        `json:"promptType"`
	Lang       string        `json:"lang"`
	Text       DecoratedText `json:"text"`
	Image      string        `json:"image"`
}
type DecoratedText struct {
	Element string `json:"element"`
}
type Text struct {
	InnerString string `json:"innerString"`
}
type Markup struct {
	Tag        string          `json:"tag"`
	Attributes KVList          `json:"attributes"`
	Children   []DecoratedText `json:"children"`
}
type Caption struct {
	PromptType string `json:"promptType"`
	Text       string `json:"text"`
	E164       bool   `json:"isE164"`
	Language   string `json:"language"`
}

//HTTP POST - /{domain_id}/sessions/{session_id}/action
func ActionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domain_id"]
	campaignID := vars["campaign_id"]

	rs := IVRSessionResp{}
	rs.Count = 1
	ivrSess := IVRSession{
		ID:         12345,
		CampaignID: campaignID,
		DomainID:   domainID,
		Languages:  []string{"ru_RU", "us_US"},
		Theme:      "a",
	}
	rs.Items = append(rs.Items, ivrSess)

	j, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
