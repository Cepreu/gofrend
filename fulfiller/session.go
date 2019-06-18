package fulfiller

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Cepreu/gofrend/ivr/vars"
)

// Session contains data stored by webhook between sessions.
type Session struct {
	// Script variables.
	Vars []*vars.Variable
}

func sessionExists(ID string) bool {
	filename := fmt.Sprintf("sessions/%s.json", ID)
	f, err := os.Open(filename)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

func (session *Session) store(ID string) {
	filename := fmt.Sprintf("sessions/%s.json", ID)
	f, err := os.Open(filename)
	if err != nil {
		f, _ = os.Create(filename)
	}
	bytes, _ := json.Marshal(session)
	f.Write(bytes)
	f.Close()
}

// Not safe to call if session does not exist.
func loadSession(ID string) (session *Session) {
	filename := fmt.Sprintf("sessions/%s.json", ID)
	f, _ := os.Open(filename)
	finfo, _ := f.Stat()
	data := make([]byte, finfo.Size())
	f.Read(data)
	f.Close()
	json.Unmarshal(data, session)
	return
}

func deleteSession(ID string) {
	filename := fmt.Sprintf("sessions/%s.json", ID)
	os.Remove(filename)
}
