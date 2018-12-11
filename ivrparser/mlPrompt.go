package ivrparser

import (
	"encoding/xml"
	"fmt"
	"io"
)

type langCode string

type multilingualPrompt struct {
	ID          string
	Name        string
	Description string
	Type        string
	Prompts     map[langCode][]promptID
	DefLanguage langCode
	//	IsPersistent bool       `xml:"isPersistent"`
}

type multilanguageMenuChoice struct {
	ID          string
	Name        string
	Description string
	Type        string
	AudPrompts  map[langCode]promptID
	VisPrompts  map[langCode]promptID
	TxtPrompts  map[langCode]promptID
	//	DefLanguage  string     `xml:"defaultLanguage"`
	//	IsPersistent bool       `xml:"isPersistent"`
}

func (s *IVRScript) parseMultilanguagePrompts(decoder *xml.Decoder, v *xml.StartElement) (*multilingualPrompt, error) {
	pml := new(multilingualPrompt)
	pml.Prompts = make(map[langCode][]promptID)
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	inDescription, inType, inName, inPromptID, inPrompts, inDefaultLanguage := false, false, false, false, false, false

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "promptId" {
				inPromptID = true
			} else if v.Name.Local == "name" {
				inName = true
			} else if v.Name.Local == "type" {
				inType = true
			} else if v.Name.Local == "description" {
				inDescription = true
			} else if v.Name.Local == "defaultLanguage" {
				inDefaultLanguage = true
			} else if v.Name.Local == "prompts" {
				inPrompts = true
			} else if v.Name.Local == "entry" && inPrompts {
				for _, attr := range v.Attr {
					if attr.Name.Local == "key" {
						lang := langCode(attr.Value)
						pids, err := s.parseVoicePrompt(decoder, &v, fmt.Sprintf("%s_%s", pml.ID, lang))
						if err == nil {
							pml.Prompts[lang] = pids
						}
						break
					}
				}
			}
		case xml.EndElement:
			if v.Name.Local == "promptId" {
				inPromptID = false
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "type" {
				inType = false
			} else if v.Name.Local == "description" {
				inDescription = false
			} else if v.Name.Local == "defaultLanguage" {
				inDefaultLanguage = false
			} else if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			}
		case xml.CharData:
			if inPromptID {
				pml.ID = string(v)
			} else if inDescription {
				pml.Description = string(v)
			} else if inName {
				pml.Name = string(v)
			} else if inType {
				pml.Type = string(v)
			} else if inDefaultLanguage {
				pml.DefLanguage = langCode(v)
			}
		}
	}
	return pml, nil
}

func (s *IVRScript) parseMultilanguageMenuElements(decoder *xml.Decoder, v *xml.StartElement) (*multilanguageMenuChoice, error) {
	pml := new(multilanguageMenuChoice)
	pml.AudPrompts = make(map[langCode]promptID)
	pml.VisPrompts = make(map[langCode]promptID)
	pml.TxtPrompts = make(map[langCode]promptID)

	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	var lang langCode
	var prefix promptID

	inDescription, inType, inName, inPromptID, inPrompts := false, false, false, false, false
	inVoicePrompt, inVisualPrompt, inTextPrompt := false, false, false
F:
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			// io.EOF should not be here
			return nil, err
		}
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "promptId" {
				inPromptID = true
			} else if v.Name.Local == "name" {
				inName = true
			} else if v.Name.Local == "type" {
				inType = true
			} else if v.Name.Local == "description" {
				inDescription = true
			} else if v.Name.Local == "prompts" {
				inPrompts = true
			} else if v.Name.Local == "entry" && inPrompts {
				for _, attr := range v.Attr {
					if attr.Name.Local == "key" {
						lang = langCode(attr.Value)
					}
				}
			} else if v.Name.Local == "voicePrompt" {
				inVoicePrompt = true
				prefix = promptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "A"))
			} else if v.Name.Local == "visualPrompt" {
				inVisualPrompt = true
				prefix = promptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "V"))
			} else if v.Name.Local == "textPrompt" {
				inTextPrompt = true
				prefix = promptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "T"))
			} else if v.Name.Local == "xml" {
				innerText, err := decoder.Token()
				if err == nil {
					p, err := cmdUnzip(string(innerText.(xml.CharData)))
					if err == nil {
						var pp prompt = &ttsPrompt{TTSPromptXML: p}
						s.Prompts[prefix] = pp

						if inVoicePrompt {
							pml.AudPrompts[lang] = prefix
						} else if inVisualPrompt {
							pml.VisPrompts[lang] = prefix
						} else if inTextPrompt {
							pml.TxtPrompts[lang] = prefix
						}
					}
				}
			}
		case xml.EndElement:
			if v.Name.Local == "promptId" {
				inPromptID = false
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "type" {
				inType = false
			} else if v.Name.Local == "description" {
				inDescription = false
			} else if v.Name.Local == "voicePrompt" {
				inVoicePrompt = false
			} else if v.Name.Local == "visualPrompt" {
				inVisualPrompt = false
			} else if v.Name.Local == "textPrompt" {
				inTextPrompt = false
			} else if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			}
		case xml.CharData:
			if inPromptID {
				pml.ID = string(v)
			} else if inDescription {
				pml.Description = string(v)
			} else if inName {
				pml.Name = string(v)
			} else if inType {
				pml.Type = string(v)
			}
		}
	}
	return pml, nil
}
