package ivrparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"golang.org/x/net/html/charset"
)

func TestForeignScript(t *testing.T) {
	var xmlData = `
	<foreignScript>
	<ascendants>EE46D4FA17844064B679BBCABB45CDE8</ascendants>
	<singleDescendant>A917886FAA054CF581C4BA3798FCB836</singleDescendant>
	<moduleName>ForeignScript19</moduleName>
	<locationX>152</locationX>
	<locationY>161</locationY>
	<moduleId>A39E830442574C8998F727E2171BFA1D</moduleId>
	<data>
		<ivrScript>
			<id>44235</id>
			<name>EXAMPLES..Five9..EmailRoutedByRulesEngine</name>
		</ivrScript>
		<passCRM>false</passCRM>
		<returnCRM>false</returnCRM>
		<params>
			<entry>
				<key>__ExtContactFields__</key>
				<value>
					<isVarSelected>true</isVarSelected>
					<variableName>__ExtContactFields__</variableName>
				</value>
			</entry>
		</params>
		<returnVals>
			<entry>
				<key>__ExtContactFields__</key>
				<value>__ExtContactFields__</value>
			</entry>
		</returnVals>
		<isConsistent>true</isConsistent>
	</data>
</foreignScript>
`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	res := newForeignScriptModule(decoder)
	if res == nil {
		t.Errorf("ForeignScript module wasn't parsed...")
		return
	}
<<<<<<< HEAD:ivrparser/mForeignScript_test.go
	//res.normalize()
	var m = res.(*ForeignScriptModule)
=======

	var m = (res.(xmlForeignScriptModule)).s
>>>>>>> 8ffe0c172a3900f3340cffda428678c03bc5cb36:ivr/xmlparser/xForeignScript_test.go

	expected := &ForeignScriptModule{
		GeneralInfo: GeneralInfo{
			ID:         "A39E830442574C8998F727E2171BFA1D",
			Ascendants: []ModuleID{"EE46D4FA17844064B679BBCABB45CDE8"},
			Descendant: "A917886FAA054CF581C4BA3798FCB836",
			Name:       "ForeignScript19",
		},
		IvrScript: "EXAMPLES..Five9..EmailRoutedByRulesEngine",
		PassCRM:   false,
		ReturnCRM: false,
		Parameters: []keyValueParametrized{
			{
				Key: "__ExtContactFields__",
				Value: &parametrized{
					VariableName: "__ExtContactFields__",
				},
			},
		},
		ReturnParameters: []keyValue{
			{Key: "__ExtContactFields__", Value: "__ExtContactFields__"},
		},
		IsConsistent: true,
	}
<<<<<<< HEAD:ivrparser/mForeignScript_test.go
=======
	expected.SetGeneralInfo("ForeignScript19", "A39E830442574C8998F727E2171BFA1D",
		[]ivr.ModuleID{"EE46D4FA17844064B679BBCABB45CDE8"}, "A917886FAA054CF581C4BA3798FCB836", "",
		"", "false")
>>>>>>> 8ffe0c172a3900f3340cffda428678c03bc5cb36:ivr/xmlparser/xForeignScript_test.go

	exp, err1 := json.MarshalIndent(expected, "", "  ")
	setv, err2 := json.MarshalIndent(m, "", "  ")

	if err1 != nil || err2 != nil || string(exp) != string(setv) {
		t.Errorf("\nMenu module: \n%s \n\nwas expected, in reality: \n\n%s", string(exp), string(setv))
	}
<<<<<<< HEAD:ivrparser/mForeignScript_test.go
	if false == reflect.DeepEqual(expected.GeneralInfo, m.GeneralInfo) {
		t.Errorf("\nForeignScript module, general info: \n%v \nwas expected, in reality: \n%v",
			expected.GeneralInfo, m.GeneralInfo)
	}
	// more sanity checking...
=======
>>>>>>> 8ffe0c172a3900f3340cffda428678c03bc5cb36:ivr/xmlparser/xForeignScript_test.go
}
