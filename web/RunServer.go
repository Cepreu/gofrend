package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CampaignStateResp struct {
	Count int             `json:"count"`
	Items []CampaignState `json:"items"`
	Error *ApiError       `json:"error"`
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
	IsVisualIVREnabled bool   `json:"is_visual_ivr_enabled"`
}

type IVRSessionResp struct {
	Count int             `json:"count"`
	Items []IVRSession `json:"items"`
	Error *ApiError       `json:"error"`
}
type IVRSession struct {
	ID	int `json:"id"`
	CampaignID string `json:"campaignId"`
	DomainID string `json:"domainId"`
	Languages []string `json:"languages"`
	Theme string `json:"theme"`
}

type UserAction struct {
	Name string `json:"name"`
	ScriptID string `json:"scriptId"`
	ModuleID string `json:"moduleId"`
	BranchID string `json:"branchId"`
	Args	KVList	`json:"args"`
}

type KVList map[string]interface{}

//HTTP Get - /domains/{domain_name}/campaigns?name={campaign_name}
func GetDomainAndCampaignIDsHandler(w http.ResponseWriter, r *http.Request) {
	campName := r.URL.Query().Get("name")
	rs := CampaignStateResp{}
	if campName == "" {
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

func RunServer() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/domains/{domain_name}/campaigns", GetDomainAndCampaignIDsHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
