package fulfiller

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Cepreu/gofrend/cloud"
	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/vars"
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

func loadSession(sessionID string, script *ivr.IVRScript) (*Session, error) { // Eventually should split into load/init
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
	err = client.Get(ctx, key, session.Data)
	if err == datastore.ErrNoSuchEntity {
		session.initializeVariables(script.Variables)
		session.initializeDefaultVariables()
	} else if err != nil {
		return nil, err
	}
	return session, nil
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
	switch v := variable.value().(type) {
	case *vars.Integer:
		v.Value = int(value.GetNumberValue())
	case *vars.String:
		v.Value = value.GetStringValue()
	default:
		panic("Not implemented")
	}
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

func (session *Session) initializeVariables(variables map[string]*vars.Variable) {
	var storageVar *StorageVariable
	for name, variable := range variables {
		storageVar = &StorageVariable{
			Name: name,
		}
		switch v := variable.Value.(type) {
		case *vars.Integer:
			storageVar.Type = "Integer"
			storageVar.IntegerValue = v
		case *vars.String:
			storageVar.Type = "String"
			storageVar.StringValue = v
		default:
			panic("Not implemented")
		}
		session.Data.Variables = append(session.Data.Variables, storageVar)
	}
}

func (session *Session) initializeDefaultVariables() {
	defaults := map[string]*vars.Variable{
		"__BUFFER__": &vars.Variable{
			Value: &vars.String{
				Value: "",
			},
		},
		// TODO __ExtContactFields__ once kvlist variables are implemented
	}
	session.initializeVariables(defaults)
}

func makeKey(sessionID string) *datastore.Key {
	return datastore.NameKey("SessionData", sessionID, nil)
}
