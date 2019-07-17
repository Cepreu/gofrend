package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"reflect"

	"github.com/Cepreu/gofrend/ivr"
)

const (
	menuActionTypeEVENT  string = "EVENT"
	menuActionTypeBRANCH string = "BRANCH"
)

const (
	menuEventHELP    string = "HELP"
	menuEventNOINPUT string = "NO_INPUT"
	menuEventNOMATCH string = "NO_MATCH"
)

// type MenuItem struct {
// 	Prompt     AttemptPrompts
// 	ShowInVivr bool
// 	MatchExact bool
// 	Dtmf       string
// 	Action     struct {
// 		Type ActionType
// 		Name string
// 	}
// }

func newMenuItem(decoder *xml.Decoder, sp ivr.ScriptPrompts, itemPromptID ivr.PromptID) *ivr.MenuItem {
	var (
		choice struct {
			cType        string
			cValue       string
			cVarName     string
			cMlItem      ivr.PromptID
			cModule      ivr.ModuleID
			cModuleField string
		}
		pItem    = new(ivr.MenuItem)
		inChoice = false
		//		inThumbnail = false
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
				//			} else if v.Name.Local == "thumbnail" {
				//				inThumbnail = true
			} else if v.Name.Local == "type" {
				innerText, err := decoder.Token()
				if err == nil {
					if inChoice {
						choice.cType = string(innerText.(xml.CharData))
					}
					// } else if inThumbnail {
					// 	pItem.Thumbnail.cType = string(innerText.(xml.CharData))
					// }
				}
			} else if v.Name.Local == "value" {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					if inChoice {
						choice.cValue = string(innerText.(xml.CharData))
						// else if inThumbnail {
						// 	pItem.Thumbnail.cValue = string(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "varName" {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					if inChoice {
						choice.cVarName = string(innerText.(xml.CharData))
						// } else if inThumbnail {
						// 	pItem.Thumbnail.cVarName = string(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "module" {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					if inChoice {
						choice.cModule = ivr.ModuleID(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "moduleField" {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					if inChoice {
						choice.cModuleField = string(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "mlItem" {
				innerText, err := decoder.Token()
				if err == nil && reflect.TypeOf(innerText).String() == "xml.CharData" {
					if inChoice {
						choice.cMlItem = ivr.PromptID(innerText.(xml.CharData))
						// } else if inThumbnail {
						// 	pItem.Thumbnail.cMlItem = string(innerText.(xml.CharData))
					}
				}
			} else if v.Name.Local == "showInVivr" {
				innerText, err := decoder.Token()
				if err == nil {
					if inChoice {
						pItem.ShowInVivr = string(innerText.(xml.CharData)) == "true"
						// } else if inThumbnail {
						// 	pItem.Thumbnail.cShowInVivr = string(innerText.(xml.CharData)) == "true"
					}
				}
			} else if v.Name.Local == "match" {
				innerText, err := decoder.Token()
				if err == nil {
					pItem.MatchExact = string(innerText.(xml.CharData)) == "EXACT"
				}
			} else if v.Name.Local == "dtmf" {
				innerText, err := decoder.Token()
				if err == nil {
					pItem.Dtmf = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "ActionType" {
				innerText, err := decoder.Token()
				if err == nil {
					pItem.Action.Type = ivr.ActionType(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "actionName" {
				innerText, err := decoder.Token()
				if err == nil {
					pItem.Action.Name = string(innerText.(xml.CharData))
				}
			}
		case xml.EndElement:
			if v.Name.Local == cMenuItems {
				break F /// <----------------------------------- Return should be HERE!
			} else if v.Name.Local == "choice" {
				inChoice = false
				// } else if v.Name.Local == "thumbnail" {
				// 	inThumbnail = false
			}
		}
	}

	const tmplVar = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<speakElement>
    <attributes>
        <langAttr>
            <name>xml:lang</name>
            <attributeValueBase value="{{.Language}}"/>
       </langAttr>
    </attributes>
    <items>
        <variableElement>
            <attributes/>
            <items>
     			<textElement>
                    <attributes/>
                    <items/>
                    <body></body>
                </textElement>
            </items>
            <variableName>{{.PromptVariable}}</variableName>
        </variableElement>
    </items>
</speakElement>`

	const tmplValue = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<speakElement>
	<attributes>
		<langAttr>
			<name>xml:lang</name>
			<attributeValueBase value="{{.Language}}"/>
		</langAttr>
	</attributes>
	<items>
		<sayAsElement>
			<attributes/>
			<items>
				<textElement>
					<attributes/>
					<items/>
					<body>{{.PromptValue}}</body>
				</textElement>
			</items>
		</sayAsElement>
	</items>
</speakElement>`
	type PromptData struct {
		Language       string
		PromptValue    string
		PromptVariable string
	}

	var pp *ivr.TtsPrompt
	switch choice.cType {
	case "VALUE":
		promptdata := PromptData{Language: "en-US", PromptValue: choice.cValue}
		tmpl, err := template.New("promptTemplate").Parse(tmplValue)
		if err != nil {
			panic(err)
		}
		var doc bytes.Buffer
		err = tmpl.Execute(&doc, promptdata)
		if err != nil {
			panic(err)
		}
		pp = &ivr.TtsPrompt{TTSPromptXML: doc.String()}
		sp[itemPromptID] = pp
	case "VARIABLE":
		promptdata := PromptData{Language: "en-US", PromptVariable: choice.cVarName}
		tmpl, err := template.New("promptTemplateVar").Parse(tmplVar)
		if err != nil {
			panic(err)
		}
		var doc bytes.Buffer
		err = tmpl.Execute(&doc, promptdata)
		if err != nil {
			panic(err)
		}
		pp = &ivr.TtsPrompt{TTSPromptXML: doc.String()}
		sp[itemPromptID] = pp
	case "MODULE":
		promptdata := PromptData{Language: "en-US", PromptVariable: string(choice.cModule) + ":" + choice.cModuleField}
		//TBD: Attention: normally "ModuleName:ModuleField", here "ModuleID:ModuleField"
		tmpl, err := template.New("promptTemplateVar").Parse(tmplVar)
		if err != nil {
			panic(err)
		}
		var doc bytes.Buffer
		err = tmpl.Execute(&doc, promptdata)
		if err != nil {
			panic(err)
		}
		pp = &ivr.TtsPrompt{TTSPromptXML: doc.String()}
		sp[itemPromptID] = pp
	case "ML_ITEM":
		itemPromptID = choice.cMlItem
	}
	pItem.Prompt = ivr.AttemptPrompts{
		LangPrArr: []ivr.LanguagePrompts{
			{PrArr: []ivr.PromptID{itemPromptID}, Language: defaultLang},
		}, Count: 1}
	return pItem
}
