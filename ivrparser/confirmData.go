package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

const (
	CONF_REQUIRED     string = "REQUIRED"
	CONF_NOT_REQUIRED string = "NOT_REQUIRED"
)

type confirmData struct {
	ConfirmRequired      string
	RequiredConfidence   int
	MaxAttemptsToConfirm int
	NoInputTimeout       int
	VoicePromptIDs       modulePrompts
	Events               []*recoEvent
}

func newConfirmData(decoder *xml.Decoder, sp scriptPrompts, prefix string) *confirmData {
	var pCD = new(confirmData)

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			///// prompts -->
			if v.Name.Local == "prompt" {
				if res, err := parseVoicePrompt(decoder, &v, sp, prefix); err == nil {
					pCD.VoicePromptIDs, _ = newModulePrompts(1, res)
				}
			} else if v.Name.Local == "recoEvents" {
				pRE := newEvent(decoder, sp, fmt.Sprintf("%s_", prefix))
				if pRE != nil {
					pCD.Events = append(pCD.Events, pRE)
				}
			} else if v.Name.Local == "confirmRequired" {
				innerText, err := decoder.Token()
				if err == nil {
					pCD.ConfirmRequired = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "requiredConfidence" {
				innerText, err := decoder.Token()
				if err == nil {
					pCD.RequiredConfidence, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "maxAttemptsToConfirm" {
				innerText, err := decoder.Token()
				if err == nil {
					pCD.MaxAttemptsToConfirm, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "noInputTimeout" {
				innerText, err := decoder.Token()
				if err == nil {
					pCD.NoInputTimeout, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			}
		case xml.EndElement:
			if v.Name.Local == cConfirmData {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return pCD
}
