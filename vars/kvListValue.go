package vars

import (
	"encoding/json"
)

type KVListValue struct {
	secure bool
	value  map[string]string
}

//NewKVListValue - returns pointer to a new KVListValue struct, or <nil> for an error
func NewKVListValue(s string) *KVListValue {
	f := KVListValue{secure: false, value: make(map[string]string)}
	if err := json.Unmarshal([]byte(s), &f.value); err != nil {
		return nil
	}
	return &f
}

func (sval *KVListValue) isSecure() bool { return sval.secure }

func (sval *KVListValue) new(secure bool, strValue string) error {
	return nil
}

func (sval *KVListValue) String() string {
	bs, _ := json.Marshal(sval.value)
	return string(bs)
}

func (*KVListValue) getType() VarType {
	return VarString
}
