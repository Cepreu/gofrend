package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

// Module - represents IVR modules
// //
// type Module interface {
// 	GetID() ModuleID
// 	GetDescendant() ModuleID
// 	normalize(*IVRScript) error
// }

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

func parseModules(s *ivr.IVRScript, decoder *xml.Decoder, v *xml.StartElement) (ms map[ivr.ModuleID]ivr.Module) {
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	ms = make(map[ivr.ModuleID]ivr.Module)
F:
	for {
		var m ivr.Module

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
				m = newHangupModule(decoder)
			case cPlay:
				m = newPlayModule(decoder, s.Prompts)
			case cInput:
				m = newInputModule(decoder, s.Prompts)
			case cVoiceInput:
				m = newVoiceInput(decoder, s.Prompts)
			case cMenu:
				m = newMenuModule(decoder, s.Prompts)
				s.Menus = append(s.Menus, m.GetID())
			case cGetDigits:
				m = newGetDigitsModule(decoder, s.Prompts)
			case cQuery:
				m = newQueryModule(decoder, s.Prompts)
			case cSetVariables:
				m = newSetVariablesModule(decoder)
			case cIfElse:
				m = newIfElseModule(decoder)
			case cForeignScript:
				m = newForeignScriptModule(decoder)
			default:
				fmt.Printf("Warning: unsupported module '%s'\n", v.Name.Local)
				m = newUnknownModule(decoder, &v)
			}
			if m != nil {
				ms[m.GetID()] = m
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
		asc         = []ivr.ModuleID{}
		desc        ivr.ModuleID
		excDesc     ivr.ModuleID
		name        string
		id          ivr.ModuleID
		dispo       string
		collapsible bool
	)
	if v.Name.Local == "ascendants" {
		innerText, err := decoder.Token()
		if err == nil {
			asc = append(asc, ivr.ModuleID(innerText.(xml.CharData)))
		}
	} else if v.Name.Local == "singleDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			desc = ivr.ModuleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "exceptionalDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			excDesc = ivr.ModuleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "moduleName" {
		innerText, err := decoder.Token()
		if err == nil {
			name = string(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "moduleId" {
		innerText, err := decoder.Token()
		if err == nil {
			id = ivr.ModuleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "collapsible" {
		innerText, err := decoder.Token()
		if err == nil {
			collapsible = (string(innerText.(xml.CharData)) == "true")
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
					}
				}
			case xml.EndElement:
				if v.Name.Local == "dispo" {
					return true, nil
				}
			}
		}
		gi.SetGeneralInfo(name, id, asc, desc, excDesc, dispo, collapsible)
	} else {
		return false, nil
	}

	return true, nil
}

////////////////
type unknownModule struct {
	ivr.Module
}

func (*unknownModule) normalize(*ivr.IVRScript) error {
	return nil
}

func newUnknownModule(decoder *xml.Decoder, v *xml.StartElement) ivr.Module {
	var lastElement = v.Name.Local
	var pUnknown = new(ivr.IncomingCallModule)

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			parseGeneralInfo(pUnknown, decoder, &v)
		case xml.EndElement:
			if lastElement == v.Name.Local {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return pUnknown
}
