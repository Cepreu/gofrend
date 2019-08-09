package preparer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

var filename1 = "test_files/comparison_test.five9ivr"

func TestPrepare(t *testing.T) {
	err := Prepare("LoanScoreRouting", "sergei_inbound", "sam2@007.com", "pwd1234567")
	checkNil(err, t)
}

func TestPrepareCloud(t *testing.T) {
	URL := "https://us-central1-f9-dialogflow-converter.cloudfunctions.net/handle-preparer-webhook"
	form := url.Values{
		cScriptName:        {"sk_text"},
		cCampaignName:      {"sergei_inbound"},
		cUsername:          {"sam2@007.com"},
		cTemporaryPassword: {"pwd1234567"},
	}
	resp, err := http.PostForm(URL, form)
	checkNil(err, t)
	data, err := ioutil.ReadAll(resp.Body)
	checkNil(err, t)
	fmt.Print(string(data))
}
