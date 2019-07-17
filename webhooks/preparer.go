package webhook

import (
	"net/http"

	"github.com/Cepreu/gofrend/preparer"
)

// HandlePreparerWebhook handles the fulfiller webhook
func HandlePreparerWebhook(w http.ResponseWriter, r *http.Request) {
	preparer.HandleWebhook(w, r)
}
