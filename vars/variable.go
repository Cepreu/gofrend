package vars

import "fmt"

const (
	attrSystem         uint8 = 1
	attrCRM            uint8 = 2
	attrExternal             = 4
	attrInternal             = 8
	attrUserPredefined       = 16
	attrTTSEnumeration       = 32
	attrUserDefined          = attrExternal | attrInternal | attrUserPredefined
	attrWritable             = attrCRM | attrUserDefined
	attrAny                  = attrSystem | attrCRM | attrUserDefined
)

type VarType int

const (
	VarUndefined VarType = iota
	VarString
	VarInteger
	VarError
	VarDate
	VarNumeric
	VarCurrency
	VarTime
	VarCurrencyPound
	VarCurrencyEuro
	VarKVList
)

func typeName(t VarType) string {
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

// Value - interface type for internal ivr values
type Value interface {
	new(bool, string) error
	SetValue(string, string) error
	SetSecured(bool)
	IsSecured() bool
	fmt.Stringer
}

type secured struct {
	secured bool
}

func (s secured) IsSecured() bool {
	return s.secured
}
func (s secured) SetSecured(v bool) {
	s.secured = v
}

// Variable - ivr variable description
type Variable struct {
	name        string
	value       Value
	description string
	attributes  int
	isNullValue bool
}

// NewVariable - Returns address of a new user variable, or <nil> if error
func NewVariable(name, descr string, attr int, isNull bool) *Variable {
	if name == "" {
		return nil
	}
	return &Variable{name: name, description: descr, attributes: attr, isNullValue: isNull}
}

// SetValue - Assigns datatype-specific info to the variable
func (vv *Variable) SetValue(theVal Value) error {
	vv.value = theVal
	return nil
}

func (v Variable) IsNull() bool {
	return v.isNullValue
}

func (pv *Variable) AssignNull() {
	pv.isNullValue = true
}

func (v *Variable) String() string {
	re := "NULL"
	if v.value != nil {
		re = v.value.String()
	}
	return fmt.Sprintf("{{name=\"%s\"}{description=\"%s\"} %s}", v.name, v.description, re)
}

// func (pv *Variable) Assign(val *Value) {
// 	//	if val.isEmpty() {
// 	//		pv.isNullValue = true
// 	//	} else {
// 	pv.value.assign(val)
// 	pv.isNullValue = false
// 	//	}
// }

// func (v Variable) compareTo(arg Variable) (int, error) {
// 	if v.isNullValue && arg.isNullValue {
// 		return 0, nil
// 	} else if v.isNullValue || arg.isNullValue {
// 		return 0, fmt.Errorf("Cannot compare variable to NULL value : %v and %v", v.value, arg.value)
// 	}
// 	return v.value.compareTo(arg.value)
// }

// func (v Variable) IsExternal() bool {
// 	return !((v.attributes & attrEXTERNAL) == 0)
// }

// func (v Variable) IsCrm() bool {
// 	return (v.attributes & attrCRM) != 0
// }

// func (v Variable) IsCav() bool {
// 	return false //TBD (this instanceof CavVariable);
// }

// func (v Variable) HashCode() uint64 {
// 	var (
// 		prime  uint64 = 31
// 		result uint64 = 1
// 	)
// 	result = prime*result + uint64(v.attributes)
// 	result = prime*result + utils.HashCode(v.description)
// 	if v.IsNull() {
// 		result = prime*result + 1231
// 	} else {
// 		result = prime*result + 1237
// 	}
// 	result = prime*result + utils.HashCode(v.name)
// 	result = prime*result + utils.HashCode(v.value.String())
// 	return result
// }
