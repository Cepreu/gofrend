package xmlparser

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/Cepreu/gofrend/ivr"
	"github.com/Cepreu/gofrend/utils"
)

type xmlQueryModule struct {
	m *ivr.QueryModule
}

func newQueryModule(decoder *xml.Decoder, sp ivr.ScriptPrompts) normalizer {
	var (
		immersion      = 1
		pQM            = new(ivr.QueryModule)
		inURLParts     = false
		inURL          = false
		inMethod       = false
		inBodyType     = false
		inFetchTimeout = false
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in Query with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			//			fmt.Println(immersion, " <<<<", v.Name.Local)
			if v.Name.Local == "fetchTimeout" {
				inFetchTimeout = true
			} else if v.Name.Local == "url" {
				inURL = true
			} else if v.Name.Local == "method" {
				inMethod = true
			} else if v.Name.Local == "urlParts" {
				inURLParts = true
			} else if v.Name.Local == "holder" && inURLParts {
				thePart := new(ivr.Parametrized)
				parse(thePart, decoder)
				immersion--
				pQM.URLParts = append(pQM.URLParts, thePart)
			} else if v.Name.Local == "parameters" {
				pQM.Parameters, _ = parseKeyValueListParmetrized(decoder)
				immersion--
			} else if v.Name.Local == "headers" {
				pQM.Headers, _ = parseKeyValueListParmetrized(decoder)
				immersion--
			} else if v.Name.Local == "requestInfo" {
				parseRequestInfo(&pQM.RequestInfo, decoder)
				immersion--
			} else if v.Name.Local == "returnValues" {
				pQM.ReturnValues, _ = parseKeyValueList(decoder)
				immersion--
			} else if v.Name.Local == "responseInfos" {
				ri, _ := parseResponseInfo(decoder)
				immersion--
				pQM.ResponseInfos = append(pQM.ResponseInfos, ri)
			} else if v.Name.Local == "requestBodyType" {
				inBodyType = true

				///// prompts -->
			} else if v.Name.Local == cPrompt {
				if res, err := parseVoicePrompt(decoder, &v, sp, fmt.Sprintf("%s_%s_", pQM.ID, "P")); err == nil {
					pQM.VoicePromptIDs, _ = newModulePrompts(1, res)
				}
				immersion--
			} else {
				parseGeneralInfo(pQM, decoder, &v)
			}

		case xml.CharData:
			if inURL {
				pQM.URL = string(v)
			} else if inMethod {
				pQM.Method = string(v)
			} else if inBodyType {
				pQM.RequestBodyType = string(v)
			} else if inFetchTimeout {
				pQM.FetchTimeout, _ = strconv.Atoi(string(v))
			}

		case xml.EndElement:
			immersion--
			//			fmt.Println(immersion, " ", v.Name.Local, ">>>>")
			if v.Name.Local == "urlParts" {
				inURLParts = false
			} else if v.Name.Local == "url" {
				inURL = false
			} else if v.Name.Local == "method" {
				inMethod = false
			} else if v.Name.Local == "requestBodyType" {
				inBodyType = false
			} else if v.Name.Local == "FetchTimeout" {
				inFetchTimeout = false
			}
		}
	}

	return xmlQueryModule{pQM}
}

///-------------
func parseRequestInfo(p *ivr.RequestInfo, decoder *xml.Decoder) error {
	var (
		immersion      = 1
		tmp            *ivr.Replacement
		inTemplate     = false
		inReplacements = false
		inPosition     = false
		inVariableName = false
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			return err
		}

		switch v := t.(type) {

		case xml.StartElement:
			immersion++
			if v.Name.Local == "base64" {
				inTemplate = true
			} else if v.Name.Local == "replacements" {
				inReplacements = true
				tmp = new(ivr.Replacement)
			} else if v.Name.Local == "position" {
				inPosition = true
			} else if v.Name.Local == "variableName" {
				inVariableName = true
			}

		case xml.CharData:
			if inTemplate {
				p.Base64 = string(v)
			} else if inPosition && inReplacements {
				tmp.Position, _ = strconv.Atoi(string(v))
			} else if inVariableName && inReplacements {
				tmp.VariableName = string(v)
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "base64" {
				inTemplate = false
			} else if v.Name.Local == "replacements" {
				p.Replacements = append(p.Replacements, tmp)
			} else if v.Name.Local == "position" {
				inPosition = false
			} else if v.Name.Local == "variableName" {
				inVariableName = false
			}
		}
	}
	return nil
}

/////////////
func parseResponseInfo(decoder *xml.Decoder) (p *ivr.ResponseInfo, err error) {
	var (
		immersion = 1
	)
	p = new(ivr.ResponseInfo)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "from" {
				innerText, err := decoder.Token()
				if err == nil {
					p.HTTPCodeFrom, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "to" {
				innerText, err := decoder.Token()
				if err == nil {
					p.HTTPCodeTo, _ = strconv.Atoi(string(innerText.(xml.CharData)))
				}
			} else if v.Name.Local == "method" {
				innerText, err := decoder.Token()
				if err == nil {
					p.ParsingMethod = string(innerText.(xml.CharData))
				}
			} else if v.Name.Local == "regexp" {
				parseRegexpInfo(p, decoder)
				immersion--
			} else if v.Name.Local == "func" {
				parseFunctionInfo(p, decoder)
				immersion--
			}

		case xml.EndElement:
			immersion--
		}
	}
	return p, nil
}

//---->
func parseFunctionInfo(p *ivr.ResponseInfo, decoder *xml.Decoder) error {
	var (
		immersion    = 1
		inReturnType = false
		inName       = false
		inArguments  = false
		inVariable   = false
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			return err
		}

		switch v := t.(type) {

		case xml.StartElement:
			immersion++
			if v.Name.Local == "returnType" {
				inReturnType = true
			} else if v.Name.Local == "name" {
				inName = true
			} else if v.Name.Local == "arguments" {
				inArguments = true
			} else if v.Name.Local == "variable" {
				inVariable = true
			}

		case xml.CharData:
			if inReturnType {
				p.Function.ReturnType = string(v)
			} else if inName {
				p.Function.Name = string(v)
			} else if inArguments {
				p.Function.Arguments = string(v)
			} else if inVariable {
				p.TargetVariables = append(p.TargetVariables, string(v))
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "returnType" {
				inReturnType = false
			} else if v.Name.Local == "name" {
				inName = false
			} else if v.Name.Local == "arguments" {
				inArguments = false
			} else if v.Name.Local == "variable" {
				inVariable = false
			}
		}
	}

	return nil
}

//+++++++++
func parseRegexpInfo(p *ivr.ResponseInfo, decoder *xml.Decoder) error {
	var (
		immersion                        = 1
		inRegexpParameters, inGroupIndex = false, false
		inRegexp, inRegexpFlags, inValue = false, false, false
		key                              int
		value                            string
		tmpvars                          []struct {
			ind     int
			varname string
		}
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			return err
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "regexp" {
				inRegexp = true
			} else if v.Name.Local == "regexpFlags" {
				inRegexpFlags = true
			} else if v.Name.Local == "regexpParameters" {
				inRegexpParameters = true
			} else if v.Name.Local == "groupIndex" && inRegexpParameters {
				inGroupIndex = true
			} else if v.Name.Local == "variable" && inRegexpParameters {
				inValue = true
			}

		case xml.CharData:
			if inRegexp {
				p.Regexp.RegexpBody = string(v)
			} else if inRegexpFlags {
				p.Regexp.RegexpFlags, _ = strconv.Atoi(string(v))
			} else if inGroupIndex {
				key, _ = strconv.Atoi(string(v))
			} else if inValue {
				value = string(v)
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "regexp" {
				inRegexp = false
			} else if v.Name.Local == "regexpFlags" {
				inRegexpFlags = false
			} else if v.Name.Local == "regexpParameters" {
				inRegexpParameters = false
				tmpvars = append(tmpvars, struct {
					ind     int
					varname string
				}{key, value})
			} else if v.Name.Local == "groupIndex" && inRegexpParameters {
				inGroupIndex = false
			} else if v.Name.Local == "variable" && inRegexpParameters {
				inValue = false
			}
		}
	}
	p.TargetVariables = make([]string, len(tmpvars), len(tmpvars))
	for _, s := range tmpvars {
		if s.ind < len(tmpvars) {
			p.TargetVariables[s.ind] = s.varname
		}
	}
	return nil
}

/////////////
func parseKeyValueList(decoder *xml.Decoder) (p []ivr.KeyValue, err error) {

	var (
		immersion               = 1
		inEntry, inKey, inValue = false, false, false
		key                     string
		value                   string
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "entry" {
				inEntry = true
			} else if v.Name.Local == "key" && inEntry {
				inKey = true
			} else if v.Name.Local == "value" && inEntry {
				inValue = true
			}

		case xml.CharData:
			if inKey {
				key = string(v)
			} else if inValue {
				value = string(v)
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "entry" {
				inEntry = false
				p = append(p, ivr.KeyValue{key, value})
			} else if v.Name.Local == "key" && inEntry {
				inKey = false
			} else if v.Name.Local == "value" && inEntry {
				inValue = false
			}
		}
	}
	return p, nil
}

func (module xmlQueryModule) normalize(s *ivr.IVRScript) error {
	err := normalizePrompt(s, module.m.VoicePromptIDs)
	if err != nil {
		return err
	}
	module.m.RequestInfo.Template, err = utils.CmdUnzip(module.m.RequestInfo.Base64)
	return err
}
