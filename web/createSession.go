package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type IVRSessionResp struct {
	Count int          `json:"count"`
	Items []IVRSession `json:"items"`
	Error *ApiError    `json:"error"`
}
type IVRSession struct {
	ID         int      `json:"id"`
	CampaignID string   `json:"campaignId"`
	DomainID   string   `json:"domainId"`
	Languages  []string `json:"languages"`
	Theme      string   `json:"theme"`
}

//HTTP GET - /{domain_id}/campaigns/{campaign_id}/create_session?{args}
func CreateSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domain_id"]
	campaignID := vars["campaign_id"]
	params := r.URL.Query()

	for k, v := range params {
		for i, vi := range v {
			fmt.Printf("k=%s, v[%d]=%s\n", k, i, vi)
		}
	}

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
