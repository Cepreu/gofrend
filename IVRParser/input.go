package ivrparser

import (
	"encoding/xml"
	"fmt"
	"strings"
)

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

type inputModule struct {
	Ascendant       string
	Descendant      string
	Name            string
	ID              string
	VoicePromptIDs  [][]PromptID
	VisualPromptIDs []PromptID
	TextPromptIDs   []PromptID
	Grammar         xInputGrammar
	Dispo           int16
}

func parseInput(decoder *xml.Decoder, v *xml.StartElement) (*inputModule, error) {
	var m xInputModule
	err := decoder.DecodeElement(&m, v)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n\n%+#v\n", m)

	var pIM = new(inputModule)
	pIM.Ascendant, pIM.Descendant = m.Ascendant, m.Descendant
	pcount := len(m.ModuleData.Prompts)
	pIM.VoicePromptIDs = make([][]PromptID, pcount)
	for i, v := range m.ModuleData.Prompts {
		if v.Count <= pcount {
			fp := strings.NewReader(v.Prompt.FullPrompt)
			pIM.VoicePromptIDs[i], _ = parseVoicePrompt(fp)
		}
	}

	return pIM, err
}
