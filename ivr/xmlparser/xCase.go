package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

type xmlCaseModule struct {
	m *ivr.CaseModule
}

func newCaseModule(script *ivr.IVRScript, decoder *xml.Decoder) normalizer {
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
				pCase.Branches = parseBranches(script, decoder)
			} else {
				parseGeneralInfo(pCase, decoder, &v)
			}

		case xml.EndElement:
			if v.Name.Local == cCase {
				inModule = false
			}
		}
	}
	return xmlCaseModule{pCase}
}

func (module xmlCaseModule) normalize(s *ivr.IVRScript) error {
	s.Modules[module.m.ID] = module.m
	return nil
}
