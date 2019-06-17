package ivr

// OutputBranch - the Menu,IfElse, and Case modules' branch
type OutputBranch struct {
	Name       string
	Descendant ModuleID
	Cond       *ComplexCondition
}

// ComplexCondition - a combinedcondition
type ComplexCondition struct {
	CustomCondition   string
	ConditionGrouping groupingType
	Conditions        []*Condition
}

//Condition - ifElse module's condition
type Condition struct {
	comparisonType string
	rightOperand   parametrized
	leftOperand    parametrized
}
