package xmlparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
	"golang.org/x/net/html/charset"
)

func TestSetVariables(t *testing.T) {
	var xmlData = `<setVariable>
	<ascendants>7E06FB60DFAA4F3E89946625F8677CB3</ascendants>
	<singleDescendant>4065B9160FE54714AEFAB56B9F9CB9C6</singleDescendant>
	<moduleName>SetVariable107</moduleName>
	<locationX>152</locationX>
	<locationY>118</locationY>
	<moduleId>05FDBF862B6F493AAAAFEC26E9FC8733</moduleId>
	<data>
		<expressions>
			<variableName>__BUFFER__</variableName>
			<isFunction>false</isFunction>
			<constant>
				<isVarSelected>false</isVarSelected>
				<stringValue>
					<value>CONST STRING</value>
					<id>0</id>
				</stringValue>
			</constant>
		</expressions>
		<expressions>
			<variableName>mystring</variableName>
			<isFunction>false</isFunction>
			<constant>
				<isVarSelected>true</isVarSelected>
				<variableName>__BUFFER__</variableName>
			</constant>
		</expressions>
		<expressions>
			<variableName>myint</variableName>
			<isFunction>false</isFunction>
			<constant>
				<isVarSelected>false</isVarSelected>
				<numericValue>
					<value>1000000</value>
				</numericValue>
			</constant>
		</expressions>
		<expressions>
			<variableName>__BUFFER__</variableName>
			<isFunction>true</isFunction>
			<function>
				<returnType>STRING</returnType>
				<name>TOSTRING</name>
				<arguments>INTEGER</arguments>
			</function>
			<functionArgs>
				<isVarSelected>true</isVarSelected>
				<variableName>Contact.order_id</variableName>
			</functionArgs>
		</expressions>
		<expressions>
			<variableName>mydate</variableName>
			<isFunction>false</isFunction>
			<constant>
				<isVarSelected>false</isVarSelected>
				<dateValue>
					<year>2019</year>
					<month>1</month>
					<day>11</day>
				</dateValue>
			</constant>
		</expressions>
		<expressions>
			<variableName>mycurr_dlr</variableName>
			<isFunction>false</isFunction>
			<constant>
				<isVarSelected>false</isVarSelected>
				<currencyValue>
					<value>599.99</value>
				</currencyValue>
			</constant>
		</expressions>
		<expressions>
			<variableName>mycurr_dlr</variableName>
			<isFunction>true</isFunction>
			<function>
				<returnType>CURRENCY</returnType>
				<name>SUM</name>
				<arguments>CURRENCY</arguments>
				<arguments>CURRENCY</arguments>
			</function>
			<functionArgs>
				<isVarSelected>true</isVarSelected>
				<variableName>mycurr_dlr</variableName>
			</functionArgs>
			<functionArgs>
				<isVarSelected>false</isVarSelected>
				<currencyValue>
					<value>3000</value>
				</currencyValue>
			</functionArgs>
		</expressions>
		<expressions>
			<variableName>mytime</variableName>
			<isFunction>false</isFunction>
			<constant>
				<isVarSelected>false</isVarSelected>
				<timeValue>
					<minutes>1267</minutes>
				</timeValue>
			</constant>
		</expressions>
		<expressions>
			<variableName>mydate</variableName>
			<isFunction>false</isFunction>
			<constant>
				<isVarSelected>true</isVarSelected>
				<variableName>__DATE__</variableName>
			</constant>
		</expressions>
	</data>
</setVariable>
`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel
	_, _ = decoder.Token()

	res := newSetVariablesModule(decoder)
	if res == nil {
		t.Fatal("nSetVariables module wasn't parsed...")
	}
	var mhu = (res.(xmlSetVariablesModule)).m
	costr, _ := ivr.NewStringValue("CONST STRING")
	mln, _ := ivr.NewNumericValue(1000000)
	nov, _ := ivr.NewDateValue(2019, 1, 11)
	fiveninenine, _ := ivr.NewUSCurrencyValue(599.99)
	onetwo, _ := ivr.NewTimeValue(1267)
	threethousand, _ := ivr.NewUSCurrencyValue(3000)

	var expected = ivr.SetVariableModule{
		Exprs: []*ivr.Expression{
			{
				Lval: "__BUFFER__",
				Rval: ivr.Assigner{
					P: &ivr.Parametrized{Value: costr},
				},
			},

			{
				Lval: "mystring",
				Rval: ivr.Assigner{
					P: &ivr.Parametrized{
						VariableName: "__BUFFER__",
					},
				},
			},

			{
				Lval: "myint",
				Rval: ivr.Assigner{
					P: &ivr.Parametrized{Value: mln},
				},
			},

			{
				Lval:   "__BUFFER__",
				IsFunc: true,
				Rval: ivr.Assigner{
					F: &ivr.IvrFuncInvocation{
						FuncDef: ivr.IvrFunc{
							Name:       "TOSTRING",
							ReturnType: "STRING",
							ArgTypes:   []string{"INTEGER"},
						},
						Params: []*ivr.Parametrized{{VariableName: "Contact.order_id"}},
					},
				},
			},

			{
				Lval: "mydate",
				Rval: ivr.Assigner{
					P: &ivr.Parametrized{Value: nov},
				},
			},

			{
				Lval: "mycurr_dlr",
				Rval: ivr.Assigner{
					P: &ivr.Parametrized{Value: fiveninenine},
				},
			},

			{
				Lval:   "mycurr_dlr",
				IsFunc: true,
				Rval: ivr.Assigner{
					F: &ivr.IvrFuncInvocation{
						FuncDef: ivr.IvrFunc{
							Name:       "SUM",
							ReturnType: "CURRENCY",
							ArgTypes:   []string{"CURRENCY", "CURRENCY"},
						},
						Params: []*ivr.Parametrized{
							{VariableName: "mycurr_dlr"},
							{Value: threethousand},
						},
					},
				},
			},
			{
				Lval: "mytime",
				Rval: ivr.Assigner{
					P: &ivr.Parametrized{Value: onetwo},
				},
			},

			{
				Lval: "mydate",
				Rval: ivr.Assigner{
					P: &ivr.Parametrized{VariableName: "__DATE__"},
				},
			},
		},
	}
	expected.SetGeneralInfo("SetVariable107", "05FDBF862B6F493AAAAFEC26E9FC8733",
		[]ivr.ModuleID{"7E06FB60DFAA4F3E89946625F8677CB3"}, "4065B9160FE54714AEFAB56B9F9CB9C6", "",
		"", "false")

	exp, err1 := json.MarshalIndent(expected, "", "  ")
	setv, err2 := json.MarshalIndent(mhu, "", "  ")

	if err1 != nil || err2 != nil || string(exp) != string(setv) {
		t.Errorf("\nSetVariables module: \n%s \nwas expected, in reality: \n%s", string(exp), string(setv))
	}

}
