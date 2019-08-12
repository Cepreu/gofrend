package xmlparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	ivr "github.com/Cepreu/gofrend/ivr"
	"golang.org/x/net/html/charset"
)

func TestQuery(t *testing.T) {
	var xmlData = `<query>
<ascendants>E966C8048939446B8680C6DC86B603FD</ascendants>
<exceptionalDescendant>A56CB456D56A4F44B06009CD833CE4E5</exceptionalDescendant>
<singleDescendant>4E8C591A3E5A495DBFE688E80850A5DF</singleDescendant>
<moduleName>Query24</moduleName>
<locationX>301</locationX>
<locationY>132</locationY>
<moduleId>3889A8058A794604863B3AE44BB88BE1</moduleId>
<data>
	<prompt>
		<interruptible>false</interruptible>
		<canChangeInterruptableOption>true</canChangeInterruptableOption>
		<ttsEnumed>false</ttsEnumed>
		<exitModuleOnException>false</exitModuleOnException>
	</prompt>
	<vivrPrompts>
		<interruptible>false</interruptible>
		<canChangeInterruptableOption>true</canChangeInterruptableOption>
		<ttsEnumed>false</ttsEnumed>
		<exitModuleOnException>false</exitModuleOnException>
	</vivrPrompts>
	<vivrHeader>
		<interruptible>false</interruptible>
		<canChangeInterruptableOption>true</canChangeInterruptableOption>
		<ttsEnumed>false</ttsEnumed>
		<exitModuleOnException>false</exitModuleOnException>
	</vivrHeader>
	<textChannelData>
		<textPrompts>
			<interruptible>false</interruptible>
			<canChangeInterruptableOption>true</canChangeInterruptableOption>
			<ttsEnumed>false</ttsEnumed>
			<exitModuleOnException>false</exitModuleOnException>
		</textPrompts>
		<isUsedVivrPrompts>true</isUsedVivrPrompts>
		<isTextOnly>true</isTextOnly>
	</textChannelData>
	<url>https://api.five9.com/wsadmin/v11/AdminWebService</url>
	<method>POST</method>
	<fetchTimeout>5</fetchTimeout>
	<parameters/>
	<returnValues>
		<entry>
			<key>__BUFFER__</key>
			<value>__BUFFER__</value>
		</entry>
	</returnValues>
	<storeNumberOfArrayElementsInVariable>false</storeNumberOfArrayElementsInVariable>
	<requestInfo>
		<template>
			<base64>H4sIAAAAAAAAAG2RTW7DIBBG9zkF8gGYdS3KIlGlrHOBCsHnBBUzFhCa3L71n9KkXgHvzegbQGU2
A2JtP2JF4AHi1oeY2wW/N5dShpYo2wt6k+WvHZXkdKZxQ1j6qFk7kR5dSNVbSON6H+V3lp2veJOW
e2r0Tgih1vgjjEOiZ7hnd5/IBJHaM8qBYzG2nGA5uayFWPxUE5i/rsMh+YLkjf6jRmu3+ag6j+B0
NgG542Tx6Z2iGf4vriZcoRXN60sIbaUo2p5M0fatdrN7eobHef0q/QNhcqWBvQEAAA==</base64>
		</template>
		<replacements>
			<position>314</position>
			<variableName>sergei.qwerty</variableName>
		</replacements>
	</requestInfo>
	<headers>
		<entry>
			<key>Content-Type</key>
			<value>
				<isVarSelected>false</isVarSelected>
				<stringValue>
					<value>text/xml; charset=utf-8</value>
					<id>0</id>
				</stringValue>
			</value>
		</entry>
	</headers>
	<requestBodyType>XML_JSON</requestBodyType>
	<authProfileHolder>
		<extrnalObj>
			<id>31</id>
			<name>SystemAccount</name>
		</extrnalObj>
		<varSelected>false</varSelected>
	</authProfileHolder>
	<responseInfos>
		<from>200</from>
		<to>200</to>
		<method>REG_EXP</method>
		<regexp>
			<regexp>&lt;data(&gt;[^/&gt;]*|/&gt;)(&lt;/data&gt;)?&lt;data(&gt;[^/&gt;]*|/&gt;)(&lt;/data&gt;)?&lt;data(&gt;[^/&gt;]*|/&gt;)(&lt;/data&gt;)?&lt;/values&gt;</regexp>
			<regexpParameters>
				<groupIndex>0</groupIndex>
				<variable>__BUFFER__</variable>
			</regexpParameters>
			<regexpParameters>
				<groupIndex>1</groupIndex>
			</regexpParameters>
			<regexpParameters>
				<groupIndex>2</groupIndex>
			</regexpParameters>
			<regexpParameters>
				<groupIndex>3</groupIndex>
			</regexpParameters>
			<regexpParameters>
				<groupIndex>4</groupIndex>
			</regexpParameters>
			<regexpParameters>
				<groupIndex>5</groupIndex>
			</regexpParameters>
			<regexpFlags>0</regexpFlags>
		</regexp>
	</responseInfos>
	<responseInfos>
		<from>200</from>
		<to>500</to>
		<method>REG_EXP</method>
		<regexp>
			<regexp>(.+)</regexp>
			<regexpParameters>
				<groupIndex>0</groupIndex>
				<variable>__BUFFER__</variable>
			</regexpParameters>
			<regexpFlags>0</regexpFlags>
		</regexp>
	</responseInfos>
	<saveStatusCode>false</saveStatusCode>
	<saveReasonPhrase>false</saveReasonPhrase>
</data>
</query>`

	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel
	_, _ = decoder.Token()

	s := &ivr.IVRScript{
		Variables: make(ivr.Variables),
		Modules:   make(map[ivr.ModuleID]ivr.Module),
	}
	res := newQueryModule(decoder, s)
	if res == nil {
		t.Fatal("nQuery module wasn't parsed...")
	}
	if err := res.normalize(s); err != nil {
		t.Fatal("nQuery module wasn't normalized...")
	}

	var mq = (res.(xmlQueryModule)).m
	contentType, _ := addStringConstant(s, "text/xml; charset=utf-8")

	var expected = ivr.QueryModule{
		VoicePromptIDs: ivr.ModulePrompts{{
			LangPrArr: []ivr.LanguagePrompts{{PrArr: []ivr.PromptID{}, Language: "Default"}},
			Count:     1,
		}},
		URL:          "https://api.five9.com/wsadmin/v11/AdminWebService",
		Method:       "POST",
		ReturnValues: []ivr.KeyValue{{Key: "__BUFFER__", Value: "__BUFFER__"}},
		RequestInfo: ivr.RequestInfo{
			Template:     "\u003csoapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:ser=\"http://service.admin.ws.five9.com/\"\u003e\n   \u003csoapenv:Header/\u003e\n   \u003csoapenv:Body\u003e\n      \u003cser:getContactRecords\u003e  \n          \u003clookupCriteria\u003e\n             \u003ccriteria\u003e\n                \u003cfield\u003esalesforce_id\u003c/field\u003e\n                \u003cvalue\u003e\u003c/value\u003e\n             \u003c/criteria\u003e\n          \u003c/lookupCriteria\u003e\n      \u003c/ser:getContactRecords\u003e\n   \u003c/soapenv:Body\u003e\n\u003c/soapenv:Envelope\u003e",
			Replacements: []*ivr.Replacement{&ivr.Replacement{Position: 314, VariableName: "sergei.qwerty"}},
		},
		Headers:         []ivr.KeyValue{{Key: "Content-Type", Value: contentType}},
		RequestBodyType: "XML_JSON",
		ResponseInfos: []*ivr.ResponseInfo{
			&ivr.ResponseInfo{
				HTTPCodeFrom:  200,
				HTTPCodeTo:    200,
				ParsingMethod: "REG_EXP",
				Regexp: struct {
					RegexpBody  string
					RegexpFlags int
				}{
					RegexpBody:  "\u003cdata(\u003e[^/\u003e]*|/\u003e)(\u003c/data\u003e)?\u003cdata(\u003e[^/\u003e]*|/\u003e)(\u003c/data\u003e)?\u003cdata(\u003e[^/\u003e]*|/\u003e)(\u003c/data\u003e)?\u003c/values\u003e",
					RegexpFlags: 0,
				},
				Function: struct {
					Name       string
					ReturnType string
					Arguments  string
				}{},
				TargetVariables: []string{"__BUFFER__", "", "", "", "", ""},
			},
			&ivr.ResponseInfo{
				HTTPCodeFrom:  200,
				HTTPCodeTo:    500,
				ParsingMethod: "REG_EXP",
				Regexp: struct {
					RegexpBody  string
					RegexpFlags int
				}{"(.+)", 0},
				Function: struct {
					Name       string
					ReturnType string
					Arguments  string
				}{},
				TargetVariables: []string{"__BUFFER__"},
			},
		},
	}
	expected.SetGeneralInfo("Query24", "3889A8058A794604863B3AE44BB88BE1",
		[]ivr.ModuleID{"E966C8048939446B8680C6DC86B603FD"}, "4E8C591A3E5A495DBFE688E80850A5DF",
		"A56CB456D56A4F44B06009CD833CE4E5", "", "false")

	eqm, err2 := json.MarshalIndent(mq, "", "  ")
	exp, err1 := json.MarshalIndent(expected, "", "  ")
	if err1 != nil || err2 != nil || string(exp) != string(eqm) {
		t.Errorf("\nQuery module: \n%s \nwas expected, in reality: \n%s", string(exp), string(eqm))
	}
}
