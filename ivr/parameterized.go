package ivr

type Parametrized struct {
	VariableName string
	Value        *Value
}

func (p *Parametrized) IsVarSelected() bool { return p.VariableName != "" }

type KeyValueParametrized struct {
	Key   string
	Value *Parametrized
}
