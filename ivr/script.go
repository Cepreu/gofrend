package ivr

type (
	IVRScript struct {
		Domain          string
		Properties      string
		Modules         map[ModuleID]Module
		ModulesOnHangup map[ModuleID]Module
		Prompts         ScriptPrompts
		MLChoices       []*MultilanguageMenuChoice
		Variables       Variables
		Input           []VariableID
		Output          []VariableID
		Languages       []Language
		Functions       []*Function
		Menus           []ModuleID
	}

	ScriptPrompts map[PromptID]prompt

	Variables map[VariableID]*Variable

	Languages struct {
		Langs []Language `xml:"languages"`
	}
	Language struct {
		Lang     LangCode `xml:"lang"`
		TtsLang  LangCode `xml:"ttsLanguage"`
		TtsVoice string   `xml:"ttsVoice"`
	}

	//FuncID - ID of the function
	FuncID string
	//FuncType - Type of the function (build-in, javascript, or copy operator)
	FuncType int
	Function struct {
		ID          FuncID
		Type        FuncType
		Description string
		ReturnType  string //varType
		Name        string
		Arguments   []*FuncArgument
		Body        string
	}
	FuncArgument struct {
		Name        string
		Description string
		ArgType     string
	}

	FuncInvocation struct {
		FuncDef FuncID
		Params  []VariableID
	}
)

const (
	FuncUndefined FuncType = iota
	FuncStandard
	FuncJS
	FuncCopy
)
