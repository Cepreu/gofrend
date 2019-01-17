package ivrparser

import (
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

//IVRScript - parsed IVR script
type IVRScript struct {
	Domain          string
	Properties      string
	Modules         []Module
	ModulesOnHangup []Module
	Prompts         scriptPrompts
	MLPrompts       []*multilingualPrompt
	MLChoices       []*multilanguageMenuChoice
	Variables       variables
	Languages       []language
}
type scriptPrompts map[promptID]prompt

func newIVRScript() *IVRScript {
	return &IVRScript{
		Prompts:   make(scriptPrompts),
		Variables: make(variables),
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
		inDomainID      = false
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
			if v.Name.Local == "domainId" {
				inDomainID = true
			} else if v.Name.Local == "modules" {
				s.Modules = s.parseModules(decoder, &v)
			} else if v.Name.Local == "modulesOnHangup" {
				s.ModulesOnHangup = s.parseModules(decoder, &v)
			} else if v.Name.Local == "userVariables" {
				s.Variables.parse(decoder)
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
		case xml.CharData:
			if inDomainID {
				s.Domain = string(v)
			}
		case xml.EndElement:
			if v.Name.Local == "multiLanguagesVIVRPrompts" {
				//				inMLVIVRPrompts = false
			} else if v.Name.Local == "multiLanguagesPrompts" {
				inMLPrompts = false
			} else if v.Name.Local == "multiLanguagesMenuChoices" {
				inMLMenuChoices = false
			} else if v.Name.Local == "domainId" {
				inDomainID = false
			}
		}
	}
	s.finalization()
	///////TBD - debug, delete/////
	for vname, vval := range s.Variables {
		fmt.Println("=====", vname, vval)
	}

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
