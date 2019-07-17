package webhook

import (
	"net/http"

	"github.com/Cepreu/gofrend/fulfiller"
)

// HandleFulfillerWebhook handles the fulfiller webhook
func HandleFulfillerWebhook(w http.ResponseWriter, r *http.Request) {
	fulfiller.HandleWebhook(w, r)
}
