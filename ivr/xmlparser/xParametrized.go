package xmlparser

import (
	"encoding/xml"
	"strconv"

	"github.com/Cepreu/gofrend/ivr"
)

type parametrized struct {
	variableName string
	vType        ivr.ValType
	value        string
}

func (p *parametrized) isVarSelected() bool { return p.variableName != "" }

func isVarSelected(p *parametrized) bool {
	return p.variableName != ""
}

func parse(p *parametrized, decoder *xml.Decoder) (err error) {
	var (
		immersion                                                  = 1
		inValue, inID, inVariable, inIsVarSelected                 = false, false, false, false
		inYear, inMonth, inDay, inMinutes                          = false, false, false, false
		IsVarSelected                                              = false
		inIValue, inCValue, inNValue, inSValue, inDValue, inTValue = false, false, false, false, false, false
		numericVal                                                 float64
		integerVal                                                 int
		stringVal                                                  string
		day, month, year                                           int
	)
	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			return err
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "isVarSelected" {
				inIsVarSelected = true
			} else if v.Name.Local == "integerValue" {
				inIValue = true
			} else if v.Name.Local == "currencyValue" {
				inCValue = true
			} else if v.Name.Local == "numericValue" {
				inNValue = true
			} else if v.Name.Local == "stringValue" {
				inSValue = true
			} else if v.Name.Local == "dateValue" {
				inDValue = true
			} else if v.Name.Local == "timeValue" {
				inTValue = true
			} else if v.Name.Local == "value" {
				inValue = true
			} else if v.Name.Local == "id" {
				inID = true
			} else if v.Name.Local == "year" {
				inYear = true
			} else if v.Name.Local == "month" {
				inMonth = true
			} else if v.Name.Local == "day" {
				inDay = true
			} else if v.Name.Local == "minutes" {
				inMinutes = true
			} else if v.Name.Local == "variableName" {
				inVariable = true
			}
		case xml.CharData:
			if inValue {
				if inIValue {
					integerVal, err = strconv.Atoi(string(v))
				} else if inCValue || inNValue {
					numericVal, err = strconv.ParseFloat(string(v), 64)
				} else if inSValue {
					stringVal = string(v)
				}
			} else if inTValue && inMinutes {
				integerVal, err = strconv.Atoi(string(v))
			} else if inID {
				integerVal, err = strconv.Atoi(string(v))
			} else if inYear && inDValue {
				year, err = strconv.Atoi(string(v))
			} else if inMonth && inDValue {
				month, err = strconv.Atoi(string(v))
			} else if inDay {
				day, err = strconv.Atoi(string(v))
			} else if inVariable && IsVarSelected {
				p.variableName = string(v)
			} else if inIsVarSelected {
				IsVarSelected = string(v) == "true"
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "isVarSelected" {
				inIsVarSelected = false
			} else if v.Name.Local == "integerValue" {
				p.value, err = ivr.NewIntegerValue(integerVal)
				p.vType = ivr.ValInteger
				integerVal = 0
				inIValue = false
			} else if v.Name.Local == "currencyValue" {
				p.value, err = ivr.NewUSCurrencyValue(numericVal)
				p.vType = ivr.ValCurrency
				numericVal = 0
				inCValue = false
			} else if v.Name.Local == "numericValue" {
				p.value, err = ivr.NewNumericValue(numericVal)
				p.vType = ivr.ValNumeric
				numericVal = 0
				inNValue = false
			} else if v.Name.Local == "stringValue" {
				p.value, err = ivr.NewStringValue(stringVal)
				p.vType = ivr.ValString
				stringVal = ""
				inSValue = false
			} else if v.Name.Local == "dateValue" {
				p.value, err = ivr.NewDateValue(year, month, day)
				p.vType = ivr.ValDate
				year, month, day = 0, 0, 0
				inDValue = false
			} else if v.Name.Local == "timeValue" {
				p.value, err = ivr.NewTimeValue(integerVal)
				p.vType = ivr.ValTime
				integerVal = 0
				inTValue = false
			} else if v.Name.Local == "id" {
				p.value, err = ivr.NewIDValue(integerVal)
				p.vType = ivr.ValID
				integerVal = 0
				inID = false
			} else if v.Name.Local == "value" {
				inValue = false
			} else if v.Name.Local == "year" {
				inYear = false
			} else if v.Name.Local == "month" {
				inMonth = false
			} else if v.Name.Local == "day" {
				inDay = false
			} else if v.Name.Local == "minutes" {
				inMinutes = false
			} else if v.Name.Local == "variableName" {
				inVariable = false
			}
		}
	}
	return err
}

func parseKeyValueListParmetrized(decoder *xml.Decoder, script *ivr.IVRScript) (p []ivr.KeyValue, err error) {
	var (
		immersion      = 1
		inEntry, inKey = false, false
		key            string
		value          ivr.VariableID
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "entry" {
				inEntry = true
			} else if v.Name.Local == "key" && inEntry {
				inKey = true
			} else if v.Name.Local == "value" && inEntry {
				pv := new(parametrized)
				parse(pv, decoder)
				immersion--
				value = toID(script, pv)
			}

		case xml.CharData:
			if inKey {
				key = string(v)
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "entry" {
				inEntry = false
				p = append(p, ivr.KeyValue{Key: key, Value: value})
			} else if v.Name.Local == "key" && inEntry {
				inKey = false
			}
		}
	}

	return p, nil
}

func toID(s *ivr.IVRScript, pv *parametrized) ivr.VariableID {
	if isVarSelected(pv) {
		return ivr.VariableID(pv.variableName)
	}
	return addConstantVar(s, pv.vType, pv.value)
}
