package IVRParser

import "encoding/xml"

type InputModule struct {
	XMLName    xml.Name  `xml:"input"`
	Ascendant  string    `xml:"ascendants"`
	Descendant string    `xml:singleDescendant"`
	Name       string    `xml:"moduleName"`
	X          int32     `xml:"locationX"`
	Y          int32     `xml:"locationY"`
	Id         string    `xml:"moduleId"`
	ModuleData DataInput `xml:"data"`
}

type DataInput struct {
	DispoName string       `xml:"dispo>name"`
	DispoId   int32        `xml:"dispo>id"`
	VPrompt   VivrPrompt   `xml:"vivrPrompts"`
	VHeader   VivrHeader   `xml:"vivrHeader"`
	Prompt    Prompts      `xml:"prompts"`
	Grammar   inputGrammar `xml:"grammar"`
	//	Events    RecoEvents   `xml:"recoEvents"`
}

type inputGrammar struct {
	Type           string          `xml:"type,attr"`
	MRVname        string          `xml:"mainReturnValue>name"`
	MRVtype        string          `xml:"mainReturnValue>type"`
	MRVvariable    string          `xml:"mainReturnValue>varName"`
	StrProperties  []grammPropStr  `xml:"stringProperty"`
	ListProperties []grammPropList `xml:"listProperty"`
	ARVname        string          `xml:"additionalReturnValues>name"`
	ARVtype        string          `xml:"additionalReturnValues>type"`
}
type grammPropStr struct {
	PropType    string `xml:"type"`
	PropValue   string `xml:"value"`
	PropEnabled bool   `xml:"enabled"`
}
type grammPropList struct {
	PropType    string   `xml:"type"`
	PropValEnum []string `xml:"list"`
	PropValue   string   `xml:"value"`
	PropEnabled bool     `xml:"enabled"`
}
