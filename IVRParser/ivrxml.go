package IVRParser

import (
	"encoding/xml"
	"fmt"
)

var xmlData = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<ivrScript>
    <domainId>1</domainId>
    <properties/>
    <modules>
		<hangup>
            <ascendants>702E95609FF5429E8E9E3DEED30934AB</ascendants>
            <moduleName>Hangup2</moduleName>
            <locationX>386</locationX>
            <locationY>99</locationY>
            <moduleId>96A0B72AC33844A98D82F286B605F103</moduleId>
            <data>
                <dispo>
                    <id>0</id>
                    <name>No Disposition</name>
                </dispo>
                <returnToCallingModule>true</returnToCallingModule>
                <errCode>
                    <isVarSelected>false</isVarSelected>
                    <integerValue>
                        <value>0</value>
                    </integerValue>
                </errCode>
                <overwriteDisposition>true</overwriteDisposition>
            </data>
        </hangup>

        <incomingCall>
            <singleDescendant>D0009CF178F84C41B2495B7534771518</singleDescendant>
            <moduleName>IncomingCall1</moduleName>
            <locationX>53</locationX>
            <locationY>99</locationY>
            <moduleId>5C8BADD9D6944D1395E9B4876AEB030D</moduleId>
            <data/>
        </incomingCall>
		<hangup>
            <ascendants>702E95609FF5429E8E9E3DEED30934AB</ascendants>
            <moduleName>Hangup5</moduleName>
            <locationX>386</locationX>
            <locationY>99</locationY>
            <moduleId>96A0B72AC33844A98D82F286B605F103</moduleId>
            <data>
                <dispo>
                    <id>0</id>
                    <name>No Disposition</name>
                </dispo>
                <returnToCallingModule>true</returnToCallingModule>
                <errCode>
                    <isVarSelected>false</isVarSelected>
                    <integerValue>
                        <value>0</value>
                    </integerValue>
                </errCode>
                <overwriteDisposition>true</overwriteDisposition>
            </data>
        </hangup>
			</modules>
				</ivrScript>
`)

type S struct {
	XMLName    xml.Name `xml:"ivrScript"`
	Domain     int32    `xml:"domainId"`
	Properties string   `xml:"properties"`
	Modules    Modules  `xml:"modules"`
}
type Modules struct {
	XMLName   xml.Name       `xml:"modules"`
	HModules  []Hangup       `xml:"hangup"`
	ICModules []IncomingCall `xml:"incomingCall"`
}

/*
func toMap(vars ...Module) map[string]interface{} {
	m := make(map[string]interface{})
	for _, v := range vars {
		if len(v.Children) > 0 {
			m[v.Key] = toMap(v.Children...)
		} else {
			m[v.Key] = v.Value
		}
	}
	return m
}
*/
func Test() {
	s := &S{}
	if err := xml.Unmarshal(xmlData, s); err != nil {
		panic(err)
	}
	fmt.Printf("s: %#v\n", *s)

	//	m := toMap(s.Test.Vars...)
	//	fmt.Printf("map: %v\n", m)

	/*	js, err := json.MarshalIndent(m, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("json: %s", js)
	*/
}
