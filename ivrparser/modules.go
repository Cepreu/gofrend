package ivrparser

import (
	"encoding/xml"
	"fmt"
)

// Module - represents IVR modules
//
type Module interface {
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

	cPrompt      string = "prompt"
	cConfirmData string = "confirmData"
	cMenuItems   string = "items"
)

type moduleID string

type generalInfo struct {
	ID              moduleID
	Ascendants      []moduleID
	Descendant      moduleID
	ExceptionalDesc moduleID
	Name            string
	Dispo           string
	Collapsible     bool
}

func (s *IVRScript) parseModules(decoder *xml.Decoder, v *xml.StartElement) (ms []Module) {
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}

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
			case cGetDigits:
				m = newGetDigitsModule(decoder, s.Prompts)
			case cQuery:
				m = newQueryModule(decoder, s.Prompts)
			default:
				fmt.Printf("Warning: unsupported module '%s'\n", v.Name.Local)
				m = newUnknownModule(decoder, &v)
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

func (gi *generalInfo) parseGeneralInfo(decoder *xml.Decoder, v *xml.StartElement) (bool, error) {
	if v.Name.Local == "ascendants" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Ascendants = append(gi.Ascendants, moduleID(innerText.(xml.CharData)))
		}
	} else if v.Name.Local == "singleDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Descendant = moduleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "exceptionalDescendant" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.ExceptionalDesc = moduleID(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "moduleName" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.Name = string(innerText.(xml.CharData))
		}
	} else if v.Name.Local == "moduleId" {
		innerText, err := decoder.Token()
		if err == nil {
			gi.ID = moduleID(innerText.(xml.CharData))
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
	generalInfo
}

func (*unknownModule) normalize(*IVRScript) error {
	return nil
}

func newUnknownModule(decoder *xml.Decoder, v *xml.StartElement) Module {
	var lastElement = v.Name.Local
	var pUnknown = new(incomingCallModule)

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
