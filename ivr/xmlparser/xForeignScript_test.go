package xmlparser

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
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
	//res.normalize()
	var m = (res.(xmlForeignScriptModule)).s

	expected := &ivr.ForeignScriptModule{
		IvrScript: "EXAMPLES..Five9..EmailRoutedByRulesEngine",
		PassCRM:   false,
		ReturnCRM: false,
		Parameters: []ivr.KeyValueParametrized{
			{
				Key: "__ExtContactFields__",
				Value: &ivr.Parametrized{
					VariableName: "__ExtContactFields__",
				},
			},
		},
		ReturnParameters: []ivr.KeyValue{
			{Key: "__ExtContactFields__", Value: "__ExtContactFields__"},
		},
		IsConsistent: true,
	}
	expected.SetGeneralInfo("ForeignScript19", "A39E830442574C8998F727E2171BFA1D",
		[]ivr.ModuleID{"EE46D4FA17844064B679BBCABB45CDE8"}, "A917886FAA054CF581C4BA3798FCB836", "",
		"No Disposition", "false")

	if false == reflect.DeepEqual(expected.Parameters[0], m.Parameters[0]) {
		t.Errorf("\nForeignScript module: \n%v \nwas expected, in reality: \n%v", expected.Parameters[0], m.Parameters[0])
	}
	if false == reflect.DeepEqual(expected.ReturnParameters[0], m.ReturnParameters[0]) {
		t.Errorf("\nForeignScript module: \n%v \nwas expected, in reality: \n%v", expected.ReturnParameters[0], m.ReturnParameters[0])
	}
	// if false == reflect.DeepEqual(expected.GeneralInfo, m.GeneralInfo) {
	// 	t.Errorf("\nForeignScript module, general info: \n%v \nwas expected, in reality: \n%v",
	// 		expected.GeneralInfo, m.GeneralInfo)
	// }
	// more sanity checking...
}
