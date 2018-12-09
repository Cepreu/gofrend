package ivrparser

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

const (
	menuMatchAPPR  string = "APPR"
	menuMatchEXACT string = "EXACT"
)
const (
	menuActionTypeEVENT  string = "EVENT"
	menuActionTypeBRANCH string = "BRANCH"
)
const (
	menuChoiceTypeVALUE    string = "VALUE"
	menuChoiceTypeVARIABLE string = "VARIABLE"
	menuChoiceTypeMLITEM   string = "ML_ITEM"
)
const (
	menuEventHELP    string = "HELP"
	menuEventNOINPUT string = "NO_INPUT"
	menuEventNOMATCH string = "NO_MATCH"
)

type menuChoice struct {
	Type       string
	Value      string
	VarName    string
	MlItem     string
	ShowInVivr bool
}

type menuItem struct {
	Choice    menuChoice
	Match     string
	Thumbnail menuChoice
	Dtmf      string
	Action    struct {
		Type actionType
		Name string
	}
}

func (s *IVRScript) newMenuItem(decoder *xml.Decoder, v *xml.StartElement, prefix string) *menuItem {
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}
	var (
		pItem       = new(menuItem)
		inChoice    = false
		inThumbnail = false
	)
F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "choice" {
				inChoice = true
			} else if v.Name.Local == "thumbnail" {
				inThumbnail = true
			} else if v.Name.Local == "type" {
				innerText, err := decoder.Token()
				if err == nil {
					if inChoice {
						pItem.Choice.Type = string(innerText.(xml.CharData))
					} else if inThumbnail {
						pItem.Thumbnail.Type = string(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "value" {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					if inChoice {
						pItem.Choice.Value = string(innerText.(xml.CharData))
					} else if inThumbnail {
						pItem.Thumbnail.Value = string(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "varName" {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					if inChoice {
						pItem.Choice.VarName = string(innerText.(xml.CharData))
					} else if inThumbnail {
						pItem.Thumbnail.VarName = string(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "mlItem" {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					if inChoice {
						pItem.Choice.MlItem = string(innerText.(xml.CharData))
					} else if inThumbnail {
						pItem.Thumbnail.MlItem = string(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "showInVivr" {
				innerText, err := decoder.Token()
				if err == nil {
					if inChoice {
						pItem.Choice.ShowInVivr = string(innerText.(xml.CharData)) == "true"
					} else if inThumbnail {
						pItem.Thumbnail.ShowInVivr = string(innerText.(xml.CharData)) == "true"
					}
				}
			} else if v.Name.Local == "match" {
				innerText, err := decoder.Token()
				if err == nil {
					pItem.Match = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "dtmf" {
				innerText, err := decoder.Token()
				if err == nil {
					pItem.Dtmf = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "actionType" {
				innerText, err := decoder.Token()
				if err == nil {
					pItem.Action.Type = actionType(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "actionName" {
				innerText, err := decoder.Token()
				if err == nil {
					pItem.Action.Name = string(innerText.(xml.CharData))
				}
			}
		case xml.EndElement:
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!
			} else if v.Name.Local == "choice" {
				inChoice = false
			} else if v.Name.Local == "thumbnail" {
				inThumbnail = false
			}
		}
	}
	return pItem
}
