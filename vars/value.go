package vars

import (
	"fmt"
)

// type comparator interface {
// 	CompareTo(value) (int, error)
// }

// type logger interface {
// 	Log(value) (int, error)
// }

type Value interface {
	new(bool, string) error
	isSecure() bool
	getType() VarType
	fmt.Stringer
}

// func (ival *defaultValueImpl) isSecure() bool {
// 	return ival.secure
// }

// func (ival *defaultValueImpl) assign(that *Value) error {
// 	ival.secure = (*that).isSecure()
// 	return nil
// }

// func (ival *defaultValueImpl) new(secure bool, strValue string) error {
// 	ival.secure = secure
// 	return nil
// }
