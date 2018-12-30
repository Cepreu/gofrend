package ivrparser

import (
	"encoding/xml"
	"strconv"
)

type parametrized struct {
	DataType    string
	IntValue    int
	IntID       string
	StringValue string
	StringID    string

	VariableName  string
	IsVarSelected bool
}

func (p *parametrized) parse(decoder *xml.Decoder) (err error) {
	var (
		immersion                 = 1
		inValue, inID, inVariable = false, false, false
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
				innerText, err := decoder.Token()
				if err == nil {
					p.IsVarSelected = string(innerText.(xml.CharData)) == "true"
				}
			} else if v.Name.Local == "integerValue" {
				p.DataType = "INTEGER"
			} else if v.Name.Local == "stringValue" {
				p.DataType = "STRING"
			} else if v.Name.Local == "value" {
				inValue = true
			} else if v.Name.Local == "id" {
				inID = true
			} else if v.Name.Local == "variableName" {
				inVariable = true
			}
		case xml.CharData:
			if inValue && p.DataType == "STRING" {
				p.StringValue = string(v)
			} else if inValue && p.DataType == "INTEGER" {
				p.IntValue, err = strconv.Atoi(string(v))
			} else if inID && p.DataType == "STRING" {
				p.StringID = string(v)
			} else if inID && p.DataType == "INTEGER" {
				p.IntID = string(v)
			} else if inVariable {
				p.VariableName = string(v)
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "value" {
				inValue = false
			} else if v.Name.Local == "id" {
				inID = false
			} else if v.Name.Local == "variableName" {
				inVariable = false
			}
		}
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
