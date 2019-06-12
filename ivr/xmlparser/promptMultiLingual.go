package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
)

func parseMultilanguagePrompts(s *ivr.IVRScript, decoder *xml.Decoder) (*ivr.MultilingualPrompt, error) {
	pml := new(ivr.MultilingualPrompt)
	pml.Prompts = make(map[ivr.LangCode][]ivr.PromptID)
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
				lang := ivr.LangCode(v.Attr[0].Value)
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
				pml.DefLanguage = ivr.LangCode(v)
			}
		}
	}
	return pml, nil
}

func parseMultilanguageMenuElements(s *ivr.IVRScript, decoder *xml.Decoder) (*ivr.MultilanguageMenuChoice, error) {
	pml := new(ivr.MultilanguageMenuChoice)
	pml.AudPrompts = make(map[ivr.LangCode][]ivr.PromptID)
	pml.VisPrompts = make(map[ivr.LangCode][]ivr.PromptID)
	pml.TxtPrompts = make(map[ivr.LangCode][]ivr.PromptID)

	var immersion = 1

	var lang ivr.LangCode
	var prefix ivr.PromptID
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
						lang = ivr.LangCode(attr.Value)
					}
				}
			} else if v.Name.Local == "voicePrompt" || v.Name.Local == "vivrPrompt" || v.Name.Local == "textPrompt" {
				pp = nil
			} else if v.Name.Local == "xml" {
				innerText, err := decoder.Token()
				if err == nil {
					p, err := utils.CmdUnzip(string(innerText.(xml.CharData)))
					if err == nil {
						pp = &ivr.TtsPrompt{TTSPromptXML: p}
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
					prefix = ivr.PromptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "A"))
					s.Prompts[prefix] = pp
					pml.AudPrompts[lang] = append(pml.AudPrompts[lang], prefix)
				}
			} else if v.Name.Local == "vivrPrompt" {
				if pp != nil {
					prefix = ivr.PromptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "V"))
					s.Prompts[prefix] = pp
					pml.VisPrompts[lang] = append(pml.VisPrompts[lang], prefix)
				}
			} else if v.Name.Local == "textPrompt" {
				if pp != nil {
					prefix = ivr.PromptID(fmt.Sprintf("%s_%s_%s", pml.ID, lang, "T"))
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
				pml.DefLanguage = ivr.LangCode(v)
			}
		}
	}
	return pml, nil
}
