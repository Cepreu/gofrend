package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type playModule struct {
	GeneralInfo
	VoicePromptIDs ModulePrompts
	//	VisualPromptIDs modulePrompt
	//	TextPromptIDs   modulePrompt
	AuxInfo struct {
		NumberOfDigits   int
		TerminateDigit   string
		ClearDigitBuffer bool
	}
}

func (module *playModule) normalize(s *IVRScript) error {
	return s.normalizePrompt(module.VoicePromptIDs)
}

func newPlayModule(decoder *xml.Decoder, sp ScriptPrompts) Module {
	var pPM = new(playModule)

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
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
			} else if v.Name.Local == cPrompt {
				if res, err := parseVoicePrompt(decoder, &v, sp, fmt.Sprintf("%s_%s_", pPM.ID, "P")); err == nil {
					pPM.VoicePromptIDs, _ = newModulePrompts(1, res)
				}

			} else {
				pPM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == cPlay {
				break F /// <----------------------------------- Return should be HERE!

			}
		}
	}

	return pPM
}
