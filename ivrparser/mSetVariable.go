package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type (
	setVariableModule struct {
		generalInfo
		Exprs []*expression
	}

	expression struct {
		Lval   string
		IsFunc bool
		Rval   assigner
	}

	assigner struct {
		P *parametrized
		F *ivrFuncInvocation
	}

	ivrFunc struct {
		Name       string
		ReturnType string
		ArgTypes   []string
	}

	ivrFuncInvocation struct {
		FuncDef ivrFunc
		Params  []*parametrized
	}
)

func (module *setVariableModule) normalize(s *IVRScript) (err error) {
	return nil
}

func newSetVariablesModule(decoder *xml.Decoder) Module {
	var (
		immersion = 1
		pSVM      = new(setVariableModule)
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in SetVariable with '%s', imm=%d\n", err, immersion)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "expressions" {
				e := parseAssignment(decoder)
				immersion--
				if e != nil {
					pSVM.Exprs = append(pSVM.Exprs, e)
				}
			} else {
				pSVM.parseGeneralInfo(decoder, &v)
			}

		case xml.EndElement:
			immersion--
		}
	}
	return pSVM
}

func parseAssignment(decoder *xml.Decoder) *expression {
	var (
		immersion      = 1
		inVariableName = false
		inIsFunction   = false
		inFunction     = false
		inReturnType   = false
		inName         = false
		inArguments    = false

		e = new(expression)
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in SetVariable with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "variableName" {
				inVariableName = true
			} else if v.Name.Local == "isFunction" {
				inIsFunction = true
			} else if v.Name.Local == "constant" {
				e.Rval.P = new(parametrized)
				e.Rval.P.parse(decoder)
				immersion--
			} else if v.Name.Local == "function" {
				inFunction = true
				e.Rval.F = new(ivrFuncInvocation)
			} else if v.Name.Local == "returnType" {
				inReturnType = true
			} else if v.Name.Local == "name" {
				inName = true
			} else if v.Name.Local == "arguments" {
				inArguments = true

			} else if v.Name.Local == "functionArgs" {
				p := new(parametrized)
				p.parse(decoder)
				immersion--
				e.Rval.F.Params = append(e.Rval.F.Params, p)
			}

		case xml.CharData:
			if inVariableName {
				e.Lval = string(v)
			} else if inIsFunction {
				e.IsFunc = "true" == string(v)
			} else if inFunction {
				if inName {
					e.Rval.F.FuncDef.Name = string(v)
				} else if inReturnType {
					e.Rval.F.FuncDef.ReturnType = string(v)
				} else if inArguments {
					e.Rval.F.FuncDef.ArgTypes = append(e.Rval.F.FuncDef.ArgTypes, string(v))
				}
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "variableName" {
				inVariableName = false
			} else if v.Name.Local == "function" {
				inFunction = false
			} else if v.Name.Local == "returnType" {
				inReturnType = false
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "arguments" {
				inArguments = false
			} else if v.Name.Local == "isFunction" {
				inIsFunction = false
			}
		}
	}
	return e
}
