package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type xHangupModule struct {
	XMLName    xml.Name    `xml:"hangup"`
	Ascendant  string      `xml:"ascendants"`
	Name       string      `xml:"moduleName"`
	X          int32       `xml:"locationX"`
	Y          int32       `xml:"locationY"`
	ID         string      `xml:"moduleId"`
	ModuleData xDataHangup `xml:"data"`
}
type xDataHangup struct {
	DispoName        string `xml:"dispo>name"`
	DispoID          int32  `xml:"dispo>id"`
	Return2Caller    bool   `xml:"returnToCallingModule"`
	ErrIsVarSelected bool   `xml:"errCode>isVarSelected"`
	ErrIntegerValue  int32  `xml:"errCode>integerValue>value"`
	OverwriteDisp    bool   `xml:"overwriteDisposition"`
}

func parseHangup(decoder *xml.Decoder, v *xml.StartElement, inModulesOnHangup bool) error {
	var m xHangupModule
	err := decoder.DecodeElement(&m, v)
	if err == nil {
		fmt.Println(inModulesOnHangup)
		fmt.Printf("\n\n%+#v\n", m)
	}
	return err
}
