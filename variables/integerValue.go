package variables

import (
	"errors"
	"fmt"
	"strconv"
)

type IntegerValue struct {
	secure bool
	value  int64
}

func (ival *IntegerValue) isSecure() bool { return ival.secure }

func (ival *IntegerValue) assign(that IVRValue) error {
	ival.secure = that.isSecure()
	v, err := that.toLong()
	ival.value = v
	return err
}

func (ival *IntegerValue) new(secure bool, strValue string) error {
	ival.secure = secure
	i64, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return errors.New("Cannot convert string to long")
	}
	ival.value = i64
	return nil
}

func (ival *IntegerValue) compareTo(value2 IVRValue) (int, error) {
	if value2.getType() == NUMERIC {
		return value2.compareTo(ival)
	}
	res := 0
	toCompare, err := value2.toLong()
	if err != nil {
		return res, errors.New("Variable to compare must be of IntegerValue type")
	}
	if ival.value > toCompare {
		res = 1
	} else if ival.value < toCompare {
		res = -1
	}
	return res, nil
}

func (ival *IntegerValue) toLong() (int64, error) {
	return ival.value, nil
}

func (ival *IntegerValue) toString() string {
	str := "*****"
	if !ival.secure {
		str, _ = ival.convertToString()
	}
	return fmt.Sprintf("{type=IntegerValue}{value=%s}", str)
}

func (ival *IntegerValue) convertToString() (string, error) {
	return strconv.FormatInt(ival.value, 10), nil
}

func (ival *IntegerValue) toBigDecimal() (float64, error) {
	return float64(ival.value), nil
}
func (*IntegerValue) toDate() ([]int32, error) {
	return nil, errors.New("Data casting error")
}
func (*IntegerValue) toTime() ([]int32, error) {
	return nil, errors.New("Data casting error")
}

func (*IntegerValue) getType() Type {
	return INTEGER
}

///////////
func getSum(args []IVRValue) (IntegerValue, error) {
	var value int64
	secure := false

	for _, v := range args {
		add, err := v.toLong()
		if err != nil {
			return IntegerValue{}, err
		}
		value += add
		secure = secure || v.isSecure()
	}

	return IntegerValue{secure, value}, nil
}

func getDifference(arg1 IVRValue, arg2 IVRValue) (IntegerValue, error) {
	v1, err := arg1.toLong()
	if err == nil {
		if v2, err := arg2.toLong(); err == nil {
			return IntegerValue{arg1.isSecure() || arg2.isSecure(), v1 - v2}, nil
		}
	}
	return IntegerValue{}, err
}

func getProduct(args []IVRValue) (IntegerValue, error) {

	value := int64(1)
	secure := false

	for _, v := range args {
		prod, err := v.toLong()
		if err != nil {
			return IntegerValue{}, err
		}
		value *= prod
		secure = secure || v.isSecure()
	}

	return IntegerValue{secure, value}, nil

}

func getQuotation(arg1 IVRValue, arg2 IVRValue) (IntegerValue, error) {
	v1, err := arg1.toLong()
	if err == nil {
		if v2, err := arg2.toLong(); err == nil {
			if v2 == 0 {
				return IntegerValue{}, errors.New("Division by zero")
			}
			return IntegerValue{arg1.isSecure() || arg2.isSecure(), v1 / v2}, nil
		}
	}
	return IntegerValue{}, err
}
