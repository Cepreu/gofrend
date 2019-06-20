package ivr

import "github.com/Cepreu/gofrend/ivr/vars"

type IVRScript struct {
	Domain          string
	Properties      string
	Modules         map[ModuleID]Module
	ModulesOnHangup map[ModuleID]Module
	Prompts         ScriptPrompts
	MLPrompts       []*multilingualPrompt
	MLChoices       []*multilanguageMenuChoice
	Variables       variables
	Languages       []language
	Menus           []ModuleID
}

type ScriptPrompts map[PromptID]prompt

type variables map[string]*vars.Variable

type xUserVariable struct {
	Key           string `xml:"key"`
	Name          string `xml:"value>name"`
	Description   string `xml:"value>description"`
	StringValue   string `xml:"value>stringValue>value"`
	StringValueID int32  `xml:"value>stringValue>id"`
	Attributes    int32  `xml:"value>attributes"`
	IsNullValue   bool   `xml:"isNullValue"`
}

type languages struct {
	Langs []language `xml:"languages"`
}
type language struct {
	Lang     langCode `xml:"lang"`
	TtsLang  langCode `xml:"ttsLanguage"`
	TtsVoice string   `xml:"ttsVoice"`
}
