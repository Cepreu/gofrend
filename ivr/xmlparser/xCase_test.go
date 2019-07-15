package xmlparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
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

	s := &ivr.IVRScript{
		Variables: make(ivr.Variables),
	}
	res := newCaseModule(s, decoder)
	if res == nil {
		t.Fatal("Case module wasn't parsed...")
	}
	var mCase = (res.(xmlCaseModule)).m
	qwerty, _ := addStringConstant(s, "qwerty")
	var expected = &ivr.CaseModule{
		Branches: []*ivr.OutputBranch{
			&ivr.OutputBranch{
				Name:       "bA",
				Descendant: "D7F8916C13384EE08A6109F54109307E",
				Cond: &ivr.ComplexCondition{
					CustomCondition: "1",
					Conditions: []*ivr.Condition{
						{
							ComparisonType: "EQUALS",
							RightOperand:   qwerty,
							LeftOperand:    "__BUFFER__",
						},
					},
				},
			},
			&ivr.OutputBranch{
				Name:       "bB",
				Descendant: "D7F8916C13384EE08A6109F54109307E",
				Cond: &ivr.ComplexCondition{
					CustomCondition: "1",
					Conditions: []*ivr.Condition{
						{
							ComparisonType: "LIKE",
							RightOperand:   "Contact.city",
							LeftOperand:    "__BUFFER__",
						},
					},
				},
			},
			&ivr.OutputBranch{
				Name:       "bC",
				Descendant: "D7F8916C13384EE08A6109F54109307E",
				Cond: &ivr.ComplexCondition{
					CustomCondition: "1",
					Conditions: []*ivr.Condition{
						{
							ComparisonType: "IS_NULL",
							LeftOperand:    "__BUFFER__",
						},
					},
				},
			},
			&ivr.OutputBranch{Name: "No Match", Descendant: "D7F8916C13384EE08A6109F54109307E", Cond: nil},
		},
	}
	expected.SetGeneralInfo("Case3", "D2CC05B0F6FC44F29B04C1C9E42DF732",
		[]ivr.ModuleID{"368A8C40D5AD48668FB2DC7ED894B3BA"}, "", "", "", "false")

	exp, err1 := json.MarshalIndent(expected, "", "  ")
	setv, err2 := json.MarshalIndent(mCase, "", "  ")

	if err1 != nil || err2 != nil || string(exp) != string(setv) {
		t.Errorf("\nCase module: \n%s \n\nwas expected, in reality: \n\n%s", string(exp), string(setv))
	}
}
