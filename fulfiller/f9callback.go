package fulfiller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type CampaignStateResp struct {
	Count int             `json:"count"`
	Items []CampaignState `json:"items"`
	Error *APIError       `json:"error"`
}
type CampaignState struct {
	SelfURL            string `json:"selfURL"` // format: url
	Name               string `json:"name"`
	ID                 int    `json:"id"`
	DomainID           int    `json:"domainId"`
	IsVisualIVREnabled bool   `json:"is_visual_ivr_enabled"`
}
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type IVRSessionResp struct {
	Count int          `json:"count"`
	Items []IVRSession `json:"items"`
	Error *APIError    `json:"error"`
}
type IVRSession struct {
	ID         int      `json:"id"`
	CampaignID string   `json:"campaignId"`
	DomainID   string   `json:"domainId"`
	Languages  []string `json:"languages"`
	Theme      string   `json:"theme"`
}

type UserAction struct {
	Name     string                 `json:"name"`
	ScriptID string                 `json:"scriptId"`
	ModuleID string                 `json:"moduleId"`
	BranchID string                 `json:"branchId"`
	Args     map[string]interface{} `json:"args"`
}

type SessionStateResp struct {
	Count int            `json:"count"`
	Items []SessionState `json:"items"`
	Error *APIError      `json:"error"`
}
type SessionState struct {
	SessionURL          string                 `json:"sessionURL"`
	ModuleID            string                 `json:"moduleId"`
	ScriptID            string                 `json:"scriptId"`
	Stage               int                    `json:"stage"`
	IsFinal             bool                   `json:"isFinal"`
	IsFeedbackRequested bool                   `json:"isFeedbackRequested"`
	IsBackAvailable     bool                   `json:"isBackAvailable"`
	Variables           map[string]interface{} `json:"variables"`
	Module              ModuleDescription      `json:"module"`
}
type ModuleDescription struct {
	ModuleID   string    `json:"moduleId"`
	ModuleType string    `json:"moduleType"`
	ModuleName string    `json:"moduleName"`
	Prompts    []Prompt  `json:"prompts"`
	Captions   []Caption `json:"caption"`
	Languages  []string  `json:"languages"`
}
type Prompt struct {
	PromptType string        `json:"promptType"`
	Lang       string        `json:"lang"`
	Text       DecoratedText `json:"text"`
	Image      string        `json:"image"`
}
type DecoratedText struct {
	Element string `json:"element"`
}
type Text struct {
	InnerString string `json:"innerString"`
}

type KVList map[string]interface{}

type ScriptArgs struct {
	ContactFields KVList `json:"contactFields"`
	ClientRecord  KVList `json:"clientRecord"`
	ExternalVars  KVList `json:"externalVars"`
}
type Markup struct {
	Tag        string          `json:"tag"`
	Attributes KVList          `json:"attributes"`
	Children   []DecoratedText `json:"children"`
}
type Caption struct {
	PromptType string `json:"promptType"`
	Text       string `json:"text"`
	E164       bool   `json:"isE164"`
	Language   string `json:"language"`
}

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func (c *Client) newRequest(method, path string, queryParams map[string]string, body interface{}) (*http.Request, error) {
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
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println("+++++", resp)
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) getDomainCampaignIDs(domainName, campaignName string) (domainID, campaignID string, err error) {
	req, err := c.newRequest("GET", fmt.Sprintf("ivr/1/domains/%s/campaigns", domainName), map[string]string{"name": campaignName}, nil)
	if err == nil {
		var cs = CampaignStateResp{}
		_, err = c.do(req, &cs)
		if err == nil {
			domainID = strconv.Itoa(cs.Items[0].DomainID)
			campaignID = strconv.Itoa(cs.Items[0].ID)
			return
		}
	}
	return "", "", err
}

func (c *Client) newIVRSession(domainID, campaignID string) (sessionID string, err error) {
	req, err := c.newRequest("GET", fmt.Sprintf("ivr/1/%s/campaigns/%s/new_ivr_session", domainID, campaignID), map[string]string{}, nil)
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
		return "", err
	}
	parts := strings.Split(location.Path, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("Incorrect location: %s", location.Path)
	}
	sessionID = parts[len(parts)-2]
	return
}

func (c *Client) getSessionState(domainID, sessionID string, stage int64) (*SessionState, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("ivr/1/%s/sessions/%s/", domainID, sessionID),
		map[string]string{"stage": strconv.FormatInt(stage, 10), "_": strconv.FormatInt(time.Now().UnixNano()/1000000, 10)}, nil)
	if err == nil {
		var ss = SessionStateResp{}
		_, err = c.do(req, &ss)
		fmt.Println("--------", req)
		if err == nil {
			return &ss.Items[0], nil
		}
	}
	return nil, err
}