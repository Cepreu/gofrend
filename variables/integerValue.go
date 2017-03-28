package variables

import (
	"errors"
	"strconv"
)

type IntegerValue struct {
	secure bool
	value  int64
}

func (this *IntegerValue) Assign(secure bool, strValue string) error {
	this.secure = secure
	if i64, err := strconv.ParseInt(strValue, 10, 64); err != nil {
		return errors.New("Cannot convert string to long")
	}
	return nil
}

func (this IntegerValue) compareTo(value2 Value) (int, error) {
	if value2.getType() == "NumericValue" {
		return value2.compareTo(this)
	}
	res := 0
	if toCompare, err := value2.toLong(); err != nil {
		return res, errors.New("Variable to compare must be of IntegerValue type:")
	}
	if this.value > toCompare {
		res = 1
	} else if value < toCompare {
		res = -1
	}
	return res, nil
}

func (this IntegerValue) toLong() (int64, error) {
	return this.value, nil
}

func (this IntegerValue) convertToString() string {
	return strconv.FormatInt(this.value, 16)
}

func (IntegerValue) getType() {
	return Type.INTEGER
}

///////////
func getSum(args []Value) (IntegerValue, error) {
	value := 0
	secure := false

	for _, v := range args {
		if add, err := v.toLong(); err != nil {
			return nil, err
		}
		value += add
		secure |= v.isSecure()
	}

	return IntegerValue{secure, value}, nil
}

func getDifference(arg1 Value, arg2 Value) (IntegerValue, error) {
	if v1, err := arg1.toLong(); err == nil {
		if v2, err := arg2.toLong(); err == nil {
			return IntegerValue{arg1.isSecure() || arg2.isSecure(), v1 - v2}, nil
		}
	}
	return nil, err
}

func getProduct(args []Value) (IntegerValue, error) {

	value := 1
	secure := false

	for _, v := range args {
		if prod, err := v.toLong(); err != nil {
			return nil, err
		}
		value *= prod
		secure |= v.isSecure()
	}

	return IntegerValue{secure, value}, nil

}

func getQuotation(arg1 Value, arg2 Value) (IntegerValue, error) {
	if v1, err := arg1.toLong(); err == nil {
		if v2, err := arg2.toLong(); err == nil {
			if v2 == 0 {
				return nil, errors.New("Division by zero")
			}
			return IntegerValue{arg1.isSecure() || arg2.isSecure(), v1 / v2}, nil
		}
	}
	return nil, err
}
