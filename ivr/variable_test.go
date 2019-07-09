package ivr

import (
	"testing"
)

func TestValues(t *testing.T) {
	var (
		pVal     []*Value
		err      []error
		expected []*Value
	)
	p, e := NewStringValue("hello, test!")
	pVal = append(pVal, p)
	err = append(err, e)
	expected = append(expected, &Value{ValString, "hello, test!"})

	p, e = NewDateValue(2017, 2, 28)
	pVal = append(pVal, p)
	err = append(err, e)
	expected = append(expected, &Value{ValDate, "2017-02-28"})

	p, e = NewDateValue(17, 2, 28)
	pVal = append(pVal, p)
	err = append(err, e)
	expected = append(expected, nil)

	p, e = NewTimeValue(256)
	pVal = append(pVal, p)
	err = append(err, e)
	expected = append(expected, &Value{ValTime, "04:16"})

	p, e = NewEUCurrencyValue(1000.45)
	pVal = append(pVal, p)
	err = append(err, e)
	expected = append(expected, &Value{ValCurrencyEuro, "EU$1000.45"})

	p, e = NewKeyValue("{\"testkey\":\"testvalue\"}")
	pVal = append(pVal, p)
	err = append(err, e)
	expected = append(expected, &Value{ValKVList, `{"testkey":"testvalue"}`})

	for i, ex := range expected {
		if err[i] != nil {
			if ex != nil {
				t.Errorf("\nVariables: %d was not created: \"%s\", expected %s", i, err[i], expected[i])
			}
		} else if pVal[i].StringValue != ex.StringValue {
			t.Errorf("\nVariables[%d]: %v was expected, in reality: %v", i, expected[i], pVal[i])
		}
	}
}
