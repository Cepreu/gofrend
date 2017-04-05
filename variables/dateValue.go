package variables

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const timeRe = "^((0?[1-9])|(1[0-2])):([0-5][0-9]) ?(AM|PM|am|pm)|(([01][0-9])|(2[0-3])):([0-5][0-9])$"

type DateValue struct {
	defaultValueImpl
	value int64
}

func (dval *DateValue) isSecure() bool { return dval.secure }

func (dval *DateValue) assign(that IVRValue) error {
	dval.defaultValueImpl.assign(that)
	v, err := that.toTime()
	if err == nil {
		dval.value = v
	}
	return err
}

func (dval *DateValue) new(secure bool, strValue string) error {
	dval.secure = secure
	r, _ := regexp.Compile(timeRe)
	res := r.FindAllStringSubmatch(strValue, -1)
	if len(res) == 10 {
		i64, _ := strconv.ParseInt(res[0][1], 10, 64)
		dval.value += i64
		i64, _ = strconv.ParseInt(res[0][7], 10, 64)
		dval.value += i64
		if res[0][5] == "PM" || res[0][5] == "pm" {
			dval.value += 12
		}
		dval.value *= 60
		i64, _ = strconv.ParseInt(res[0][4], 10, 64)
		dval.value += i64
		i64, _ = strconv.ParseInt(res[0][9], 10, 64)
		dval.value += i64
	}
	return errors.New("Cannot convert string to time")
}

func (dval *DateValue) compareTo(value2 IVRValue) (int, error) {
	if value2.getType() == NUMERIC {
		return value2.compareTo(dval)
	}
	res := 0
	toCompare, err := value2.toLong()
	if err != nil {
		return res, errors.New("Variable to compare must be of DateValue type")
	}
	if dval.value > toCompare {
		res = 1
	} else if dval.value < toCompare {
		res = -1
	}
	return res, nil
}

func (dval *DateValue) toLong() (int64, error) {
	return dval.value, nil
}

func (dval *DateValue) toString() string {
	str := "*****"
	if !dval.secure {
		str, _ = dval.convertToString()
	}
	return fmt.Sprintf("{type=DateValue}{value=%s}", str)
}

func (dval *DateValue) convertToString() (string, error) {
	return strconv.FormatInt(dval.value, 10), nil
}

func (dval *DateValue) toBigDecimal() (float64, error) {
	return float64(dval.value), nil
}

func (*DateValue) getType() Type {
	return INTEGER
}

///////////
func getSum(args []IVRValue) (DateValue, error) {
	var val int64
	scr := false

	for _, v := range args {
		add, err := v.toLong()
		if err != nil {
			return DateValue{}, err
		}
		val += add
		scr = scr || v.isSecure()
	}

	return DateValue{defaultValueImpl{scr}, val}, nil
}

func getDifference(arg1 IVRValue, arg2 IVRValue) (DateValue, error) {
	v1, err := arg1.toLong()
	if err == nil {
		if v2, err := arg2.toLong(); err == nil {
			return DateValue{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 - v2}, nil
		}
	}
	return DateValue{}, err
}

func getProduct(args []IVRValue) (DateValue, error) {

	value := int64(1)
	secure := false

	for _, v := range args {
		prod, err := v.toLong()
		if err != nil {
			return DateValue{}, err
		}
		value *= prod
		secure = secure || v.isSecure()
	}

	return DateValue{defaultValueImpl{secure}, value}, nil

}

func getQuotation(arg1 IVRValue, arg2 IVRValue) (DateValue, error) {
	v1, err := arg1.toLong()
	if err == nil {
		if v2, err := arg2.toLong(); err == nil {
			if v2 == 0 {
				return DateValue{}, errors.New("Division by zero")
			}
			return DateValue{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 / v2}, nil
		}
	}
	return DateValue{}, err
}
