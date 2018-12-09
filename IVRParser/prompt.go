package ivrparser

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type promptID string

/*
type modulePrompt map[language][]promptID

type moduleBigPrompt struct {
	PromptIDArr modulePrompt
	Count       int
}
*/
type bigTempPrompt struct {
	PrArr []promptID
	Count int
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
