package variables

import (
	"errors"
	"fmt"
	"strconv"
)

type TimeValue struct {
	defaultValueImpl
	value int64
}

func (tval *TimeValue) isSecure() bool { return tval.secure }

func (tval *TimeValue) assign(that IVRValue) error {
	tval.defaultValueImpl.assign(that)
	v, err := that.toTime()
	if err == nil {tval.value = v}
	return err
}

func (tval *TimeValue) new(secure bool, strValue string) error {
	tval.secure = secure
	tval.value, err = vuStringToMinutes(strValue)
	return err
}

func (tval *TimeValue) compareTo(value2 IVRValue) (int, error) {
	if value2.getType() == NUMERIC {
		return value2.compareTo(tval)
	}
	res := 0
	toCompare, err := value2.toLong()
	if err != nil {
		return res, errors.New("Variable to compare must be of TimeValue type")
	}
	if tval.value > toCompare {
		res = 1
	} else if tval.value < toCompare {
		res = -1
	}
	return res, nil
}

func (tval *TimeValue) toLong() (int64, error) {
	return tval.value, nil
}

func (tval *TimeValue) toString() string {
	str := "*****"
	if !tval.secure {
		str, _ = tval.convertToString()
	}
	return fmt.Sprintf("{type=TimeValue}{value=%s}", str)
}

func (tval *TimeValue) convertToString() (string, error) {
	return strconv.FormatInt(tval.value, 10), nil
}

func (tval *TimeValue) toBigDecimal() (float64, error) {
	return float64(tval.value), nil
}

func (*TimeValue) getType() Type {
	return INTEGER
}

///////////
func getSum(args []IVRValue) (TimeValue, error) {
	var val int64
	scr := false

	for _, v := range args {
		add, err := v.toLong()
		if err != nil {
			return TimeValue{}, err
		}
		val += add
		scr = scr || v.isSecure()
	}

	return TimeValue{defaultValueImpl{scr}, val}, nil
}

func getDifference(arg1 IVRValue, arg2 IVRValue) (TimeValue, error) {
	v1, err := arg1.toLong()
	if err == nil {
		if v2, err := arg2.toLong(); err == nil {
			return TimeValue{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 - v2}, nil
		}
	}
	return TimeValue{}, err
}

func getProduct(args []IVRValue) (TimeValue, error) {

	value := int64(1)
	secure := false

	for _, v := range args {
		prod, err := v.toLong()
		if err != nil {
			return TimeValue{}, err
		}
		value *= prod
		secure = secure || v.isSecure()
	}

	return TimeValue{defaultValueImpl{secure}, value}, nil

}

func getQuotation(arg1 IVRValue, arg2 IVRValue) (TimeValue, error) {
	v1, err := arg1.toLong()
	if err == nil {
		if v2, err := arg2.toLong(); err == nil {
			if v2 == 0 {
				return TimeValue{}, errors.New("Division by zero")
			}
			return TimeValue{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 / v2}, nil
		}
	}
	return TimeValue{}, err
}
