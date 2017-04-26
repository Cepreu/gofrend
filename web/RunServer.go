package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CampaignStateResp struct {
	Count int             `json:"count"`
	Items []CampaignState `json:"items"`
	Error ApiError        `json:"error"`
}
type ApiError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type CampaignState struct {
	SelfURL            string `json:"selfURL"` // format: url
	Name               string `json:"name"`
	ID                 int32  `json:"id"`
	DomainID           int32  `json:"domainId"`
	IsVisualIVREnabled bool   `json:"isVisualIVREnabled"`
}

//HTTP Get - /domains/{domain_name}/campaigns?name={campaign_name}
func GetDomainAndCampaignIDs(w http.ResponseWriter, r *http.Request) {
	campName := r.URL.Query().Get("name")
	rs := CampaignStateResp{}

	if campName != "" {
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(rs)
	if err == nil {
		panic(err)
	}
	fmt.Println(j)
}
