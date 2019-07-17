package fulfiller

import "github.com/Cepreu/gofrend/ivr"

// StorageVariable is an un-interfaced struct for storing variables between sessions
type StorageVariable struct {
	Type        string
	Name        string
	StringValue string
}

func makeStorageVar(v ivr.Variable) *StorageVariable {
	return &StorageVariable{
		Name:        string(v.ID),
		Type:        v.ValType.String(),
		StringValue: v.Value.StringValue,
	}
}

func (storageVar *StorageVariable) value() string {
	return storageVar.StringValue
}

func (storageVar *StorageVariable) valueAsString() string {
	return storageVar.StringValue
}
