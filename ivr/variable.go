package ivr

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ValType int

const (
	ValUndefined ValType = iota
	ValString
	ValInteger
	ValID
	ValError
	ValNumeric
	ValDate
	ValTime
	ValCurrency
	ValCurrencyPound
	ValCurrencyEuro
	ValKVList
)

func (t ValType) String() string {
	var typeStr = [...]string{
		"Undefined",
		"String",
		"Integer",
		"ID",
		"ErrorCode",
		"Numeric",
		"Date",
		"Time",
		"Currency (dollar)",
		"Currency (pound)",
		"Currency (euro)",
		"List of key-value pairs",
	}
	if int(t) < len(typeStr) {
		return typeStr[t]
	}
	return "Unknown"
}

// VariableID - Auto-generated UUID of constant, or variable name for non-constant variable.
type VariableID string

// VarType - Is it a variable, constant, CAV, of a contact
type VarType int

const (
	VarConstant VarType = iota
	VarUserVariable
	VarContactField
	VarCallVariable
)

// Variable - ivr variable description
type Variable struct {
	ID          VariableID
	Value       string
	Description string
	ValType     ValType
	VarType     VarType
	Secured     bool
}

type KeyVal [2]string

type KVList []KeyVal

func (kvlist *KVList) ToString() string {
	data, err := json.Marshal(kvlist)
	if err != nil {
		panic("Error marshalling KVList: " + err.Error())
	}
	return string(data)
}

func StringToKVList(s string) *KVList {
	kvlist := new(KVList)
	err := json.Unmarshal([]byte(s), kvlist)
	if err != nil {
		panic("Error unmarshalling KVList: " + err.Error())
	}
	return kvlist
}

func (kvlist *KVList) Get(key string) string {
	for _, keyval := range *kvlist {
		if keyval[0] == key {
			return keyval[1]
		}
	}
	return ""
}

func (kvlist *KVList) GetKey(index int) string {
	return (*kvlist)[index][0]
}

func (kvlist *KVList) Put(key, value string) string {
	for _, keyval := range *kvlist {
		if keyval[0] == key {
			temp := keyval[1]
			keyval[1] = value
			return temp
		}
	}
	*kvlist = append(*kvlist, [2]string{key, value})
	return ""
}

func (kvlist *KVList) Remove(key string) string {
	for i, keyval := range *kvlist {
		if keyval[0] == key {
			*kvlist = append((*kvlist)[:i], (*kvlist)[i+1:]...)
			return keyval[1]
		}
	}
	return ""
}

func NewIntegerValue(intValue int) (string, error) {
	return strconv.Itoa(intValue), nil
}

func NewIDValue(intValue int) (string, error) {
	return strconv.Itoa(intValue), nil
}

func NewStringValue(strValue string) (string, error) {
	return strValue, nil
}

func NewNumericValue(numValue float64) (string, error) {
	return fmt.Sprintf("%f", numValue), nil
}

func NewUSCurrencyValue(currValue float64) (string, error) {
	return fmt.Sprintf("US$%.2f", currValue), nil
}

func NewEUCurrencyValue(currValue float64) (string, error) {
	return fmt.Sprintf("EU$%.2f", currValue), nil
}

func NewUKCurrencyValue(currValue float64) (string, error) {
	return fmt.Sprintf("UK$%.2f", currValue), nil
}

func NewDateValue(y int, m int, d int) (string, error) {
	if 1899 < y && y < 2101 && 0 < m && m < 13 && 0 < d && d < 31 {
		return fmt.Sprintf("%4d-%02d-%02d", y, m, d), nil
	}
	return "", fmt.Errorf("Incorrect date: %4d-%02d-%02d", y, m, d)
}

func NewTimeValue(minutes int) (string, error) {
	return fmt.Sprintf("%02d:%02d", minutes/60, minutes%60), nil
}

func NewKeyValue(kv string) (string, error) {
	// r, e := regexp.MatchString(`\{.+)\}`, kv)
	// if !r && e == nil {
	// 	return "nil", fmt.Errorf("Value %s is not of %s type", kv, ValKVList)
	// }
	// return kv, nil
	return "[]", nil
}

// NewVariable - Returns address of a new user variable, or <nil> if error
func NewVariable(name, descr string, t ValType, val string) *Variable {
	if name == "" {
		return nil
	}
	return &Variable{
		ID:          VariableID(name),
		Description: descr,
		VarType:     VarUserVariable,
		ValType:     t,
		Value:       val,
	}
}
