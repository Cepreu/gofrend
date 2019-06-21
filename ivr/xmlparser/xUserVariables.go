package xmlparser

import (
	"encoding/base64"
	"encoding/xml"
	"strconv"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/ivr/vars"
)

func parseVars(vs ivr.Variables, decoder *xml.Decoder) (err error) {
	var (
		immersion                                        = 1
		userVar                                          *vars.Variable
		val                                              vars.Value
		inEntry                                          = false
		inName, inDescription, inAttributes, inNullValue = false, false, false, false
		name, description                                string
		attrs                                            int
		nullVal                                          bool
		inDateValue, inYear, inMonth, inDay              = false, false, false, false
		date                                             struct{ y, m, d int }
		inTimeValue, inMinutes                           = false, false
		minutes                                          int
		inValue                                          = false
		inNumericValue                                   = false
		numValue                                         float64
		inIntegerValue                                   = false
		intValue                                         int
		inStringValue, inID                              = false, false
		strValue                                         string
		id                                               int
		inCurrencyValue                                  = false
		currValue                                        float64
		inKVListValue                                    = false
	)
	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			return err
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "entry" {
				inEntry = true
			} else if v.Name.Local == "name" && inEntry {
				inName = true
			} else if v.Name.Local == "description" && inEntry {
				inDescription = true
			} else if v.Name.Local == "attributes" && inEntry {
				inAttributes = true
			} else if v.Name.Local == "isNullValue" && inEntry {
				inNullValue = true
			} else if v.Name.Local == "value" && inEntry {
				inValue = true

			} else if v.Name.Local == "dateValue" {
				inDateValue = true
			} else if v.Name.Local == "day" {
				inDay = true
			} else if v.Name.Local == "month" {
				inMonth = true
			} else if v.Name.Local == "year" {
				inYear = true

			} else if v.Name.Local == "timeValue" {
				inTimeValue = true
			} else if v.Name.Local == "minutes" {
				inMinutes = true

			} else if v.Name.Local == "numericValue" {
				inNumericValue = true
			} else if v.Name.Local == "integerValue" {
				inIntegerValue = true
			} else if v.Name.Local == "stringValue" {
				inStringValue = true
			} else if v.Name.Local == "id" {
				inID = true
			} else if v.Name.Local == "currencyValue" {
				inCurrencyValue = true
			} else if v.Name.Local == "kvListValue" {
				inKVListValue = true
			}

		case xml.CharData:
			if inName {
				name = string(v)
			} else if inDescription {
				description = string(v)
			} else if inAttributes {
				attrs, _ = strconv.Atoi(string(v))
			} else if inNullValue {
				nullVal = string(v) == "true"

			} else if inDateValue && inDay {
				date.d, _ = strconv.Atoi(string(v))
			} else if inDateValue && inMonth {
				date.m, _ = strconv.Atoi(string(v))
			} else if inDateValue && inYear {
				date.y, _ = strconv.Atoi(string(v))

			} else if inTimeValue && inMinutes {
				minutes, _ = strconv.Atoi(string(v))

			} else if inNumericValue && inValue {
				numValue, _ = strconv.ParseFloat(string(v), 32)
			} else if inIntegerValue && inValue {
				intValue, _ = strconv.Atoi(string(v))
			} else if inCurrencyValue && inValue {
				currValue, _ = strconv.ParseFloat(string(v), 32)
			} else if inStringValue && inValue {
				strValue = string(v)
			} else if inKVListValue && inValue {
				strValue = string(v)
			} else if inStringValue && inID {
				id, _ = strconv.Atoi(string(v))
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "entry" {
				inEntry = false

				userVar = vars.NewVariable(name, description, attrs, nullVal)
				if val != nil {
					userVar.SetValue(val)
				}
				vs[name] = userVar

				name, description = "", ""
				attrs = 0
				nullVal = false
				val = nil
			} else if v.Name.Local == "name" && inEntry {
				inName = false
			} else if v.Name.Local == "description" && inEntry {
				inDescription = false
			} else if v.Name.Local == "attributes" && inEntry {
				inAttributes = false
			} else if v.Name.Local == "isNullValue" && inEntry {
				inNullValue = false
			} else if v.Name.Local == "value" && inEntry {
				inValue = false

			} else if v.Name.Local == "dateValue" {
				inDateValue = false
				val = vars.NewDate(date.y, date.m, date.d)
				date.m, date.d, date.y = 0, 0, 0
			} else if v.Name.Local == "day" {
				inDay = false
			} else if v.Name.Local == "month" {
				inMonth = false
			} else if v.Name.Local == "year" {
				inYear = false

			} else if v.Name.Local == "timeValue" {
				inTimeValue = false
				val = vars.NewTime(minutes)
				minutes = 0
			} else if v.Name.Local == "minutes" {
				inMinutes = false

			} else if v.Name.Local == "numericValue" {
				inNumericValue = false
				val = vars.NewNumeric(numValue)
				numValue = 0

			} else if v.Name.Local == "integerValue" {
				inIntegerValue = false
				val = vars.NewInteger(intValue)
				intValue = 0

			} else if v.Name.Local == "stringValue" {
				inStringValue = false
				val = vars.NewString(strValue, id)
				strValue = ""
				id = 0
			} else if v.Name.Local == "id" {
				inID = false

			} else if v.Name.Local == "currencyValue" {
				inCurrencyValue = false
				val = vars.NewCurrency(currValue)
				currValue = 0.0

			} else if v.Name.Local == "kvListValue" {
				sDec, _ := base64.StdEncoding.DecodeString(strValue)
				val = vars.NewKVList(string(sDec))

				inKVListValue = false
				strValue = ""
			}
		}
	}
	return err
}
