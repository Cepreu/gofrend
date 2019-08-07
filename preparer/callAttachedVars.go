package preparer

import (
	"bytes"
	"encoding/xml"
	"strconv"
	"strings"
	"text/template"

	ivr "github.com/Cepreu/gofrend/ivr"
)

func getCallVarsDescrContent(callVarName string) string {
	type QueryData struct {
		CallVarsName string
		GroupName    string
	}
	const getCallVarReq = `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
<soapenv:Body>
	<ser:getCallVariables>  
		<namePattern>{{.CallVarsName}}</namePattern>
	</ser:getCallVariables>
	<groupName>{{ .GroupName }}</groupName>
</soapenv:Body>
</soapenv:Envelope>`

	const getAllCallVarReq = `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
<soapenv:Body>
	<ser:getCallVariables>  
		<namePattern>.+</namePattern>
	</ser:getCallVariables>
</soapenv:Body>
</soapenv:Envelope>`

	if args := strings.Split(callVarName, "."); len(args) == 2 {
		querydata := QueryData{args[0], args[1]}
		tmpl, err := template.New("getIVRScriptsTemplate").Parse(getCallVarReq)
		if err != nil {
			panic(err)
		}
		var doc bytes.Buffer
		err = tmpl.Execute(&doc, querydata)
		if err != nil {
			panic(err)
		}
		return doc.String()
	}
	return getAllCallVarReq
}

func getAllCallVarsDescr(auth string, s *ivr.IVRScript) (err error) {
	resp, err := queryF9(auth, func() string { return getCallVarsDescrContent("") })
	if err == nil {
		err = parseCAVResponse(resp, s)
	}
	return
}

func parseCAVResponse(resp []byte, s *ivr.IVRScript) (err error) {
	type restrictions struct {
		Type  string `xml:"type"`
		Value string `xml:"value"`
	}
	type callvar struct {
		DefaultValue  string         `xml:"defaultValue"`
		Description   string         `xml:"description"`
		Group         string         `xml:"group"`
		Name          string         `xml:"name"`
		Reporting     bool           `xml:"reporting"`
		Restrictions  []restrictions `xml:"restrictions"`
		SensitiveData bool           `xml:"sensitiveData"`
		Type          string         `xml:"type"`
	}
	type cavResp struct {
		Cavs []callvar `xml:"Body>getCallVariablesResponse>return"`
	}
	var (
		Cavs  cavResp
		val   string
		vtype ivr.ValType
	)
	if err = xml.Unmarshal(resp, &Cavs); err == nil {
		for _, v := range Cavs.Cavs {
			switch v.Type {
			case "TIME":
				//TBD				val, _ = ivr.NewTimeValue(v.DefaultValue)
				vtype = ivr.ValTime
			case "DATE":
				//TBD				val, _ = ivr.NewTimeValue(v.DefaultValue)
				vtype = ivr.ValDate
			case "NUMBER":
				scale := 0
				for _, r := range v.Restrictions {
					if r.Type == "Scale" {
						scale, _ = strconv.Atoi(r.Value)
					}
				}
				if scale == 0 {
					intVal, _ := strconv.Atoi(v.DefaultValue)
					val, _ = ivr.NewIntegerValue(intVal)
					vtype = ivr.ValInteger
				} else {
					numVal, _ := strconv.ParseFloat(v.DefaultValue, 64)
					val, _ = ivr.NewNumericValue(numVal)
					vtype = ivr.ValNumeric
				}
			case "CURRENCY":
				numVal, _ := strconv.ParseFloat(v.DefaultValue, 64)
				val, _ = ivr.NewUSCurrencyValue(numVal)
				vtype = ivr.ValCurrency
				for _, r := range v.Restrictions {
					if r.Type == "CurrencyType" {
						if r.Value == "£" {
							val, _ = ivr.NewUKCurrencyValue(numVal)
							vtype = ivr.ValCurrencyPound
						} else if r.Value == "€" {
							val, _ = ivr.NewEUCurrencyValue(numVal)
							vtype = ivr.ValCurrencyEuro
						}
						break
					}
				}
			default:
				val, _ = ivr.NewStringValue(v.DefaultValue)
				vtype = ivr.ValString
			}
			name := v.Group + "." + v.Name
			userVar := ivr.NewVariable(name, v.Description, vtype, val)
			s.Variables[ivr.VariableID(name)] = userVar
		}
	}
	return
}
