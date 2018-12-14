package ivrparser

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
)

const defaultLang = "Default"

type promptID string

type modulePrompts []attemptPrompts

type attemptPrompts struct {
	LangPrArr []languagePrompts
	Count     int
}

type languagePrompts struct {
	PrArr    []promptID
	Language langCode
}

func newModulePrompts(c int, p []promptID) (modulePrompts, error) {
	return modulePrompts{newAttemptPrompts(c, p)}, nil
}
func newAttemptPrompts(c int, p []promptID) attemptPrompts {
	pmp := attemptPrompts{Count: c, LangPrArr: []languagePrompts{{Language: defaultLang, PrArr: p}}}
	return pmp
}

func (s *IVRScript) normalizeAttemptPrompt(ap *attemptPrompts) error {
	var prsdef = []promptID{}
	for _, l := range s.Languages {
		ap.LangPrArr = append(ap.LangPrArr, languagePrompts{Language: l.Lang, PrArr: nil})
	}
	for _, pid := range ap.LangPrArr[0].PrArr {
		if _, found := s.Prompts[pid]; found {
			for j := range s.Languages {
				ap.LangPrArr[j+1].PrArr = append(ap.LangPrArr[j+1].PrArr, pid)
			}
			prsdef = append(prsdef, pid)
		} else {
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
		}
	}
	ap.LangPrArr[0].PrArr = prsdef
	return nil
}
func (s *IVRScript) normalizePrompt(mp modulePrompts) error {
	for i := range mp {
		s.normalizeAttemptPrompt(&mp[i])
	}
	return nil
}

var counter int32

func getPromptID(prefix string, t string) promptID {
	counter++
	return promptID(fmt.Sprintf("%s%s%d", prefix, t, counter))
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

type ttsPrompt struct {
	TTSPromptXML    string
	PromptTTSEnumed bool
}

func (t ttsPrompt) TransformToAI() string { return t.TTSPromptXML }

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

func (s *IVRScript) parseVoicePromptS(decoder *xml.Decoder, v *xml.StartElement, prefix string) (attemptPrompts, error) {
	var lastElement string
	var (
		PrArr []promptID
		count int
	)
	if v != nil {
		lastElement = v.Name.Local
	}
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return attemptPrompts{}, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "prompt" {
				PrArr, _ = s.parseVoicePrompt(decoder, &v, prefix)
			} else if v.Name.Local == "count" {
				innerText, err := decoder.Token()
				if err == nil {
					count, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			}

		case xml.EndElement:
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return newAttemptPrompts(count, PrArr), nil
}

func (s *IVRScript) parseVoicePrompt(decoder *xml.Decoder, v *xml.StartElement, prefix string) (pids []promptID, err error) {
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
				s.Prompts[pid] = fp
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
				s.Prompts[pid] = pp
			} else if v.Name.Local == "multiLanguagesPromptItem" {
				var pp xMultiLanguagesPromptItem
				err = decoder.DecodeElement(&pp, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement(multiLanguagesPromptItem>) failed with '%s'\n", err)
					break
				}
				pids = append(pids, promptID(pp.MLPromptID))
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
				p, err := cmdUnzip(string(v))
				if err == nil {
					pid := getPromptID(prefix, "T")
					pids = append(pids, pid)
					var pp prompt = &ttsPrompt{TTSPromptXML: p}
					s.Prompts[pid] = pp
				}
			}
		}
	}
	return pids, nil
}

func cmdUnzip(encoded string) (string, error) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	base64.StdEncoding.Decode(base64Text, []byte(encoded))

	grepCmd := exec.Command("gunzip")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write(base64Text)
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	//	fmt.Println("> grep hello")
	return string(grepBytes), nil
}
