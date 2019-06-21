package vars

import (
	"errors"
	"fmt"
	"strconv"
)

// CurrencyVar - Currency variable
type CurrencyVar struct {
	Variable
	value *Currency
}

// Currency - value of Currrency type
type Currency struct {
	value float64
}

func (fval *Currency) new(strValue string) error {
	return fval.SetValue("value", strValue)
}

//SetValue - helper function for parsing xml
func (fval Currency) SetValue(fieldName string, fieldStrValue string) (err error) {
	switch fieldName {
	case "value":
		f64, err := strconv.ParseFloat(fieldStrValue, 64)
		if err != nil {
			return errors.New("Cannot convert string to float64")
		}
		fval.value = f64
	default:
		err = fmt.Errorf("Unknown field '%s' for Currency value", fieldName)
	}
	return err
}

// NewCurrency - returns pointer to a new Numeric value struct, or <nil> for an error
func NewCurrency(v float64) *Currency {
	return &Currency{v}
}

func (fval *Currency) String() string {
	str, _ := fval.convertToString()
	return fmt.Sprintf("{type=Currency}{value=%s}", str)
}

func (fval *Currency) convertToString() (string, error) {
	return strconv.FormatFloat(fval.value, 'f', 2, 64), nil
}
func (*Currency) getType() VarType {
	return VarCurrency
}
