package preparer

import (
	"testing"
)

var filename1 = "test_files/comparison_test.five9ivr"

func TestPrepareFile(t *testing.T) {
	err := Prepare("sk_text", "sergei_inbound", "sam2@007.com", "pwd1234567")
	checkNil(err, t)
}

// func TestPrepareCloud(t *testing.T) {
// 	url := "https://us-central1-f9-dialogflow-converter.cloudfunctions.net/handle-preparer-webhook"
// 	form := url.Values{
// 		cScriptName: {"sk_text"},
// 		cCampaignName: {"sergei_inbound"},
// 		cUsername: {"sam2@007.com"}
// 	}
// 	resp, err := http.PostForm()
// }
