package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Cepreu/gofrend/ivrparser"
	"github.com/clbanning/mxj"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
)

type IvrScriptDef struct {
	Description   string
	XMLDefinition string
	Name          string
}

func getIvrFromF9(name string) (string, error) {
	contents, err := queryF9(func() string { return generateIVRRequestContent(name) })
	if err != nil {
		fmt.Println("........", err)
		return "", err
	}
	m, _ := mxj.NewMapXml(contents, true)
	//	PrettyPrint(m)
	fmt.Println("==================================================")
	ivr, error := convertIVRResults(&m)
	if error != nil {
		return "", error
	}
	return spew.Sdump(ivr), nil
}

func generateIVRRequestContent(ivrName string) string {
	type QueryData struct {
		IvrName string
	}
	const getIvrReq = `<?xml version="1.0" encoding="utf-8"?>
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
	<soapenv:Body>
	   <ser:getIVRScripts>  
		   <namePattern>{{.IvrName}}</namePattern>
	   </ser:getIVRScripts>
	</soapenv:Body>
    </soapenv:Envelope>`

	querydata := QueryData{IvrName: ivrName}
	tmpl, err := template.New("getIVRScriptsTemplate").Parse(getIvrReq)
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

func convertIVRResults(soapResponse *mxj.Map) (*IvrScriptDef, error) {
	ivrResult, err := soapResponse.ValueForPath("Envelope.Body.getIVRScriptsResponse.return")
	if err != nil {
		return nil, err
	}
	var result IvrScriptDef
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
		// add a DecodeHook here if you need complex Decoding of results -> DecodeHook: yourfunc,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(ivrResult); err != nil {
		return nil, err
	}
	go func() {
		IVR, err := ivrparser.NewIVRScript(strings.NewReader(result.XMLDefinition))
		err = PrettyPrint(IVR)
		fmt.Println(">========== PrettyPrint =========>>>", err)
	}()

	return &result, nil
}
