package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

func newCaseModule(decoder *xml.Decoder) ivr.Module {
	var (
		pCase    = new(ivr.CaseModule)
		inModule = true
	)
	for inModule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in IfElse with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "branches" {
				pCase.Branches = parseBranches(decoder)
			} else {
				parseGeneralInfo(pCase, decoder, &v)
			}

		case xml.EndElement:
			if v.Name.Local == cCase {
				inModule = false
			}
		}
	}
	return pCase
}
