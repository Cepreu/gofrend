package fulfiller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	ivr "github.com/Cepreu/gofrend/ivrparser"
	"github.com/Cepreu/gofrend/vars"
	structpb "github.com/golang/protobuf/ptypes/struct"
)

// Session contains data stored by webhook between sessions.
type Session struct {
	// Script variables.
	Variables map[string]*vars.Variable
}

func initSession(script *ivr.IVRScript) *Session {
	return &Session{
		Variables: script.Variables,
	}
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

func (session *Session) setParameter(name string, value *structpb.Value) {
	var variable *vars.Variable
	for _, v := range session.Variables {
		if v.Name() == name {
			variable = v
		}
	}
	if variable == nil {
		log.Fatalf("Could not find session variable with name: %s", name)
	}
	switch v := variable.Value.(type) {
	case *vars.Integer:
		v.Value = int(value.GetNumberValue())
	}
}
