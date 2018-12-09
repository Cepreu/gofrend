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

func (p *parametrized) parse(decoder *xml.Decoder, v *xml.StartElement) (err error) {
	var (
		lastElement               = v.Name.Local
		inValue, inID, inVariable = false, false, false
	)
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			return err
		}

		switch v := t.(type) {
		case xml.StartElement:
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
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			} else if v.Name.Local == "value" {
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
