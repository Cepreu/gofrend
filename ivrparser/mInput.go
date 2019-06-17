package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type xInputGrammar struct {
	Type           string           `xml:"type,attr"`
	MRVname        string           `xml:"mainReturnValue>name"`
	MRVtype        string           `xml:"mainReturnValue>type"`
	MRVvariable    string           `xml:"mainReturnValue>varName"`
	StrProperties  []xGrammPropStr  `xml:"stringProperty"`
	ListProperties []xGrammPropList `xml:"listProperty"`
	ARVname        string           `xml:"additionalReturnValues>name"`
	ARVtype        string           `xml:"additionalReturnValues>type"`
}
type xGrammPropStr struct {
	PropType    string `xml:"type"`
	PropValue   string `xml:"value"`
	PropEnabled bool   `xml:"enabled"`
}
type xGrammPropList struct {
	PropType    string   `xml:"type"`
	PropValEnum []string `xml:"list"`
	PropValue   string   `xml:"value"`
	PropEnabled bool     `xml:"enabled"`
}

///////////////////////////////////////////////

type InputModule struct {
	GeneralInfo
	VoicePromptIDs ModulePrompts
	//	VisualPromptIDs  modulePrompt
	//	TextPromptIDs    modulePrompt
	Grammar  xInputGrammar
	Events   []*RecoEvent
	ConfData *ConfirmData
}

func (module *InputModule) normalize(s *IVRScript) error {
	return s.normalizePrompt(module.VoicePromptIDs)
}

//////////////////////////////////////
func newInputModule(decoder *xml.Decoder, sp ScriptPrompts) Module {
	var pIM = new(InputModule)

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
				pIM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == cInput {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}

	return pIM
}
