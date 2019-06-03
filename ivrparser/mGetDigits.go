package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type GetDigitsModule struct {
	GeneralInfo
	VoicePromptIDs     ModulePrompts
	VisualPromptIDs    []PromptID
	TextPromptIDs      []PromptID
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

func (module *GetDigitsModule) normalize(s *IVRScript) error {
	return s.normalizePrompt(module.VoicePromptIDs)
}

func newGetDigitsModule(decoder *xml.Decoder, sp ScriptPrompts) Module {
	var pModule = new(GetDigitsModule)

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
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
				if res, err := parseVoicePrompt(decoder, &v, sp, fmt.Sprintf("%s_%s_", pModule.ID, "G")); err == nil {
					//					s.TempAPrompts[pModule.ID] = []*AttemptPrompts{{res, 1}}
					pModule.VoicePromptIDs, _ = newModulePrompts(1, res)
				}

			} else {
				pModule.parseGeneralInfo(decoder, &v)
			}

		case xml.EndElement:
			if v.Name.Local == cGetDigits {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}

	return pModule
}
