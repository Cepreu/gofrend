package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
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

type inputModule struct {
	generalInfo
	VoicePromptIDs modulePrompts
	//	VisualPromptIDs  modulePrompt
	//	TextPromptIDs    modulePrompt
	Grammar  xInputGrammar
	Events   []*recoEvent
	ConfData *confirmData
}

//////////////////////////////////////
func (s *IVRScript) parseInput2(decoder *xml.Decoder, v *xml.StartElement) error {
	var pIM = new(inputModule)
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	var inPrompts = false
	var pWrk *attemptPrompts
	var pTempBP = make([]*attemptPrompts, 0)

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return err
		}

		switch v := t.(type) {
		case xml.StartElement:
			///// prompts -->
			if v.Name.Local == "prompts" {
				inPrompts = true
				pWrk = new(attemptPrompts)
			} else if v.Name.Local == "prompt" && inPrompts {
				pWrk.PrArr, _ = s.parseVoicePrompt(decoder, &v, fmt.Sprintf("%s_%s_", pIM.ID, "I"))
			} else if v.Name.Local == "count" && inPrompts {
				innerText, err := decoder.Token()
				if err == nil {
					pWrk.Count, _ = strconv.Atoi(string(innerText.(xml.CharData)))
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
				pRE := s.newEvent(decoder, &v, fmt.Sprintf("%s_", pIM.ID))
				if pRE != nil {
					pIM.Events = append(pIM.Events, pRE)
				}

				///// -->confirmData
			} else if v.Name.Local == "confirmData" {
				pIM.ConfData = s.newConfirmData(decoder, &v, fmt.Sprintf("%s_%s_", pIM.ID, "CD"))

			} else {
				pIM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == "prompts" {
				inPrompts = false
				pTempBP = append(pTempBP, pWrk)
			} else if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}

	s.InputModules = append(s.InputModules, pIM)
	s.TempAPrompts[pIM.ID] = pTempBP
	return nil
}
