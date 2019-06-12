package xmlparser

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
)

const defaultLang = "Default"

// // PromptID - Unique prompt ID
// type PromptID string

// // ModulePrompts - Complete prompt structure
// type ModulePrompts []AttemptPrompts

// // AttemptPrompts - Single-attempt prompt structure
// type AttemptPrompts struct {
// 	LangPrArr []LanguagePrompts
// 	Count     int
// }

// // // TransformToAI - Transforms to a form suitable for Dialogflow
// func (mp ModulePrompts) TransformToAI(sp ScriptPrompts) (res []string) {
// 	for _, ap := range mp {
// 		res = append(res, ap.TransformToAI(sp))
// 	}
// 	return
// }

// // TransformToAI - Transforms to a form suitable for Dialogflow
// func (ap AttemptPrompts) TransformToAI(sp ScriptPrompts) (txt string) {
// 	for _, id := range ap.LangPrArr[0].PrArr {
// 		txt = txt + sp[id].TransformToAI()
// 	}
// 	return
// }

// // LanguagePrompts - prompt for a specified language
// type LanguagePrompts struct {
// 	PrArr    []PromptID
// 	Language langCode
// }

func newModulePrompts(c int, p []ivr.PromptID) (ivr.ModulePrompts, error) {
	return ivr.ModulePrompts{newAttemptPrompts(c, p)}, nil
}
func newAttemptPrompts(c int, p []ivr.PromptID) ivr.AttemptPrompts {
	pmp := ivr.AttemptPrompts{Count: c, LangPrArr: []ivr.LanguagePrompts{{Language: defaultLang, PrArr: p}}}
	return pmp
}

// func normalizeAttemptPrompt(s *ivr.IVRScript, ap *ivr.AttemptPrompts, mlPrompTrueMlItemFalse bool) error {
// 	var prsdef = []ivr.PromptID{}
// 	for _, l := range s.Languages {
// 		ap.LangPrArr = append(ap.LangPrArr, ivr.LanguagePrompts{Language: l.Lang, PrArr: nil})
// 	}
// 	for _, pid := range ap.LangPrArr[0].PrArr {
// 		if _, found := s.Prompts[pid]; found {
// 			for j := range s.Languages {
// 				ap.LangPrArr[j+1].PrArr = append(ap.LangPrArr[j+1].PrArr, pid)
// 			}
// 			prsdef = append(prsdef, pid)
// 		} else if mlPrompTrueMlItemFalse {
// 			for k := range s.MLPrompts {
// 				if s.MLPrompts[k].ID == string(pid) {
// 					// Found!
// 					for j, l := range s.Languages {
// 						ap.LangPrArr[j+1].PrArr = append(ap.LangPrArr[j+1].PrArr, s.MLPrompts[k].Prompts[l.Lang]...)
// 					}
// 					prsdef = append(prsdef, s.MLPrompts[k].Prompts[s.MLPrompts[k].DefLanguage]...)
// 					break
// 				}
// 			}
// 		} else {
// 			for k := range s.MLChoices {
// 				if s.MLChoices[k].ID == string(pid) {
// 					// Found!
// 					for j, l := range s.Languages {
// 						ap.LangPrArr[j+1].PrArr = append(ap.LangPrArr[j+1].PrArr, s.MLChoices[k].AudPrompts[l.Lang]...)
// 					}
// 					prsdef = append(prsdef, s.MLChoices[k].AudPrompts[s.MLChoices[k].DefLanguage]...)
// 					break
// 				}
// 			}

// 		}
// 	}
// 	ap.LangPrArr[0].PrArr = prsdef
// 	return nil
// }

// func normalizePrompt(s *ivr.IVRScript, mp ivr.ModulePrompts) error {
// 	for i := range mp {
// 		NormalizeAttemptPrompt(s, &mp[i], true)
// 	}
// 	return nil
// }

var counter int32

func getPromptID(prefix string, t string) ivr.PromptID {
	counter++
	return ivr.PromptID(fmt.Sprintf("%s_%s_%d", prefix, t, counter))
}

type prompt interface {
	TransformToAI() string
}

func parseVoicePromptS(decoder *xml.Decoder, sp ivr.ScriptPrompts, prefix string) (ivr.AttemptPrompts, error) {
	var (
		PrArr []ivr.PromptID
		count int
	)
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return ivr.AttemptPrompts{}, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "prompt" {
				PrArr, err = parseVoicePrompt(decoder, &v, sp, prefix)
				if err != nil {
					fmt.Printf("parseVoicePrompt failed with '%s'\n", err)
					return ivr.AttemptPrompts{}, err
				}

			} else if v.Name.Local == "count" {
				innerText, err := decoder.Token()
				if err == nil {
					count, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			}

		case xml.EndElement:
			if v.Name.Local == "prompts" {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return newAttemptPrompts(count, PrArr), nil
}

func parseVoicePrompt(decoder *xml.Decoder, v *xml.StartElement, sp ivr.ScriptPrompts, prefix string) (pids []ivr.PromptID, err error) {
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	inTTS, inXML := false, false
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "filePrompt" {
				var fp ivr.XFilePrompt
				err = decoder.DecodeElement(&fp, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement(filePrompt) failed with '%s'\n", err)
					break
				}
				pid := getPromptID(prefix, "F")
				pids = append(pids, pid)
				sp[pid] = fp
			} else if v.Name.Local == "ttsPrompt" {
				inTTS = true
			} else if v.Name.Local == "pausePrompt" {
				var pp ivr.XPausePrompt
				err = decoder.DecodeElement(&pp, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement(pausePrompt) failed with '%s'\n", err)
					break
				}
				pid := getPromptID(prefix, "P")
				pids = append(pids, pid)
				sp[pid] = pp
			} else if v.Name.Local == "multiLanguagesPromptItem" {
				var pp ivr.XMultiLanguagesPromptItem
				err = decoder.DecodeElement(&pp, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement(multiLanguagesPromptItem>) failed with '%s'\n", err)
					break
				}
				pids = append(pids, ivr.PromptID(pp.MLPromptID))
			} else if v.Name.Local == "xml" {
				inXML = true
			}

		case xml.EndElement:
			if v.Name.Local == "ttsPrompt" {
				inTTS = false
			} else if v.Name.Local == "xml" {
				inXML = false
			} else if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return is HERE!
			}
		case xml.CharData:
			if inTTS && inXML {
				p, err := utils.CmdUnzip(string(v))
				if err == nil {
					pid := getPromptID(prefix, "T")
					pids = append(pids, pid)
					var pp prompt = &ivr.TtsPrompt{TTSPromptXML: p}
					sp[pid] = pp
				}
			}
		}
	}
	return pids, nil
}
