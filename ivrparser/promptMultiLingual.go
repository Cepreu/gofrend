package ivrparser

import (
	"encoding/xml"
	"fmt"
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
	AudPrompts  map[langCode][]promptID
	VisPrompts  map[langCode][]promptID
	TxtPrompts  map[langCode][]promptID
	DefLanguage langCode
	//	IsPersistent bool       `xml:"isPersistent"`
}

func (s *IVRScript) parseMultilanguagePrompts(decoder *xml.Decoder) (*multilingualPrompt, error) {
	pml := new(multilingualPrompt)
	pml.Prompts = make(map[langCode][]promptID)
	inDescription, inType, inName, inPromptID, inPrompts, inDefaultLanguage := false, false, false, false, false, false

	var immersion = 1
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
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
				lang := langCode(v.Attr[0].Value)
				pids, err := parseVoicePrompt(decoder, &v, s.Prompts, fmt.Sprintf("%s_%s", pml.ID, lang))
				if err == nil {
					pml.Prompts[lang] = pids
					immersion--
				}
			}
		case xml.EndElement:
			immersion--
			if immersion == 0 {
				break F /// <----------------------------------- Return should be HERE!
			} else if v.Name.Local == "promptId" {
				inPromptID = false
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "type" {
				inType = false
			} else if v.Name.Local == "description" {
				inDescription = false
			} else if v.Name.Local == "defaultLanguage" {
				inDefaultLanguage = false
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

func (s *IVRScript) parseMultilanguageMenuElements(decoder *xml.Decoder) (*multilanguageMenuChoice, error) {
	pml := new(multilanguageMenuChoice)
	pml.AudPrompts = make(map[langCode][]promptID)
	pml.VisPrompts = make(map[langCode][]promptID)
	pml.TxtPrompts = make(map[langCode][]promptID)

	var immersion = 1

	var lang langCode
	var prefix promptID
	var pp prompt

	inDescription, inType, inName, inPromptID := false, false, false, false
	inDefaultLanguage := false
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
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
			} else if v.Name.Local == "entry" {
				for _, attr := range v.Attr {
					if attr.Name.Local == "key" {
						lang = langCode(attr.Value)
					}
				}
			} else if v.Name.Local == "voicePrompt" || v.Name.Local == "vivrPrompt" || v.Name.Local == "textPrompt" {
				pp = nil
			} else if v.Name.Local == "xml" {
				innerText, err := decoder.Token()
				if err == nil {
					p, err := cmdUnzip(string(innerText.(xml.CharData)))
					if err == nil {
						pp = &ttsPrompt{TTSPromptXML: p}
					}
				}
			}
		case xml.EndElement:
			immersion--
			if immersion == 0 {
				break F /// <----------------------------------- Return should be HERE!
			} else if v.Name.Local == "promptId" {
				inPromptID = false
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "type" {
				inType = false
			} else if v.Name.Local == "description" {
				inDescription = false
			} else if v.Name.Local == "voicePrompt" {
				if pp != nil {
					prefix = promptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "A"))
					s.Prompts[prefix] = pp
					pml.AudPrompts[lang] = append(pml.AudPrompts[lang], prefix)
				}
			} else if v.Name.Local == "vivrPrompt" {
				if pp != nil {
					prefix = promptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "V"))
					s.Prompts[prefix] = pp
					pml.VisPrompts[lang] = append(pml.VisPrompts[lang], prefix)
				}
			} else if v.Name.Local == "textPrompt" {
				if pp != nil {
					prefix = promptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "T"))
					s.Prompts[prefix] = pp
					pml.TxtPrompts[lang] = append(pml.TxtPrompts[lang], prefix)
				}
			} else if v.Name.Local == "defaultLanguage" {
				inDefaultLanguage = false
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
