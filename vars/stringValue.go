package vars

import (
	"fmt"
	"strconv"
)

type StringValue struct {
	secure       bool
	value        string
	fileNameType FileNameType
	id           int
}
type FileNameType int

const (
	fnUndefined FileNameType = iota
	fnGreeting
	fnVoiceMail
	fnRecording
)

//NewStringValue - returns pointer to a new StringValue struct, or <nil> for an error
func NewStringValue(v string, i int) *StringValue {
	return &StringValue{secure: false, value: v, id: i}
}

func (sval *StringValue) isSecure() bool { return sval.secure }

// func (sval *StringValue) assign(that Value) error {
// 	sval.defaultValueImpl.assign(that)
// 	if that.getType() == STRING {
// 		t, ok := that.(*StringValue)
// 		if !ok {
// 			return fmt.Errorf("Corrupted IVRValue: %v", that)
// 		}
// 		if t.isFileName() {
// 			sval.fileNameType = t.fileNameType
// 			sval.id = t.id
// 			sval.value = fmt.Sprintf("%v, id=%v)", sval.fileNameType, sval.id)
// 			return nil
// 		}
// 	}
// 	sval.value = that.String()
// 	return nil
// }

func (sval *StringValue) new(secure bool, strValue string) error {
	sval.secure = secure
	sval.value = strValue
	return nil
}

// func (sval *StringValue) compareTo(value2 value) (int, error) {
// 	res := 0
// 	toCompare := value2.String()
// 	if sval.value > toCompare {
// 		res = 1
// 	} else if sval.value < toCompare {
// 		res = -1
// 	}
// 	return res, nil
// }

func (sval *StringValue) toLong() (int64, error) {
	return strconv.ParseInt(sval.value, 10, 64)
}

func (sval *StringValue) String() string {
	return sval.value
}

func (sval *StringValue) toBigDecimal() (float64, error) {
	return strconv.ParseFloat(sval.value, 64)
}

func (sval *StringValue) toTime() (int, error) {
	return vuStringToMinutes(sval.value)
}
func (sval *StringValue) toDate() ([]int, error) {
	if y, m, d, e := vuStringToDate(sval.value); e == nil {
		return []int{y, m, d}, e
	}
	return nil, fmt.Errorf("Cannot convert \" %s\" to Date", sval.value)
}
func (*StringValue) getType() VarType {
	return VarString
}

///////////
func (sval StringValue) isFileName() bool {
	return (sval.fileNameType != fnUndefined)
}
func (sval StringValue) getFileNameType() FileNameType {
	return sval.fileNameType
}
func (sval StringValue) getRecordingID() int {
	return sval.id
}
