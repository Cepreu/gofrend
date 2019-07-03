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
	new(string) error
	SetValue(string, string) error
	String() string
}

// Variable - ivr variable description
type Variable struct {
	Name        string
	Value       Value
	Description string
	Attributes  int
	IsNullValue bool
	Secured     bool
	Vtype       VarType
}

// NewVariable - Returns address of a new user variable, or <nil> if error
func NewVariable(name, descr string, attr int, isNull bool) *Variable {
	if name == "" {
		return nil
	}
	return &Variable{Name: name, Description: descr, Attributes: attr, IsNullValue: isNull}
}

// SetValue - Assigns datatype-specific info to the variable
func (vv *Variable) SetValue(theVal Value) error {
	vv.Value = theVal
	return nil
}

// GetValue - Returns pointer to the variable's value
func (vv *Variable) GetValue() Value {
	return vv.Value
}

func (v *Variable) IsNull() bool {
	return v.IsNullValue
}

func (pv *Variable) AssignNull() {
	pv.IsNullValue = true
}

func (v *Variable) String() string {
	if v.Value != nil {
		re := v.GetValue()
		return fmt.Sprintf("{{name=\"%s\"}{description=\"%s\"} %s}", v.Name, v.Description, re)
	}
	return "NILL"
}

// func (pv *Variable) Assign(val *Value) {
// 	//	if val.isEmpty() {
// 	//		pv.IsNullValue = true
// 	//	} else {
// 	pv.Value.assign(val)
// 	pv.IsNullValue = false
// 	//	}
// }

// func (v Variable) compareTo(arg Variable) (int, error) {
// 	if v.IsNullValue && arg.IsNullValue {
// 		return 0, nil
// 	} else if v.IsNullValue || arg.IsNullValue {
// 		return 0, fmt.Errorf("Cannot compare variable to NULL value : %v and %v", v.Value, arg.Value)
// 	}
// 	return v.Value.compareTo(arg.Value)
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
// 	result = prime*result + utils.HashCode(v.Description)
// 	if v.IsNull() {
// 		result = prime*result + 1231
// 	} else {
// 		result = prime*result + 1237
// 	}
// 	result = prime*result + utils.HashCode(v.Name)
// 	result = prime*result + utils.HashCode(v.Value.String())
// 	return result
// }