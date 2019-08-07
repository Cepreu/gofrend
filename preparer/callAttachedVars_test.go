package preparer

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
)

func TestParseCAVResponse(t *testing.T) {
	const resp = `<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/"><env:Header/><env:Body><ns2:getCallVariablesResponse xmlns:ns2="http://service.admin.ws.five9.com/"><return><applyToAllDispositions>false</applyToAllDispositions><defaultValue>300</defaultValue><description></description><group>sergei</group><name>varNumber</name><reporting>false</reporting><restrictions><type>MaxValue</type><value>6000</value></restrictions><restrictions><type>Precision</type><value>5</value></restrictions><restrictions><type>MinValue</type><value>100</value></restrictions><restrictions><type>Scale</type><value>0</value></restrictions><restrictions><type>Required</type><value></value></restrictions><sensitiveData>false</sensitiveData><type>NUMBER</type></return><return><applyToAllDispositions>false</applyToAllDispositions><defaultValue>3.1245</defaultValue><description></description><group>sergei</group><name>varNumber2</name><reporting>false</reporting><restrictions><type>Required</type><value></value></restrictions><restrictions><type>Precision</type><value>9</value></restrictions><restrictions><type>Scale</type><value>4</value></restrictions><sensitiveData>false</sensitiveData><type>NUMBER</type></return></ns2:getCallVariablesResponse></env:Body></env:Envelope>`
	var script = ivr.IVRScript{Domain: "qwerty", Name: "rememberme",
		Modules:   make(map[ivr.ModuleID]ivr.Module),
		Variables: make(map[ivr.VariableID]*ivr.Variable),
	}

	err := parseCAVResponse([]byte(resp), &script)

	exp, err := json.MarshalIndent(script.Variables, "", "  ")

	fmt.Println(string(exp))

	if err == nil {
		t.Errorf("parseCAVResponse: \n%v", err)
	}
}
