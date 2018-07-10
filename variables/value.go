package variables

import (
	"fmt"
)

type Value interface {
	String() string
	compareTo(Value) (int, error)
	assign(Value) error
	new(bool, string) error
	toLong() (int64, error)
	toBigDecimal() (float64, error)
	toDate() ([]int, error)
	toTime() (int32, error)
	isSecure() bool
	getType() Type
}

type defaultValueImpl struct {
	secure bool
}

func (ival *defaultValueImpl) isSecure() bool {
	return ival.secure
}

func (ival *defaultValueImpl) assign(that Value) error {
	ival.secure = that.isSecure()
	return nil
}

func (ival *defaultValueImpl) new(secure bool, strValue string) error {
	ival.secure = secure
	return nil
}

func (ival *defaultValueImpl) compareTo(value2 Value) (int, error) {
	return 0, nil
}

func (ival *defaultValueImpl) toLong() (int64, error) {
	return 0, fmt.Errorf("toLong() is not supported for\" %v\"", ival.getType())
}

func (ival *defaultValueImpl) String() (string, error) {
	return "", fmt.Errorf("String() is not supported for \"%v\"", ival.getType())
}

func (ival *defaultValueImpl) toBigDecimal() (float64, error) {
	return 0, fmt.Errorf("toBigDecimal() is not supported for \"%v\"", ival.getType())
}
func (ival *defaultValueImpl) toDate() ([]int, error) {
	return nil, fmt.Errorf("toDate() is not supported for \"%v\"", ival.getType())
}
func (ival *defaultValueImpl) toTime() (int32, error) {
	return 0, fmt.Errorf("toTime() is not supported for \"%v\"", ival.getType())
}

func (*defaultValueImpl) getType() Type {
	return UNDEFINED
}
