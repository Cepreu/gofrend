package variables

type IVRValue interface {
	compareTo(IVRValue) (int, error)
	assign(IVRValue) error
	new(bool, string) error
	toString() string
	convertToString() (string, error)
	toLong() (int64, error)
	toBigDecimal() (float64, error)
	toDate() ([]int32, error)
	toTime() ([]int32, error)
	isSecure() bool
	getType() Type
}
