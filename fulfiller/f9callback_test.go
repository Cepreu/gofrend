package fulfiller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestSessionJSON(t *testing.T) {
	str := `
{
	"count":1,
	"items":[
	   {
		  "moduleDescription":{
			 "moduleId":"73594C84757A47B9AD1ED2CD1FCFD5A8",
			 "moduleType":"skillTransfer",
			 "moduleName":"SkillTransfer2",
			 "prompts":[
 
			 ],
			 "caption":[
 
			 ],
			 "languages":[
 
			 ],
			 "restriction":{
				"isE164":"false",
				"type":"phone"
			 },
			 "priority":10,
			 "skillTransferVars":{
				"skills":[
				   "sk_skill1"
				],
				"callVars":{
 
				}
			 },
			 "callbackAllowInternational":false,
			 "isTcpaConsentEnabled":false,
			 "tcpaConsentText":"I consent to receive calls from [Business Name] at the telephone number provided above, including promotional calls.I understand that consent is not a condition of purchase.",
			 "isCallbackEnabled":true,
			 "isChatEnabled":false,
			 "isEmailEnabled":false,
			 "isVideoEnabled":false
		  },
		  "stage":2,
		  "scriptId":"44332",
		  "moduleId":"73594C84757A47B9AD1ED2CD1FCFD5A8",
		  "sessionURL":"https://api.five9.com:443/ivr/1/115675/sessions/40FB30F1B2CD4E9BA8F5C24AB5CC9E21_6",
		  "isFinal":false,
		  "variables":{
			 "$EWT":"600"
		  },
		  "isBackAvailable":true,
		  "isCallbackRequested":false,
		  "isFeedbackRequested":false
	   }
	],
	"error":null
 }
`

	res := sessionStateResp{}
	err := json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainAndCampaignIDs(t *testing.T) {
	domainName := "Product Management DW"
	campaignName := "sergei_inbound"
	baseURL := ""

	f9client := client{
		httpClient: &http.Client{},
		BaseURL:    &url.URL{Path: baseURL},
	}
	domainID, campaignID, err := f9client.getDomainCampaignIDs(domainName, campaignName)

	if err != nil {
		t.Errorf("\ngetDomainCampaignIDs: Error: %v, DomainID: %s, CampaignID: %s\n", err, domainID, campaignID)
		t.Fatal(err)
	}
	sessionID, err := f9client.newIVRSession(domainID, campaignID, map[string]string{})
	if err != nil {
		t.Errorf("\newIVRSession: Error: %v, SessionID: %s, CampaignID: %s\n", err, sessionID, campaignID)
		t.Fatal(err)
	}

	sss, err := f9client.getSessionState(domainID, sessionID, -1)
	if err != nil {
		t.Errorf("\ngetSessionState: Error: %v, SessionID: %s, \nsession: %v\n", err, sessionID, sss)
		t.Fatal(err)
	}

	body := &userAction{
		Name:     "userAnswer",
		ScriptID: sss.ScriptID,
		ModuleID: sss.ModuleID,
	}
	sss, err = f9client.postAction(domainID, sessionID, body, sss.Stage)
	if err != nil {
		t.Errorf("\npostAction: Error: %v, SessionID: %s, \nsession: %v\n", err, sessionID, sss)
		t.Fatal(err)
	}

	body = &userAction{
		Name:     "requestCallback",
		ScriptID: sss.ScriptID,
		ModuleID: sss.ModuleID,
		Args:     map[string]interface{}{"callbackNumber": "6502437004"},
	}
	sss, err = f9client.postAction(domainID, sessionID, body, sss.Stage)
	if err != nil {
		t.Errorf("\npostActionCB: Error: %v, SessionID: %s, \nsession: %v\n", err, sessionID, sss)
		t.Fatal(err)
	}
	_, err = f9client.finalize(domainID, sessionID)
	if err != nil {
		t.Errorf("\nfinalize: Error: %v, SessionID: %s\n", err, sessionID)
		t.Fatal(err)
	}
}

func TestCallback(t *testing.T) {
	domainName := "Product Management DW"
	campaignName := "sergei_inbound"
	cbPhoneNumber := "7818641521"
	params := map[string]string{"MODULE_ID": "73594C84757A47B9AD1ED2CD1FCFD5A8"}
	err := createCallback(domainName, campaignName, cbPhoneNumber, params)
	if err != nil {
		t.Fatal(err)
	}
}
