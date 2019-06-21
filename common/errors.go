package common

import "fmt"

type variableNotExternalErr struct {
	varName    string
	scriptName string
}

func (e *variableNotExternalErr) Error() string {
	return fmt.Sprintf("The variable \"  %s \" in IVR script \" %s \" is not external", e.varName, e.scriptName)
}

type BuiltinGrammarErrEnum int

const (
	EXPECTED_EXCEED_MAXALLOWED BuiltinGrammarErrEnum = iota
	MINALLOWED_EXCEED_MAXALLOWED
	MINEXPECTED_EXCEED_MAXEXPECTED
	MINLENGTH_EXCEED_MAXLENGTH
	BOOLEAN_PROPS_EQUALS
	CARD_LIST_EMPTY
	INVALID_URI
	RULEREF_URI_INVALID
	RULEID_INVALID
	VARIABLE_INVALID
)

type BuiltinGrammarErr struct {
	cause   BuiltinGrammarErrEnum
	grammar string
}

func (e *BuiltinGrammarErr) Error() string {
	return fmt.Sprintf("The cause: %i in the grammar: %s", e.cause, e.grammar)
}
