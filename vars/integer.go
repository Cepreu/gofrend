package vars

import (
	"errors"
	"fmt"
	"strconv"
)

// IntegerVar - Integer variable
type IntegerVar struct {
	Variable
	value *Integer
}

// Integer - value of Integer type
type Integer struct {
	value int
}

//SetValue - helper function for parsing xml
func (ival Integer) SetValue(fieldName string, fieldStrValue string) (err error) {
	switch fieldName {
	case "value":
		ival.value, err = strconv.Atoi(fieldStrValue)
	default:
		err = fmt.Errorf("Unknown field '%s' for Integer value", fieldName)
	}
	return err
}

func (ival Integer) new(strValue string) error {
	i, err := strconv.Atoi(strValue)
	if err != nil {
		return errors.New("Cannot convert string to long")
	}
	ival.value = i
	return nil
}

func (ival Integer) String() string {
	return fmt.Sprintf("{type=Integer}{value=%s}", ival.value)
}

func (Integer) getType() VarType {
	return VarInteger
}

//NewInteger - returns pointer to a new Date struct, or <nil> for an error
func NewInteger(v int) *Integer {
	return &Integer{v}
}
