package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type KVList map[string]interface{}

type ScriptArgs struct {
	ContactFields KVList `json:"contactFields"`
	ClientRecord  KVList `json:"clientRecord"`
	ExternalVars  KVList `json:"externalVars"`
}

////////////////////////////////////////////////////////////////////////////////////

func RunServer() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/domains/{domain_name}/campaigns", GetDomainAndCampaignIDsHandler).Methods("GET")
	r.HandleFunc("/api/{domain_id}/campaigns/{campaign_id}/create_session", CreateSession).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
