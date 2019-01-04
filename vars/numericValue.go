package vars

import (
	"errors"
	"fmt"
	"strconv"
)

type NumericValue struct {
	secure bool
	value  float64
}

//NewNumericValue - returns pointer to a new Numeric value struct, or <nil> for an error
func NewNumericValue(v float64) *NumericValue {

	return &NumericValue{secure: false, value: v}
}

func (fval *NumericValue) isSecure() bool { return fval.secure }

func (fval *NumericValue) String() string {
	str := "*****"
	if !fval.secure {
		str, _ = fval.convertToString()
	}
	return fmt.Sprintf("{type=NumericValue}{value=%s}", str)
}

func (fval *NumericValue) new(secure bool, strValue string) error {
	fval.secure = secure
	f64, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return errors.New("Cannot convert string to float64")
	}
	fval.value = f64
	return nil
}

func (fval *NumericValue) convertToString() (string, error) {
	return strconv.FormatFloat(fval.value, 'f', 2, 64), nil
}

func (*NumericValue) getType() VarType {
	return VarNumeric
}

// func (fval *NumericValue) assign(that Value) error {
// 	fval.defaultValueImpl.assign(that)
// 	v, err := that.toBigDecimal()
// 	fval.value = v
// 	return err
// }

// func (fval *NumericValue) compareTo(value2 Value) (int, error) {
// 	res := 0
// 	toCompare, err := value2.toBigDecimal()
// 	if err != nil {
// 		return res, errors.New("Variable to compare must be of NumericValue type")
// 	}
// 	if fval.value > toCompare {
// 		res = 1
// 	} else if fval.value < toCompare {
// 		res = -1
// 	}
// 	return res, nil
// }

// func (fval *NumericValue) toLong() (int64, error) {
// 	if fval.value > 0.0 { // UpRounded
// 		return int64(fval.value + 1.0), nil
// 	}
// 	return int64(fval.value), nil
// }

// func (fval *NumericValue) toBigDecimal() (float64, error) {
// 	return float64(fval.value), nil
// }

///////////
// func getSumN(args []Value) (NumericValue, error) {
// 	var val float64
// 	scr := false

// 	for _, v := range args {
// 		add, err := v.toBigDecimal()
// 		if err != nil {
// 			return NumericValue{}, err
// 		}
// 		val += add
// 		scr = scr || v.isSecure()
// 	}

// 	return NumericValue{defaultValueImpl{scr}, val}, nil
// }

// func getDifferenceN(arg1 Value, arg2 Value) (NumericValue, error) {
// 	v1, err := arg1.toBigDecimal()
// 	if err == nil {
// 		if v2, err := arg2.toBigDecimal(); err == nil {
// 			return NumericValue{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 - v2}, nil
// 		}
// 	}
// 	return NumericValue{}, err
// }

// func getProductN(args []Value) (NumericValue, error) {

// 	value := 1.0
// 	secure := false

// 	for _, v := range args {
// 		prod, err := v.toBigDecimal()
// 		if err != nil {
// 			return NumericValue{}, err
// 		}
// 		value *= prod
// 		secure = secure || v.isSecure()
// 	}

// 	return NumericValue{defaultValueImpl{secure}, value}, nil

// }

// func getQuotationN(arg1 Value, arg2 Value) (NumericValue, error) {
// 	v1, err := arg1.toBigDecimal()
// 	if err == nil {
// 		if v2, err := arg2.toBigDecimal(); err == nil {
// 			if v2 == 0 {
// 				return NumericValue{}, errors.New("Division by zero")
// 			}
// 			return NumericValue{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 / v2}, nil
// 		}
// 	}
// 	return NumericValue{}, err
// }
