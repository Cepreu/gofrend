package variables

import (
	"errors"
	"fmt"
)

type DateValue struct {
	defaultValueImpl
	value struct {
		day   uint8
		month uint8
		year  uint16
	}
}

func (dval *DateValue) isSecure() bool { return dval.secure }

func (dval *DateValue) assign(that *Value) error {
	dval.defaultValueImpl.assign(that)
	v := that.toDate()
	if v == nil {
		return errors.New("Variable of this type cannot be assigned to Date")
	}
	dval.value.day = v.value.day
	dval.value.month = v.value.month
	dval.value.year = v.value.year

	return nil
}

func (dval *DateValue) new(secure bool, strValue string) error {
	dval.secure = secure
	var err error
	dval.value[day], dval.value[month], dval.value[year], err = vuStringToDate(strValue)
	return err
}

func (dval *DateValue) compareTo(value2 *DateValue) (int, error) {
	res := 0

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

func (dval *DateValue) String() string {
	return fmt.Sprintf("%4d-%02d-%02d", dval.value[year], dval.value[month], dval.value[day])
}

func (dval *DateValue) toDate() *DateValue {
	return dval.value[:], nil
}

func (*DateValue) getType() Type {
	return DATE
}
