package ivrparser

import (
	"encoding/xml"
	"fmt"
)

// Module - represents IVR modules
//
type Module interface {
	GetID() ModuleID
	GetDescendant() ModuleID
	normalize(*IVRScript) error
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

type (
	// ModuleID - the IVR module's ID, string
	ModuleID string

	// GeneralInfo - Set of fields common for all modules
	GeneralInfo struct {
		ID              ModuleID
		Ascendants      []ModuleID
		Descendant      ModuleID
		ExceptionalDesc ModuleID
		Name            string
		Dispo           string
		Collapsible     bool
	}
)

// GetID - returns ID if the module
func (gi *GeneralInfo) GetID() ModuleID {
	return gi.ID
}

// GetDescendant - returns ID if the module's descendant
func (gi *GeneralInfo) GetDescendant() ModuleID {
	return gi.Descendant
}

func (s *IVRScript) parseModules(decoder *xml.Decoder, v *xml.StartElement) (ms map[ModuleID]Module) {
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	ms = make(map[ModuleID]Module)
F:
	for {
		var m Module

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

func (gi *GeneralInfo) parseGeneralInfo(decoder *xml.Decoder, v *xml.StartElement) (bool, error) {
	if v.Name.Local == "ascendants" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Ascendants = append(gi.Ascendants, ModuleID(innerText.(xml.CharData)))
		}
	} else if v.Name.Local == "singleDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Descendant = ModuleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "exceptionalDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.ExceptionalDesc = ModuleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "moduleName" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Name = string(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "moduleId" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.ID = ModuleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "collapsible" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Collapsible = (string(innerText.(xml.CharData)) == "true")
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
						gi.Dispo = string(innerText.(xml.CharData))
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
type unknownModule struct {
	GeneralInfo
}

func (*unknownModule) normalize(*IVRScript) error {
	return nil
}

func newUnknownModule(decoder *xml.Decoder, v *xml.StartElement) Module {
	var lastElement = v.Name.Local
	var pUnknown = new(IncomingCallModule)

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			pUnknown.parseGeneralInfo(decoder, &v)
		case xml.EndElement:
			if lastElement == v.Name.Local {
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return pUnknown
}
