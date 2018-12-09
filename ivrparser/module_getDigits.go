package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type getDigitsModule struct {
	generalInfo
	//	VoicePromptIDs     modulePrompt
	VisualPromptIDs    []promptID
	TextPromptIDs      []promptID
	TargetVariableName string
	InputInfo          struct {
		NumberOfDigits   int
		TerminateDigit   string
		ClearDigitBuffer bool
		MaxTime          int
		MaxSilence       int
		Format           string
	}
}

func (s *IVRScript) newGetDigitsModule(decoder *xml.Decoder, v *xml.StartElement) error {
	var pModule = new(getDigitsModule)
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return err
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "targetVariableName" {
				innerText, err := decoder.Token()
				if err == nil {
					pModule.TargetVariableName = string(innerText.(xml.CharData))
				}
				///// inputInfo -->
			} else if v.Name.Local == "maxTime" {
				innerText, err := decoder.Token()
				if err == nil {
					pModule.InputInfo.MaxTime, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "maxSilence" {
				innerText, err := decoder.Token()
				if err == nil {
					pModule.InputInfo.MaxSilence, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "numberOfDigits" {
				innerText, err := decoder.Token()
				if err == nil {
					pModule.InputInfo.NumberOfDigits, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "terminateDigit" {
				innerText, err := decoder.Token()
				if err == nil {
					pModule.InputInfo.TerminateDigit = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "clearDigitBuffer" {
				innerText, err := decoder.Token()
				if err == nil {
					pModule.InputInfo.ClearDigitBuffer = (string(innerText.(xml.CharData)) == "true")
				}
			} else if v.Name.Local == "format" {
				innerText, err := decoder.Token()
				if err == nil {
					pModule.InputInfo.Format = string(innerText.(xml.CharData))
				}

				///// prompts -->
			} else if v.Name.Local == "prompt" {
				if res, err := s.parseVoicePrompt(decoder, &v, fmt.Sprintf("%s_%s_", pModule.ID, "G")); err == nil {
					s.TempAPrompts[pModule.ID] = []*bigTempPrompt{{res, 1}}
				}

			} else {
				pModule.parseGeneralInfo(decoder, &v)
			}

		case xml.EndElement:
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}

	s.Modules.GetDigitsModules = append(s.Modules.GetDigitsModules, pModule)
	return nil
}
