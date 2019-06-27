package fulfiller

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/vars"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/api/option"
)

// Session contains data stored by webhook between sessions.
type Session struct {
	// Script variables.
	Client    *datastore.Client
	Key       *datastore.Key
	Ctx       context.Context
	Variables map[string]*vars.Variable
}

func (session *Session) store() error {
	_, err := session.Client.Put(session.Ctx, session.Key, session.Variables)
	return err
}

func loadSession(sessionID string, script *ivr.IVRScript) (*Session, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "f9-dialogflow-converter", option.WithCredentialsFile("credentials.json"))
	if err != nil {
		return nil, err
	}
	key := makeKey(sessionID)
	session := &Session{
		Client: client,
		Key:    key,
		Ctx:    ctx,
	}
	variables := make(map[string]*vars.Variable)
	err = client.Get(ctx, key, variables)
	if err == datastore.ErrNoSuchEntity {
		session.Variables = script.Variables
		return session, nil
	} else if err != nil {
		return nil, err
	}
	session.Variables = variables
	return session, nil
}

func (session *Session) delete() error {
	return session.Client.Delete(session.Ctx, session.Key)
}

func (session *Session) setParameter(name string, value *structpb.Value) error {
	variable, ok := session.Variables[name]
	if !ok {
		return fmt.Errorf("Could not find session variable with name: %s", name)
	}
	switch v := variable.Value.(type) {
	case *vars.Integer:
		v.Value = int(value.GetNumberValue())
	}
	return nil
}

func (session *Session) getParameter(name string) (*vars.Variable, bool) {
	v, ok := session.Variables[name]
	return v, ok
}

func makeKey(sessionID string) *datastore.Key {
	return datastore.NameKey("Session", sessionID, nil)
}
