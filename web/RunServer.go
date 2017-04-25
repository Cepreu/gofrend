package web

import (
	"encoding/json"
	"net/http"
)

type campaignStateResp struct {
	Count int             `json:"count"`
	Items []CampaignState `json:"items"`
	Error APIError        `json:"error"`
}

//HTTP Get - /domains/{domain_name}/campaigns?name={campaign_name}
func GetDomainAndCampaignIDs(w http.ResponseWriter, r *http.Request) {
	campName := r.URL.Query().Get("name")
	rs := campaignStateResp{}

	if campName != "" {
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(rs)
	if err == nil {
		panic(err)
	}
}
