package vars

import (
	"errors"
	"fmt"
	"strconv"
)

type CurrencyValue struct {
	secure bool
	value  float64
}

//NewCurrencyValue - returns pointer to a new Numeric value struct, or <nil> for an error
func NewCurrencyValue(v float64) *CurrencyValue {

	return &CurrencyValue{secure: false, value: v}
}

func (fval *CurrencyValue) isSecure() bool { return fval.secure }

func (fval *CurrencyValue) String() string {
	str := "*****"
	if !fval.secure {
		str, _ = fval.convertToString()
	}
	return fmt.Sprintf("{type=CurrencyValue}{value=%s}", str)
}

func (fval *CurrencyValue) new(secure bool, strValue string) error {
	fval.secure = secure
	f64, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return errors.New("Cannot convert string to float64")
	}
	fval.value = f64
	return nil
}
func (fval *CurrencyValue) convertToString() (string, error) {
	return strconv.FormatFloat(fval.value, 'f', 2, 64), nil
}
func (*CurrencyValue) getType() VarType {
	return VarCurrency
}
