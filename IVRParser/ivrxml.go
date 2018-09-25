package ivrparser

import (
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

type xScript struct {
	XMLName         xml.Name             `xml:"ivrScript"`
	Domain          int32                `xml:"domainId"`
	Properties      string               `xml:"properties"`
	Modules         xModules             `xml:"modules"`
	ModulesOnHangup xModulesOnHangup     `xml:"modulesOnHangup"`
	MLPrompts       xMultilingualPrompts `xml:"multiLanguagesPrompts"`
	Variables       []xUserVariable      `xml:"userVariables>entry"`
}
type xModules struct {
	XMLName       xml.Name            `xml:"modules"`
	IncomingCall  xIncomingCallModule `xml:"incomingCall"`
	HangupModules []xHangupModule     `xml:"hangup"`
	PlayModules   []xPlayModule       `xml:"play"`
	InputModules  []xInputModule      `xml:"input"`
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
func ParseIVR(src io.Reader) (*xScript, error) {
	s := &xScript{}
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
