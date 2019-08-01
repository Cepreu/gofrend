package preparer

import (
	"log"
	"net/http"

	"github.com/Cepreu/gofrend/cloud"
	"github.com/Cepreu/gofrend/utils"
)

const (
	cScriptName        string = "SCRIPT_NAME"
	cCampaignName      string = "CAMPAIGN_NAME"
	cUsername          string = "USERNAME"
	cTemporaryPassword string = "TEMPORARY_PASSWORD"
	cPassword          string = "PASSWORD"
)

// HandleWebhook ---
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(cloud.GcpConfigAccessTokenKeyString)
	config, err := cloud.GetConfig()
	if err != nil {
		log.Panicf("Error getting config from cloud: %v", err)
	}
	if config[cloud.GcpConfigAccessTokenKeyString] != token {
		log.Printf("Permission denied: incorrect access token")
		w.Write([]byte("Permission denied: incorrect access token"))
		return
	}
	err = r.ParseForm()
	if err != nil {
		log.Panicf("Error parsing form: %v", err)
	}
	f := r.Form
	utils.PrettyLog(f)
	err = Prepare(f[cScriptName][0], f[cCampaignName][0], f[cUsername][0], f[cTemporaryPassword][0])

	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte("SUCCESS\n"))

	if err != nil {
		log.Panicf("Error writing response: %v", err)
	}
}
