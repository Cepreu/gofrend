package xmlparser

import (
	"encoding/base64"
	"encoding/xml"
	"strconv"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
)

func parseVars(s *ivr.IVRScript, decoder *xml.Decoder) (err error) {
	var (
		immersion                                        = 1
		vs                                               = s.Variables
		userVar                                          *ivr.Variable
		val                                              string
		inEntry                                          = false
		inName, inDescription, inAttributes, inNullValue = false, false, false, false
		name, description                                string
		attrs                                            int
		nullVal                                          bool
		inDateValue, inYear, inMonth, inDay              = false, false, false, false
		y, m, d                                          int
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
		vtype                                            = ivr.ValUndefined
	)
	const (
		attrSystem         uint8 = 1
		attrCRM            uint8 = 2
		attrExternal             = 4
		attrInternal             = 8
		attrUserPredefined       = 16
		attrTTSEnumeration       = 32
		attrInput                = 64
		attrOutput               = 128
		attrUserDefined          = attrExternal | attrInternal | attrUserPredefined
		attrWritable             = attrCRM | attrUserDefined
		attrAny                  = attrSystem | attrCRM | attrUserDefined
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
				d, _ = strconv.Atoi(string(v))
			} else if inDateValue && inMonth {
				m, _ = strconv.Atoi(string(v))
			} else if inDateValue && inYear {
				y, _ = strconv.Atoi(string(v))

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
				if nullVal {
					val = ""
				}
				userVar = ivr.NewVariable(name, description, vtype, val)
				vs[ivr.VariableID(name)] = userVar
				if attrs&attrInput == attrInput {
					s.Input = append(s.Input, userVar.ID)
				}
				if attrs&attrOutput == attrOutput {
					s.Output = append(s.Output, userVar.ID)
				}

				name, description = "", ""
				attrs = 0
				nullVal = false
				val = ""
				vtype = ivr.ValUndefined

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
				val, _ = ivr.NewDateValue(y, m, d)
				m, d, y = 0, 0, 0
				vtype = ivr.ValDate
			} else if v.Name.Local == "day" {
				inDay = false
			} else if v.Name.Local == "month" {
				inMonth = false
			} else if v.Name.Local == "year" {
				inYear = false

			} else if v.Name.Local == "timeValue" {
				inTimeValue = false
				val, _ = ivr.NewTimeValue(minutes)
				minutes = 0
				vtype = ivr.ValTime
			} else if v.Name.Local == "minutes" {
				inMinutes = false

			} else if v.Name.Local == "numericValue" {
				inNumericValue = false
				val, _ = ivr.NewNumericValue(numValue)
				numValue = 0
				vtype = ivr.ValNumeric

			} else if v.Name.Local == "integerValue" {
				inIntegerValue = false
				val, _ = ivr.NewIntegerValue(intValue)
				intValue = 0
				vtype = ivr.ValInteger

			} else if v.Name.Local == "stringValue" {
				inStringValue = false
				val, _ = ivr.NewStringValue(strValue)
				strValue = ""
				vtype = ivr.ValString

			} else if v.Name.Local == "id" {
				inID = false
				val, _ = ivr.NewIDValue(id)
				id = 0
				vtype = ivr.ValID

			} else if v.Name.Local == "currencyValue" {
				inCurrencyValue = false
				val, _ = ivr.NewUSCurrencyValue(currValue)
				currValue = 0.0
				vtype = ivr.ValCurrency

			} else if v.Name.Local == "kvListValue" {
				sDec, _ := base64.StdEncoding.DecodeString(strValue)
				val, _ = ivr.NewKeyValue(string(sDec))
				inKVListValue = false
				strValue = ""
				vtype = ivr.ValKVList
			}
		}
	}
	return err
}

func addConstantVar(s *ivr.IVRScript, t ivr.ValType, v string) ivr.VariableID {
	for oldID, oldV := range s.Variables {
		if oldV.VarType == ivr.VarConstant &&
			oldV.ValType == t &&
			oldV.Value == v {
			return oldID
		}
	}
	newVariable := &ivr.Variable{
		ID:      ivr.VariableID(utils.GenUUIDv4()),
		VarType: ivr.VarConstant,
		ValType: t,
		Value:   v,
	}
	s.Variables[newVariable.ID] = newVariable
	return newVariable.ID
}

func addIntegerConstant(s *ivr.IVRScript, intValue int) (ivr.VariableID, error) {
	v, err := ivr.NewIntegerValue(intValue)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValInteger, v), nil
}

func addIDConstant(s *ivr.IVRScript, intValue int) (ivr.VariableID, error) {
	v, err := ivr.NewIDValue(intValue)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValID, v), nil
}

func addStringConstant(s *ivr.IVRScript, strValue string) (ivr.VariableID, error) {
	v, err := ivr.NewStringValue(strValue)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValString, v), nil
}

func addNumericConstant(s *ivr.IVRScript, numValue float64) (ivr.VariableID, error) {
	v, err := ivr.NewNumericValue(numValue)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValNumeric, v), nil
}

func addUSCurrencyConstant(s *ivr.IVRScript, currValue float64) (ivr.VariableID, error) {
	v, err := ivr.NewUSCurrencyValue(currValue)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValCurrency, v), nil
}

func addEUCurrencyConstant(s *ivr.IVRScript, currValue float64) (ivr.VariableID, error) {
	v, err := ivr.NewEUCurrencyValue(currValue)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValCurrencyEuro, v), nil
}

func addUKCurrencyConstant(s *ivr.IVRScript, currValue float64) (ivr.VariableID, error) {
	v, err := ivr.NewUKCurrencyValue(currValue)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValCurrencyPound, v), nil
}

func addDateConstant(s *ivr.IVRScript, y int, m int, d int) (ivr.VariableID, error) {
	v, err := ivr.NewDateValue(y, m, d)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValDate, v), nil
}

func addTimeConstant(s *ivr.IVRScript, minutes int) (ivr.VariableID, error) {
	v, err := ivr.NewTimeValue(minutes)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValTime, v), nil
}

func addSKeyValueConstant(s *ivr.IVRScript, kv string) (ivr.VariableID, error) {
	v, err := ivr.NewKeyValue(kv)
	if err != nil {
		return "", err
	}
	return addConstantVar(s, ivr.ValKVList, v), nil
}
