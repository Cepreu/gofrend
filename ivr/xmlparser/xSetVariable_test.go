package xmlparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
	"golang.org/x/net/html/charset"
)

func TestPizzaVariables(t *testing.T) {
	var xmlData = `<setVariable>
            <ascendants>0D646BA00EA84E279ECE93C532DFC200</ascendants>
            <singleDescendant>E86971EB3DBA4993B376DC71F2F66BE9</singleDescendant>
            <moduleName>add</moduleName>
            <locationX>522</locationX>
            <locationY>362</locationY>
            <moduleId>DC04339BF98A49BD9046AEA568C78A7A</moduleId>
            <data>
                <expressions>
                    <variableName>__BUFFER__</variableName>
                    <isFunction>true</isFunction>
                    <function>
                        <returnType>STRING</returnType>
                        <name>PUT</name>
                        <arguments>KVLIST</arguments>
                        <arguments>STRING</arguments>
                        <arguments>STRING</arguments>
                    </function>
                    <functionArgs>
                        <isVarSelected>true</isVarSelected>
                        <variableName>toppings</variableName>
                    </functionArgs>
                    <functionArgs>
                        <isVarSelected>true</isVarSelected>
                        <stringValue>
                            <value></value>
                            <id>0</id>
                        </stringValue>
                        <variableName>__BUFFER__</variableName>
                    </functionArgs>
                    <functionArgs>
                        <isVarSelected>false</isVarSelected>
                        <stringValue>
                            <value>1</value>
                            <id>0</id>
                        </stringValue>
                    </functionArgs>
                </expressions>
                <expressions>
                    <variableName>NumOfToppings</variableName>
                    <isFunction>true</isFunction>
                    <function>
                        <returnType>INTEGER</returnType>
                        <name>SIZE</name>
                        <arguments>KVLIST</arguments>
                    </function>
                    <functionArgs>
                        <isVarSelected>true</isVarSelected>
                        <variableName>toppings</variableName>
                    </functionArgs>
                </expressions>
            </data>
        </setVariable>
	`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel
	_, _ = decoder.Token()

	s := &ivr.IVRScript{
		Variables: make(ivr.Variables),
	}
	res := newSetVariablesModule(decoder, s)
	if res == nil {
		t.Fatal("nSetVariables module wasn't parsed...")
	}
	var mhu = (res.(xmlSetVariablesModule)).m
	utils.PrettyPrint(mhu)
}

func TestPizzaVariables2(t *testing.T) {
	var xmlData = `<setVariable>
	<ascendants>AC1FBAC578A744FE9B610EA20C1328A9</ascendants>
	<ascendants>3616BB19FDD0498FB671BC3ECCDD1EA0</ascendants>
	<ascendants>DC04339BF98A49BD9046AEA568C78A7A</ascendants>
	<singleDescendant>AC1FBAC578A744FE9B610EA20C1328A9</singleDescendant>
	<moduleName>SetVariable33</moduleName>
	<locationX>617</locationX>
	<locationY>319</locationY>
	<moduleId>E86971EB3DBA4993B376DC71F2F66BE9</moduleId>
	<data>
		<expressions>
			<variableName>__BUFFER__</variableName>
			<isFunction>true</isFunction>
			<function>
				<returnType>STRING</returnType>
				<name>GET_KEY</name>
				<arguments>KVLIST</arguments>
				<arguments>INTEGER</arguments>
			</function>
			<functionArgs>
				<isVarSelected>true</isVarSelected>
				<variableName>toppings</variableName>
			</functionArgs>
			<functionArgs>
				<isVarSelected>true</isVarSelected>
				<variableName>i</variableName>
			</functionArgs>
		</expressions>
		<expressions>
			<variableName>topping_list</variableName>
			<isFunction>true</isFunction>
			<function>
				<returnType>STRING</returnType>
				<name>CONCAT</name>
				<arguments>STRING</arguments>
				<arguments>STRING</arguments>
			</function>
			<functionArgs>
				<isVarSelected>true</isVarSelected>
				<stringValue>
					<value></value>
					<id>0</id>
				</stringValue>
				<variableName>topping_list</variableName>
			</functionArgs>
			<functionArgs>
				<isVarSelected>true</isVarSelected>
				<stringValue>
					<value></value>
					<id>0</id>
				</stringValue>
				<variableName>__BUFFER__</variableName>
			</functionArgs>
		</expressions>
		<expressions>
			<variableName>topping_list</variableName>
			<isFunction>true</isFunction>
			<function>
				<returnType>STRING</returnType>
				<name>CONCAT</name>
				<arguments>STRING</arguments>
				<arguments>STRING</arguments>
			</function>
			<functionArgs>
				<isVarSelected>true</isVarSelected>
				<stringValue>
					<value></value>
					<id>0</id>
				</stringValue>
				<variableName>topping_list</variableName>
			</functionArgs>
			<functionArgs>
				<isVarSelected>false</isVarSelected>
				<stringValue>
					<value>","</value>
					<id>0</id>
				</stringValue>
			</functionArgs>
		</expressions>
		<expressions>
			<variableName>i</variableName>
			<isFunction>true</isFunction>
			<function>
				<returnType>INTEGER</returnType>
				<name>SUM</name>
				<arguments>INTEGER</arguments>
				<arguments>INTEGER</arguments>
			</function>
			<functionArgs>
				<isVarSelected>true</isVarSelected>
				<variableName>i</variableName>
			</functionArgs>
			<functionArgs>
				<isVarSelected>false</isVarSelected>
				<integerValue>
					<value>1</value>
				</integerValue>
			</functionArgs>
		</expressions>
	</data>
</setVariable>`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel
	_, _ = decoder.Token()

	s := &ivr.IVRScript{
		Variables: make(ivr.Variables),
	}
	res := newSetVariablesModule(decoder, s)
	if res == nil {
		t.Fatal("nSetVariables module wasn't parsed...")
	}
	var mhu = (res.(xmlSetVariablesModule)).m
	utils.PrettyPrint(mhu)
}

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

	s := &ivr.IVRScript{
		Variables: make(ivr.Variables),
	}
	res := newSetVariablesModule(decoder, s)
	if res == nil {
		t.Fatal("nSetVariables module wasn't parsed...")
	}
	var mhu = (res.(xmlSetVariablesModule)).m
	utils.PrettyPrint(mhu)
	costr, _ := addStringConstant(s, "CONST STRING")
	mln, _ := addNumericConstant(s, 1000000)
	nov, _ := addDateConstant(s, 2019, 1, 11)
	fiveninenine, _ := addUSCurrencyConstant(s, 599.99)
	onetwo, _ := addTimeConstant(s, 1267)
	threethousand, _ := addUSCurrencyConstant(s, 3000)

	var expected = ivr.SetVariableModule{
		Exprs: []*ivr.Expression{
			{
				Lval: "__BUFFER__",
				Rval: ivr.FuncInvocation{FuncDef: "__COPY__", Params: []ivr.VariableID{costr}},
			},
			{
				Lval: "mystring",
				Rval: ivr.FuncInvocation{FuncDef: "__COPY__", Params: []ivr.VariableID{"__BUFFER__"}},
			},
			{
				Lval: "myint",
				Rval: ivr.FuncInvocation{FuncDef: "__COPY__", Params: []ivr.VariableID{mln}},
			},
			{
				Lval: "__BUFFER__",
				Rval: ivr.FuncInvocation{FuncDef: "STRING#TOSTRING#INTEGER",
					Params: []ivr.VariableID{"Contact.order_id"}},
			},
			{
				Lval: "mydate",
				Rval: ivr.FuncInvocation{FuncDef: "__COPY__", Params: []ivr.VariableID{nov}},
			},
			{
				Lval: "mycurr_dlr",
				Rval: ivr.FuncInvocation{FuncDef: "__COPY__", Params: []ivr.VariableID{fiveninenine}},
			},
			{
				Lval: "mycurr_dlr",
				Rval: ivr.FuncInvocation{FuncDef: "CURRENCY#SUM#CURRENCY#CURRENCY",
					Params: []ivr.VariableID{"mycurr_dlr", threethousand}},
			},
			{
				Lval: "mytime",
				Rval: ivr.FuncInvocation{FuncDef: "__COPY__", Params: []ivr.VariableID{onetwo}},
			},
			{
				Lval: "mydate",
				Rval: ivr.FuncInvocation{FuncDef: "__COPY__", Params: []ivr.VariableID{"__DATE__"}},
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
