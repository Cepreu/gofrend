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
		JSFunctions     []*JsFunction
		Menus           []ModuleID
	}

	ScriptPrompts map[PromptID]prompt

	Variables map[string]*Variable

	Languages struct {
		Langs []Language `xml:"languages"`
	}
	Language struct {
		Lang     LangCode `xml:"lang"`
		TtsLang  LangCode `xml:"ttsLanguage"`
		TtsVoice string   `xml:"ttsVoice"`
	}
	JsFunction struct {
		JsFunctionID string
		Description  string
		ReturnType   string //varType
		Name         string
		Arguments    []*FuncArgument
		FuncBody     string
	}
	FuncArgument struct {
		Name        string
		Description string
		ArgType     string
	}
)
