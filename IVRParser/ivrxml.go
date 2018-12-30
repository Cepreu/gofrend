package ivrparser

import (
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

//IVRScript - Unparsed IVR script structure
type IVRScript struct {
	Domain          int32
	Properties      string
	Modules         []Module
	ModulesOnHangup []Module
	Prompts         scriptPrompts
	MLPrompts       []*multilingualPrompt
	MLChoices       []*multilanguageMenuChoice
	//Variables map[string]*Variable
	Languages []language
}
type scriptPrompts map[promptID]prompt

func newIVRScript() *IVRScript {
	return &IVRScript{
		Prompts: make(scriptPrompts),
		//Variables: make(map[string]*Variable),
	}
}

type modulesOnHangup struct {
	StartOnHangup *incomingCallModule
	HangupModules []*hangupModule
}

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

//NewIVRScript - Parsing of the getIVRResponse received from Five9 Config web service
func NewIVRScript(src io.Reader) (*IVRScript, error) {
	s := newIVRScript()

	decoder := xml.NewDecoder(src)
	decoder.CharsetReader = charset.NewReaderLabel

	var (
		//		inVariables     = false
		inMLPrompts = false
		//		inMLVIVRPrompts = false
		inMLMenuChoices = false
	)
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

		switch v := t.(type) {

		case xml.StartElement:
			if v.Name.Local == "modules" {
				s.Modules = s.parseModules(decoder, &v)
			} else if v.Name.Local == "modulesOnHangup" {
				s.ModulesOnHangup = s.parseModules(decoder, &v)
			} else if v.Name.Local == "userVariables" {
				//				inVariables = true
			} else if v.Name.Local == "multiLanguagesVIVRPrompts" {
				//				inMLVIVRPrompts = true
			} else if v.Name.Local == "multiLanguagesPrompts" {
				inMLPrompts = true
			} else if v.Name.Local == "value" && inMLPrompts {
				mlp, err := s.parseMultilanguagePrompts(decoder)
				if err != nil {
					fmt.Printf("parseMLPrompt() failed with '%s'\n", err)
					break
				} else {
					s.MLPrompts = append(s.MLPrompts, mlp)
				}
				// mlMenuChoices <---
			} else if v.Name.Local == "multiLanguagesMenuChoices" {
				inMLMenuChoices = true
			} else if v.Name.Local == "value" && inMLMenuChoices {
				mlp, err := s.parseMultilanguageMenuElements(decoder)
				if err != nil {
					fmt.Printf("parseMLPrompt() failed with '%s'\n", err)
					break
				} else {
					s.MLChoices = append(s.MLChoices, mlp)
				}

			} else if v.Name.Local == "languages" {
				var m languages
				err := decoder.DecodeElement(&m, &v)
				if err == nil {
					if len(m.Langs) > 0 {
						s.Languages = m.Langs
					} else {
						s.Languages = []language{{Lang: "en-US", TtsLang: "en-US", TtsVoice: "Samanta"}}
					}
				}
			}
		case xml.EndElement:
			if v.Name.Local == "userVariables" {
				//				inVariables = false
			} else if v.Name.Local == "multiLanguagesVIVRPrompts" {
				//				inMLVIVRPrompts = false
			} else if v.Name.Local == "multiLanguagesPrompts" {
				inMLPrompts = false
			} else if v.Name.Local == "multiLanguagesMenuChoices" {
				inMLMenuChoices = false
			}
		}
	}
	s.finalization()
	return s, nil
}

func (s *IVRScript) finalization() error {
	for _, module := range s.Modules {
		module.normalize(s)
	}
	for _, module := range s.ModulesOnHangup {
		module.normalize(s)
	}
	return nil
}
