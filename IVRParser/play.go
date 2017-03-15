package IVRParser

type PlayModule struct {
	Ascendant  string   `xml:"ascendants"`
	Descendant string   `xml:"singleDescendant"`
	Name       string   `xml:"moduleName"`
	X          int32    `xml:"locationX"`
	Y          int32    `xml:"locationY"`
	Id         string   `xml:"moduleId"`
	Data       DataPlay `xml:"data"`
}

type DataPlay struct {
	VoicePrompts     Prompts     `xml:"prompt"`
	VivrPrompts      VivrPrompts `xml:"vivrPrompts"`
	VivrHeader       VivrHeader  `xml:"vivrHeader"`
	DispoName        string      `xml:"dispo>name"`
	DispoId          int32       `xml:"dispo>id"`
	NumberOfDigits   int32       `xml:"numberOfDigits"`
	TerminateDigit   string      `xml:"terminateDigit"`
	ClearDigitBuffer bool        `xml:"clearDigitBuffer"`
	Collapsible      bool        `xml:"collapsible"`
}
