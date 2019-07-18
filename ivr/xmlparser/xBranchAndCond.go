package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

// parseBranches - is used in IfElse, Case, and Menu modules
func parseBranches(script *ivr.IVRScript, decoder *xml.Decoder) (b []*ivr.OutputBranch) {
	var (
		inName     = false
		inDesc     = false
		inBranches = true
		pBranch    *ivr.OutputBranch
	)
	for inBranches {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in Query with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "entry" {
				pBranch = new(ivr.OutputBranch)
			} else if v.Name.Local == "name" {
				inName = true
			} else if v.Name.Local == "desc" {
				inDesc = true
			} else if v.Name.Local == "conditions" {
				if pCond, err := parseCondition(script, decoder); pCond != nil && err == nil {
					pBranch.Cond = &ivr.ComplexCondition{
						CustomCondition: "ALL",
						Conditions:      []*ivr.Condition{pCond},
					}
				}
			}

		case xml.CharData:
			if inName {
				pBranch.Name = string(v)
			} else if inDesc {
				pBranch.Descendant = ivr.ModuleID(v)
			}

		case xml.EndElement:
			if v.Name.Local == "entry" {
				b = append(b, pBranch)
				pBranch = nil
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "desc" {
				inDesc = false
			} else if v.Name.Local == "branches" {
				inBranches = false
			}
		}
	}
	return
}

// parseCondition - is used in IfElse and Case modules
func parseCondition(script *ivr.IVRScript, decoder *xml.Decoder) (*ivr.Condition, error) {
	var (
		inCondition = true
		pC          *ivr.Condition

		inComparisonType = false
	)
	for inCondition {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in parseCondition with '%s'\n", err)
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			if pC == nil {
				pC = new(ivr.Condition)
			}
			if v.Name.Local == "comparisonType" {
				inComparisonType = true
			} else if v.Name.Local == "rightOperand" {
				ro := new(parametrized)
				parse(ro, decoder)
				pC.RightOperand = toID(script, ro)
			} else if v.Name.Local == "leftOperand" {
				lo := new(parametrized)
				parse(lo, decoder)
				pC.LeftOperand = toID(script, lo)
			}

		case xml.CharData:
			if inComparisonType {
				pC.ComparisonType = string(v)
			}

		case xml.EndElement:
			if v.Name.Local == "comparisonType" {
				inComparisonType = false
			} else if v.Name.Local == "conditions" {
				inCondition = false
			}
		}
	}
	return pC, nil
}
