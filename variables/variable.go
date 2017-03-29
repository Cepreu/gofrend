package variables

import (
	"fmt"

	"github.com/Cepreu/gofrend/utils"
)

const (
	ATTR_SYSTEM          uint8 = 1
	ATTR_CRM             uint8 = 2
	ATTR_EXTERNAL              = 4
	ATTR_INTERNAL              = 8
	ATTR_USER_PREDEFINED       = 16
	ATTR_TTS_ENUMERATION       = 32
	ATTR_USER_DEFINED          = ATTR_EXTERNAL | ATTR_INTERNAL | ATTR_USER_PREDEFINED
	ATTR_WRITABLE              = ATTR_CRM | ATTR_USER_DEFINED
	ATTR_ANY                   = ATTR_SYSTEM | ATTR_CRM | ATTR_USER_DEFINED
)

type Type int

const (
	UNDEFINED Type = iota
	STRING
	INTEGER
	ERROR
	NUMERIC
	CURRENCY
	DATE
	TIME
	CURRENCY_POUND
	CURRENCY_EURO
	KVLIST
)

func typeName(t Type) string {
	var TypeStr = [...]string{
		"String",
		"Integer",
		"ErrorCode",
		"Numeric",
		"Currency (dollar)",
		"Date",
		"Time",
		"Currency (pound)",
		"Currency (euro)",
		"List of key-value pairs"}
	return TypeStr[t]
}

type Variable struct {
	value       IVRValue
	name        string
	description string
	attributes  uint8
	isNullValue bool
}

func (v Variable) IsNull() bool {
	return v.isNullValue
}
func (pv *Variable) AssignNull() {
	pv.isNullValue = true
}
func (v Variable) compareTo(arg Variable) (int, error) {
	if v.isNullValue && arg.isNullValue {
		return 0, nil
	} else if v.isNullValue || arg.isNullValue {
		return 0, fmt.Errorf("Cannot compare variable to NULL value : %v and %v", v.value, arg.value)
	}
	return v.value.compareTo(arg.value)
}

func (pv *Variable) assign(val IVRValue) {
	//	if val.isEmpty() {
	//		pv.isNullValue = true
	//	} else {
	pv.value.assign(val)
	pv.isNullValue = false
	//	}
}
func (v Variable) IsExternal() bool {
	return !((v.attributes & ATTR_EXTERNAL) == 0)
}

func (v Variable) IsCrm() bool {
	return (v.attributes & ATTR_CRM) != 0
}

func (v Variable) IsCav() bool {
	return false //TBD (this instanceof CavVariable);
}

func (v Variable) ToString() string {
	re := "NULL"
	if !v.isNullValue {
		re = v.value.toString()
	}
	return fmt.Sprintf("{{name=\"%s\"}{description=\"%s\"} %s}", v.name, v.description, re)
}

func (v Variable) HashCode() uint64 {
	var (
		prime  uint64 = 31
		result uint64 = 1
	)
	result = prime*result + uint64(v.attributes)
	result = prime*result + utils.HashCode(v.description)
	if v.IsNull() {
		result = prime*result + 1231
	} else {
		result = prime*result + 1237
	}
	result = prime*result + utils.HashCode(v.name)
	result = prime*result + utils.HashCode(v.value.toString())
	return result
}
