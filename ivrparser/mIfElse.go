package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type groupingType string //"ALL", "ANY", or "CUSTOM"

type IfElseModule struct {
	GeneralInfo
	BranchIf   OutputBranch
	BranchElse OutputBranch
}

func (*IfElseModule) transformToAI() string {
	return ""
}

func (*IfElseModule) normalize(*IVRScript) error { return nil }

func newIfElseModule(decoder *xml.Decoder) Module {
	var (
		pIE      = new(IfElseModule)
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
						pIE.BranchIf.Cond = new(ComplexCondition)
					} else if pb.Name == "ELSE" {
						pIE.BranchElse = *pb
					}
				}
			} else if v.Name.Local == "conditions" {
				if pCond, err := parseCondition(decoder); err == nil {
					pIE.BranchIf.Cond.Conditions = append(pIE.BranchIf.Cond.Conditions, pCond)
				}
			} else {
				pIE.parseGeneralInfo(decoder, &v)
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
	return pIE
}
