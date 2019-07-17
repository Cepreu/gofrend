package xmlparser

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
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

	res := &ivr.IVRScript{
		Variables: make(ivr.Variables),
	}
	parseVars(res, decoder)

	if len(res.Variables) != 6 {
		t.Errorf("UserVariables weren't parsed...")
		return
	}

	//	exp := make(map[string]*variables)

	for _, vVar := range res.Variables {
		switch vVar.ValType {
		case ivr.ValTime:
			val, _ := ivr.NewTimeValue(60)
			vExp := ivr.NewVariable("var_time", "time var", vVar.ValType, val)
			if false == reflect.DeepEqual(vVar, vExp) {
				t.Errorf("\nTimeVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case ivr.ValDate:
			val, _ := ivr.NewDateValue(2019, 6, 7)
			vExp := ivr.NewVariable("var_date", "date var", vVar.ValType, val)

			if false == reflect.DeepEqual(vVar, vExp) {
				t.Errorf("\nDateVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case ivr.ValInteger:
			val, _ := ivr.NewIntegerValue(120)
			vExp := ivr.NewVariable("var_int", "integer var", vVar.ValType, val)
			if false == reflect.DeepEqual(vVar, vExp) {
				t.Errorf("\nIntegerVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case ivr.ValNumeric:
			val, _ := ivr.NewNumericValue(3.14)
			vExp := ivr.NewVariable("var_num", "numeric var", vVar.ValType, val)
			if false == reflect.DeepEqual(vVar, vExp) {
				t.Errorf("\nNumericVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case ivr.ValCurrency:
			val, _ := ivr.NewUSCurrencyValue(100.50)
			vExp := ivr.NewVariable("var_currency", "currency evro", vVar.ValType, val)
			if false == reflect.DeepEqual(vVar, vExp) {
				t.Errorf("\nCurrencyVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		case ivr.ValKVList:
			val, _ := ivr.NewKeyValue(`{"a":"sdfg","q":"werty"}`)
			vExp := ivr.NewVariable("var_kvlist", "KV list", vVar.ValType, val)
			if false == reflect.DeepEqual(vVar, vExp) {
				t.Errorf("\nKVListVariables: \n%v \nwas expected, in reality: \n%v", vExp, vVar)
			}
		default:
			t.Errorf("\nVariables: unknown type: %v", vVar)
		}
	}
}
