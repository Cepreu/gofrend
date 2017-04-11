package variables

import (
	"errors"
	"fmt"
)

type DateValue struct {
	defaultValueImpl
	day   int
	month int
	year  int
}

func (dval *DateValue) isSecure() bool { return dval.secure }

func (dval *DateValue) assign(that Value) error {
	dval.defaultValueImpl.assign(that)
	v, err := that.toDate()
	if err == nil {
		dval.day = v[2]
		dval.month = v[1]
		dval.year = v[0]
	}
	return err
}

func (dval *DateValue) new(secure bool, strValue string) error {
	dval.secure = secure
	var err error
	dval.day, dval.month, dval.year, err = vuStringToDate(strValue)
	return err
}

func (dval *DateValue) compareTo(value2 Value) (int, error) {
	res := 0
	toCompare, err := value2.toDate()
	if err != nil {
		return res, errors.New("Variable to compare must be of DateValue type")
	}
	d1 := dval.day + dval.month*100 + dval.year*10000
	d2 := toCompare[2] + toCompare[1]*100 + toCompare[0]*10000
	if d1 > d2 {
		res = 1
	} else if d1 < d2 {
		res = -1
	}
	return res, nil
}

func (dval *DateValue) toLong() (int64, error) {
	return int64(dval.day + dval.month*100 + dval.year*10000), nil
}

func (dval *DateValue) toString() string {
	str := "*****"
	if !dval.secure {
		str, _ = dval.convertToString()
	}
	return fmt.Sprintf("{type=DateValue}{value=%s}", str)
}

func (dval *DateValue) convertToString() (string, error) {
	return fmt.Sprintf("%4d-%02d-%02d", dval.year, dval.month, dval.day), nil
}

func (dval *DateValue) toDate() ([]int, error) {
	return []int{dval.year, dval.month, dval.day}, nil
}

func (*DateValue) getType() Type {
	return DATE
}
