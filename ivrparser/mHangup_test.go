package ivrparser

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/vars"
	"golang.org/x/net/html/charset"
)

func TestHangup(t *testing.T) {
	var xmlData = `
<hangup>
	<ascendants>ED132095BE1E4F47B51DA0BB842C3EEF</ascendants>
	<ascendants>F1E142D8CF27471D8940713A637A1C1D</ascendants>
	<moduleName>Hangup13</moduleName>
	<locationX>533</locationX>
	<locationY>69</locationY>
	<moduleId>A96A2609FDDE4C499773122F6C6296A1</moduleId>
	<data>
		<dispo>
			<id>0</id>
			<name>No Disposition</name>
		</dispo>
		<returnToCallingModule>true</returnToCallingModule>
		<errCode>
			<isVarSelected>true</isVarSelected>
			<integerValue>
				<value>94</value>
			</integerValue>
			<variableName>duration</variableName>
		</errCode>
		<errDescription>
			<isVarSelected>false</isVarSelected>
			<stringValue>
				<value>Hello, World!!!</value>
				<id>0</id>
			</stringValue>
			<variableName>Contact.city</variableName>
		</errDescription>
		<overwriteDisposition>true</overwriteDisposition>
	</data>
</hangup>
`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	res := newHangupModule(decoder)
	if res == nil {
		t.Fatal("Hangup module wasn't parsed...")
	}
	var mhu = res.(*hangupModule)
	var expected = hangupModule{
		generalInfo: generalInfo{
			ID: "A96A2609FDDE4C499773122F6C6296A1",
			Ascendants: []moduleID{
				"ED132095BE1E4F47B51DA0BB842C3EEF",
				"F1E142D8CF27471D8940713A637A1C1D"},
			Descendant:      "",
			ExceptionalDesc: "",
			Name:            "Hangup13",
			Dispo:           "No Disposition",
			Collapsible:     false,
		},
		Return2Caller: true,
		ErrCode:       parametrized{VariableName: "duration"},
		ErrDescr:      parametrized{Value: vars.NewString("Hello, World!!!", 0)},
		OverwriteDisp: true,
	}

	if false == reflect.DeepEqual(&expected, mhu) {
		t.Errorf("\nHangup module: \n%v \nwas expected, in reality: \n%v", expected, mhu)
	}

}
