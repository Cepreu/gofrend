package web

//HTTP Get - /{domain_id}/sessions/{session_id}
func getStateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domain_id"]
	campaignID := vars["campaign_id"]

	rs := CampaignStateResp{}
	rs.Count = 1
	campState := campState{
		selfURL:    "http://www.google.com",
		name: "sergei_inb",
		id:   "1234",
		domainID: 1234,
		isVisualIVREnabled: true,
	}
	rs.Items = append(rs.Items, campState)

	j, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
