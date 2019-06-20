package ivrparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

<<<<<<< HEAD:ivrparser/mCase_test.go
	"github.com/Cepreu/gofrend/vars"
=======
	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/vars"
>>>>>>> 8ffe0c172a3900f3340cffda428678c03bc5cb36:ivr/xmlparser/xCase_test.go
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
	var mCase = res.(*CaseModule)
	var expected = &CaseModule{
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
							ComparisonType: "EQUALS",
							RightOperand:   parametrized{Value: vars.NewString("qwerty", 0)},
							LeftOperand:    parametrized{VariableName: "__BUFFER__"},
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
							ComparisonType: "LIKE",
							RightOperand:   parametrized{VariableName: "Contact.city"},
							LeftOperand:    parametrized{VariableName: "__BUFFER__"},
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
							ComparisonType: "IS_NULL",
							LeftOperand:    parametrized{VariableName: "__BUFFER__"},
						},
					},
				},
			},
			&OutputBranch{"No Match", "D7F8916C13384EE08A6109F54109307E", nil},
		},
	}
<<<<<<< HEAD:ivrparser/mCase_test.go

	if !reflect.DeepEqual(expected.GeneralInfo, mCase.GeneralInfo) ||
		!reflect.DeepEqual(expected.Branches[3], mCase.Branches[3]) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[3].Name, expected.Branches[3], mCase.Branches[3])
	}

	if !reflect.DeepEqual(expected.Branches[0].Cond.Conditions[0].RightOperand, mCase.Branches[0].Cond.Conditions[0].RightOperand) {
		t.Errorf("\nCase module, branch \"%s\": \n%v \nwas expected, in reality: \n%v", mCase.Branches[0].Name, expected.Branches[0].Cond.Conditions[0].RightOperand, mCase.Branches[0].Cond.Conditions[0].RightOperand)
	}
=======
	expected.SetGeneralInfo("Case3", "D2CC05B0F6FC44F29B04C1C9E42DF732",
		[]ivr.ModuleID{"368A8C40D5AD48668FB2DC7ED894B3BA"}, "", "", "", "false")

	exp, err1 := json.MarshalIndent(expected, "", "  ")
	setv, err2 := json.MarshalIndent(mCase, "", "  ")
>>>>>>> 8ffe0c172a3900f3340cffda428678c03bc5cb36:ivr/xmlparser/xCase_test.go

	if err1 != nil || err2 != nil || string(exp) != string(setv) {
		t.Errorf("\nCase module: \n%s \n\nwas expected, in reality: \n\n%s", string(exp), string(setv))
	}
}
