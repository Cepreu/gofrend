package fulfiller

import (
	"net/http"
	"net/url"
	"testing"
)

func TestDomainAndCampaignIDs(t *testing.T) {
	domainName := "Product Management DW"
	campaignName := "sergei_inbound"
	baseURL := "https://api.five9.com/ivr/1"

	f9client := Client{
		httpClient: &http.Client{},
		BaseURL:    &url.URL{Path: baseURL},
	}
	domainID, campaignID, err := f9client.getDomainCampaignIDs(domainName, campaignName)

	if err != nil {
		t.Errorf("getDomainCampaignIDs: \nDomainID: %s, CampaignID: %s\n", domainID, campaignID)
	}
	//	check(err)
}
