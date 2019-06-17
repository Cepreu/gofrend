package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type (
	CaseModule struct {
		GeneralInfo
		Branches []*OutputBranch
	}
)

func (*CaseModule) transformToAI() string {
	return ""
}

func (*CaseModule) normalize(*IVRScript) error { return nil }

func newCaseModule(decoder *xml.Decoder) Module {
	var (
		pCase    = new(CaseModule)
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
				pCase.parseGeneralInfo(decoder, &v)
			}

		case xml.EndElement:
			if v.Name.Local == cCase {
				inModule = false
			}
		}
	}
	return pCase
}
