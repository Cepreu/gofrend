package ivrparser

import (
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

type Script struct {
	Domain          int32            `xml:"domainId"`
	Properties      string           `xml:"properties"`
	Modules         modules          `xml:"modules"`
	ModulesOnHangup xModulesOnHangup `xml:"modulesOnHangup"`
	//	MLPrompts       xMultilingualPrompts `xml:"multiLanguagesPrompts"`
	Prompts map[string]*prompt
	//	Variables map[string]*Variable
}
type modules struct {
	IncomingCall  *xIncomingCallModule
	HangupModules []*xHangupModule
	PlayModules   []*xPlayModule
	InputModules  []*inputModule
}
type xModulesOnHangup struct {
	XMLName       xml.Name            `xml:"modulesOnHangup"`
	StartOnHangup xIncomingCallModule `xml:"startOnHangup"`
	HangupModules []xHangupModule     `xml:"hangup"`
}

type xUserVariable struct {
	Key           string `xml:"key"`
	Name          string `xml:"value>name"`
	Description   string `xml:"value>description"`
	StringValue   string `xml:"value>stringValue>value"`
	StringValueID int32  `xml:"value>stringValue>id"`
	Attributes    int32  `xml:"value>attributes"`
	IsNullValue   bool   `xml:"isNullValue"`
}

//ParseIVR - Parsing of the getIVRResponse received from Five9 Config web service
func ParseIVR(src io.Reader) (*Script, error) {
	s := &Script{}
	decoder := xml.NewDecoder(src)
	decoder.CharsetReader = charset.NewReaderLabel

	inModules := false
	inVariables := false
	inMLPrompts := false
	inMLVIVRPrompts := false
	inMLMenuChoices := false
	inModulesOnHangup := false

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			// io.EOF is a successful end
			break
		}
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {

		case xml.StartElement:
			if v.Name.Local == "modules" {
				inModules = true
			} else if v.Name.Local == "modulesOnHangup" {
				inModulesOnHangup = true
			} else if v.Name.Local == "userVariables" {
				inVariables = true
			} else if v.Name.Local == "multiLanguagesVIVRPrompts" {
				inMLVIVRPrompts = true
			} else if v.Name.Local == "multiLanguagesPrompts" {
				inMLPrompts = true
			} else if v.Name.Local == "multiLanguagesMenuChoices" {
				inMLMenuChoices = true

			} else if v.Name.Local == "incomingCall" || v.Name.Local == "startOnHangup" {
				if err = parseIncomingCall(decoder, &v, inModulesOnHangup); err != nil {
					fmt.Printf("parseIncomingCall() failed with '%s'\n", err)
					fmt.Println(inMLPrompts, inMLMenuChoices, inMLVIVRPrompts, inVariables, inModules)
					break
				}
			} else if v.Name.Local == "hangup" {
				if err = parseHangup(decoder, &v, inModulesOnHangup); err != nil {
					fmt.Printf("parseHangup() failed with '%s'\n", err)
					break
				}
			} else if v.Name.Local == "play" {
				m, err := parsePlay(decoder, &v)
				if err != nil {
					fmt.Printf("parsePlay() failed with '%s'\n", err)
					break
				}
				s.Modules.PlayModules = append(s.Modules.PlayModules, m)

			} else if v.Name.Local == "input" {
				m, err := parseInput(decoder, &v)
				if err != nil {
					fmt.Printf("parsePlay() failed with '%s'\n", err)
					break
				}
				s.Modules.InputModules = append(s.Modules.InputModules, m)
			}
		case xml.EndElement:
			if v.Name.Local == "modules" {
				inModules = false
			} else if v.Name.Local == "modulesOnHangup" {
				inModulesOnHangup = false
			} else if v.Name.Local == "userVariables" {
				inVariables = false
			} else if v.Name.Local == "multiLanguagesVIVRPrompts" {
				inMLVIVRPrompts = false
			} else if v.Name.Local == "multiLanguagesPrompts" {
				inMLPrompts = false
			} else if v.Name.Local == "multiLanguagesMenuChoices" {
				inMLMenuChoices = false
			}
		}
	}

	return s, nil
}
