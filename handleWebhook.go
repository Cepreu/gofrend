package webhook

import (
	"net/http"

	"github.com/Cepreu/gofrend/fulfiller"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fulfiller.HandleWebhook(w, r)
}
