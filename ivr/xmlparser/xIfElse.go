package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

type xmlIfElseModule struct {
	s *ivr.IfElseModule
}

func newIfElseModule(decoder *xml.Decoder) normalizer {
	var (
		pIE      = new(ivr.IfElseModule)
		inModule = true

		inCustomCondition = false
	)

	for inModule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in IfElse with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "customCondition" {
				inCustomCondition = true
			} else if v.Name.Local == "branches" {
				p := parseBranches(decoder)
				for _, pb := range p {
					if pb.Name == "IF" {
						pIE.BranchIf = *pb
						pIE.BranchIf.Cond = new(ivr.ComplexCondition)
					} else if pb.Name == "ELSE" {
						pIE.BranchElse = *pb
					}
				}
			} else if v.Name.Local == "conditions" {
				if pCond, err := parseCondition(decoder); err == nil {
					pIE.BranchIf.Cond.Conditions = append(pIE.BranchIf.Cond.Conditions, pCond)
				}
			} else {
				parseGeneralInfo(pIE, decoder, &v)
			}
		case xml.CharData:
			if inCustomCondition {
				pIE.BranchIf.Cond.CustomCondition = string(v)
			}

		case xml.EndElement:
			if v.Name.Local == "customCondition" {
				inCustomCondition = false
			} else if v.Name.Local == cIfElse {
				inModule = false
			}
		}
	}
	return xmlIfElseModule{pIE}
}

func (module xmlIfElseModule) normalize(s *ivr.IVRScript) error {
	s.Modules[module.s.ID] = module.s
	return nil
}
