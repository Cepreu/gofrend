package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type playModule struct {
	generalInfo
	VoicePromptIDs modulePrompts
	//	VisualPromptIDs modulePrompt
	//	TextPromptIDs   modulePrompt
	AuxInfo struct {
		NumberOfDigits   int
		TerminateDigit   string
		ClearDigitBuffer bool
	}
}

func (s *IVRScript) newPlayModule(decoder *xml.Decoder, v *xml.StartElement) error {
	var pPM = new(playModule)
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
			if v.Name.Local == "numberOfDigits" {
				innerText, err := decoder.Token()
				if err == nil {
					pPM.AuxInfo.NumberOfDigits, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "terminateDigit" {
				innerText, err := decoder.Token()
				if err == nil {
					pPM.AuxInfo.TerminateDigit = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "clearDigitBuffer" {
				innerText, err := decoder.Token()
				if err == nil {
					pPM.AuxInfo.ClearDigitBuffer = (string(innerText.(xml.CharData)) == "true")
				}

				///// prompts -->
			} else if v.Name.Local == "prompt" {
				if res, err := s.parseVoicePrompt(decoder, &v, fmt.Sprintf("%s_%s_", pPM.ID, "P")); err == nil {
					s.TempAPrompts[pPM.ID] = []*attemptPrompts{{res, 1}}
				}

			} else {
				pPM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!

			}
		}
	}

	s.PlayModules = append(s.PlayModules, pPM)
	return nil
}
