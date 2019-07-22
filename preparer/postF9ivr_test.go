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
