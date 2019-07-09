package ivr

import (
	"fmt"
	"regexp"

	"github.com/Cepreu/gofrend/utils"
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
		"String",
		"Integer",
		"ID",
		"ErrorCode",
		"Numeric",
		"Currency (dollar)",
		"Date",
		"Time",
		"Currency (pound)",
		"Currency (euro)",
		"List of key-value pairs",
	}
	if int(t) < len(typeStr) {
		return typeStr[t]
	}
	return "Unknown"
}

// Value - Represents Variable's type and default value
type Value struct {
	VType       ValType //KSA: To be removed
	StringValue string
}

// VariableID - Auto-generated UUID of a variable or a constant
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
	Name        string
	Value       *Value
	Description string
	ValType     ValType
	VarType     VarType
	Secured     bool
}

func NewIntegerValue(intValue int) (*Value, error) {
	return &Value{ValInteger, fmt.Sprintf("%d", intValue)}, nil
}

func NewIDValue(intValue int) (*Value, error) {
	return &Value{ValID, fmt.Sprintf("%d", intValue)}, nil
}

func NewStringValue(strValue string) (*Value, error) {
	return &Value{ValString, strValue}, nil
}

func NewNumericValue(numValue float64) (*Value, error) {
	return &Value{ValNumeric, fmt.Sprintf("%f", numValue)}, nil
}

func NewUSCurrencyValue(currValue float64) (*Value, error) {
	return &Value{ValNumeric, fmt.Sprintf("US$%.2f", currValue)}, nil
}

func NewEUCurrencyValue(currValue float64) (*Value, error) {
	return &Value{ValCurrencyEuro, fmt.Sprintf("EU$%.2f", currValue)}, nil
}

func NewUKCurrencyValue(currValue float64) (*Value, error) {
	return &Value{ValCurrencyPound, fmt.Sprintf("UK$%.2f", currValue)}, nil
}

func NewDateValue(y int, m int, d int) (*Value, error) {
	if 1899 < y && y < 2101 && 0 < m && m < 13 && 0 < d && d < 31 {
		return &Value{ValDate, fmt.Sprintf("%4d-%02d-%02d", y, m, d)}, nil
	}
	return nil, fmt.Errorf("Incorrect date: %4d-%02d-%02d", y, m, d)
}

func NewTimeValue(minutes int) (*Value, error) {
	return &Value{ValTime, fmt.Sprintf("%02d:%02d", minutes/60, minutes%60)}, nil
}

func NewKeyValue(kv string) (*Value, error) {
	r, e := regexp.MatchString(`\{.+)\}`, kv)
	if !r && e == nil {
		return nil, fmt.Errorf("Value %s is not of %s type", kv, ValKVList)
	}
	return &Value{ValKVList, kv}, nil
}

// NewVariable - Returns address of a new user variable, or <nil> if error
func NewVariable(name, descr string, t ValType, val *Value) *Variable {
	if name == "" {
		return nil
	}
	return &Variable{
		ID:          VariableID(utils.GenUUIDv4()),
		Name:        name,
		Description: descr,
		VarType:     VarUserVariable,
		ValType:     t,
		Value:       val,
	}
}
