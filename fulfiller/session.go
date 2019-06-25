package fulfiller

import (
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/vars"
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

func sessionExists(ID string) (bool, error) {
	return false, nil
}

func (session *Session) store(ID string) error {
	return nil
}

func loadSession(ID string) (*Session, error) {
	return nil, nil
}

func deleteSession(ID string) error {
	return nil
}

func (session *Session) setParameter(name string, value *structpb.Value) error {
	var variable *vars.Variable
	for _, v := range session.Variables {
		if v.Name == name {
			variable = v
		}
	}
	if variable == nil {
		return fmt.Errorf("Could not find session variable with name: %s", name)
	}
	switch v := variable.Value.(type) {
	case *vars.Integer:
		v.Value = int(value.GetNumberValue())
	}
	return nil
}
