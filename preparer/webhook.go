package preparer

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

// HandleWebhook ---
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panicf("Error reading request body: %v", err)
	}
	err = r.Body.Close()
	if err != nil {
		log.Panicf("Error reading request body: %v", err)
	}
	unescapedData := []byte(html.UnescapeString(string(data)))
	log.Print(unescapedData)

	err = Prepare(unescapedData)
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte("SUCCESS\n"))

	if err != nil {
		log.Panicf("Error writing response: %v", err)
	}
}
