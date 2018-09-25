package IVRParser

type xIncomingCallModule struct {
	Descendant string `xml:"singleDescendant"`
	Name       string `xml:"moduleName"`
	X          int32  `xml:"locationX"`
	Y          int32  `xml:"locationY"`
	ID         string `xml:"moduleId"`
	Data       string `xml:"data"`
}
