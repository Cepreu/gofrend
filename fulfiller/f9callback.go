package fulfiller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	cModuleIDParamKey      string = "MODULE_ID"
	cDomainNameConfigKey   string = "DOMAIN_NAME"
	cCampaignNameConfigKey string = "CAMPAIGN_NAME"
)

type campaignStateResp struct {
	Count int             `json:"count"`
	Items []campaignState `json:"items"`
	Error *apiError       `json:"error"`
}
type campaignState struct {
	SelfURL            string `json:"selfURL"` // format: url
	Name               string `json:"name"`
	ID                 int    `json:"id"`
	DomainID           int    `json:"domainId"`
	IsVisualIVREnabled bool   `json:"is_visual_ivr_enabled"`
}
type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ivrSessionResp struct {
	Count int          `json:"count"`
	Items []ivrSession `json:"items"`
	Error *apiError    `json:"error"`
}
type ivrSession struct {
	ID         int      `json:"id"`
	CampaignID string   `json:"campaignId"`
	DomainID   string   `json:"domainId"`
	Languages  []string `json:"languages"`
	Theme      string   `json:"theme"`
}

type userAction struct {
	Name     string                 `json:"name"`
	ScriptID string                 `json:"scriptId"`
	ModuleID string                 `json:"moduleId"`
	BranchID *string                `json:"branchId"`
	Args     map[string]interface{} `json:"args"`
}

type sessionStateResp struct {
	Count int            `json:"count"`
	Items []sessionState `json:"items"`
	Error *apiError      `json:"error"`
}
type finalizeResp struct {
	RedirectURL string
}

type sessionState struct {
	SessionURL          string                 `json:"sessionURL"`
	ModuleID            string                 `json:"moduleId"`
	ScriptID            string                 `json:"scriptId"`
	Stage               int                    `json:"stage"`
	IsFinal             bool                   `json:"isFinal"`
	IsFeedbackRequested bool                   `json:"isFeedbackRequested"`
	IsBackAvailable     bool                   `json:"isBackAvailable"`
	Variables           map[string]interface{} `json:"variables"`
	Module              moduleDescription      `json:"moduleDescription"`
}
type moduleDescription struct {
	ModuleID                   string                 `json:"moduleId"`
	ModuleType                 string                 `json:"moduleType"`
	ModuleName                 string                 `json:"moduleName"`
	Prompts                    []prompt               `json:"prompts"`
	Captions                   []caption              `json:"caption"`
	Languages                  []string               `json:"languages"`
	Restrictions               map[string]interface{} `json:"restriction"`
	Priority                   int                    `json:"priority"`
	SkillTransferVars          map[string]interface{} `json:"skillTransferVars"`
	CallbackAllowInternational bool                   `json:"callbackAllowInternational"`
	IsTcpaConsentEnabled       bool                   `json:"isTcpaConsentEnabled"`
	TcpaConsentText            string                 `json:"tcpaConsentText"`
	IsCallbackEnabled          bool                   `json:"isCallbackEnabled"`
	IsChatEnabled              bool                   `json:"isChatEnabled"`
	IsEmailEnabled             bool                   `json:"isEmailEnabled"`
	IsVideoEnabled             bool                   `json:"isVideoEnabled"`
}

type prompt struct {
	PromptType string        `json:"promptType"`
	Lang       string        `json:"lang"`
	Text       decoratedText `json:"text"`
	Image      string        `json:"image"`
}
type decoratedText struct {
	Element string `json:"element"`
}
type text struct {
	InnerString string `json:"innerString"`
}

type kvList map[string]interface{}

type scriptArgs struct {
	ContactFields kvList `json:"contactFields"`
	ClientRecord  kvList `json:"clientRecord"`
	ExternalVars  kvList `json:"externalVars"`
}
type markup struct {
	Tag        string          `json:"tag"`
	Attributes kvList          `json:"attributes"`
	Children   []decoratedText `json:"children"`
}
type caption struct {
	PromptType string `json:"promptType"`
	Text       string `json:"text"`
	E164       bool   `json:"isE164"`
	Language   string `json:"language"`
}

type client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func (c *client) newRequest(method, path string, queryParams map[string]string, body interface{}) (*http.Request, error) {
	var rel = &url.URL{}
	rel.Scheme = "https"
	rel.Host = "api.five9.com"
	rel.Path = path
	q := rel.Query()
	for p, v := range queryParams {
		q.Set(p, v)
	}
	rel.RawQuery = q.Encode()
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}
func (c *client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func timestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}

func (c *client) getDomainCampaignIDs(domainName, campaignName string) (domainID, campaignID string, err error) {
	req, err := c.newRequest(
		"GET",
		fmt.Sprintf("ivr/1/domains/%s/campaigns", domainName),
		map[string]string{"name": campaignName},
		nil)
	if err == nil {
		var cs = campaignStateResp{}
		_, err = c.do(req, &cs)
		if err == nil {
			domainID = strconv.Itoa(cs.Items[0].DomainID)
			campaignID = strconv.Itoa(cs.Items[0].ID)
			return
		}
	}
	return "", "", err
}

func (c *client) newIVRSession(domainID, campaignID string, params map[string]string) (sessionID string, err error) {
	req, err := c.newRequest("GET",
		fmt.Sprintf("ivr/1/%s/campaigns/%s/new_ivr_session", domainID, campaignID),
		params,
		nil)
	if err != nil {
		return "", err
	}
	c.httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	location, err := resp.Location()
	if err != nil {
		log.Panicf("Error reading response location: %v", err)
	}
	parts := strings.Split(location.Path, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("Incorrect location: %s", location.Path)
	}
	sessionID = parts[len(parts)-2]
	return
}

func (c *client) getSessionState(domainID, sessionID string, stage int64) (*sessionState, error) {
	req, err := c.newRequest(
		"GET",
		fmt.Sprintf("ivr/1/%s/sessions/%s/", domainID, sessionID),
		map[string]string{
			"stage": strconv.FormatInt(stage, 10),
			"_":     timestamp(),
		},
		nil)
	if err == nil {
		var ss = sessionStateResp{}
		_, err = c.do(req, &ss)
		if err == nil {
			return &ss.Items[0], nil
		}
	}
	return nil, err
}
func (c *client) postAction(domainID, sessionID string, ua *userAction, stage int) (*sessionState, error) {
	req, err := c.newRequest("POST",
		fmt.Sprintf("ivr/1/%s/sessions/%s/stages/%d/action", domainID, sessionID, stage),
		map[string]string{},
		ua)
	if err == nil {
		var ss = sessionStateResp{}
		_, err = c.do(req, &ss)
		if err == nil {
			return &ss.Items[0], nil
		}
	}
	return nil, err
}

func (c *client) finalize(domainID, sessionID string) (redirectURL string, err error) {
	req, err := c.newRequest(
		"GET",
		fmt.Sprintf("ivr/1/%s/sessions/%s/finalize", domainID, sessionID),
		map[string]string{"_": timestamp()},
		nil)
	if err == nil {
		var fr = finalizeResp{}
		_, err = c.do(req, &fr)
		if err == nil {
			return fr.RedirectURL, nil
		}
	}
	return "", err
}

func createTermination(domainName, campaignName string, params map[string]string) error {
	return createCallback(domainName, campaignName, "", params)
}

func createCallback(domainName, campaignName, phone string, params map[string]string) (err error) {
	var baseURL = ""
	var (
		f9client = client{
			httpClient: &http.Client{},
			BaseURL:    &url.URL{Path: baseURL},
		}
		domainID, campaignID, sessionID string
		body                            *userAction
	)
	domainID, campaignID, err = f9client.getDomainCampaignIDs(domainName, campaignName)
	if err != nil {
		return err
	}
	if sessionID, err = f9client.newIVRSession(domainID, campaignID, params); err == nil {
		if sss, err := f9client.getSessionState(domainID, sessionID, -1); err == nil {
			body = &userAction{
				Name:     "userAnswer",
				ScriptID: sss.ScriptID,
				ModuleID: sss.ModuleID,
			}
			if sss, err = f9client.postAction(domainID, sessionID, body, sss.Stage); err == nil {
				err = fmt.Errorf("Improper IVR script: expected skillTransfer, received %s", sss.Module.ModuleType)
				if sss.Module.ModuleType == "skillTransfer" {
					body = &userAction{
						Name:     "requestCallback",
						ScriptID: sss.ScriptID,
						ModuleID: sss.ModuleID,
						Args:     map[string]interface{}{"callbackNumber": phone},
					}
					sss, err = f9client.postAction(domainID, sessionID, body, sss.Stage)
				}
			}
		}
	}
	f9client.finalize(domainID, sessionID) //try to finalize anyway
	return err
}
