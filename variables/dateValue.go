package variables

import (
	"errors"
	"fmt"
)

const (
	year int = iota
	month
	day
)

type DateValue struct {
	defaultValueImpl
	value [3]int
}

func (dval *DateValue) isSecure() bool { return dval.secure }

func (dval *DateValue) assign(that Value) error {
	dval.defaultValueImpl.assign(that)
	v, err := that.toDate()
	if err == nil {
		dval.value[day] = v[day]
		dval.value[month] = v[month]
		dval.value[year] = v[year]
	}
	return err
}

func (dval *DateValue) new(secure bool, strValue string) error {
	dval.secure = secure
	var err error
	dval.value[day], dval.value[month], dval.value[year], err = vuStringToDate(strValue)
	return err
}

func (dval *DateValue) compareTo(value2 Value) (int, error) {
	res := 0
	if value2.getType() != DATE {
		return res, errors.New("Variable to compare must be of DateValue type")
	}
	d1, _ := dval.toLong()
	d2, _ := value2.toLong()
	if d1 > d2 {
		res = 1
	} else if d1 < d2 {
		res = -1
	}
	return res, nil
}

func (dval *DateValue) toLong() (int64, error) {
	return int64(dval.value[day] + dval.value[month]*100 + dval.value[year]*10000), nil
}

func (dval *DateValue) toString() string {
	str := "*****"
	if !dval.secure {
		str, _ = dval.convertToString()
	}
	return fmt.Sprintf("{type=DateValue}{value=%s}", str)
}

func (dval *DateValue) convertToString() (string, error) {
	return fmt.Sprintf("%4d-%02d-%02d", dval.value[year], dval.value[month], dval.value[day]), nil
}

func (dval *DateValue) toDate() ([]int, error) {
	return dval.value[:], nil
}

func (*DateValue) getType() Type {
	return DATE
}
