package ivrparser

import (
	"encoding/xml"
	"fmt"
	"io"
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
	Prompts              []promptID
	Events               []*recoEvent
}

func (s *IVRScript) newConfirmData(decoder *xml.Decoder, v *xml.StartElement, prefix string) *confirmData {
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	var pCD = new(confirmData)

F:
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			// io.EOF is a successful end
			break
		}
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "prompt" {
				pCD.Prompts, _ = s.parseVoicePrompt(decoder, &v, fmt.Sprintf("%s_%s_", prefix, "CD"))
			} else if v.Name.Local == "recoEvents" {
				pRE := s.newEvent(decoder, &v, fmt.Sprintf("%s_", prefix))
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
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return pCD
}
