package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/Cepreu/gofrend/ivr"
	"golang.org/x/net/html/charset"
)

var xMLPrompts []*MultilingualPrompt

//NewIVRScript - Parsing of the getIVRResponse received from Five9 Config web service
func NewIVRScript(src io.Reader) (*ivr.IVRScript, error) {

	var (
		s = &ivr.IVRScript{
			Prompts:   make(ivr.ScriptPrompts),
			Variables: make(ivr.Variables),
			Modules:   make(map[ivr.ModuleID]ivr.Module),
		}
		xmlModules         []normalizer
		xmlModulesOnHangup []normalizer
		decoder            = xml.NewDecoder(src)
		inMLPrompts        = false
		inMLMenuChoices    = false
		inDomainID         = false
	)
	decoder.CharsetReader = charset.NewReaderLabel
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
				xmlModules = parseModules(s, decoder, &v)
			} else if v.Name.Local == "modulesOnHangup" {
				xmlModulesOnHangup = parseModules(s, decoder, &v)
			} else if v.Name.Local == "userVariables" {
				parseVars(s, decoder)
			} else if v.Name.Local == "multiLanguagesVIVRPrompts" {
				//				inMLVIVRPrompts = true
			} else if v.Name.Local == "multiLanguagesPrompts" {
				inMLPrompts = true
			} else if v.Name.Local == "value" && inMLPrompts {
				mlp, err := parseMultilanguagePrompts(s, decoder)
				if err != nil {
					fmt.Printf("parseMLPrompt() failed with '%s'\n", err)
					break
				} else {
					xMLPrompts = append(xMLPrompts, mlp)
				}
				// mlMenuChoices <---
			} else if v.Name.Local == "multiLanguagesMenuChoices" {
				inMLMenuChoices = true
			} else if v.Name.Local == "value" && inMLMenuChoices {
				mlp, err := parseMultilanguageMenuElements(s, decoder)
				if err != nil {
					fmt.Printf("parseMLPrompt() failed with '%s'\n", err)
					break
				} else {
					s.MLChoices = append(s.MLChoices, mlp)
				}

			} else if v.Name.Local == "languages" {
				var m ivr.Languages
				err := decoder.DecodeElement(&m, &v)
				if err == nil {
					if len(m.Langs) > 0 {
						s.Languages = m.Langs
					} else {
						s.Languages = []ivr.Language{{Lang: "en-US", TtsLang: "en-US", TtsVoice: "Samanta"}}
					}
				}

			} else if v.Name.Local == "functions" {
				s.JSFunctions = newJSFunctions(decoder)
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

	for _, module := range xmlModules {
		module.normalize(s)
	}
	for _, module := range xmlModulesOnHangup {
		module.normalize(s)
	}

	return s, nil
}
