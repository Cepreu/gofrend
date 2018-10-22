package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/clbanning/mxj"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
)

type SkillInfo struct {
	Description     string
	ID              int64
	MessageOfTheDay string
	Name            string
	RouteVoiceMails bool
}

func getSkillsFromF9(name string) (string, error) {
	contents, err := queryF9(func() string { return generateSkillsRequestContent(name) })
	if err != nil {
		fmt.Println("===========================================  ", err)

		return "", err
	}
	m, _ := mxj.NewMapXml(contents, true)
	PrettyPrint(m)
	fmt.Println("============================================================")
	skills, error := convertSkillsResults(&m)
	if error != nil {
		return "", error
	}
	return spew.Sdump(skills), nil
}

func generateSkillsRequestContent(name string) string {
	type QueryData struct {
		SkillName string
	}
	const getSkillsReq = `<?xml version="1.0" encoding="utf-8"?>
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
	<soapenv:Body>
	   <ser:getSkills>  
		   <skillNamePattern>{{.SkillName}}</skillNamePattern>
	   </ser:getSkills>
	</soapenv:Body>
    </soapenv:Envelope>`

	//	querydata := QueryData{PostalCode: postalCode}
	querydata := QueryData{SkillName: name}
	tmpl, err := template.New("skillsTemplate").Parse(getSkillsReq)
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

func convertSkillsResults(soapResponse *mxj.Map) ([]SkillInfo, error) {
	skillsResult, err := soapResponse.ValuesForPath("Envelope.Body.getSkillsResponse.return")
	if err != nil {
		return nil, err
	}
	var result []SkillInfo
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
		// add a DecodeHook here if you need complex Decoding of results -> DecodeHook: yourfunc,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(skillsResult); err != nil {
		return nil, err
	}
	return result, nil
}
