package fulfiller

import (
	"net/http"
	"net/url"
	"testing"
)

func TestDomainAndCampaignIDs(t *testing.T) {
	domainName := "Product Management DW"
	campaignName := "sergei_inbound"
	baseURL := ""

	f9client := Client{
		httpClient: &http.Client{},
		BaseURL:    &url.URL{Path: baseURL},
	}
	domainID, campaignID, err := f9client.getDomainCampaignIDs(domainName, campaignName)

	if err != nil {
		t.Errorf("\ngetDomainCampaignIDs: Error: %v, DomainID: %s, CampaignID: %s\n", err, domainID, campaignID)
		t.Fatal(err)
	}
	sessionID, err := f9client.newIVRSession(domainID, campaignID)
	if err != nil {
		t.Errorf("\newIVRSession: Error: %v, SessionID: %s, CampaignID: %s\n", err, sessionID, campaignID)
		t.Fatal(err)
	}

	ss, err := f9client.getSessionState(domainID, sessionID, -1)
	if err != nil {
		t.Errorf("\ngetSessionState: Error: %v, SessionID: %s, \nsession: %v\n", err, sessionID, ss)
		t.Fatal(err)
	}

}
