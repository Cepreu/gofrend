package variables

import (
	"errors"
	"fmt"
	"strconv"
)

type StringValue struct {
	defaultValueImpl
	value        string
	fileNameType FileNameType
	id           int64
}
type FileNameType int8

const (
	FN_UNDEFINED FileNameType = iota
	FN_GREETING
	FN_VOICEMAIL
	FN_RECORDING
)

func (sval *StringValue) isSecure() bool { return sval.secure }

func (sval *StringValue) assign(that IVRValue) error {
	sval.defaultValueImpl.assign(that)
	if that.getType() == STRING {
		t, ok := that.(*StringValue)
		if !ok {
			return fmt.Errorf("Corrupted IVRValue: %v", that)
		}
		if t.isFileName() {
			sval.fileNameType = t.fileNameType
			sval.id = t.id
			sval.value = fmt.Sprintf("%v, id=%v)", sval.fileNameType, sval.id)
			return nil
		}
	}
	v, err := that.convertToString()
	if err == nil {
		sval.value = v
	}
	return err
}

func (sval *StringValue) new(secure bool, strValue string) error {
	sval.secure = secure
	sval.value = strValue
	return nil
}

func (sval *StringValue) compareTo(value2 IVRValue) (int, error) {
	if value2.getType() == NUMERIC {
		return value2.compareTo(sval)
	}
	res := 0
	toCompare, err := value2.convertToString()
	if err != nil {
		return res, errors.New("Variable to compare must be of StringValue type")
	}
	if sval.value > toCompare {
		res = 1
	} else if sval.value < toCompare {
		res = -1
	}
	return res, nil
}

func (sval *StringValue) toLong() (int64, error) {
	return strconv.ParseInt(sval.value, 10, 64)
}

func (sval *StringValue) toString() string {
	str := "*****"
	if !sval.secure {
		str, _ = sval.convertToString()
	}
	return fmt.Sprintf("{type=StringValue}{value=%s}", str)
}

func (sval *StringValue) convertToString() (string, error) {
	return sval.value, nil
}

func (sval *StringValue) toBigDecimal() (float64, error) {
	return strconv.ParseFloat(sval.value, 64)
}

func (*StringValue) getType() Type {
	return STRING
}

///////////
func (sval StringValue) isFileName() bool {
	return (sval.fileNameType != FN_UNDEFINED)
}
func (sval StringValue) getFileNameType() FileNameType {
	return sval.fileNameType
}
func (sval StringValue) getRecordingID() int64 {
	return sval.id
}
