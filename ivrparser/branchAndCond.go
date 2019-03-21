package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type (
	// OutputBranch - the Menu,IfElse, and Case modules' branch
	OutputBranch struct {
		Name       string
		Descendant ModuleID
		Cond       *ComplexCondition
	}

	// ComplexCondition - a combinedcondition
	ComplexCondition struct {
		CustomCondition   string
		ConditionGrouping groupingType
		Conditions        []*Condition
	}

	//Condition - ifElse module's condition
	Condition struct {
		comparisonType string
		rightOperand   parametrized
		leftOperand    parametrized
	}
)

// parseBranches - is used in IfElse, Case, and Menu modules
func parseBranches(decoder *xml.Decoder) (b []*OutputBranch) {
	var (
		inName     = false
		inDesc     = false
		inBranches = true
		pBranch    *OutputBranch
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
				pBranch = new(OutputBranch)
			} else if v.Name.Local == "name" {
				inName = true
			} else if v.Name.Local == "desc" {
				inDesc = true
			} else if v.Name.Local == "conditions" {
				if pCond, err := parseCondition(decoder); pCond != nil && err == nil {
					pBranch.Cond = &ComplexCondition{
						CustomCondition:   "1",
						ConditionGrouping: "AND",
						Conditions:        []*Condition{pCond},
					}
				}
			}

		case xml.CharData:
			if inName {
				pBranch.Name = string(v)
			} else if inDesc {
				pBranch.Descendant = ModuleID(v)
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
func parseCondition(decoder *xml.Decoder) (*Condition, error) {
	var (
		inCondition = true
		pC          *Condition

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
				pC = new(Condition)
			}
			if v.Name.Local == "comparisonType" {
				inComparisonType = true
			} else if v.Name.Local == "rightOperand" {
				pC.rightOperand.parse(decoder)
			} else if v.Name.Local == "leftOperand" {
				pC.leftOperand.parse(decoder)
			}

		case xml.CharData:
			if inComparisonType {
				pC.comparisonType = string(v)
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
