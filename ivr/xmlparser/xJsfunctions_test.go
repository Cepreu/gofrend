package xmlparser

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"

	"golang.org/x/net/html/charset"
)

func TestJSFunctions(t *testing.T) {
	var xmlData = `<functions>
<entry>
	<key>01622849F1F54399A411F5DBC9BC80B7</key>
	<value>
		<jsFunctionId>01622849F1F54399A411F5DBC9BC80B7</jsFunctionId>
		<description></description>
		<returnType>KVLIST</returnType>
		<name>days_of_week</name>
		<arguments>
			<arguments>
				<name>month</name>
				<description></description>
				<type>INTEGER</type>
			</arguments>
			<arguments>
				<name>year</name>
				<description></description>
				<type>INTEGER</type>
			</arguments>
		</arguments>
		<functionBody>H4sIAAAAAAAAAHWQTwuCQBDF7/spBk+KEtpfyDwE0a1TQQfxsLBTRrTCqImI3711NyM2uu2+37yd
fY89OYEUvC0hAYkN7HiFboucAngUssoDCL3JFSutezFjg6FBvH95tkS8dVeKjiANM8WcYz087Xzp
kdYPha1PtX6qsbTATIMzCvmD5saT12SThSZ7uln60vyKVzUZwi4FgSuSKAaxSXQP6uT7HnQMAIas
/G8xQgV+D5Ea+mzhpq/W9bKY9YwRqn0S0s6JnPWYvQ/MdYzcZy/qB7vqiwEAAA==</functionBody>
	</value>
</entry>
<entry>
	<key>D3E76AA32B644F3AB11707B5B26CAF31</key>
	<value>
		<jsFunctionId>D3E76AA32B644F3AB11707B5B26CAF31</jsFunctionId>
		<description>Calculates Sinus</description>
		<returnType>NUMERIC</returnType>
		<name>sin</name>
		<arguments>
			<arguments>
				<name>angle</name>
				<description>Angle (degrees)</description>
				<type>NUMERIC</type>
			</arguments>
		</arguments>
		<functionBody>H4sIAAAAAAAAAEvMS89JVbBV8E0sydAL8FTQUkgEiegbWhjoGShwFaWWlBblQWSLM/M0wJKaAEH3
Sio1AAAA</functionBody>
	</value>
</entry>
</functions>`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	res := newJSFunctions(decoder)
	if res == nil {
		t.Errorf("JSFunsctions  weren't parsed...")
		return
	}

	expected := []*ivr.JsFunction{
		{
			JsFunctionID: "01622849F1F54399A411F5DBC9BC80B7",
			ReturnType:   "KVLIST",
			Name:         "days_of_week",
			Arguments: []*ivr.FuncArgument{
				{
					Name:    "month",
					ArgType: "INTEGER",
				},
				{
					Name:    "year",
					ArgType: "INTEGER",
				},
			},
			FuncBody: `
var ndays = new Date(year, month, 0).getDate();

var weekdays = new Array(7);
weekdays[0] = "Sunday";
weekdays[1] = "Monday";
weekdays[2] = "Tuesday";
weekdays[3] = "Wednesday";
weekdays[4] = "Thursday";
weekdays[5] = "Friday";
weekdays[6] = "Saturday";

for (d=1; d<=ndays; d++) {
	var a = new Date(year, month, d);
	var r = weekdays[a.getDay()];
}

return [{"1":"Monday"},{"1":"Tuesday"}]`,
		},
		{
			JsFunctionID: "D3E76AA32B644F3AB11707B5B26CAF31",
			ReturnType:   "NUMERIC",
			Name:         "Calculates Sinus",
			Arguments: []*ivr.FuncArgument{
				{
					Name:        "angle",
					ArgType:     "NUMERIC",
					Description: "Angle (degrees)",
				},
			},
			FuncBody: "",
		},
	}
	if len(expected) != len(res) {
		t.Errorf("\nJSFunsctions length: \n%v \nwas expected, in reality: \n%v", len(expected), len(res))
		return
	}
	if len(utils.StripSpaces(expected[0].FuncBody)) != len(utils.StripSpaces(res[0].FuncBody)) {
		t.Errorf("\nJSFunsctions length: \n%v \nwas expected, in reality: \n%v", len(utils.StripSpaces(expected[0].FuncBody)), len(utils.StripSpaces(res[0].FuncBody)))
		return
	}
	if expected[0].Name != res[0].Name ||
		expected[0].ReturnType != res[0].ReturnType ||
		expected[0].JsFunctionID != res[0].JsFunctionID {
		t.Errorf("\nJSFunsctions: \n%v \nwas expected, in reality: \n%v", expected[0], res[0])
		return
	}
	if false == reflect.DeepEqual(expected[0].Arguments[0], res[0].Arguments[0]) {
		t.Errorf("\nJSFunsctions: \n%v \nwas expected, in reality: \n%v", expected[0].Arguments[0], res[0].Arguments[0])
	}
	// if false == reflect.DeepEqual(expected.ReturnParameters[0], m.ReturnParameters[0]) {
	// 	t.Errorf("\nJSFunsctions: \n%v \nwas expected, in reality: \n%v", expected.ReturnParameters[0], m.ReturnParameters[0])
	// }
	// if false == reflect.DeepEqual(expected.GeneralInfo, m.GeneralInfo) {
	// 	t.Errorf("\nJSFunsctions general info: \n%v \nwas expected, in reality: \n%v",
	// 		expected.GeneralInfo, m.GeneralInfo)
	// }
	// more sanity checking...
}
