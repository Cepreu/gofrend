package fulfiller

import "github.com/Cepreu/gofrend/ivr/vars"

// StorageVariable is an un-interfaced struct for
type StorageVariable struct {
	Type         string
	Name         string
	IntegerValue *vars.Integer
}

func makeStorageVar(val vars.Value) *StorageVariable {
	switch v := val.(type) {
	case *vars.Integer:
		return &StorageVariable{
			Type:         "Integer",
			IntegerValue: v,
		}
	default:
		panic("Not implemented")
	}
}

func (storageVar *StorageVariable) value() vars.Value {
	switch storageVar.Type {
	case "Integer":
		return storageVar.IntegerValue
	default:
		panic("Not implemented")
	}
}
