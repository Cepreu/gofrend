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

func TestUserVariables(t *testing.T) {
	var xmlData = `
<userVariables>
<entry>
	<key>var_time</key>
	<value>
		<name>var_time</name>
		<description>time var</description>
		<timeValue>
			<minutes>60</minutes>
		</timeValue>
		<attributes>64</attributes>
		<isNullValue>false</isNullValue>
	</value>
</entry>
<entry>
	<key>var_num</key>
	<value>
		<name>var_num</name>
		<description>numeric var</description>
		<numericValue>
			<value>3.1415</value>
		</numericValue>
		<attributes>128</attributes>
		<isNullValue>false</isNullValue>
	</value>
</entry>
<entry>
	<key>var_currency</key>
	<value>
		<name>var_currency</name>
		<description>currency evro</description>
		<currencyValue>
			<value>100.50</value>
		</currencyValue>
		<attributes>8</attributes>
		<isNullValue>false</isNullValue>
	</value>
</entry>
<entry>
	<key>var_kvlist</key>
	<value>
		<name>var_kvlist</name>
		<description>KV list</description>
		<kvListValue>
			<value>eyJxIjoid2VydHkiLCJhIjoic2RmZyJ9</value>
		</kvListValue>
		<attributes>192</attributes>
		<isNullValue>false</isNullValue>
	</value>
</entry>
<entry>
	<key>var_int</key>
	<value>
		<name>var_int</name>
		<description>integer var</description>
		<integerValue>
			<value>120</value>
		</integerValue>
		<attributes>8</attributes>
		<isNullValue>false</isNullValue>
	</value>
</entry>
<entry>
	<key>var_date</key>
	<value>
		<name>var_date</name>
		<description>date var</description>
		<dateValue>
			<year>2019</year>
			<month>6</month>
			<day>7</day>
		</dateValue>
		<attributes>8</attributes>
		<isNullValue>false</isNullValue>
	</value>
</entry>
</userVariables>
`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	res := make(ivr.Variables)
	parseVars(res, decoder)

	if len(res) != 6 {
		t.Errorf("ForeignScript module wasn't parsed...")
		return
	}

	//	exp := make(map[string]*variables)

	for _, vVar := range res {
		switch vVal := vVar.GetValue().(type) {
		case *vars.Time:
			vExp := vars.NewVariable("var_time", "time var", 0, false)
			vExp.SetValue(vars.NewTime(60))
			if false == reflect.DeepEqual(vVar.GetValue(), vExp.GetValue()) {
				t.Errorf("\nTimeVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case *vars.Date:
			vExp := vars.NewVariable("var_date", "date var", 0, false)
			vExp.SetValue(vars.NewDate(2019, 6, 7))
			if false == reflect.DeepEqual(vVar.GetValue(), vExp.GetValue()) {
				t.Errorf("\nDateVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case *vars.Integer:
			vExp := vars.NewVariable("var_int", "integer var", 0, false)
			vExp.SetValue(vars.NewInteger(120))
			if vVal.String() != vExp.GetValue().String() {
				t.Errorf("\nIntegerVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case *vars.Numeric:
			vExp := vars.NewVariable("var_num", "numeric var", 128, false)
			vExp.SetValue(vars.NewNumeric(3.14))
			if vVal.String() != vExp.GetValue().String() {
				t.Errorf("\nNumericVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case *vars.Currency:
			vExp := vars.NewVariable("var_currency", "currency evro", 128, false)
			vExp.SetValue(vars.NewCurrency(100.50))
			if vVal.String() != vExp.GetValue().String() {
				t.Errorf("\nCurrencyVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case *vars.KVList:
			vExp := vars.NewVariable("var_kvlist", "KV list", 128, false)
			vExp.SetValue(vars.NewKVList(`{"a":"sdfg","q":"werty"}`))
			if vVal.String() != vExp.GetValue().String() {
				t.Errorf("\nKVListVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		default:
			vExp := vars.NewVariable("var_int", "integer var", 0, false)
			vExp.SetValue(vars.NewInteger(120))
			if false == reflect.DeepEqual(vVar, vExp) ||
				vVal.String() != vExp.GetValue().String() {
				t.Errorf("\nVariables: \n%v \nwas expected, in reality: \n%v", reflect.TypeOf(vExp.GetValue()), reflect.TypeOf(vVar.GetValue()))
				t.Errorf("\nVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		}
	}

	// more sanity checking...
}
