package ivr

import (
	"testing"
)

func checkNil(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeyValue(t *testing.T) {
	kvliststring, err := NewKeyValue("{}")
	checkNil(err, t)
	kvlist := StringToKVList(kvliststring)
	kvlist.Put("A", "a")
	kvlist.Put("B", "b")
	store := kvlist.ToString()
	kvlist = StringToKVList(store)
	if kvlist.Get("B") != "b" {
		panic("KVList storage broken")
	}
}
