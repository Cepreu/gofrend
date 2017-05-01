package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserAction struct {
	Name string `json:"name"`
	ScriptID string `json:"scriptId"`
	ModuleID string `json:"moduleId"`
	BranchID string `json:"branchId"`
	Args	map[string] interface{} `json:"args"`
}

type SessionStateResp struct {
	Count int          `json:"count"`
	Items []SessionState `json:"items"`
	Error *ApiError    `json:"error"`
}
type SessionState struct {
	SessionURL string `json:"sessionURL"`
	ModuleID string   `json:"moduleId"`
	ScriptID int   `json:scriptId"`
	Stage int `json:"stage"`
	IsFinal bool `json:"isFinal"`
	IsFeedbackRequested bool `json:"isFeedbackRequested"`
	IsBackAvailable bool `json:"isBackAvailable"`
	Variables map[string]interface{} `json:"variables"`
}

//HTTP POST - /{domain_id}/sessions/{session_id}/action
func Action(w http.ResponseWriter, r *http.Request) {
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
