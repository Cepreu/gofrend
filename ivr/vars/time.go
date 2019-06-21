package vars

import (
	"fmt"
	"strconv"
)

// TimeVar - Time variable
type TimeVar struct {
	Variable
	value *Time
}

// Time - value of Time type
type Time struct {
	minutes int
}

//SetValue - helper function for parsing xml
func (tval *Time) SetValue(fieldName string, fieldStrValue string) (err error) {
	switch fieldName {
	case "minutes":
		tval.minutes, err = strconv.Atoi(fieldStrValue)
	default:
		err = fmt.Errorf("Unknown field '%s' for Time value", fieldName)
	}
	return err
}

func NewTime(minutes int) *Time {
	return &Time{minutes}
}

func (tval *Time) String() string {
	return fmt.Sprintf("%d", tval.minutes)
}

func (tval *Time) new(strValue string) error {
	var err error
	tval.minutes, err = vuStringToMinutes(strValue)
	return err
}

// func (tval *Time) assign(that Value) error {
// 	tval.defaultValueImpl.assign(that)
// 	v, err := that.toTime()
// 	if err == nil {
// 		tval.value = v
// 	}
// 	return err
// }

// func (tval *Time) compareTo(value2 Value) (int, error) {
// 	res := 0
// 	toCompare, err := value2.toTime()
// 	if err != nil {
// 		return res, errors.New("Variable to compare must be of Time type")
// 	}
// 	if tval.value > toCompare {
// 		res = 1
// 	} else if tval.value < toCompare {
// 		res = -1
// 	}
// 	return res, nil
// }

// func (tval *Time) toLong() (int64, error) {
// 	return int64(tval.minutes), nil
// }

// func (tval *Time) toBigDecimal() (float64, error) {
// 	return float64(tval.minutes), nil
// }

func (*Time) getType() VarType {
	return VarTime
}

///////
// func (tval *Time) getDifference(t2 Time) (iv Integer) {
// 	iv = Integer{}
// 	iv.new(fmt.Sprintf("%d", tval.minutes-t2.minutes))
// 	return iv
// }

// func (tval Time) getAddition(i2 IntegerValue) (iv IntegerValue, e error) {
// 	iv = IntegerValue{}
// 	i64 := int64(tval.value) + i2.value
// 	if i64 >= 24*60 {
// 		i64 %= 24 * 60
// 		e = errors.New("Time in the next day")
// 	}
// 	iv.new(true, fmt.Sprintf("%d", i64))
// 	return iv, e
// }
// func (tval Time) getHours() int32 {
// 	return tval.value / 60
// }
// func (tval Time) getMinutes() int32 {
// 	return tval.value % 60
// }
