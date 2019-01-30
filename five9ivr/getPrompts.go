package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Cepreu/gofrend/utils"
	"github.com/clbanning/mxj"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
)

type promptInfo struct {
	Description string
	Languages   []string
	Name        string
	Type        string
}

func getPromptFromF9(name string) (string, error) {
	var contents []byte
	var err error
	var path string

	if name != "" {
		contents, err = queryF9(func() string { return generatePromptRequestContent(name) })
		path = "Envelope.Body.getPromptResponse.return"
	} else {
		contents, err = queryF9(func() string { return generatePromptsRequestContent() })
		path = "Envelope.Body.getPromptsResponse.prompts"
	}
	if err != nil {
		return "", err
	}
	m, _ := mxj.NewMapXml(contents, true)
	utils.PrettyPrint(m)
	prompt, error := convertPromptResults(&m, path)
	if error != nil {
		return "", error
	}
	return spew.Sdump(prompt), nil
}

func generatePromptRequestContent(name string) string {
	type QueryData struct {
		PromptName string
	}
	const getPromptReq = `<?xml version="1.0" encoding="utf-8"?>
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
	<soapenv:Body>
	   <ser:getPrompt>  
		   <promptName>{{.PromptName}}</promptName>
	   </ser:getPrompt>
	</soapenv:Body>
    </soapenv:Envelope>`

	querydata := QueryData{PromptName: name}
	tmpl, err := template.New("promptTemplate").Parse(getPromptReq)
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

func generatePromptsRequestContent() string {
	const getPromptsReq = `<?xml version="1.0" encoding="utf-8"?>
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
	<soapenv:Body>
	   <ser:getPrompts/>  
	</soapenv:Body>
    </soapenv:Envelope>`

	tmpl, err := template.New("promptsTemplate").Parse(getPromptsReq)
	if err != nil {
		panic(err)
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(doc.String())
	return doc.String()
}

func convertPromptResults(soapResponse *mxj.Map, path string) ([]promptInfo, error) {
	promptResult, err := soapResponse.ValuesForPath(path)
	if err != nil {
		return nil, err
	}
	var result []promptInfo
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
		// add a DecodeHook here if you need complex Decoding of results -> DecodeHook: yourfunc,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(promptResult); err != nil {
		return nil, err
	}
	return result, nil
}
