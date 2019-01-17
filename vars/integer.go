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
	secured
	value int
}

//SetValue - helper function for parsing xml
func (ival *Integer) SetValue(fieldName string, fieldStrValue string) (err error) {
	switch fieldName {
	case "value":
		ival.value, err = strconv.Atoi(fieldStrValue)
	default:
		err = fmt.Errorf("Unknown field '%s' for Integer value", fieldName)
	}
	return err
}

//NewInteger - returns pointer to a new Integer struct, or <nil> for an error
func NewInteger(v int) *Integer {
	return &Integer{secured{secured: false}, v}
}

// func (ival *Integer) assign(that *Value) error {
// 	ival.defaultValueImpl.assign(that)
// 	v, err := (*that).toLong()
// 	ival.value = v
// 	return err
// }

func (ival *Integer) new(secure bool, strValue string) error {
	ival.SetSecured(secure)
	i, err := strconv.Atoi(strValue)
	if err != nil {
		return errors.New("Cannot convert string to long")
	}
	ival.value = i
	return nil
}

// func (ival *Integer) CompareTo(value2 Value) (int, error) {
// 	if value2.getType() == NUMERIC {
// 		return value2.CompareTo(ival)
// 	}
// 	res := 0
// 	toCompare, err := value2.toLong()
// 	if err != nil {
// 		return res, errors.New("Variable to compare must be of Integer type")
// 	}
// 	if ival.value > toCompare {
// 		res = 1
// 	} else if ival.value < toCompare {
// 		res = -1
// 	}
// 	return res, nil
// }

func (ival *Integer) toLong() (int, error) {
	return ival.value, nil
}
func (ival *Integer) toDate() (Date, error) {
	return Date{}, nil
}

func (ival *Integer) String() string {
	str := "*****"
	if !ival.IsSecured() {
		str, _ = ival.convertToString()
	}
	return fmt.Sprintf("{type=Integer}{value=%s}", str)
}

func (ival *Integer) convertToString() (string, error) {
	return fmt.Sprintf("%d", ival.value), nil
}

func (ival *Integer) toBigDecimal() (float64, error) {
	return float64(ival.value), nil
}

func (*Integer) getType() VarType {
	return VarInteger
}

/*
func getSum(args []Value) (Integer, error) {
	var val int64
	scr := false

	for _, v := range args {
		add, err := v.toLong()
		if err != nil {
			return Integer{}, err
		}
		val += add
		scr = scr || v.isSecure()
	}

	return Integer{defaultValueImpl{scr}, val}, nil
}

func getDifference(arg1 Value, arg2 Value) (Integer, error) {
	v1, err := arg1.toLong()
	if err == nil {
		if v2, err := arg2.toLong(); err == nil {
			return Integer{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 - v2}, nil
		}
	}
	return Integer{}, err
}

func getProduct(args []Value) (Integer, error) {

	value := int64(1)
	secure := false

	for _, v := range args {
		prod, err := v.toLong()
		if err != nil {
			return Integer{}, err
		}
		value *= prod
		secure = secure || v.isSecure()
	}

	return Integer{defaultValueImpl{secure}, value}, nil

}

func getQuotation(arg1 Value, arg2 Value) (Integer, error) {
	v1, err := arg1.toLong()
	if err == nil {
		if v2, err := arg2.toLong(); err == nil {
			if v2 == 0 {
				return Integer{}, errors.New("Division by zero")
			}
			return Integer{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 / v2}, nil
		}
	}
	return Integer{}, err
}
*/
