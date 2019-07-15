package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

type normalizer interface {
	normalize(*ivr.IVRScript) error
}

const (
	cIncomingCall  string = "incomingCall"
	cStartOnHangup string = "startOnHangup"
	cHangup        string = "hangup"
	cGetDigits     string = "getDigits"
	cPlay          string = "play"
	cInput         string = "input"
	cVoiceInput    string = "recording"
	cMenu          string = "menu"
	cQuery         string = "query"
	cSetVariables  string = "setVariable"
	cIfElse        string = "ifElse"
	cCase          string = "case"
	cForeignScript string = "foreignScript"

	cPrompt      string = "prompt"
	cConfirmData string = "ConfirmData"
	cMenuItems   string = "items"
)

func parseModules(s *ivr.IVRScript, decoder *xml.Decoder, v *xml.StartElement) (ms []normalizer) {
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	ms = []normalizer{}
F:
	for {
		var m normalizer

		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			switch v.Name.Local {
			case cIncomingCall, cStartOnHangup:
				m = newIncomingCallModule(decoder)
			case cHangup:
				m = newHangupModule(s, decoder)
			case cPlay:
				m = newPlayModule(decoder, s.Prompts)
			case cInput:
				m = newInputModule(decoder, s.Prompts)
			case cVoiceInput:
				m = newVoiceInput(decoder, s.Prompts)
			case cMenu:
				m = newMenuModule(s, decoder)
				//				s.Menus = append(s.Menus, m.GetID())
			case cGetDigits:
				m = newGetDigitsModule(decoder, s.Prompts)
			case cQuery:
				m = newQueryModule(decoder, s)
			case cSetVariables:
				m = newSetVariablesModule(decoder, s)
			case cIfElse:
				m = newIfElseModule(s, decoder)
			case cCase:
				m = newCaseModule(s, decoder)
			case cForeignScript:
				m = newForeignScriptModule(decoder, s)
			default:
				fmt.Printf("Warning: unsupported module '%s'\n", v.Name.Local)
				//				m = newUnknownModule(decoder, &v)
			}
			if m != nil {
				ms = append(ms, m)
			}
		case xml.EndElement:
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return ms
}

func parseGeneralInfo(gi ivr.Module, decoder *xml.Decoder, v *xml.StartElement) (bool, error) {
	var (
		asc         []ivr.ModuleID
		desc        ivr.ModuleID
		excDesc     ivr.ModuleID
		name        string
		id          ivr.ModuleID
		dispo       string
		collapsible string
	)
	if v.Name.Local == "ascendants" {
		innerText, err := decoder.Token()
		if err == nil {
			asc = []ivr.ModuleID{ivr.ModuleID(innerText.(xml.CharData))}
			gi.SetGeneralInfo("", "", asc, "", "", "", "")
		}
	} else if v.Name.Local == "singleDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			desc = ivr.ModuleID(innerText.(xml.CharData))
			gi.SetGeneralInfo("", "", asc, desc, "", "", "")
		}
	} else if v.Name.Local == "exceptionalDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			excDesc = ivr.ModuleID(innerText.(xml.CharData))
			gi.SetGeneralInfo("", "", asc, "", excDesc, "", "")
		}
	} else if v.Name.Local == "moduleName" {
		innerText, err := decoder.Token()
		if err == nil {
			name = string(innerText.(xml.CharData))
			gi.SetGeneralInfo(name, "", asc, "", "", "", "")
		}
	} else if v.Name.Local == "moduleId" {
		innerText, err := decoder.Token()
		if err == nil {
			id = ivr.ModuleID(innerText.(xml.CharData))
			gi.SetGeneralInfo("", id, asc, "", "", "", "")
		}
	} else if v.Name.Local == "collapsible" {
		innerText, err := decoder.Token()
		if err == nil {
			collapsible = string(innerText.(xml.CharData))
			gi.SetGeneralInfo("", "", asc, "", "", "", collapsible)
		}
	} else if v.Name.Local == "dispo" {
		for {
			t, err := decoder.Token()
			if err != nil {
				fmt.Printf("decoder.Token() failed with '%s'\n", err)
				return false, err
			}
			switch v := t.(type) {
			case xml.StartElement:
				if v.Name.Local == "name" {
					innerText, err := decoder.Token()
					if err == nil {
						dispo = string(innerText.(xml.CharData))
						gi.SetGeneralInfo("", "", asc, "", "", dispo, "")
					}
				}
			case xml.EndElement:
				if v.Name.Local == "dispo" {
					return true, nil
				}
			}
		}
	} else {
		return false, nil
	}

	return true, nil
}

////////////////
// type unknownModule struct {
// 	ivr.Module
// }

// func (*unknownModule) normalize(*ivr.IVRScript) error {
// 	return nil
// }

// func newUnknownModule(decoder *xml.Decoder, v *xml.StartElement) ivr.Module {
// 	var lastElement = v.Name.Local
// 	var pUnknown = new(ivr.IncomingCallModule)

// F:
// 	for {
// 		t, err := decoder.Token()
// 		if err != nil {
// 			fmt.Printf("decoder.Token() failed with '%s'\n", err)
// 			return nil
// 		}

// 		switch v := t.(type) {
// 		case xml.StartElement:
// 			parseGeneralInfo(pUnknown, decoder, &v)
// 		case xml.EndElement:
// 			if lastElement == v.Name.Local {
// 				break F /// <----------------------------------- Return should be HERE!
// 			}
// 		}
// 	}
// 	return pUnknown
// }
