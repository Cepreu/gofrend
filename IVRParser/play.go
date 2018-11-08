package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type xPlayModule struct {
	Ascendant  string    `xml:"ascendants"`
	Descendant string    `xml:"singleDescendant"`
	Name       string    `xml:"moduleName"`
	X          int32     `xml:"locationX"`
	Y          int32     `xml:"locationY"`
	ID         string    `xml:"moduleId"`
	ModuleData xDataPlay `xml:"data"`
}

type xDataPlay struct {
	VoicePrompts     xPrompts     `xml:"prompt"`
	VivrPrompts      xVivrPrompts `xml:"vivrPrompts"`
	VivrHeader       xVivrHeader  `xml:"vivrHeader"`
	DispoName        string       `xml:"dispo>name"`
	DispoID          int32        `xml:"dispo>id"`
	NumberOfDigits   int32        `xml:"numberOfDigits"`
	TerminateDigit   string       `xml:"terminateDigit"`
	ClearDigitBuffer bool         `xml:"clearDigitBuffer"`
	Collapsible      bool         `xml:"collapsible"`
}

func parsePlay(decoder *xml.Decoder, v *xml.StartElement) (*xPlayModule, error) {
	var m = new(xPlayModule)
	err := decoder.DecodeElement(m, v)
	if err == nil {
		fmt.Printf("\n\n%+#v\n", &m)
	}
	return m, err
}
