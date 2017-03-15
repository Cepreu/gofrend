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
		       <play>
            <ascendants>5C8BADD9D6944D1395E9B4876AEB030D</ascendants>
            <singleDescendant>702E95609FF5429E8E9E3DEED30934AB</singleDescendant>
            <moduleName>Play2</moduleName>
            <locationX>158</locationX>
            <locationY>99</locationY>
            <moduleId>D0009CF178F84C41B2495B7534771518</moduleId>
            <data>
                <prompt>
                    <ttsPrompt>
                        <xml>H4sIAAAAAAAAAMWUTU+DQBCG7/0V697r6s2YhaYmGk9eavU87U6RuB+G3ZLy7+XDQrfA0vQiBwIz
w7zPvEyWLw5Kkhwzmxod0fvbO0pQb41IdRLR9fvL/IES60ALkEZjRAu0dBHPuP1B+H6WqFC7eEbK
i4NzWbrZO7RNoA5K0MmyTHShOqxBYVxKP1Z5zupXv6Lt9gFyj09gkeTVU0RRz9cryk5EWF/lX4Q5
O/eApw7VqR+5Sbfo+dYXtn4iMM7lY10zXsjfNpegFphNoTVVQbCuUYd3RNuhAonDbCGAxu63UnQK
sLpd6dsKFGj3BcN4AYL+urSZs7Vp4xaKpR1cnz6oHeAJNW/zDg8uqHG5lq85VbQxoohfUUpDPk0m
xQ1ndWgclU2ycjbmJRs3s/fR33/0D7xjUdnJOw1/AW4p0G1VBQAA</xml>
                        <promptTTSEnumed>false</promptTTSEnumed>
                    </ttsPrompt>
                    <interruptible>false</interruptible>
                    <canChangeInterruptableOption>true</canChangeInterruptableOption>
                    <ttsEnumed>false</ttsEnumed>
                    <exitModuleOnException>false</exitModuleOnException>
                </prompt>
                <dispo>
                    <id>-17</id>
                    <name>Caller Disconnected</name>
                </dispo>
                <vivrPrompts>
                    <vivrPrompt>
                        <xml>H4sIAAAAAAAAALWSsW4CMQyGd54izQ4pG0Pu2Kp2RZTdcIZG8sVq4kPw9oUcAu4C1XVoJsefY/+R
fzs/1KT2GKJjX+jp5FUr9BuunN8V+nP5Np5pFQV8BcQeC33EqOflyAZmKUfqdCyIBLduBKO5ZJxg
Hds43b8Qzg1vmd67LujCnCXuocby0teadHtcuAdqTpVTa9ooH2V+mTVMx4aJwwAV8N3Af+oQPMgY
yO38ADGEW/m7mDvQW1t/69f8WdUTJZl1sorU9Blcc3Us35GI1TZwrVYuNkDqY7V4sSbBBx/L5VjT
N6zpOPbKrWlt/wPx/+hdNgMAAA==</xml>
                    </vivrPrompt>
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
                <numberOfDigits>0</numberOfDigits>
                <terminateDigit>N/A</terminateDigit>
                <clearDigitBuffer>false</clearDigitBuffer>
                <collapsible>false</collapsible>
            </data>
        </play>
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
	XMLName   xml.Name             `xml:"modules"`
	HModules  []HangupModule       `xml:"hangup"`
	ICModules []IncomingCallModule `xml:"incomingCall"`
	PModules  []PlayModule         `xml:"play"`
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
