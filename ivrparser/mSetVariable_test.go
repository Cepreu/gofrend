package ivrparser

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/vars"
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
	var mhu = res.(*SetVariableModule)
	var expected = SetVariableModule{
		GeneralInfo: GeneralInfo{
			ID:              "05FDBF862B6F493AAAAFEC26E9FC8733",
			Ascendants:      []ModuleID{"7E06FB60DFAA4F3E89946625F8677CB3"},
			Descendant:      "4065B9160FE54714AEFAB56B9F9CB9C6",
			ExceptionalDesc: "",
			Name:            "SetVariable107",
			Dispo:           "",
			Collapsible:     false,
		},
		Exprs: []*expression{
			{
				Lval: "__BUFFER__",
				Rval: assigner{
					P: &parametrized{Value: vars.NewString("CONST STRING", 0)},
				},
			},

			{
				Lval: "mystring",
				Rval: assigner{
					P: &parametrized{
						VariableName: "__BUFFER__",
					},
				},
			},

			{
				Lval: "myint",
				Rval: assigner{
					P: &parametrized{Value: vars.NewNumeric(1000000)},
				},
			},

			{
				Lval:   "__BUFFER__",
				IsFunc: true,
				Rval: assigner{
					F: &ivrFuncInvocation{
						FuncDef: ivrFunc{
							Name:       "TOSTRING",
							ReturnType: "STRING",
							ArgTypes:   []string{"INTEGER"},
						},
						Params: []*parametrized{{VariableName: "Contact.order_id"}},
					},
				},
			},

			{
				Lval: "mydate",
				Rval: assigner{
					P: &parametrized{Value: vars.NewDate(2019, 1, 11)},
				},
			},

			{
				Lval: "mycurr_dlr",
				Rval: assigner{
					P: &parametrized{Value: vars.NewCurrency(599.99)},
				},
			},

			{
				Lval:   "mycurr_dlr",
				IsFunc: true,
				Rval: assigner{
					F: &ivrFuncInvocation{
						FuncDef: ivrFunc{
							Name:       "SUM",
							ReturnType: "CURRENCY",
							ArgTypes:   []string{"CURRENCY", "CURRENCY"},
						},
						Params: []*parametrized{
							{VariableName: "mycurr_dlr"},
							{Value: vars.NewCurrency(3000)},
						},
					},
				},
			},
			{
				Lval: "mytime",
				Rval: assigner{
					P: &parametrized{Value: vars.NewTime(1267)},
				},
			},

			{
				Lval: "mydate",
				Rval: assigner{
					P: &parametrized{VariableName: "__DATE__"},
				},
			},
		},
	}

	if false == reflect.DeepEqual(&expected, mhu) {
		t.Errorf("\nSetVariables module: \n%v \nwas expected, in reality: \n%v", expected, mhu)
	}

}
