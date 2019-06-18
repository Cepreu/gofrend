package ivr

import "github.com/Cepreu/gofrend/ivr/vars"

type (
	IVRScript struct {
		Domain          string
		Properties      string
		Modules         map[ModuleID]Module
		ModulesOnHangup map[ModuleID]Module
		Prompts         ScriptPrompts
		//		MLPrompts       []*MultilingualPrompt
		MLChoices   []*MultilanguageMenuChoice
		Variables   Variables
		Languages   []Language
		JSFunctions []*JsFunction
		Menus       []ModuleID
	}

	ScriptPrompts map[PromptID]prompt

	Variables map[string]*vars.Variable

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
