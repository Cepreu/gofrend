package ivrparser

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

type PromptID string

type xPrompts struct {
	Prompt xPrompt `xml:"prompt"`
	Count  int     `xml:"count"`
}
type xPrompt struct {
	//	TTSes                        []xTTSPrompt   `xml:"ttsPrompt"`
	//	Files                        []xFilePrompt  `xml:"filePrompt"`
	//	Pauses                       []xPausePrompt `xml:"pausePrompt"`
	FullPrompt                   string `xml:",innerxml"`
	Interruptible                bool   `xml:"interruptible"`
	CanChangeInterruptableOption bool   `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool   `xml:"ttsEnumed"`
	ExitModuleOnException        bool   `xml:"exitModuleOnException"`
}
type xFilePrompt struct {
	PromptDirectly     bool   `xml:"promptData>promptSelected"`
	PromptID           int32  `xml:"promptData>prompt>id"`
	PromptName         string `xml:"promptData>prompt>name"`
	PromptVariableName string `xml:"promptData>promptVariableName"`
	IsRecordedMessage  bool   `xml:"promptData>isRecordedMessage"`
}
type xPausePrompt struct {
	Timeout int32 `xml:"timeout"`
}
type xTTSPrompt struct {
	TtsPromptXML    string `xml:"xml"`
	PromptTTSEnumed bool   `xml:"promptTTSEnumed"`
}
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

type xMultilingualPrompts struct {
	Entries []xMLPromptEntry `xml:"entry"`
}
type xMLPromptEntry struct {
	Key          string     `xml:"key"`
	ID           string     `xml:"value>promptId"`
	Name         string     `xml:"value>name"`
	Description  string     `xml:"value>description"`
	Type         string     `xml:"value>type"`
	Prompts      xMLPrompts `xml:"value>prompts"`
	DefLanguage  string     `xml:"defaultLanguage"`
	IsPersistent bool       `xml:"isPersistent"`
}
type xMLPrompts struct {
	Entries []xMLanguageEntry `xml:"entry"`
}
type xMLanguageEntry struct {
	Language string         `xml:"key,attr"`
	TTSes    []xTTSPrompt   `xml:"ttsPrompt"`
	Files    []xFilePrompt  `xml:"filePrompt"`
	Pauses   []xPausePrompt `xml:"pausePrompt"`
}

type gPrompt struct{}

type prompt interface {
	TransformToAI() string
}

func parseVoicePrompt(fullPrompt io.Reader) (pids []PromptID, err error) {
	decoder := xml.NewDecoder(fullPrompt)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			// io.EOF is a successful end
			break
		}
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		inFile := false
		inPause := false
		inTTS := true
		inXML := true

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "filePrompt" {
				inFile = true
			} else if v.Name.Local == "ttsPrompt" {
				inTTS = true
			} else if v.Name.Local == "pausePrompt" {
				inPause = true
			} else if v.Name.Local == "xml" {
				inXML = true

			}
		case xml.EndElement:
			if v.Name.Local == "filePrompt" {
				inFile = false
			} else if v.Name.Local == "ttsPrompt" {
				inTTS = false
			} else if v.Name.Local == "pausePrompt" {
				inPause = false
			} else if v.Name.Local == "xml" {
				inXML = false
			}

		case xml.CharData:
			if inTTS && inXML {
				fmt.Printf("XML: %s\n", cmdUnzip(string(v))
			}

		}
	}
	return
}

func cmdUnzip(encoded string) string {
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
	fmt.Println("> grep hello")
	return string(grepBytes)
}
