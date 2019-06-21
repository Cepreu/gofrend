package ivr

import (
	"fmt"
	"strings"
)

const defaultLang = "Default"

type prompt interface {
	TransformToAI() string
}

// PromptID - Unique prompt ID
type PromptID string

// ModulePrompts - Complete prompt structure
type ModulePrompts []AttemptPrompts

// AttemptPrompts - Single-attempt prompt structure
type AttemptPrompts struct {
	LangPrArr []LanguagePrompts
	Count     int
}

// TransformToAI - Transforms to a form suitable for Dialogflow
func (mp ModulePrompts) TransformToAI(sp ScriptPrompts) (res []string) {
	for _, ap := range mp {
		res = append(res, ap.TransformToAI(sp))
	}
	return
}

// TransformToAI - Transforms to a form suitable for Dialogflow
func (ap AttemptPrompts) TransformToAI(sp ScriptPrompts) (txt string) {
	for _, id := range ap.LangPrArr[0].PrArr {
		txt = txt + sp[id].TransformToAI()
	}
	return
}

// LanguagePrompts - prompt for a specified language
type LanguagePrompts struct {
	PrArr    []PromptID
	Language LangCode
}

type TtsPrompt struct {
	TTSPromptXML    string
	PromptTTSEnumed bool
}

// TransformToAI is a recursive function that calls itself for every child
func (t TtsPrompt) TransformToAI() string {
	n, _ := parseTtsPrompt(strings.NewReader(t.TTSPromptXML))
	return n.transformToAI()
}
func (n *ttsNode) transformToAI() (s string) {
	for _, child := range n.children {
		s = strings.TrimPrefix(s+" "+child.transformToAI(), " ")
	}
	switch n.nodeType {
	case "textElement":
		s = n.body
	case "variableElement":
		s = "@" + n.variableName + "@"
	case "breakElement":
		s = "\n"
	case "sentenceElement":
		s = strings.TrimSuffix(s, " ") + "."
	case "paragraphElement":
		for {
			sz := len(s)
			if sz > 0 && (s[sz-1] == '.' || s[sz-1] == '\n' || s[sz-1] == ' ') {
				s = s[:sz-1]
			} else {
				break
			}
		}
		s += ".\n"
	}
	return s
}

type ttsNode struct {
	nodeType     string
	parent       *ttsNode
	children     []*ttsNode
	attributes   []ttsAttr
	body         string
	variableName string
}
type ttsAttr struct {
	name  string
	value string
}

type XFilePrompt struct {
	PromptDirectly     bool   `xml:"promptData>promptSelected"`
	PromptID           int32  `xml:"promptData>prompt>id"`
	PromptName         string `xml:"promptData>prompt>name"`
	PromptVariableName string `xml:"promptData>promptVariableName"`
	IsRecordedMessage  bool   `xml:"promptData>isRecordedMessage"`
}

func (t XFilePrompt) TransformToAI() string {
	if t.IsRecordedMessage {
		return fmt.Sprintf("%s", t.PromptName)
	}
	return fmt.Sprintf("%s", t.PromptVariableName)
}

type XPausePrompt struct {
	Timeout int32 `xml:"timeout"`
}

func (t XPausePrompt) TransformToAI() string { return fmt.Sprintf("%d", t.Timeout) }

type XVivrPrompts struct {
	VivrPrompts                  []XVivrPrompt  `xml:"vivrPrompt"`
	ImagePrompts                 []XImagePrompt `xml:"imagePrompt"`
	Interruptible                bool           `xml:"interruptible"`
	CanChangeInterruptableOption bool           `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool           `xml:"ttsEnumed"`
	ExitModuleOnException        bool           `xml:"exitModuleOnException"`
}

type XVivrHeader struct {
	VPrompt                      XVivrPrompt `xml:"vivrPrompt"`
	Interruptible                bool        `xml:"interruptible"`
	CanChangeInterruptableOption bool        `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool        `xml:"ttsEnumed"`
	ExitModuleOnException        bool        `xml:"exitModuleOnException"`
}

type XVivrPrompt struct {
	VivrXML string `xml:"xml"`
}

type XImagePrompt struct {
	ImageURL           string `xml:"imageURL"`
	IsVariableSelected bool   `xml:"isVariableSelected"`
}

type XTextChanneddata struct {
	TextPrompts       XTextPrompts `xml:"textPrompts"`
	IsUsedVivrPrompts bool         `xml:"isUsedVivrPrompts"`
	IsTextOnly        bool         `xml:"isTextOnly"`
}

type XTextPrompts struct {
	SimpleTextPromptItem         XSimpleTextPromptItem `xml:"simpleTextPromptItem"`
	Interruptible                bool                  `xml:"interruptible"`
	CanChangeInterruptableOption bool                  `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool                  `xml:"ttsEnumed"`
	ExitModuleOnException        bool                  `xml:"exitModuleOnException"`
}

type XSimpleTextPromptItem struct {
	TextXML string `xml:"xml"`
}

type LangCode string

type MultilanguageMenuChoice struct {
	ID          string
	Name        string
	Description string
	Type        string
	AudPrompts  map[LangCode][]PromptID
	VisPrompts  map[LangCode][]PromptID
	TxtPrompts  map[LangCode][]PromptID
	DefLanguage LangCode
	//	IsPersistent bool       `xml:"isPersistent"`
}
