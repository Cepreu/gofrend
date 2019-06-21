package ivr

import "github.com/Cepreu/gofrend/ivr/vars"

type Parametrized struct {
	VariableName string
	Value        vars.Value
}

func (p *Parametrized) IsVarSelected() bool { return p.VariableName != "" }

type KeyValueParametrized struct {
	Key   string
	Value *Parametrized
}
