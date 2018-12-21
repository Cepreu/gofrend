package ivrparser

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
)

type voiceInputModule struct {
	generalInfo
	VoicePromptIDs modulePrompts
	//	VisualPromptIDs  modulePrompt
	//	TextPromptIDs    modulePrompt
	Events   []*recoEvent
	ConfData *confirmData

	RecordingParams struct {
		MaxTime                           int
		FinalSilence                      int
		DTMFtermination                   bool
		ThrowExceptionIfMaxSilenceReached bool
		MaxAttempts                       int
	}
	PostRecording struct {
		VarToAccessRecording    string
		RecordingDurationVar    string
		TerminationCharacterVar string
	}
	AsGreeting struct {
		UseRecordingAsGreeting bool
		IsVarSelected          bool
		StringValue            string
		VariableName           string
	}
}

func (module *voiceInputModule) normalize(s *IVRScript) error {
	return s.normalizePrompt(module.VoicePromptIDs)
}

//////////////////////////////////////
func newVoiceInput(decoder *xml.Decoder, sp scriptPrompts) Module {
	var inRecordingDuration, inTerminationCharacter, inVarToAccessRecording = false, false, false

	var pIM = new(voiceInputModule)
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			///// prompts -->
			if v.Name.Local == "prompts" {
				if prmts, err := parseVoicePromptS(decoder, sp, fmt.Sprintf("%s_%s_", pIM.ID, "V")); err == nil {
					pIM.VoicePromptIDs = append(pIM.VoicePromptIDs, prmts)
				}
				// RecordingParams -->
			} else if v.Name.Local == "maxTime" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.RecordingParams.MaxTime, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "finalSilence" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.RecordingParams.FinalSilence, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "DTMFtermination" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.RecordingParams.DTMFtermination = (string(innerText.(xml.CharData)) == "true")
				}
			} else if v.Name.Local == "throwExceptionIfMaxSilenceReached" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.RecordingParams.ThrowExceptionIfMaxSilenceReached = (string(innerText.(xml.CharData)) == "true")
				}
			} else if v.Name.Local == "maxAttempts" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.RecordingParams.MaxAttempts, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
				//	PostRecording -->
			} else if v.Name.Local == "varToAccessRecording" {
				inVarToAccessRecording = true
			} else if v.Name.Local == "name" && inVarToAccessRecording {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					pIM.PostRecording.VarToAccessRecording = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "recordingDuration" {
				inRecordingDuration = true
			} else if v.Name.Local == "name" && inRecordingDuration {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					pIM.PostRecording.RecordingDurationVar = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "terminationCharacter" {
				inTerminationCharacter = true
			} else if v.Name.Local == "name" && inTerminationCharacter {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					pIM.PostRecording.TerminationCharacterVar = string(innerText.(xml.CharData))
				}

				// AsGreeting struct -- >
			} else if v.Name.Local == "useRecordingAsGreeting" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.AsGreeting.UseRecordingAsGreeting = (string(innerText.(xml.CharData)) == "true")
				}
			} else if v.Name.Local == "isVarSelected" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.AsGreeting.IsVarSelected = (string(innerText.(xml.CharData)) == "true")
				}
			} else if v.Name.Local == "stringValue" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.AsGreeting.StringValue = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "variableName" {
				innerText, err := decoder.Token()
				if err == nil {
					pIM.AsGreeting.VariableName = string(innerText.(xml.CharData))
				}
				///// -->reco-events
			} else if v.Name.Local == "recoEvents" {
				pRE := newEvent(decoder, sp, fmt.Sprintf("%s_", pIM.ID))
				if pRE != nil {
					pIM.Events = append(pIM.Events, pRE)
				}

				///// -->confirmData
			} else if v.Name.Local == "confirmData" {
				pIM.ConfData = newConfirmData(decoder, sp, fmt.Sprintf("%s_%s_", pIM.ID, "CD"))
				fmt.Println(pIM.ConfData)

			} else {
				pIM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == "varToAccessRecording" {
				inRecordingDuration = false
			} else if v.Name.Local == "recordingDuration" {
				inRecordingDuration = false
			} else if v.Name.Local == "terminationCharacter" {
				inRecordingDuration = false
			} else if v.Name.Local == cVoiceInput {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}

	return pIM
}
