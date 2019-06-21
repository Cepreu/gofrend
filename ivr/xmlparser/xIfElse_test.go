package xmlparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/vars"
	"golang.org/x/net/html/charset"
)

func TestIfElse(t *testing.T) {
	var xmlData = `
<ifElse>
	<ascendants>E1149703A8E543BE93F142456FED91F6</ascendants>
	<moduleName>IfElse5</moduleName>
	<locationX>346</locationX>
	<locationY>245</locationY>
	<moduleId>EEC8BE771FBC4E659AC3CA54F4FBEBF4</moduleId>
	<data>
		<branches>
			<entry>
				<key>IF</key>
				<value>
					<name>IF</name>
					<desc>4663D0E48FA048CF938695D75A06739D</desc>
				</value>
			</entry>
			<entry>
				<key>ELSE</key>
				<value>
					<name>ELSE</name>
					<desc>497089347DB34EB6A2C9E72518F75125</desc>
				</value>
			</entry>
		</branches>
		<customCondition>(1 or 2) AND 3</customCondition>
		<conditionGrouping>CUSTOM</conditionGrouping>
		<conditions>
			<comparisonType>EQUALS</comparisonType>
			<joinMode>AND</joinMode>
			<rightOperand>
				<isVarSelected>true</isVarSelected>
				<variableName>Call.ANI</variableName>
			</rightOperand>
			<leftOperand>
				<isVarSelected>true</isVarSelected>
				<variableName>Contact.number1</variableName>
			</leftOperand>
		</conditions>
		<conditions>
			<comparisonType>EQUALS</comparisonType>
			<joinMode>AND</joinMode>
			<rightOperand>
				<isVarSelected>true</isVarSelected>
				<variableName>Call.DNIS</variableName>
			</rightOperand>
			<leftOperand>
				<isVarSelected>true</isVarSelected>
				<variableName>Contact.number2</variableName>
			</leftOperand>
		</conditions>
		<conditions>
			<comparisonType>EQUALS</comparisonType>
			<joinMode>AND</joinMode>
			<rightOperand>
				<isVarSelected>false</isVarSelected>
				<stringValue>
					<value>9252012040</value>
					<id>0</id>
				</stringValue>
			</rightOperand>
			<leftOperand>
				<isVarSelected>true</isVarSelected>
				<variableName>Contact.number3</variableName>
			</leftOperand>
		</conditions>
	</data>
</ifElse>
`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	res := newIfElseModule(decoder)
	if res == nil {
		t.Fatal("IfElse module wasn't parsed...")
	}
	var mie = (res.(xmlIfElseModule)).s
	var expected = &ivr.IfElseModule{
		BranchIf: ivr.OutputBranch{"IF", "4663D0E48FA048CF938695D75A06739D",
			&ivr.ComplexCondition{
				CustomCondition:   "(1 or 2) AND 3",
				ConditionGrouping: "CUSTOM",
				Conditions: []*ivr.Condition{
					{ComparisonType: "EQUALS",
						RightOperand: ivr.Parametrized{VariableName: "Call.ANI"},
						LeftOperand:  ivr.Parametrized{VariableName: "Contact.number1"},
					},
					{ComparisonType: "EQUALS",
						RightOperand: ivr.Parametrized{VariableName: "Call.DNIS"},
						LeftOperand:  ivr.Parametrized{VariableName: "Contact.number2"},
					},
					{ComparisonType: "EQUALS",
						RightOperand: ivr.Parametrized{Value: vars.NewString("9252012040", 0)},
						LeftOperand:  ivr.Parametrized{VariableName: "Contact.number3"},
					},
				},
			},
		},
		BranchElse: ivr.OutputBranch{Name: "ELSE", Descendant: "497089347DB34EB6A2C9E72518F75125"},
	}

	expected.SetGeneralInfo("IfElse5", "EEC8BE771FBC4E659AC3CA54F4FBEBF4",
		[]ivr.ModuleID{"E1149703A8E543BE93F142456FED91F6"}, "", "", "", "false")

	exp, err1 := json.MarshalIndent(expected, "", "  ")
	setv, err2 := json.MarshalIndent(mie, "", "  ")

	if err1 != nil || err2 != nil || string(exp) != string(setv) {
		t.Errorf("\nIfElse module: \n%s \n\nwas expected, in reality: \n\n%s", string(exp), string(setv))
	}
}
