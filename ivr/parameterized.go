package ivr

import "github.com/Cepreu/gofrend/ivr/vars"

type parametrized struct {
	VariableName string
	Value        vars.Value
}

func (p *parametrized) IsVarSelected() bool { return p.VariableName != "" }

type keyValueParametrized struct {
	Key   string
	Value *parametrized
}
