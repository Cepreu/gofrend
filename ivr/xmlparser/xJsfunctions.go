package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"

	"github.com/Cepreu/gofrend/utils"
)

func newJSFunctions(decoder *xml.Decoder) []*ivr.Function {
	var (
		inFunctions = true
		fa          = []*ivr.Function{}
	)
	for inFunctions {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "entry" {
				fa = append(fa, parseJSFunction(decoder))
			}
		case xml.EndElement:
			if v.Name.Local == "functions" {
				inFunctions = false
			}
		}
	}
	return fa
}

func parseJSFunction(decoder *xml.Decoder) *ivr.Function {
	var (
		inEntry = true
		f       = ivr.Function{Type: ivr.FuncJS}

		inJSFunctionID = false
		inDescription  = false
		inReturnType   = false
		inName         = false
		inFunctionBody = false
		inArguments    = false
	)
	for inEntry {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "jsFunctionId" {
				inJSFunctionID = true
			} else if v.Name.Local == "description" {
				inDescription = true
			} else if v.Name.Local == "returnType" {
				inReturnType = true
			} else if v.Name.Local == "name" {
				inName = true
			} else if v.Name.Local == "functionBody" {
				inFunctionBody = true
			} else if v.Name.Local == "arguments" {
				if inArguments == false {
					inArguments = true
				} else {
					f.Arguments = append(f.Arguments, parseArgument(decoder))
				}
			}
		case xml.CharData:
			if inName {
				f.Name = string(v)
			} else if inJSFunctionID {
				f.ID = ivr.FuncID(v)
			} else if inReturnType {
				f.ReturnType = string(v)
			} else if inDescription {
				f.Description = string(v)
			} else if inFunctionBody {
				b, err := utils.CmdUnzip(string(v))
				if err == nil {
					f.Body = b
				}
			}

		case xml.EndElement:
			if v.Name.Local == "jsFunctionId" {
				inJSFunctionID = false
			} else if v.Name.Local == "description" {
				inDescription = false
			} else if v.Name.Local == "returnType" {
				inReturnType = false
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "functionBody" {
				inFunctionBody = false
			} else if v.Name.Local == "entry" {
				inEntry = false
			}
		}
	}
	return &f
}
func parseArgument(decoder *xml.Decoder) *ivr.FuncArgument {
	var (
		a          = ivr.FuncArgument{}
		inArgument = true

		inDescription = false
		inName        = false
		inArgType     = false
	)
	for inArgument {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "description" {
				inDescription = true
			} else if v.Name.Local == "type" {
				inArgType = true
			} else if v.Name.Local == "name" {
				inName = true
			}
		case xml.CharData:
			if inName {
				a.Name = string(v)
			} else if inArgType {
				a.ArgType = string(v)
			} else if inDescription {
				a.Description = string(v)
			}

		case xml.EndElement:
			if v.Name.Local == "description" {
				inDescription = false
			} else if v.Name.Local == "type" {
				inArgType = false
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "arguments" {
				inArgument = false
			}
		}
	}
	return &a
}
