package ivrparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/utils"
)

type (
	jsFunction struct {
		jsFunctionID string
		description  string
		returnType   string //varType
		name         string
		arguments    []*funcArgument
		funcBody     string
	}
	funcArgument struct {
		name        string
		description string
		argType     string
	}
)

func newJSFunctions(decoder *xml.Decoder) []*jsFunction {
	var (
		inFunctions = true
		fa          = []*jsFunction{}
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

func parseJSFunction(decoder *xml.Decoder) *jsFunction {
	var (
		inEntry = true
		f       jsFunction

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
					f.arguments = append(f.arguments, parseArgument(decoder))
				}
			}
		case xml.CharData:
			if inName {
				f.name = string(v)
			} else if inJSFunctionID {
				f.jsFunctionID = string(v)
			} else if inReturnType {
				f.returnType = string(v)
			} else if inDescription {
				f.description = string(v)
			} else if inFunctionBody {
				b, err := utils.CmdUnzip(string(v))
				if err == nil {
					f.funcBody = b
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
func parseArgument(decoder *xml.Decoder) *funcArgument {
	var (
		a          = funcArgument{}
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
				a.name = string(v)
			} else if inArgType {
				a.argType = string(v)
			} else if inDescription {
				a.description = string(v)
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
