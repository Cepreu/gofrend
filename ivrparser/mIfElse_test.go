package ivrparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

<<<<<<< HEAD:ivrparser/mIfElse_test.go
	"github.com/Cepreu/gofrend/utils"
	"github.com/Cepreu/gofrend/vars"
=======
	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/vars"
>>>>>>> 8ffe0c172a3900f3340cffda428678c03bc5cb36:ivr/xmlparser/xIfElse_test.go
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
	var mie = res.(*IfElseModule)
	utils.PrettyPrint(mie)
	var expected = &IfElseModule{
		GeneralInfo: GeneralInfo{
			ID:         "EEC8BE771FBC4E659AC3CA54F4FBEBF4",
			Ascendants: []ModuleID{"E1149703A8E543BE93F142456FED91F6"},
			Name:       "IfElse5",
		},
		BranchIf: OutputBranch{"IF", "4663D0E48FA048CF938695D75A06739D",
			&ComplexCondition{
				CustomCondition:   "(1 or 2) AND 3",
				ConditionGrouping: "CUSTOM",
				Conditions: []*Condition{
					{ComparisonType: "EQUALS",
						RightOperand: parametrized{VariableName: "Call.ANI"},
						LeftOperand:  parametrized{VariableName: "Contact.number1"},
					},
					{ComparisonType: "EQUALS",
						RightOperand: parametrized{VariableName: "Call.DNIS"},
						LeftOperand:  parametrized{VariableName: "Contact.number2"},
					},
					{ComparisonType: "EQUALS",
						RightOperand: parametrized{Value: vars.NewString("9252012040", 0)},
						LeftOperand:  parametrized{VariableName: "Contact.number3"},
					},
				},
			},
		},
		BranchElse: OutputBranch{Name: "ELSE", Descendant: "497089347DB34EB6A2C9E72518F75125"},
	}

<<<<<<< HEAD:ivrparser/mIfElse_test.go
	if !reflect.DeepEqual(expected.GeneralInfo, mie.GeneralInfo) ||
		!reflect.DeepEqual(expected.BranchElse, mie.BranchElse) ||
		!reflect.DeepEqual(expected.BranchIf.Cond.Conditions[0], mie.BranchIf.Cond.Conditions[0]) ||
		!reflect.DeepEqual(expected.BranchIf.Cond.Conditions[1], mie.BranchIf.Cond.Conditions[1]) ||
		!reflect.DeepEqual(expected.BranchIf.Cond.Conditions[2], mie.BranchIf.Cond.Conditions[2]) {
		t.Errorf("\nIfElse module: \n%v \nwas expected, in reality: \n%v", expected, mie)
=======
	expected.SetGeneralInfo("IfElse5", "EEC8BE771FBC4E659AC3CA54F4FBEBF4",
		[]ivr.ModuleID{"E1149703A8E543BE93F142456FED91F6"}, "", "", "", "false")

	exp, err1 := json.MarshalIndent(expected, "", "  ")
	setv, err2 := json.MarshalIndent(mie, "", "  ")

	if err1 != nil || err2 != nil || string(exp) != string(setv) {
		t.Errorf("\nIfElse module: \n%s \n\nwas expected, in reality: \n\n%s", string(exp), string(setv))
>>>>>>> 8ffe0c172a3900f3340cffda428678c03bc5cb36:ivr/xmlparser/xIfElse_test.go
	}
}
