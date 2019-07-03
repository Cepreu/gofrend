package fulfiller

import "github.com/Cepreu/gofrend/ivr/vars"

// StorageVariable is an un-interfaced struct for storing variables between sessions
type StorageVariable struct {
	Type         string
	Name         string
	IntegerValue *vars.Integer
	StringValue  *vars.String
}

func makeStorageVar(val vars.Value) *StorageVariable {
	switch v := val.(type) {
	case *vars.Integer:
		return &StorageVariable{
			Type:         "Integer",
			IntegerValue: v,
		}
	case *vars.String:
		return &StorageVariable{
			Type:        "String",
			StringValue: v,
		}
	default:
		panic("Not implemented")
	}
}

func (storageVar *StorageVariable) value() vars.Value {
	switch storageVar.Type {
	case "Integer":
		return storageVar.IntegerValue
	case "String":
		return storageVar.StringValue
	default:
		panic("Not implemented")
	}
}
