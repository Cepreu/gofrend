package ivr

// OutputBranch - the Menu,IfElse, and Case modules' branch
type OutputBranch struct {
	Name       string
	Descendant ModuleID
	Cond       *ComplexCondition
}

// ComplexCondition - a combinedcondition
type ComplexCondition struct {
	CustomCondition string
	Conditions      []*Condition
}

//Condition - ifElse module's condition
type Condition struct {
	ComparisonType string
	RightOperand   VariableID
	LeftOperand    VariableID
}
