package variables

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type DateValue struct {
	defaultValueImpl
	value int64
}

func (dval *DateValue) isSecure() bool { return dval.secure }

func (dval *DateValue) assign(that IVRValue) error {
	dval.defaultValueImpl.assign(that)
	v, err := that.toDate()
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
	return errors.New("Cannot convert string to date")
}

func (dval *DateValue) compareTo(value2 IVRValue) (int, error) {
	res := 0
	toCompare, err := value2.toDate()
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
	return DATE
}
