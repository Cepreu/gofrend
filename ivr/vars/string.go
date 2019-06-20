package vars

import (
	"fmt"
	"strconv"
)

// StringVar - String variable
type StringVar struct {
	Variable
	value *String
}

// String - value of String type
type String struct {
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

func (sval *String) new(strValue string) error {
	sval.value = strValue
	return nil
}

//SetValue - helper function for parsing xml
func (sval String) SetValue(fieldName string, fieldStrValue string) (err error) {
	switch fieldName {
	case "value":
		sval.value = fieldStrValue
	case "id":
		sval.id, err = strconv.Atoi(fieldStrValue)
	default:
		err = fmt.Errorf("Unknown field '%s' for String value", fieldName)
	}
	return err
}

//NewString - returns pointer to a new String struct, or <nil> for an error
func NewString(v string, i int) *String {
	return &String{v, 0, i}
}

// func (sval *String) assign(that Value) error {
// 	sval.defaultValueImpl.assign(that)
// 	if that.getType() == STRING {
// 		t, ok := that.(*String)
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

// func (sval *String) compareTo(value2 value) (int, error) {
// 	res := 0
// 	toCompare := value2.String()
// 	if sval.value > toCompare {
// 		res = 1
// 	} else if sval.value < toCompare {
// 		res = -1
// 	}
// 	return res, nil
// }
// func (sval *String) toLong() (int64, error) {
// 	return strconv.ParseInt(sval.value, 10, 64)
// }

func (sval String) String() string {
	return sval.value
}

// func (sval *String) toBigDecimal() (float64, error) {
// 	return strconv.ParseFloat(sval.value, 64)
// }

// func (sval *String) toTime() (int, error) {
// 	return vuStringToMinutes(sval.value)
// }
// func (sval String) toDate() ([]int, error) {
// 	if y, m, d, e := vuStringToDate(sval.value); e == nil {
// 		return []int{y, m, d}, e
// 	}
// 	return nil, fmt.Errorf("Cannot convert \" %s\" to Date", sval.value)
// }
// func (*String) getType() VarType {
// 	return VarString
// }

///////////
func (sval String) isFileName() bool {
	return (sval.fileNameType != fnUndefined)
}
func (sval String) getFileNameType() FileNameType {
	return sval.fileNameType
}
func (sval String) getRecordingID() int {
	return sval.id
}
