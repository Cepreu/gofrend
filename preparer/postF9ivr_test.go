package preparer

import (
	"testing"

	"github.com/Cepreu/gofrend/ivr"
)

var script = ivr.IVRScript{Domain: "qwerty", Name: "Test"}

func TestGetIVRContent(t *testing.T) {
	s := getIVRscriptContent("test")
	t.Errorf("TestGetIVRScriptContent: \n%s ", s)
}

func TestCreateIVRContent(t *testing.T) {
	s := createIVRscriptContent(&script)
	t.Errorf("TestCreateIVRScriptContent: \n%s ", s)
}

func TestModifyIVRscriptContent(t *testing.T) {
	s := modifyIVRscriptContent(&script)
	t.Errorf("TestModifyIVRscriptContent: \n%s ", s)
}

func TestSetDefaultIVRScheduleContent(t *testing.T) {
	s := setDefaultIVRScheduleContent("sergei_inbound", "Five9 SFDC", []struct {
		Name  string
		Value string
	}{{"Object", "Accounts"}, {"Fields", `{"name":"qwerty","Id":"123"}`}})
	t.Errorf("TestSetDefaultIVRScheduleContent: \n%s ", s)
}

func TestGetIvrFromF9(t *testing.T) {
	ivr, err := getIvrFromF9("sergei@007.com", "five9demo", "delme")
	if err == nil {
		t.Errorf("TestGetIvrFromF9: \n%s", ivr)
	}
}
