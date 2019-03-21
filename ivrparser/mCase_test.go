package ivrparser

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

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
	var mCase = res.(*caseModule)
	var expected = &caseModule{
		GeneralInfo: GeneralInfo{
			ID:         "D2CC05B0F6FC44F29B04C1C9E42DF732",
			Ascendants: []ModuleID{"368A8C40D5AD48668FB2DC7ED894B3BA"},
			Name:       "Case3",
		},
		Branches: []*OutputBranch{
			&OutputBranch{"bA", "D7F8916C13384EE08A6109F54109307E",
				&ComplexCondition{
					CustomCondition:   "1",
					ConditionGrouping: "CUSTOM",
					Conditions: []*Condition{
						{
							comparisonType: "EQUALS",
							rightOperand:   parametrized{Value: vars.NewString("qwerty", 0)},
							leftOperand:    parametrized{VariableName: "__BUFFER__"},
						},
					},
				},
			},
			&OutputBranch{"bB", "D7F8916C13384EE08A6109F54109307E",
				&ComplexCondition{
					CustomCondition:   "1",
					ConditionGrouping: "CUSTOM",
					Conditions: []*Condition{
						{
							comparisonType: "LIKE",
							rightOperand:   parametrized{VariableName: "Contact.city"},
							leftOperand:    parametrized{VariableName: "__BUFFER__"},
						},
					},
				},
			},
			&OutputBranch{"bC", "D7F8916C13384EE08A6109F54109307E",
				&ComplexCondition{
					CustomCondition:   "1",
					ConditionGrouping: "CUSTOM",
					Conditions: []*Condition{
						{
							comparisonType: "IS_NULL",
							leftOperand:    parametrized{VariableName: "__BUFFER__"},
						},
					},
				},
			},
			&OutputBranch{"No Match", "D7F8916C13384EE08A6109F54109307E", nil},
		},
	}

	if !reflect.DeepEqual(expected.GeneralInfo, mCase.GeneralInfo) ||
		!reflect.DeepEqual(expected.Branches[3], mCase.Branches[3]) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[3].Name, expected.Branches[3], mCase.Branches[3])
	}

	if !reflect.DeepEqual(expected.Branches[0].Cond.Conditions[0].rightOperand, mCase.Branches[0].Cond.Conditions[0].rightOperand) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[0].Name, expected.Branches[0].Cond.Conditions[0].rightOperand, mCase.Branches[0].Cond.Conditions[0].rightOperand)
	}

	if !reflect.DeepEqual(expected.Branches[1].Cond.Conditions[0], mCase.Branches[1].Cond.Conditions[0]) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[1].Name, expected.Branches[0].Cond.Conditions[0], mCase.Branches[1].Cond.Conditions[0])
	}
	if !reflect.DeepEqual(expected.Branches[2].Cond.Conditions[0], mCase.Branches[2].Cond.Conditions[0]) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[2].Name, expected.Branches[2].Cond.Conditions[0], mCase.Branches[2].Cond.Conditions[0])
	}
}
