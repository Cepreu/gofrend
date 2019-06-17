package xmlparser

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/vars"
	"golang.org/x/net/html/charset"
)

func TestCase(t *testing.T) {
	var xmlData = `
		<case>
			<ascendants>368A8C40D5AD48668FB2DC7ED894B3BA</ascendants>
			<moduleName>Case3</moduleName>
			<locationX>144</locationX>
			<locationY>62</locationY>
			<moduleId>D2CC05B0F6FC44F29B04C1C9E42DF732</moduleId>
			<data>
				<branches>
					<entry>
						<key>bA</key>
						<value>
							<name>bA</name>
							<desc>D7F8916C13384EE08A6109F54109307E</desc>
							<conditions>
								<comparisonType>EQUALS</comparisonType>
								<rightOperand>
									<isVarSelected>false</isVarSelected>
									<stringValue>
										<value>qwerty</value>
										<id>0</id>
									</stringValue>
								</rightOperand>
								<leftOperand>
									<isVarSelected>true</isVarSelected>
									<variableName>__BUFFER__</variableName>
								</leftOperand>
							</conditions>
						</value>
					</entry>
					<entry>
						<key>bB</key>
						<value>
							<name>bB</name>
							<desc>D7F8916C13384EE08A6109F54109307E</desc>
							<conditions>
								<comparisonType>LIKE</comparisonType>
								<rightOperand>
									<isVarSelected>true</isVarSelected>
									<variableName>Contact.city</variableName>
								</rightOperand>
								<leftOperand>
									<isVarSelected>true</isVarSelected>
									<variableName>__BUFFER__</variableName>
								</leftOperand>
							</conditions>
						</value>
					</entry>
					<entry>
						<key>bC</key>
						<value>
							<name>bC</name>
							<desc>D7F8916C13384EE08A6109F54109307E</desc>
							<conditions>
								<comparisonType>IS_NULL</comparisonType>
								<leftOperand>
									<isVarSelected>true</isVarSelected>
									<variableName>__BUFFER__</variableName>
								</leftOperand>
							</conditions>
						</value>
					</entry>
					<entry>
						<key>No Match</key>
						<value>
							<name>No Match</name>
							<desc>D7F8916C13384EE08A6109F54109307E</desc>
							<conditions/>
						</value>
					</entry>
				</branches>
			</data>
		</case>
		`

	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	res := newCaseModule(decoder)
	if res == nil {
		t.Fatal("Case module wasn't parsed...")
	}
	var mCase = (res.(xmlCaseModule)).m
	var expected = &ivr.CaseModule{
		Branches: []*ivr.OutputBranch{
			&ivr.OutputBranch{"bA", "D7F8916C13384EE08A6109F54109307E",
				&ivr.ComplexCondition{
					CustomCondition:   "1",
					ConditionGrouping: "CUSTOM",
					Conditions: []*ivr.Condition{
						{
							ComparisonType: "EQUALS",
							RightOperand:   ivr.Parametrized{Value: vars.NewString("qwerty", 0)},
							LeftOperand:    ivr.Parametrized{VariableName: "__BUFFER__"},
						},
					},
				},
			},
			&ivr.OutputBranch{"bB", "D7F8916C13384EE08A6109F54109307E",
				&ivr.ComplexCondition{
					CustomCondition:   "1",
					ConditionGrouping: "CUSTOM",
					Conditions: []*ivr.Condition{
						{
							ComparisonType: "LIKE",
							RightOperand:   ivr.Parametrized{VariableName: "Contact.city"},
							LeftOperand:    ivr.Parametrized{VariableName: "__BUFFER__"},
						},
					},
				},
			},
			&ivr.OutputBranch{"bC", "D7F8916C13384EE08A6109F54109307E",
				&ivr.ComplexCondition{
					CustomCondition:   "1",
					ConditionGrouping: "CUSTOM",
					Conditions: []*ivr.Condition{
						{
							ComparisonType: "IS_NULL",
							LeftOperand:    ivr.Parametrized{VariableName: "__BUFFER__"},
						},
					},
				},
			},
			&ivr.OutputBranch{"No Match", "D7F8916C13384EE08A6109F54109307E", nil},
		},
	}
	expected.SetGeneralInfo("Case3", "D2CC05B0F6FC44F29B04C1C9E42DF732",
		[]ivr.ModuleID{"368A8C40D5AD48668FB2DC7ED894B3BA"}, "", "", "No Disposition", "false")

	//	if !reflect.DeepEqual(expected.GeneralInfo, mCase.GeneralInfo) ||
	if !reflect.DeepEqual(expected.Branches[3], mCase.Branches[3]) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[3].Name, expected.Branches[3], mCase.Branches[3])
	}

	if !reflect.DeepEqual(expected.Branches[0].Cond.Conditions[0].RightOperand, mCase.Branches[0].Cond.Conditions[0].RightOperand) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[0].Name, expected.Branches[0].Cond.Conditions[0].RightOperand, mCase.Branches[0].Cond.Conditions[0].RightOperand)
	}

	if !reflect.DeepEqual(expected.Branches[1].Cond.Conditions[0], mCase.Branches[1].Cond.Conditions[0]) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[1].Name, expected.Branches[0].Cond.Conditions[0], mCase.Branches[1].Cond.Conditions[0])
	}
	if !reflect.DeepEqual(expected.Branches[2].Cond.Conditions[0], mCase.Branches[2].Cond.Conditions[0]) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[2].Name, expected.Branches[2].Cond.Conditions[0], mCase.Branches[2].Cond.Conditions[0])
	}
}
