package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

const (
	EVENT_NO_MATCH string = "NO_MATCH"
	EVENT_NO_INPUT string = "NO_INPUT"
	EVENT_HELP     string = "HELP"
)

const (
	ACTION_CONTINUE string = "CONTINUE"
	RACTION_EPROMPT string = "REPROMPT"
	EACTION_XIT     string = "EXIT"
)

type recoEvent struct {
	Event          string
	Action         string
	CountAndPrompt attemptPrompts
}

func newEvent(decoder *xml.Decoder, sp scriptPrompts, prefix string) *recoEvent {
	var (
		pRE    = new(recoEvent)
		prompt []promptID
		count  int
	)
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "compoundPrompt" {
				prompt, _ = parseVoicePrompt(decoder, &v, sp, fmt.Sprintf("%s_%s_", prefix, "RE"))
			} else if v.Name.Local == "event" {
				innerText, err := decoder.Token()
				if err == nil {
					pRE.Event = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "action" {
				innerText, err := decoder.Token()
				if err == nil {
					pRE.Action = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "count" {
				innerText, err := decoder.Token()
				if err == nil {
					count, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			}
		case xml.EndElement:
			if v.Name.Local == "recoEvents" {
				pRE.CountAndPrompt = newAttemptPrompts(count, prompt)
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return pRE
}
