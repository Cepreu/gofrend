package fulfiller

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Cepreu/gofrend/cloud"
	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
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
	Variables []*StorageVariable
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

func (session *Session) setParameterString(name string, str string) error {
	value := &structpb.Value{
		Kind: &structpb.Value_StringValue{
			StringValue: str,
		},
	}
	return session.setParameter(name, value)
}

func (session *Session) setParameter(name string, value *structpb.Value) error {
	variable, ok := session.getParameter(name)
	if !ok {
		utils.PrettyLog(session.Data)
		return fmt.Errorf("Could not find session variable with name: %s", name)
	}
	variable.StringValue = value.GetStringValue()
	return nil
}

func (session *Session) getParameter(name string) (*StorageVariable, bool) {
	var ret *StorageVariable
	found := false
	for _, storageVar := range session.Data.Variables {
		if storageVar.Name == name {
			found = true
			ret = storageVar
		}
	}
	return ret, found
}

func (session *Session) initializeVariables(variables ivr.Variables) {
	var storageVar *StorageVariable
	for name, variable := range variables {
		storageVar = &StorageVariable{
			Name:        string(name),
			Type:        variable.ValType.String(),
			StringValue: variable.Value.StringValue,
		}
		session.Data.Variables = append(session.Data.Variables, storageVar)
	}
}

func (session *Session) initializeDefaultVariables() {
	defaults := ivr.Variables{
		"__BUFFER__": &ivr.Variable{
			ID:      "__BUFFER__",
			ValType: ivr.ValString,
			VarType: ivr.VarUserVariable,
			Value:   &ivr.Value{StringValue: ""},
		},
		"__ExtContactFields__": &ivr.Variable{
			ID:      "__ExtContactFields__",
			ValType: ivr.ValKVList,
			VarType: ivr.VarUserVariable,
			Value:   &ivr.Value{StringValue: "{}"},
		},
	}
	session.initializeVariables(defaults)
}

func makeKey(sessionID string) *datastore.Key {
	return datastore.NameKey("SessionData", sessionID, nil)
}
