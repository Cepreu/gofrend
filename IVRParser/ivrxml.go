package ivrparser

import (
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

type moduleID string

//IVRScript - Unparsed IVR script structure
type IVRScript struct {
	Domain          int32
	Properties      string
	Modules         modules
	ModulesOnHangup modulesOnHangup
	Prompts         scriptPrompts
	MLPrompts       []*multilingualPrompt
	MLChoices       []*multilanguageMenuChoice
	//Variables map[string]*Variable
	Languages    []language
	TempAPrompts map[moduleID][]*bigTempPrompt
}
type scriptPrompts map[promptID]prompt

func newIVRScript() *IVRScript {
	return &IVRScript{
		Prompts: make(scriptPrompts),
		//Variables: make(map[string]*Variable),
	}
}

type modules struct {
	IncomingCall      *incomingCallModule
	HangupModules     []*hangupModule
	PlayModules       []*playModule
	InputModules      []*inputModule
	VoiceInputModules []*voiceInputModule
	MenuModules       []*menuModule
	GetDigitsModules  []*getDigitsModule
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
	XMLName xml.Name   `xml:"languages"`
	Langs   []language `xml:"languages"`
}
type language struct {
	Lang     langCode `xml:"lang"`
	TtsLang  langCode `xml:"ttsLanguage"`
	TtsVoice string   `xml:"ttsVoice"`
}

//NewIVRScript - Parsing of the getIVRResponse received from Five9 Config web service
func NewIVRScript(src io.Reader) (*IVRScript, error) {
	s := newIVRScript()
	s.TempAPrompts = make(map[moduleID][]*bigTempPrompt)

	decoder := xml.NewDecoder(src)
	decoder.CharsetReader = charset.NewReaderLabel

	inModules := false
	inVariables := false
	inMLPrompts := false
	inMLVIVRPrompts := false
	inMLMenuChoices := false
	inModulesOnHangup := false

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
				inModules = true
			} else if v.Name.Local == "modulesOnHangup" {
				inModulesOnHangup = true
			} else if v.Name.Local == "userVariables" {
				inVariables = true
			} else if v.Name.Local == "multiLanguagesVIVRPrompts" {
				inMLVIVRPrompts = true
			} else if v.Name.Local == "multiLanguagesPrompts" {
				inMLPrompts = true
			} else if v.Name.Local == "value" && inMLPrompts {
				mlp, err := s.parseMultilanguagePrompts(decoder, &v)
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
				mlp, err := s.parseMultilanguageMenuElements(decoder, &v)
				if err != nil {
					fmt.Printf("parseMLPrompt() failed with '%s'\n", err)
					break
				} else {
					s.MLChoices = append(s.MLChoices, mlp)
				}

			} else if v.Name.Local == "incomingCall" || v.Name.Local == "startOnHangup" {
				if err = s.newIncomingCallModule(decoder, &v, inModulesOnHangup); err != nil {
					fmt.Printf("parseIncomingCall() failed with '%s'\n", err)
					fmt.Println(inMLPrompts, inMLMenuChoices, inMLVIVRPrompts, inVariables, inModules)
					break
				}
			} else if v.Name.Local == "hangup" {
				if err = s.parseHangup(decoder, &v, inModulesOnHangup); err != nil {
					fmt.Printf("parseHangup() failed with '%s'\n", err)
					break
				}
			} else if v.Name.Local == "play" {
				err := s.newPlayModule(decoder, &v)
				if err != nil {
					fmt.Printf("parsePlay() failed with '%s'\n", err)
					break
				}
			} else if v.Name.Local == "input" {
				err := s.parseInput2(decoder, &v)
				if err != nil {
					fmt.Printf("parsePlay() failed with '%s'\n", err)
					break
				}
			} else if v.Name.Local == "recording" {
				err := s.parseVoiceInput(decoder, &v)
				if err != nil {
					fmt.Printf("parsePlay() failed with '%s'\n", err)
					break
				}
			} else if v.Name.Local == "menu" {
				err := s.newMenuModule(decoder, &v)
				if err != nil {
					fmt.Printf("parseVoiceInput() failed with '%s'\n", err)
					break
				}

			} else if v.Name.Local == "getDigits" {
				err := s.newGetDigitsModule(decoder, &v)
				if err != nil {
					fmt.Printf("newGetDigitsModule() failed with '%s'\n", err)
					break
				}

			} else if v.Name.Local == "languages" {
				var m languages
				err := decoder.DecodeElement(&m, &v)
				if err == nil {
					s.Languages = m.Langs
				}
			}
		case xml.EndElement:
			if v.Name.Local == "modules" {
				inModules = false
			} else if v.Name.Local == "modulesOnHangup" {
				inModulesOnHangup = false
			} else if v.Name.Local == "userVariables" {
				inVariables = false
			} else if v.Name.Local == "multiLanguagesVIVRPrompts" {
				inMLVIVRPrompts = false
			} else if v.Name.Local == "multiLanguagesPrompts" {
				inMLPrompts = false
			} else if v.Name.Local == "multiLanguagesMenuChoices" {
				inMLMenuChoices = false
			}
		}
	}

	return s, nil
}

////////////////////////////////////////////////////
type generalInfo struct {
	ID              moduleID
	Ascendants      []moduleID
	Descendant      moduleID
	ExceptionalDesc moduleID
	Name            string
	Dispo           string
	Collapsible     bool
}

func (gi *generalInfo) parseGeneralInfo(decoder *xml.Decoder, v *xml.StartElement) (bool, error) {
	if v.Name.Local == "ascendants" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Ascendants = append(gi.Ascendants, moduleID(innerText.(xml.CharData)))
		}
	} else if v.Name.Local == "singleDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Descendant = moduleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "exceptionalDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.ExceptionalDesc = moduleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "moduleName" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Name = string(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "moduleId" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.ID = moduleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "collapsible" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Collapsible = (string(innerText.(xml.CharData)) == "true")
		}
	} else if v.Name.Local == "dispo" {
		for {
			t, err := decoder.Token()
			if err != nil {
				fmt.Printf("decoder.Token() failed with '%s'\n", err)
				return false, err
			}
			switch v := t.(type) {
			case xml.StartElement:
				if v.Name.Local == "name" {
					innerText, err := decoder.Token()
					if err == nil {
						gi.Dispo = string(innerText.(xml.CharData))
					}
				}
			case xml.EndElement:
				if v.Name.Local == "dispo" {
					return true, nil
				}
			}
		}
	} else {
		return false, nil
	}

	return true, nil
}
