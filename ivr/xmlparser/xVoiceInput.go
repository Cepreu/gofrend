package xmlparser

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"

	"github.com/Cepreu/gofrend/ivr"
)

type xmlVoiceInputModule struct {
	m *ivr.VoiceInputModule
}

func newVoiceInput(decoder *xml.Decoder, sp ivr.ScriptPrompts) normalizer {
	var inRecordingDuration, inTerminationCharacter, inVarToAccessRecording = false, false, false

	var pIM = new(ivr.VoiceInputModule)
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

				///// -->ConfirmData
			} else if v.Name.Local == "ConfirmData" {
				pIM.ConfData = newConfirmData(decoder, sp, fmt.Sprintf("%s_%s_", pIM.ID, "CD"))
				fmt.Println(pIM.ConfData)

			} else {
				parseGeneralInfo(pIM, decoder, &v)
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

	return xmlVoiceInputModule{pIM}
}

func (mod xmlVoiceInputModule) normalize(s *ivr.IVRScript) error {
	return normalizePrompt(s, mod.m.VoicePromptIDs)
}
