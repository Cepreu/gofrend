package variables

import (
	"errors"
	"fmt"
)

type TimeValue struct {
	defaultValueImpl
	value int32
}

func (tval *TimeValue) isSecure() bool { return tval.secure }

func (tval *TimeValue) assign(that Value) error {
	tval.defaultValueImpl.assign(that)
	v, err := that.toTime()
	if err == nil {
		tval.value = v
	}
	return err
}

func (tval *TimeValue) new(secure bool, strValue string) error {
	tval.secure = secure
	var err error
	tval.value, err = vuStringToMinutes(strValue)
	return err
}

func (tval *TimeValue) compareTo(value2 Value) (int, error) {
	res := 0
	toCompare, err := value2.toTime()
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
	return int64(tval.value), nil
}

func (tval *TimeValue) toString() string {
	str := "*****"
	if !tval.secure {
		str, _ = tval.convertToString()
	}
	return fmt.Sprintf("{type=TimeValue}{value=%s}", str)
}

func (tval *TimeValue) convertToString() (string, error) {
	return fmt.Sprintf("%d", tval.value), nil
}

func (tval *TimeValue) toBigDecimal() (float64, error) {
	return float64(tval.value), nil
}

func (*TimeValue) getType() Type {
	return TIME
}

///////
func (tval TimeValue) getDifference(t2 TimeValue) (iv IntegerValue) {
	iv = IntegerValue{}
	iv.new(true, fmt.Sprintf("%d", tval.value-t2.value))
	return iv
}
func (tval TimeValue) getAddition(i2 IntegerValue) (iv IntegerValue, e error) {
	iv = IntegerValue{}
	i64 := int64(tval.value) + i2.value
	if i64 >= 24*60 {
		i64 %= 24 * 60
		e = errors.New("Time in the next day")
	}
	iv.new(true, fmt.Sprintf("%d", i64))
	return iv, e
}
func (tval TimeValue) getHours() int32 {
	return tval.value / 60
}
func (tval TimeValue) getMinutes() int32 {
	return tval.value % 60
}
