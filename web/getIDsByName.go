package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CampaignStateResp struct {
	Count int             `json:"count"`
	Items []CampaignState `json:"items"`
	Error *APIError       `json:"error"`
}
type CampaignState struct {
	SelfURL            string `json:"selfURL"` // format: url
	Name               string `json:"name"`
	ID                 int32  `json:"id"`
	DomainID           int32  `json:"domainId"`
	IsVisualIVREnabled bool   `json:"is_visual_ivr_enabled"`
}

//HTTP Get - /domains/{domain_name}/campaigns?name={campaign_name}
func GetDomainAndCampaignIDsHandler(w http.ResponseWriter, r *http.Request) {
	campName := r.URL.Query().Get("name")
	domainName := mux.Vars(r)["domain_name"]

	rs := CampaignStateResp{}
	if campName == "" || domainName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("query parameter 'name' is missing"))
	} else {
		rs.Count = 1
		rs.Items = append(rs.Items, CampaignState{SelfURL: "www.google.com", ID: 333333, Name: campName, DomainID: 12345, IsVisualIVREnabled: true})

		j, err := json.Marshal(rs)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
