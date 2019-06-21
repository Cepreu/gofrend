 package xmlparser

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/Cepreu/gofrend/ivr"
)

func newEvent(decoder *xml.Decoder, sp ivr.ScriptPrompts, prefix string) *ivr.RecoEvent {
	var (
		pRE    = new(ivr.RecoEvent)
		prompt []ivr.PromptID
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
					pRE.Event = ivr.RecoType(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "action" {
				innerText, err := decoder.Token()
				if err == nil {
					pRE.Action = ivr.RecoAction(innerText.(xml.CharData))
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
