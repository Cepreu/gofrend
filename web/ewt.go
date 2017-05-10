package web

//HTTP get - /{domain_id}/sessions/{session_id}/ewt
func ewtHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domain_id"]
	campaignID := vars["campaign_id"]
	lastExpiresAt := r.URL.Query().Get("last_expires_at")

	rs := CampaignStateResp{}
	rs.Count = 1
	ewt := 10
	rs.Items = append(rs.Items, ewt)

	j, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
