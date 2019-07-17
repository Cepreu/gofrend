package webhooks

import (
	"net/http"

	"github.com/Cepreu/gofrend/preparer"
)

// HandleFulfillerWebhook handles the fulfiller webhook
func HandleFulfillerWebhook(w http.ResponseWriter, r *http.Request) {
	preparer.HandleWebhook(w, r)
}
