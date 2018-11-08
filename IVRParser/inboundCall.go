package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type xIncomingCallModule struct {
	Descendant string `xml:"singleDescendant"`
	Name       string `xml:"moduleName"`
	X          int32  `xml:"locationX"`
	Y          int32  `xml:"locationY"`
	ID         string `xml:"moduleId"`
	Data       string `xml:"data"`
}

func parseIncomingCall(decoder *xml.Decoder, v *xml.StartElement, inModulesOnHangup bool) error {
	var m xIncomingCallModule
	err := decoder.DecodeElement(&m, v)
	if err == nil {
		fmt.Println(inModulesOnHangup)
		fmt.Printf("\n\n%+#v\n", m)
	}
	return err
}
