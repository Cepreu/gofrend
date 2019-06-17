package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

type xmlInputModule struct {
	m *ivr.InputModule
}

func newInputModule(decoder *xml.Decoder, sp ivr.ScriptPrompts) normalizer {
	var pIM = new(ivr.InputModule)

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
				///// -->grammar
			} else if v.Name.Local == "grammar" {
				var pg = &pIM.Grammar
				err := decoder.DecodeElement(pg, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement(Grammar) failed with '%s'\n", err)
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

			} else {
				parseGeneralInfo(pIM, decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == cInput {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}

	return xmlInputModule{pIM}
}

func (module xmlInputModule) normalize(s *ivr.IVRScript) error {
	return normalizePrompt(s, module.m.VoicePromptIDs)
}
