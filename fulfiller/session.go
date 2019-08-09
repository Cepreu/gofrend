package fulfiller

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/Cepreu/gofrend/cloud"
	"github.com/Cepreu/gofrend/ivr"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/api/option"
)

// Session contains data stored by webhook between sessions.
type Session struct {
	// Script variables.
	client *datastore.Client
	key    *datastore.Key
	ctx    context.Context
	Data   *SessionData
}

// SessionData contains info to be stored in cloud between sessions.
type SessionData struct {
	Variables []*ivr.Variable
}

func (session *Session) save() error {
	_, err := session.client.Put(session.ctx, session.key, session.Data)
	return err
}

func initSession(sessionID string, script *ivr.IVRScript) (*Session, error) {
	return getSession(sessionID, script, true)
}

func loadSession(sessionID string, script *ivr.IVRScript) (*Session, error) { // Eventually should split into load/init
	return getSession(sessionID, script, false)
}

func getSession(sessionID string, script *ivr.IVRScript, initialize bool) (*Session, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, cloud.GcpProjectID, option.WithCredentialsFile(cloud.GcpCredentialsFileName))
	if err != nil {
		return nil, err
	}
	key := makeKey(sessionID)
	session := &Session{
		client: client,
		key:    key,
		ctx:    ctx,
		Data:   new(SessionData),
	}
	if initialize {
		session.initializeVariables(script.Variables)
		session.initializeDefaultVariables()
	} else {
		err = client.Get(ctx, key, session.Data)
	}
	return session, err
}

func (session *Session) delete() error {
	return session.client.Delete(session.ctx, session.key)
}

func (session *Session) close() error {
	return session.client.Close()
}

func (session *Session) setParameter(name string, val string) {
	variable := session.getParameter(name)
	variable.Value = val
}

func (session *Session) setParameterFromPBVal(name string, value *structpb.Value) error {
	variable := session.getParameter(name)
	var err error
	switch variable.ValType {
	case ivr.ValString:
		variable.Value = value.GetStringValue()
	case ivr.ValInteger:
		variable.Value, err = ivr.NewIntegerValue(int(value.GetNumberValue()))
	case ivr.ValNumeric:
		variable.Value, err = ivr.NewNumericValue(value.GetNumberValue())
	}
	return err
}

func (session *Session) getParameter(name string) *ivr.Variable {
	var ret *ivr.Variable
	found := false
	for _, variable := range session.Data.Variables {
		if string(variable.ID) == name {
			found = true
			ret = variable
		}
	}
	if !found {
		log.Panicf("Variable not found with name: %s", name)
	}
	return ret
}

func (session *Session) initializeVariables(variables ivr.Variables) {
	for _, variable := range variables {
		session.Data.Variables = append(session.Data.Variables, variable)
	}
}

func (session *Session) initializeDefaultVariables() {
	defaults := ivr.Variables{
		"__BUFFER__": &ivr.Variable{
			ID:      "__BUFFER__",
			ValType: ivr.ValString,
			VarType: ivr.VarUserVariable,
			Value:   "",
		},
		"__ExtContactFields__": &ivr.Variable{
			ID:      "__ExtContactFields__",
			ValType: ivr.ValKVList,
			VarType: ivr.VarUserVariable,
			Value:   "{}",
		},
	}
	session.initializeVariables(defaults)
}

func makeKey(sessionID string) *datastore.Key {
	return datastore.NameKey("SessionData", sessionID, nil)
}
