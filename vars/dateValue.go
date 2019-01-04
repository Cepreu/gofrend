package vars

import (
	"fmt"
)

type DateValue struct {
	secure bool
	day    int
	month  int
	year   int
}

func (dval *DateValue) isSecure() bool { return dval.secure }

// func (dval *DateValue) assign(that Value) error {
// 	dval.defaultValueImpl.assign(&that)
// 	v, err := that.toDate()
// 	if err != nil {
// 		return errors.New("Variable of this type cannot be assigned to Date")
// 	}
// 	dval.value.day = v.value.day
// 	dval.value.month = v.value.month
// 	dval.value.year = v.value.year

// 	return nil
// }

//NewDateValue - returns pointer to a new Datevalue struct, or <nil> for an error
func NewDateValue(y, m, d int) *DateValue {
	if y == 0 || m == 0 || d == 0 {
		return nil
	}
	return &DateValue{secure: false, day: d, month: m, year: y}
}

func (dval *DateValue) new(secure bool, strValue string) error {
	dval.secure = secure
	var err error
	dval.day, dval.month, dval.year, err = vuStringToDate(strValue)
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

func (dval *DateValue) toLong() (int, error) {
	return dval.day + dval.month*100 + dval.year*10000, nil
}

func (dval *DateValue) String() string {
	return fmt.Sprintf("%4d-%02d-%02d", dval.year, dval.month, dval.day)
}

func (dval *DateValue) toDate() *DateValue {
	return dval
}

func (*DateValue) getType() VarType {
	return VarDate
}
