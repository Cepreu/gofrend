package variables

type IVRValue interface {
	compareTo(Value) (uint8, error) 
	assign(Value) 
	isEmpty() bool
	toString()
	convertToString() (string, typeCastError)
	toLong() (int64, error)
	toBigDecimal() (float64, typeCastError)
	toDate() (date, typeCastError)
	toTime() (time, typeCastError)
}

    {
        throw new TypeCastException("Cannot convert " + 
                                    this.getClass().getSimpleName() + 
                                    " to Integer");
    }
 
    func ToString()
    {
        return "{type=" + getType().getLogName() + "}{value=" + SecureValue.create(secure, convertToString()) + "}";
    }