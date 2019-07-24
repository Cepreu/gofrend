package preparer

import (
	"fmt"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
)

var script = ivr.IVRScript{Domain: "qwerty", Name: "rememberme"}
var user, pwd = "sergei@007.com", "five9demo"

func TestGetIvrFromF9(t *testing.T) {
	ivr, err := getIvrFromF9("sergei@007.com", "five9demo", "delme")
	if err != nil || len(ivr) == 0 {
		t.Errorf("TestGetIvrFromF9: \n%v", err)
	}
}

func TestCreateIvrF9(t *testing.T) {
	generatedScriptName := "rememberme_generated"
	resp, err := queryF9(user, pwd, func() string { return createIVRscriptContent(generatedScriptName) })
	if err != nil {
		fault, _ := getf9errorDescr(resp)
		exp := fmt.Sprintf("IvrScript with name \"%s\" already exists", generatedScriptName)
		if exp != fault.FaultString {
			t.Errorf("TestCreateIvrF9: expected OK or \"%s\", received \"%s\"", exp, fault.FaultString)
		}
	}
}

func TestConfigureF9(t *testing.T) {
	err := configureF9(user, pwd, "sergei_inbound", &script)
	if err != nil {
		t.Errorf("TestConfigureF9: \n%v", err)
	}
}
