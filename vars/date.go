package vars

import (
	"fmt"
	"strconv"
)

// DateVar - Date variable
type DateVar struct {
	Variable
	value *Currency
}

// Date - value of Date type
type Date struct {
	secured
	day   int
	month int
	year  int
}

//SetValue - helper function for parsing xml
func (dval *Date) SetValue(fieldName string, fieldStrValue string) (err error) {
	switch fieldName {
	case "day":
		dval.day, err = strconv.Atoi(fieldStrValue)
	case "month":
		dval.month, err = strconv.Atoi(fieldStrValue)
	case "year":
		dval.year, err = strconv.Atoi(fieldStrValue)
	default:
		err = fmt.Errorf("Unknown field '%s' for Date value", fieldName)
	}
	return err
}

// func (dval *Date) assign(that Value) error {
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

//NewDate - returns pointer to a new Date struct, or <nil> for an error
func NewDate(y, m, d int) *Date {
	if y == 0 || m == 0 || d == 0 {
		return nil
	}
	return &Date{secured{secured: false}, d, m, y}
}

func (dval *Date) new(secure bool, strValue string) error {
	dval.SetSecured(secure)
	var err error
	dval.day, dval.month, dval.year, err = vuStringToDate(strValue)
	return err
}

func (dval *Date) compareTo(value2 *Date) (int, error) {
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

func (dval *Date) toLong() (int, error) {
	return dval.day + dval.month*100 + dval.year*10000, nil
}

func (dval *Date) String() string {
	return fmt.Sprintf("%4d-%02d-%02d", dval.year, dval.month, dval.day)
}

func (dval *Date) toDate() *Date {
	return dval
}

func (*Date) getType() VarType {
	return VarDate
}
