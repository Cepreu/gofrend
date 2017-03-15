package IVRParser

import "encoding/xml"

type HangupModule struct {
	XMLName   xml.Name   `xml:"hangup"`
	Ascendant string     `xml:"ascendants"`
	Name      string     `xml:"moduleName"`
	X         int32      `xml:"locationX"`
	Y         int32      `xml:"locationY"`
	Id        string     `xml:"moduleId"`
	Data      DataHangup `xml:"data"`
}
type DataHangup struct {
	DispoName        string `xml:"dispo>name"`
	DispoId          int32  `xml:"dispo>id"`
	Return2Caller    bool   `xml:"returnToCallingModule"`
	ErrIsVarSelected bool   `xml:"errCode>isVarSelected"`
	ErrIntegerValue  int32  `xml:"errCode>integerValue>value"`
	OverwriteDisp    bool   `xml:"overwriteDisposition"`
}
