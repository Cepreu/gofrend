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
	Language langCode
}

type xFilePrompt struct {
	PromptDirectly     bool   `xml:"promptData>promptSelected"`
	PromptID           int32  `xml:"promptData>prompt>id"`
	PromptName         string `xml:"promptData>prompt>name"`
	PromptVariableName string `xml:"promptData>promptVariableName"`
	IsRecordedMessage  bool   `xml:"promptData>isRecordedMessage"`
}

func (t xFilePrompt) TransformToAI() string {
	if t.IsRecordedMessage {
		return fmt.Sprintf("%s", t.PromptName)
	}
	return fmt.Sprintf("%s", t.PromptVariableName)
}

type xPausePrompt struct {
	Timeout int32 `xml:"timeout"`
}

func (t xPausePrompt) TransformToAI() string { return fmt.Sprintf("%d", t.Timeout) }

type xMultiLanguagesPromptItem struct {
	MLPromptID string `xml:"prompt"`
}

func (t xMultiLanguagesPromptItem) TransformToAI() string { return string(t.MLPromptID) }

type xVivrPrompts struct {
	VivrPrompts                  []xVivrPrompt  `xml:"vivrPrompt"`
	ImagePrompts                 []xImagePrompt `xml:"imagePrompt"`
	Interruptible                bool           `xml:"interruptible"`
	CanChangeInterruptableOption bool           `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool           `xml:"ttsEnumed"`
	ExitModuleOnException        bool           `xml:"exitModuleOnException"`
}

type xVivrHeader struct {
	VPrompt                      xVivrPrompt `xml:"vivrPrompt"`
	Interruptible                bool        `xml:"interruptible"`
	CanChangeInterruptableOption bool        `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool        `xml:"ttsEnumed"`
	ExitModuleOnException        bool        `xml:"exitModuleOnException"`
}

type xVivrPrompt struct {
	VivrXML string `xml:"xml"`
}

type xImagePrompt struct {
	ImageURL           string `xml:"imageURL"`
	IsVariableSelected bool   `xml:"isVariableSelected"`
}

type xTextChanneddata struct {
	TextPrompts       xTextPrompts `xml:"textPrompts"`
	IsUsedVivrPrompts bool         `xml:"isUsedVivrPrompts"`
	IsTextOnly        bool         `xml:"isTextOnly"`
}

type xTextPrompts struct {
	SimpleTextPromptItem         xSimpleTextPromptItem `xml:"simpleTextPromptItem"`
	Interruptible                bool                  `xml:"interruptible"`
	CanChangeInterruptableOption bool                  `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool                  `xml:"ttsEnumed"`
	ExitModuleOnException        bool                  `xml:"exitModuleOnException"`
}

type xSimpleTextPromptItem struct {
	TextXML string `xml:"xml"`
}

type langCode string

type multilingualPrompt struct {
	ID          string
	Name        string
	Description string
	Type        string
	Prompts     map[langCode][]PromptID
	DefLanguage langCode
	//	IsPersistent bool       `xml:"isPersistent"`
}

type multilanguageMenuChoice struct {
	ID          string
	Name        string
	Description string
	Type        string
	AudPrompts  map[langCode][]PromptID
	VisPrompts  map[langCode][]PromptID
	TxtPrompts  map[langCode][]PromptID
	DefLanguage langCode
	//	IsPersistent bool       `xml:"isPersistent"`
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

func (s *IVRScript) normalizePrompt(mp ModulePrompts) error {
	for i := range mp {
		s.normalizeAttemptPrompt(&mp[i], true)
	}
	return nil
}

func (s *IVRScript) normalizeAttemptPrompt(ap *AttemptPrompts, mlPrompTrueMlItemFalse bool) error {
	var prsdef = []PromptID{}
	for _, l := range s.Languages {
		ap.LangPrArr = append(ap.LangPrArr, LanguagePrompts{Language: l.Lang, PrArr: nil})
	}
	for _, pid := range ap.LangPrArr[0].PrArr {
		if _, found := s.Prompts[pid]; found {
			for j := range s.Languages {
				ap.LangPrArr[j+1].PrArr = append(ap.LangPrArr[j+1].PrArr, pid)
			}
			prsdef = append(prsdef, pid)
		} else if mlPrompTrueMlItemFalse {
			for k := range s.MLPrompts {
				if s.MLPrompts[k].ID == string(pid) {
					// Found!
					for j, l := range s.Languages {
						ap.LangPrArr[j+1].PrArr = append(ap.LangPrArr[j+1].PrArr, s.MLPrompts[k].Prompts[l.Lang]...)
					}
					prsdef = append(prsdef, s.MLPrompts[k].Prompts[s.MLPrompts[k].DefLanguage]...)
					break
				}
			}
		} else {
			for k := range s.MLChoices {
				if s.MLChoices[k].ID == string(pid) {
					// Found!
					for j, l := range s.Languages {
						ap.LangPrArr[j+1].PrArr = append(ap.LangPrArr[j+1].PrArr, s.MLChoices[k].AudPrompts[l.Lang]...)
					}
					prsdef = append(prsdef, s.MLChoices[k].AudPrompts[s.MLChoices[k].DefLanguage]...)
					break
				}
			}

		}
	}
	ap.LangPrArr[0].PrArr = prsdef
	return nil
}
