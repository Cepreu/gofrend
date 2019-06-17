package ivrparser

import (
	"encoding/xml"

	"github.com/Cepreu/gofrend/vars"
)

type parametrized struct {
	VariableName string
	Value        vars.Value
}

func (p *parametrized) IsVarSelected() bool { return p.VariableName != "" }

func (p *parametrized) parse(decoder *xml.Decoder) (err error) {
	var (
		immersion                                  = 1
		inValue, inID, inVariable, inIsVarSelected = false, false, false, false
		inYear, inMonth, inDay, inMinutes          = false, false, false, false
		IsVarSelected                              = false

		val vars.Value
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
				val = new(vars.Integer)
			} else if v.Name.Local == "currencyValue" {
				val = new(vars.Currency)
			} else if v.Name.Local == "numericValue" {
				val = new(vars.Numeric)
			} else if v.Name.Local == "stringValue" {
				val = new(vars.String)
			} else if v.Name.Local == "dateValue" {
				val = new(vars.Date)
			} else if v.Name.Local == "timeValue" {
				val = new(vars.Time)
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
				val.SetValue("value", string(v))
			} else if inID {
				val.SetValue("id", string(v))
			} else if inYear {
				val.SetValue("year", string(v))
			} else if inMonth {
				val.SetValue("month", string(v))
			} else if inDay {
				val.SetValue("day", string(v))
			} else if inMinutes {
				val.SetValue("minutes", string(v))
			} else if inVariable && IsVarSelected {
				p.VariableName = string(v)
			} else if inIsVarSelected {
				IsVarSelected = string(v) == "true"
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "isVarSelected" {
				inIsVarSelected = false
			} else if v.Name.Local == "value" {
				inValue = false
			} else if v.Name.Local == "id" {
				inID = false
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
	if !IsVarSelected {
		p.Value = val
	}
	return err
}

///////////////////
type keyValueParametrized struct {
	Key   string
	Value *parametrized
}

func parseKeyValueListParmetrized(decoder *xml.Decoder) (p []keyValueParametrized, err error) {
	var (
		immersion      = 1
		inEntry, inKey = false, false
		key            string
		value          *parametrized
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
				value = new(parametrized)
				value.parse(decoder)
				immersion--
			}

		case xml.CharData:
			if inKey {
				key = string(v)
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "entry" {
				inEntry = false
				p = append(p, keyValueParametrized{key, value})
			} else if v.Name.Local == "key" && inEntry {
				inKey = false
			}
		}
	}

	return p, nil
}
