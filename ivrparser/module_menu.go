package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type menuModule struct {
	generalInfo

	VoicePromptIDs modulePrompts
	//	VisualPromptIDs  []promptID
	//	TextPromptIDs    []promptID

	Branches []*outputBranch
	Items    []*menuItem

	UseASR  bool
	UseDTMF bool
	//	RecordUserInput bool

	Events   []*recoEvent
	ConfData *confirmData

	RecoParams struct {
		SpeechCompleteTimeout int
		MaxTimeToEnter        int
		NoInputTimeout        int
	}
}
type actionType string

type outputBranch struct {
	Key   string
	Value struct {
		Name string
		Desc string
	}
}

//////////////////////////////////////
func (s *IVRScript) newMenuModule(decoder *xml.Decoder, v *xml.StartElement) error {
	var pMM = new(menuModule)
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	var (
		inBranches = false
		pBranch    *outputBranch
	)
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
				if prmts, err := s.parseVoicePromptS(decoder, &v, fmt.Sprintf("%s_%s_", pMM.ID, "V")); err == nil {
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
				pRE := s.newEvent(decoder, &v, fmt.Sprintf("%s_", pMM.ID))
				if pRE != nil {
					pMM.Events = append(pMM.Events, pRE)
				}

				// -->confirmData
			} else if v.Name.Local == "confirmData" {
				pMM.ConfData = s.newConfirmData(decoder, &v, fmt.Sprintf("%s_%s_", pMM.ID, "CD"))

				// -->branches
			} else if v.Name.Local == "branches" {
				inBranches = true
			} else if v.Name.Local == "entry" && inBranches {
				pBranch = new(outputBranch)
			} else if v.Name.Local == "key" && inBranches {
				innerText, err := decoder.Token()
				if err == nil {
					pBranch.Key = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "name" && inBranches {
				innerText, err := decoder.Token()
				if err == nil {
					pBranch.Value.Name = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "desc" && inBranches {
				innerText, err := decoder.Token()
				if err == nil {
					pBranch.Value.Desc = string(innerText.(xml.CharData))
				}

				// -->items
			} else if v.Name.Local == "items" {
				pItem := s.newMenuItem(decoder, &v, string(pMM.ID))
				if pItem != nil {
					pMM.Items = append(pMM.Items, pItem)
				}
			} else {
				pMM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == "branches" {
				inBranches = false
			} else if v.Name.Local == "entry" && inBranches {
				pMM.Branches = append(pMM.Branches, pBranch)
			} else if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}

	s.MenuModules = append(s.MenuModules, pMM)
	return nil
}
