package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type recoevent string

const (
	eventNoMatch recoevent = "NO_MATCH"
	eventNoInput recoevent = "NO_INPUT"
	eventHelp    recoevent = "HELP"
)

type recoaction string

const (
	actionContinue recoaction = "CONTINUE"
	actionReprompt recoaction = "REPROMPT"
	actionExit     recoaction = "EXIT"
)

// RecoEvent - recognition event
type RecoEvent struct {
	Event          recoevent
	Action         recoaction
	CountAndPrompt AttemptPrompts
}

func newEvent(decoder *xml.Decoder, sp ScriptPrompts, prefix string) *RecoEvent {
	var (
		pRE    = new(RecoEvent)
		prompt []PromptID
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
					pRE.Event = recoevent(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "action" {
				innerText, err := decoder.Token()
				if err == nil {
					pRE.Action = recoaction(innerText.(xml.CharData))
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
