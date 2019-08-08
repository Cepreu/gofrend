package preparer

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/Cepreu/gofrend/utils"

	"github.com/Cepreu/gofrend/ivr"
)

func generateIVRContent(script *ivr.IVRScript) (string, error) {
	type QueryData struct {
		DomainID       string
		CaseX          int
		CaseY          int
		SkillTransfers []struct {
			M *ivr.SkillTransferModule
			B string
			X int
			Y int
		}
		HangUps []struct {
			M *ivr.HangupModule
			B string
			X int
			Y int
		}
		DefHangupID ivr.ModuleID
		UserVars    []struct {
			VName string
			VVal  string
			VNull bool
			VAttr int
		}
	}
	querydata := QueryData{DomainID: script.Domain, CaseX: 200}

	for _, module := range script.Modules {
		if st, ok := module.(*ivr.SkillTransferModule); ok {
			bn := fmt.Sprintf("ST%d", len(querydata.SkillTransfers))
			querydata.SkillTransfers = append(querydata.SkillTransfers, struct {
				M *ivr.SkillTransferModule
				B string
				X int
				Y int
			}{M: st, B: bn})
		} else if st, ok := module.(*ivr.HangupModule); ok {
			bn := fmt.Sprintf("HU%d", len(querydata.HangUps))
			querydata.HangUps = append(querydata.HangUps, struct {
				M *ivr.HangupModule
				B string
				X int
				Y int
			}{M: st, B: bn})
		}
	}
	x := querydata.CaseX
	y := 80
	for i := range querydata.SkillTransfers {
		x += 100
		y += 20
		querydata.SkillTransfers[i].X = x
		querydata.SkillTransfers[i].Y = y
	}
	for i := range querydata.HangUps {
		x += 100
		y += 20
		querydata.HangUps[i].X = x
		querydata.HangUps[i].Y = y
	}
	querydata.CaseY = y/2 + 60

	querydata.DefHangupID = querydata.HangUps[len(querydata.HangUps)-1].M.ID
	var tmplFile = "postF9ivr.tmpl"

	for _, v := range script.Variables {
		if v.VarType == ivr.VarUserVariable {
			var vval string
			switch v.ValType {
			case ivr.ValInteger:
				vval = fmt.Sprintf("<integerValue><value>%s</value></integerValue>", v.Value)
			case ivr.ValNumeric:
				vval = fmt.Sprintf("<numericValue><value>%s</value></numericValue>", v.Value)
			case ivr.ValDate:
				vval = fmt.Sprintf("<dateValue><year>%s</year><month>%s</month><day>%s</day></dateValue>",
					v.Value[:4], v.Value[5:7], v.Value[8:])
			case ivr.ValTime:
				vval = fmt.Sprintf("<timeValue><minutes>%s</minutes></timeValue>", v.Value)
			case ivr.ValKVList:
				vval = fmt.Sprintf("<kvListValue><value>%s</value></kvListValue>",
					base64.StdEncoding.EncodeToString([]byte(v.Value)))
			case ivr.ValString:
				vval = fmt.Sprintf("<stringValue><value>%s</value><id>0</id></stringValue>", v.Value)
			case ivr.ValCurrency:
				vval = fmt.Sprintf("<currencyValue><value>%s</value></currencyValue>", v.Value[3:])
			case ivr.ValCurrencyEuro:
				vval = fmt.Sprintf("<currencyEuroValue><value>%s</value></currencyEuroValue>", v.Value[3:])
			case ivr.ValCurrencyPound:
				vval = fmt.Sprintf("<currencyPoundValue><value>%s</value></currencyPoundValue>", v.Value[3:])
			default:
				vval = fmt.Sprintf("<stringValue><value>%s</value><id>0</id></stringValue>", v.Value)
			}
			querydata.UserVars = append(querydata.UserVars, struct {
				VName string
				VVal  string
				VNull bool
				VAttr int
			}{string(v.ID), vval, v.Secured, 64}) //TBD: VNull  sould be not =Secured but = (Value==nil)
		}
	}

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

func changeUserPwdContent(userName, newPwd string) string {
	type QueryData struct {
		UserName string
		NewPwd   string
	}
	const modifyUserReq = `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
<soapenv:Body>
	<ser:modifyUser>
		<userGeneralInfo>
			<userName>{{.UserName}}</userName>
			<EMail>df@five9.com</EMail>
			<password>{{.NewPwd}}</password>
		</userGeneralInfo>
	</ser:modifyUser>
</soapenv:Body>
</soapenv:Envelope>`

	querydata := QueryData{UserName: userName, NewPwd: newPwd}
	tmpl, err := template.New("modifyUserTemplate").Parse(modifyUserReq)
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

func getIvrFromF9(auth, ivrname string) (string, error) {
	type envelope struct {
		Name          string `xml:"Body>getIVRScriptsResponse>return>name"`
		XMLDefinition string `xml:"Body>getIVRScriptsResponse>return>xmlDefinition"`
		Description   string `xml:"Body>getIVRScriptsResponse>return>description"`
	}
	content, err := queryF9(auth, func() string { return getIVRscriptContent(ivrname) })
	if err != nil {
		return "", err
	}

	var scriptDef envelope
	if err := xml.Unmarshal(content, &scriptDef); err != nil {
		return "", err
	}
	return scriptDef.XMLDefinition, nil
}

func createAuthString(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func changeUserPwdF9(oldauth string) (newauth string, err error) {
	newpwd := utils.GenUUIDv4()

	b, _ := base64.StdEncoding.DecodeString(oldauth)
	u := strings.Split(string(b), ":")

	_, err = queryF9(oldauth, func() string { return changeUserPwdContent(u[0], newpwd) })
	if err != nil {
		return "", err
	}
	newauth = base64.StdEncoding.EncodeToString([]byte(u[0] + ":" + newpwd))
	return
}

func configureF9(auth, campaign string, ivr *ivr.IVRScript) (err error) {
	generatedScriptName := ivr.Name + "_generated"
	resp, err := queryF9(auth, func() string { return createIVRscriptContent(generatedScriptName) })
	if err != nil {
		if fault, err1 := getf9errorDescr(resp); err1 != nil ||
			fault.FaultString != fmt.Sprintf("IvrScript with name \"%s\" already exists", generatedScriptName) {
			return err
		}
	}
	_, err = queryF9(auth, func() string { return modifyIVRscriptContent(generatedScriptName, ivr) })
	if err == nil {
		params := []struct {
			Name  string
			Value string
		}{}
		_, err = queryF9(auth, func() string { return setDefaultIVRScheduleContent(campaign, generatedScriptName, params) })

	}
	return err
}

func queryF9(auth string, generateRequestContent func() string) ([]byte, error) {
	//	url := conf.F9URL
	url := "https://api.five9.com/wsadmin/v11/AdminWebService"
	client := &http.Client{}
	sRequestContent := generateRequestContent()

	requestContent := []byte(sRequestContent)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestContent))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicf("Error doing request: %v", err)
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
