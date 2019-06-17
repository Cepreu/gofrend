package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

func newForeignScriptModule(decoder *xml.Decoder) ivr.Module {
	var (
		inModule       = true
		pFSM           = new(ivr.ForeignScriptModule)
		inIvrScript    = false
		inName         = false
		inPassCRM      = false
		inReturnCRM    = false
		inIsConsistent = false
	)

	for inModule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in ForeignScript with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			//			fmt.Println(immersion, " <<<<", v.Name.Local)
			if v.Name.Local == "ivrScript" {
				inIvrScript = true
			} else if v.Name.Local == "name" && inIvrScript {
				inName = true
			} else if v.Name.Local == "passCRM" {
				inPassCRM = true
			} else if v.Name.Local == "returnCRM" {
				inReturnCRM = true
			} else if v.Name.Local == "params" {
				pFSM.Parameters, _ = parseKeyValueListParmetrized(decoder)
			} else if v.Name.Local == "returnVals" {
				pFSM.ReturnParameters, _ = parseKeyValueList(decoder)
			} else if v.Name.Local == "isConsistent" {
				inIsConsistent = true
			} else {
				parseGeneralInfo(pFSM, decoder, &v)
			}

		case xml.CharData:
			if inName {
				pFSM.IvrScript = string(v)
			} else if inPassCRM {
				pFSM.PassCRM = string(v) == "true"
			} else if inReturnCRM {
				pFSM.ReturnCRM = string(v) == "true"
			} else if inIsConsistent {
				pFSM.IsConsistent = string(v) == "true"
			}

		case xml.EndElement:
			if v.Name.Local == cForeignScript {
				inModule = false
			} else if v.Name.Local == "ivrScript" {
				inIvrScript = false
			} else if v.Name.Local == "name" && inIvrScript {
				inName = false
			} else if v.Name.Local == "passCRM" {
				inPassCRM = false
			} else if v.Name.Local == "returnCRM" {
				inReturnCRM = false
			} else if v.Name.Local == "isConsistent" {
				inIsConsistent = false
			}
		}
	}
	return pFSM
}
