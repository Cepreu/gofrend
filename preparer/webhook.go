package preparer

import (
	"log"
	"net/http"

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
	r.ParseForm()
	f := r.Form
	utils.PrettyLog(f)
	err := Prepare(f[cScriptName][0], f[cCampaignName][0], f[cUsername][0], f[cTemporaryPassword][0])

	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte("SUCCESS\n"))

	if err != nil {
		log.Panicf("Error writing response: %v", err)
	}
}
