package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/Cepreu/gofrend/utils"
)

const defaultLang = "Default"

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

func newModulePrompts(c int, p []PromptID) (ModulePrompts, error) {
	return ModulePrompts{newAttemptPrompts(c, p)}, nil
}
func newAttemptPrompts(c int, p []PromptID) AttemptPrompts {
	pmp := AttemptPrompts{Count: c, LangPrArr: []LanguagePrompts{{Language: defaultLang, PrArr: p}}}
	return pmp
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

func (s *IVRScript) normalizePrompt(mp ModulePrompts) error {
	for i := range mp {
		s.normalizeAttemptPrompt(&mp[i], true)
	}
	return nil
}

var counter int32

func getPromptID(prefix string, t string) PromptID {
	counter++
	return PromptID(fmt.Sprintf("%s_%s_%d", prefix, t, counter))
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

type prompt interface {
	TransformToAI() string
}

func parseVoicePromptS(decoder *xml.Decoder, sp ScriptPrompts, prefix string) (AttemptPrompts, error) {
	var (
		PrArr []PromptID
		count int
	)
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return AttemptPrompts{}, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "prompt" {
				PrArr, err = parseVoicePrompt(decoder, &v, sp, prefix)
				if err != nil {
					fmt.Printf("parseVoicePrompt failed with '%s'\n", err)
					return AttemptPrompts{}, err
				}

			} else if v.Name.Local == "count" {
				innerText, err := decoder.Token()
				if err == nil {
					count, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			}

		case xml.EndElement:
			if v.Name.Local == "prompts" {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return newAttemptPrompts(count, PrArr), nil
}

func parseVoicePrompt(decoder *xml.Decoder, v *xml.StartElement, sp ScriptPrompts, prefix string) (pids []PromptID, err error) {
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	inTTS, inXML := false, false
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "filePrompt" {
				var fp xFilePrompt
				err = decoder.DecodeElement(&fp, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement(filePrompt) failed with '%s'\n", err)
					break
				}
				pid := getPromptID(prefix, "F")
				pids = append(pids, pid)
				sp[pid] = fp
			} else if v.Name.Local == "ttsPrompt" {
				inTTS = true
			} else if v.Name.Local == "pausePrompt" {
				var pp xPausePrompt
				err = decoder.DecodeElement(&pp, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement(pausePrompt) failed with '%s'\n", err)
					break
				}
				pid := getPromptID(prefix, "P")
				pids = append(pids, pid)
				sp[pid] = pp
			} else if v.Name.Local == "multiLanguagesPromptItem" {
				var pp xMultiLanguagesPromptItem
				err = decoder.DecodeElement(&pp, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement(multiLanguagesPromptItem>) failed with '%s'\n", err)
					break
				}
				pids = append(pids, PromptID(pp.MLPromptID))
			} else if v.Name.Local == "xml" {
				inXML = true
			}

		case xml.EndElement:
			if v.Name.Local == "ttsPrompt" {
				inTTS = false
			} else if v.Name.Local == "xml" {
				inXML = false
			} else if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return is HERE!
			}
		case xml.CharData:
			if inTTS && inXML {
				p, err := utils.CmdUnzip(string(v))
				if err == nil {
					pid := getPromptID(prefix, "T")
					pids = append(pids, pid)
					var pp prompt = &TtsPrompt{TTSPromptXML: p}
					sp[pid] = pp
				}
			}
		}
	}
	return pids, nil
}
