package variables

import (
	"errors"
	"fmt"
	"strconv"
)

type StringValue struct {
	defaultValueImpl
	value  string
        FileNameType fileNameType
	id int64
}
type FileNameType const (
	GREETING iota
        VOICEMAIL
	RECORDING
)

func (sval *StringValue) isSecure() bool { return sval.secure }

func (sval *StringValue) assign(that IVRValue) error {
	sval.defaultValueImpl.assign(that)
	v, err := that.toLong()
	sval.value = v
	return err
}

func (sval *StringValue) new(secure bool, strValue string) error {
	sval.secure = secure
	i64, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return errors.New("Cannot convert string to long")
	}
	sval.value = i64
	return nil
}

func (sval *StringValue) compareTo(value2 IVRValue) (int, error) {
	if value2.getType() == NUMERIC {
		return value2.compareTo(sval)
	}
	res := 0
	toCompare, err := value2.toLong()
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
	return sval.value, nil
}

func (sval *StringValue) toString() string {
	str := "*****"
	if !sval.secure {
		str, _ = sval.convertToString()
	}
	return fmt.Sprintf("{type=StringValue}{value=%s}", str)
}

func (sval *StringValue) convertToString() (string, error) {
	return strconv.FormatInt(sval.value, 10), nil
}

func (sval *StringValue) toBigDecimal() (float64, error) {
	return float64(sval.value), nil
}

func (*StringValue) getType() Type {
	return STRING
}

///////////
func (sval StringValue)isFileName() bool {
    return sval.fileNameType != nil
}
func (sval StringValue) getFileNameType() FileNameType {
     return fileNameType;
}
func (sval StringValue) getRecordingId() long {
     return id;
}

