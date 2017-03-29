package variables

import (
	"testing"
)

func TestInteger(t *testing.T) {
	//	v1 := "012345"
	iv := IntegerValue{false, 1234}
	strIv, err := iv.convertToString()
	if err != nil {
		t.Error(err)
		return
	}
	if strIv != "1234" {
		t.Error("Expected 1234", "got ", strIv)
	}
}
