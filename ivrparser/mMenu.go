package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

//MenuModule - Menu module definition
type MenuModule struct {
	GeneralInfo

	VoicePromptIDs ModulePrompts
	//	VisualPromptIDs  []PromptID
	//	TextPromptIDs    []PromptID

	Branches []*OutputBranch
	Items    []*MenuItem

	UseASR  bool
	UseDTMF bool
	//	RecordUserInput bool

	Events   []*RecoEvent
	ConfData *ConfirmData

	RecoParams struct {
		SpeechCompleteTimeout int
		MaxTimeToEnter        int
		NoInputTimeout        int
	}
}

func (module *MenuModule) normalize(s *IVRScript) error {
	s.normalizePrompt(module.VoicePromptIDs)
	for i := range module.Items {
		s.normalizeAttemptPrompt(&module.Items[i].Prompt, false)
	}
	return nil
}

// ActionType - Menu module item's action
type ActionType string

//////////////////////////////////////
func newMenuModule(decoder *xml.Decoder, sp ScriptPrompts) Module {
	var (
		pMM      = new(MenuModule)
		inModule = true
	)

	for inModule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			///// prompts -->
			if v.Name.Local == "prompts" {
				if prmts, err := parseVoicePromptS(decoder, sp, fmt.Sprintf("%s_%s_", pMM.ID, "V")); err == nil {
					pMM.VoicePromptIDs = append(pMM.VoicePromptIDs, prmts)
				}
			} else if v.Name.Local == "useSpeechRecognition" {
				innerText, err := decoder.Token()
				if err == nil {
					pMM.UseASR = (string(innerText.(xml.CharData)) == "true")
				}
			} else if v.Name.Local == "useDTMF" {
				innerText, err := decoder.Token()
				if err == nil {
					pMM.UseDTMF = (string(innerText.(xml.CharData)) == "true")
				}
			} else if v.Name.Local == "noInputTimeout" {
				innerText, err := decoder.Token()
				if err == nil {
					pMM.RecoParams.NoInputTimeout, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "speechCompleteTimeout" {
				innerText, err := decoder.Token()
				if err == nil {
					pMM.RecoParams.SpeechCompleteTimeout, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "maxTimeToEnter" {
				innerText, err := decoder.Token()
				if err == nil {
					pMM.RecoParams.MaxTimeToEnter, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
				// -->reco-events
			} else if v.Name.Local == "recoEvents" {
				pRE := newEvent(decoder, sp, fmt.Sprintf("%s_", pMM.ID))
				if pRE != nil {
					pMM.Events = append(pMM.Events, pRE)
				}

				// -->ConfirmData
			} else if v.Name.Local == cConfirmData {
				pMM.ConfData = newConfirmData(decoder, sp, fmt.Sprintf("%s_%s_", pMM.ID, "CD"))

				// -->branches
			} else if v.Name.Local == "branches" {
				pMM.Branches = parseBranches(decoder)
				// -->items
			} else if v.Name.Local == cMenuItems {
				pItem := newMenuItem(decoder, sp, getPromptID(string(pMM.ID), "A"))
				//				fmt.Printf("Items: .......... %v\n", pItem)
				if pItem != nil {
					pMM.Items = append(pMM.Items, pItem)
				}
			} else {
				pMM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == cMenu {
				inModule = false
			}
		}
	}
	return pMM
}
