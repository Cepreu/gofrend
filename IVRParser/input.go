package ivrparser

import "encoding/xml"

type xInputModule struct {
	XMLName    xml.Name   `xml:"input"`
	Ascendant  string     `xml:"ascendants"`
	Descendant string     `xml:"singleDescendant"`
	Name       string     `xml:"moduleName"`
	X          int32      `xml:"locationX"`
	Y          int32      `xml:"locationY"`
	ID         string     `xml:"moduleId"`
	ModuleData xDataInput `xml:"data"`
}
type xDataInput struct {
	DispoName string           `xml:"dispo>name"`
	DispoID   int32            `xml:"dispo>id"`
	VPrompts  xVivrPrompts     `xml:"vivrPrompts"`
	VHeader   xVivrHeader      `xml:"vivrHeader"`
	TChData   xTextChanneddata `xml:"textChannelData"`
	Grammar   xInputGrammar    `xml:"grammar"`
	Prompts   []xPrompts       `xml:"prompts"`
	//	Events    RecoEvents   `xml:"recoEvents"`
}
type xInputGrammar struct {
	Type           string           `xml:"type,attr"`
	MRVname        string           `xml:"mainReturnValue>name"`
	MRVtype        string           `xml:"mainReturnValue>type"`
	MRVvariable    string           `xml:"mainReturnValue>varName"`
	StrProperties  []xGrammPropStr  `xml:"stringProperty"`
	ListProperties []xGrammPropList `xml:"listProperty"`
	ARVname        string           `xml:"additionalReturnValues>name"`
	ARVtype        string           `xml:"additionalReturnValues>type"`
}
type xGrammPropStr struct {
	PropType    string `xml:"type"`
	PropValue   string `xml:"value"`
	PropEnabled bool   `xml:"enabled"`
}
type xGrammPropList struct {
	PropType    string   `xml:"type"`
	PropValEnum []string `xml:"list"`
	PropValue   string   `xml:"value"`
	PropEnabled bool     `xml:"enabled"`
}
