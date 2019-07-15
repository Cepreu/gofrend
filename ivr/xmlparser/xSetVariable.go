package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

type xmlSetVariablesModule struct {
	m *ivr.SetVariableModule
}

func newSetVariablesModule(decoder *xml.Decoder, script *ivr.IVRScript) normalizer {
	var (
		immersion = 1
		pSVM      = new(ivr.SetVariableModule)
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
				e := parseAssignment(decoder, script)
				immersion--
				if e != nil {
					pSVM.Exprs = append(pSVM.Exprs, e)
				}
			} else {
				parseGeneralInfo(pSVM, decoder, &v)
			}

		case xml.EndElement:
			immersion--
		}
	}
	return xmlSetVariablesModule{pSVM}
}

func parseAssignment(decoder *xml.Decoder, script *ivr.IVRScript) *ivr.Expression {
	var (
		immersion      = 1
		inVariableName = false
		inIsFunction   = false
		inJsFunction   = false
		inFunction     = false
		inReturnType   = false
		inName         = false
		inArguments    = false
		constVal       = new(parametrized)
		f              ivr.Function
		params         []*parametrized
		isFunction     bool

		e = new(ivr.Expression)
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
				_ = parse(constVal, decoder)
				immersion--
			} else if v.Name.Local == "function" {
				inFunction = true
				//				e.Rval.F = new(ivr.IvrFuncInvocation)
			} else if v.Name.Local == "returnType" {
				inReturnType = true
			} else if v.Name.Local == "name" || v.Name.Local == "functionName" {
				inName = true
			} else if v.Name.Local == "arguments" {
				inArguments = true

			} else if v.Name.Local == "functionArgs" {
				pp := new(parametrized)
				parse(pp, decoder)
				immersion--
				params = append(params, pp)
			}

		case xml.CharData:
			if inVariableName {
				e.Lval = string(v)
			} else if inIsFunction {
				isFunction = "true" == string(v)
			} else if inFunction {
				if inName {
					f.Name = string(v)
				} else if inReturnType {
					f.ReturnType = string(v)
				} else if inArguments {
					f.Arguments = append(f.Arguments, &ivr.FuncArgument{ArgType: string(v)})
				} else if inJsFunction {
					f.ID = ivr.FuncID(v)
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
			} else if v.Name.Local == "name" || v.Name.Local == "functionName" {
				inName = false
			} else if v.Name.Local == "arguments" {
				inArguments = false
			} else if v.Name.Local == "isFunction" {
				inIsFunction = false
			} else if v.Name.Local == "jsFunction" {
				inJsFunction = false
			}
		}
	}
	if isFunction {
		if f.ID == "" {
			id := f.ReturnType + "#" + f.Name
			for _, arg := range f.Arguments {
				id += "#" + arg.ArgType
			}
			e.Rval.FuncDef = ivr.FuncID(id)
		} else {
			e.Rval.FuncDef = f.ID
		}
		for _, p := range params {
			e.Rval.Params = append(e.Rval.Params, toID(script, p))
		}
	} else {
		e.Rval.FuncDef = "__COPY__"
		e.Rval.Params = append(e.Rval.Params, toID(script, constVal))
	}
	return e
}
func (module xmlSetVariablesModule) normalize(s *ivr.IVRScript) error {
	s.Modules[module.m.ID] = module.m
	return nil
}
