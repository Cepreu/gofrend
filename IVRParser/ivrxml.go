package IVRParser

import (
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

type Script struct {
	XMLName    xml.Name            `xml:"ivrScript"`
	Domain     int32               `xml:"domainId"`
	Properties string              `xml:"properties"`
	Modules    Modules             `xml:"modules"`
	MLPrompts  multilingualPrompts `xml:"multiLanguagesPrompts"`
	Variables  []xUserVariable     `xml:"userVariables>entry"`
}
type Modules struct {
	XMLName   xml.Name             `xml:"modules"`
	HModules  []HangupModule       `xml:"hangup"`
	ICModules []IncomingCallModule `xml:"incomingCall"`
	PModules  []PlayModule         `xml:"play"`
	IModules  []InputModule        `xml:"input"`
}

type xUserVariable struct {
	Key           string `xml:"key"`
	Name          string `xml:"value>name"`
	Description   string `xml:"value>description"`
	StringValue   string `xml:"value>stringValue>value"`
	StringValueId int32  `xml:"value>stringValue>id"`
	Attributes    int32  `xml:"value>attributes"`
	IsNullValue   bool   `xml:"isNullValue"`
}

/*
func toMap(vars ...Module) map[string]interface{} {
	m := make(map[string]interface{})
	for _, v := range vars {
		if len(v.Children) > 0 {
			m[v.Key] = toMap(v.Children...)
		} else {
			m[v.Key] = v.Value
		}
	}
	return m
}
*/
func ParseIVR(src io.Reader) (*Script, error) {
	s := &Script{}
	decoder := xml.NewDecoder(src)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("s: %#v\n", *s)

	//	m := toMap(s.Test.Vars...)
	//	fmt.Printf("map: %v\n", m)

	/*	js, err := json.MarshalIndent(m, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("json: %s", js)
	*/
	return s, nil
}
