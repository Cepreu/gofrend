package vars

import (
	"encoding/json"
	"fmt"
)

// KVListVar - KVListV variable
type KVListVar struct {
	Variable
	value *KVList
}

// KVList - value of KVList type
type KVList struct {
	value map[string]string
}

//NewKVList - returns pointer to a new KVList struct, or <nil> for an error
func NewKVList(s string) *KVList {
	f := KVList{make(map[string]string)}
	if err := json.Unmarshal([]byte(s), &f.value); err != nil {
		return nil
	}
	return &f
}

//SetValue - helper function for parsing xml
func (kval KVList) SetValue(fieldName string, fieldStrValue string) (err error) {
	switch fieldName {
	case "value":
		kval.value = make(map[string]string)
		err = json.Unmarshal([]byte(fieldStrValue), kval.value)
	default:
		err = fmt.Errorf("Unknown field '%s' for KVList value", fieldName)
	}
	return err
}

//TBD
func (sval *KVList) new(strValue string) error {
	return nil
}

func (sval *KVList) String() string {
	bs, _ := json.Marshal(sval.value)
	return string(bs)
}

func (*KVList) getType() VarType {
	return VarString
}
