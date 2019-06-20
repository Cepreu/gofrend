package vars

import (
	"errors"
	"fmt"
	"strconv"
)

// NumericVar - Numeric variable
type NumericVar struct {
	Variable
	value *Numeric
}

// Numeric - value of KVList type
type Numeric struct {
	Value float64
}

//NewNumeric - returns pointer to a new Numeric value struct, or <nil> for an error
func NewNumeric(v float64) *Numeric {
	return &Numeric{v}
}

//SetValue - helper function for parsing xml
func (fval *Numeric) SetValue(fieldName string, fieldStrValue string) (err error) {
	switch fieldName {
	case "value":
		fval.Value, err = strconv.ParseFloat(fieldStrValue, 64)
	default:
		err = fmt.Errorf("Unknown field '%s' for Numeric value", fieldName)
	}
	return err
}

func (fval *Numeric) new(strValue string) error {
	f64, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return errors.New("Cannot convert string to float64")
	}
	fval.Value = f64
	return nil
}

func (fval *Numeric) String() string {
	return strconv.FormatFloat(fval.Value, 'f', 2, 64)
}

func (*Numeric) getType() VarType {
	return VarNumeric
}

// func (fval *Numeric) assign(that Value) error {
// 	fval.defaultValueImpl.assign(that)
// 	v, err := that.toBigDecimal()
// 	fval.value = v
// 	return err
// }

// func (fval *Numeric) compareTo(value2 Value) (int, error) {
// 	res := 0
// 	toCompare, err := value2.toBigDecimal()
// 	if err != nil {
// 		return res, errors.New("Variable to compare must be of Numeric type")
// 	}
// 	if fval.value > toCompare {
// 		res = 1
// 	} else if fval.value < toCompare {
// 		res = -1
// 	}
// 	return res, nil
// }

// func (fval *Numeric) toLong() (int64, error) {
// 	if fval.value > 0.0 { // UpRounded
// 		return int64(fval.value + 1.0), nil
// 	}
// 	return int64(fval.value), nil
// }

// func (fval *Numeric) toBigDecimal() (float64, error) {
// 	return float64(fval.value), nil
// }

///////////
// func getSumN(args []Value) (Numeric, error) {
// 	var val float64
// 	scr := false

// 	for _, v := range args {
// 		add, err := v.toBigDecimal()
// 		if err != nil {
// 			return Numeric{}, err
// 		}
// 		val += add
// 		scr = scr || v.isSecure()
// 	}

// 	return Numeric{defaultValueImpl{scr}, val}, nil
// }

// func getDifferenceN(arg1 Value, arg2 Value) (Numeric, error) {
// 	v1, err := arg1.toBigDecimal()
// 	if err == nil {
// 		if v2, err := arg2.toBigDecimal(); err == nil {
// 			return Numeric{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 - v2}, nil
// 		}
// 	}
// 	return Numeric{}, err
// }

// func getProductN(args []Value) (Numeric, error) {

// 	value := 1.0
// 	secure := false

// 	for _, v := range args {
// 		prod, err := v.toBigDecimal()
// 		if err != nil {
// 			return Numeric{}, err
// 		}
// 		value *= prod
// 		secure = secure || v.isSecure()
// 	}

// 	return Numeric{defaultValueImpl{secure}, value}, nil

// }

// func getQuotationN(arg1 Value, arg2 Value) (Numeric, error) {
// 	v1, err := arg1.toBigDecimal()
// 	if err == nil {
// 		if v2, err := arg2.toBigDecimal(); err == nil {
// 			if v2 == 0 {
// 				return Numeric{}, errors.New("Division by zero")
// 			}
// 			return Numeric{defaultValueImpl{arg1.isSecure() || arg2.isSecure()}, v1 / v2}, nil
// 		}
// 	}
// 	return Numeric{}, err
// }
