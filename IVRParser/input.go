package IVRParser

import "encoding/xml"

type Module struct {
	XMLName   xml.Name   `xml:"input"`
	Ascendant string     `xml:"ascendants"`
	Descendant string `xml:singleDescendant"`
	Name      string     `xml:"moduleName"`
	X         int32      `xml:"locationX"`
	Y         int32      `xml:"locationY"`
	Id        string     `xml:"moduleId"`
	Data      TDataInput `xml:"data"`
}
type TDataInput struct {
	DispoName        string `xml:"dispo>name"`
	DispoId          int32  `xml:"dispo>id"`
	VPrompt TVivrPrompt `xml:"vivrPrompts"`
	VHeader TVivrHeader `xml:"vivrHeader"`
	Prompt TPrompt `xml:"prompts"`
	Grammar TGrammar `xml:"grammar"`
	Events TEvents `xml:"recoEvents"`

}
