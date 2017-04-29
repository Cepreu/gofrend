package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type UserAction struct {
	Name     string `json:"name"`
	ScriptID string `json:"scriptId"`
	ModuleID string `json:"moduleId"`
	BranchID string `json:"branchId"`
	Args     KVList `json:"args"`
}

type KVList map[string]interface{}

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
