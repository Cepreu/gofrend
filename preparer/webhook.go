package preparer

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// HandleWebhook ---
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	w.Write([]byte(strconv.Itoa(len(data))))
	w.Write([]byte("\n"))
	w.Write(data)
	if err != nil {
		log.Panicf("Error reading request body: %v", err)
	}
	err = r.Body.Close()
	if err != nil {
		log.Panicf("Error reading request body: %v", err)
	}
	log.Print(string(data))

	err = Prepare(data)
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte("SUCCESS\n"))

	if err != nil {
		log.Panicf("Error writing response: %v", err)
	}
}
