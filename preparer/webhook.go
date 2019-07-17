package preparer

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HandleWebhook prepares
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panicf("Error reading request body: %v", err)
	}

	err = Prepare(data)
	if err != nil {
		log.Panicf(err)
	}

	_, err = w.Write([]byte("SUCCESS"))

	if err != nil {
		log.Panicf("Error writing response: %v", err)
	}
}
