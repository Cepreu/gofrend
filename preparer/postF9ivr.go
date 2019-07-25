package preparer

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/Cepreu/gofrend/ivr"
)

func generateIVRContent(script *ivr.IVRScript) (string, error) {
	type QueryData struct {
		DomainID string
	}

	var tmplFile = "postF9ivr.tmpl"

	querydata := QueryData{DomainID: script.Domain}
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		return "", err
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, querydata)
	if err != nil {
		return "", err
	}
	return doc.String(), nil
}

func getIVRscriptContent(scriptName string) string {
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

	querydata := QueryData{IvrName: scriptName}
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

func createIVRscriptContent(generatedScriptName string) string {
	type QueryData struct {
		IvrName string
	}
	const createIvrReq = `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
<soapenv:Body>
	<ser:createIVRScript>  
		<name>{{.IvrName}}</name>
	</ser:createIVRScript>
</soapenv:Body>
</soapenv:Envelope>`

	querydata := QueryData{IvrName: generatedScriptName}
	tmpl, err := template.New("createIVRScriptTemplate").Parse(createIvrReq)
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

func modifyIVRscriptContent(generatedScriptName string, script *ivr.IVRScript) string {
	type IvrScriptDef struct {
		Name          string
		Description   string
		XMLDefinition string
	}
	ivrBody, err := generateIVRContent(script)
	if err != nil {
		panic(err)
	}
	querydata := IvrScriptDef{
		Name:          generatedScriptName,
		Description:   "Auto-generated from " + script.Name,
		XMLDefinition: html.EscapeString(ivrBody),
	}

	const modifyIVRscriptReq = `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
<soapenv:Body>
	<ser:modifyIVRScript>  
		<scriptDef>
			<description>{{.Description}}</description>
			<name>{{.Name}}</name>
			<xmlDefinition>{{.XMLDefinition}}</xmlDefinition>
		</scriptDef>
	</ser:modifyIVRScript>
</soapenv:Body>
</soapenv:Envelope>`

	tmpl, err := template.New("modifyIVRscriptTemplate").Parse(modifyIVRscriptReq)
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

func setDefaultIVRScheduleContent(campaign, generatedScriptName string, params []struct {
	Name  string
	Value string
}) string {
	type QueryData struct {
		CampaignName        string
		ScriptName          string
		IsVisualModeEnabled bool
		IsChatEnabled       bool
		Params              []struct {
			Name  string
			Value string
		}
	}
	querydata := QueryData{
		CampaignName:        campaign,
		ScriptName:          generatedScriptName,
		IsVisualModeEnabled: true,
		IsChatEnabled:       false,
		Params:              params,
	}

	const setDefaultIVRScheduleReq = `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
<soapenv:Body>
	<ser:setDefaultIVRSchedule>  
	<campaignName>{{.CampaignName}}</campaignName>
	<scriptName>{{.ScriptName}}</scriptName>
	{{range .Params}}
	<params>
		<name>{{.Name}}</name>
		<value>{{.Value}}</value>
	</params>
	{{end}}
	<isVisualModeEnabled>{{.IsVisualModeEnabled}}</isVisualModeEnabled>
	<isChatEnabled>{{.IsChatEnabled}}</isChatEnabled>
	</ser:setDefaultIVRSchedule>
</soapenv:Body>
</soapenv:Envelope>`

	tmpl, err := template.New("setDefaultIVRScheduleTemplate").Parse(setDefaultIVRScheduleReq)
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

func getIvrFromF9(user, pwd, ivrname string) (string, error) {
	type envelope struct {
		Name          string `xml:"Body>getIVRScriptsResponse>return>name"`
		XMLDefinition string `xml:"Body>getIVRScriptsResponse>return>xmlDefinition"`
		Description   string `xml:"Body>getIVRScriptsResponse>return>description"`
	}
	content, err := queryF9(user, pwd, func() string { return getIVRscriptContent(ivrname) })
	if err != nil {
		return "", err
	}

	var scriptDef envelope
	if err := xml.Unmarshal(content, &scriptDef); err != nil {
		return "", err
	}
	return scriptDef.XMLDefinition, nil
}

func configureF9(user, pwd, campaign string, ivr *ivr.IVRScript) (err error) {
	generatedScriptName := ivr.Name + "_generated"
	resp, err := queryF9(user, pwd, func() string { return createIVRscriptContent(generatedScriptName) })
	if err != nil {
		if fault, err1 := getf9errorDescr(resp); err1 != nil ||
			fault.FaultString != fmt.Sprintf("IvrScript with name \"%s\" already exists", generatedScriptName) {
			return err
		}
	}
	_, err = queryF9(user, pwd, func() string { return modifyIVRscriptContent(generatedScriptName, ivr) })
	if err == nil {
		params := []struct {
			Name  string
			Value string
		}{}
		_, err = queryF9(user, pwd, func() string { return setDefaultIVRScheduleContent(campaign, generatedScriptName, params) })

	}

	return err
}

func queryF9(user, pwd string, generateRequestContent func() string) ([]byte, error) {
	//	url := conf.F9URL
	url := "https://api.five9.com/wsadmin/v11/AdminWebService"
	client := &http.Client{}
	sRequestContent := generateRequestContent()

	requestContent := []byte(sRequestContent)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestContent))
	if err != nil {
		return nil, err
	}

	data := []byte(user + ":" + pwd)
	str := base64.StdEncoding.EncodeToString(data)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("Authorization", "Basic "+str)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(req)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		contents, _ := ioutil.ReadAll(resp.Body)
		return contents, errors.New("Error Respose " + resp.Status)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	return contents, err
}

type f9fault struct {
	FaultCode   string `xml:"Body>Fault>faultcode"`
	FaultString string `xml:"Body>Fault>faultstring"`
}

func getf9errorDescr(faultContent []byte) (*f9fault, error) {
	var fault f9fault
	if err := xml.Unmarshal(faultContent, &fault); err != nil {
		return nil, err
	}
	return &fault, nil
}
