package IVRParser

import (
	"strings"
	"testing"
)

func TestParseIVR(t *testing.T) {

	_, err := ParseIVR(strings.NewReader(xmlData))

	if err != nil {
		panic(err)
	}

	// more sanity checking...

}

var xmlData = `
<?xml version="1.0" encoding="ISO-8859-1" standalone="yes"?>
<ivrScript>
    <domainId>1</domainId>
    <properties/>
    <modules>
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
        <input>
            <ascendants>D0009CF178F84C41B2495B7534771518</ascendants>
            <singleDescendant>96A0B72AC33844A98D82F286B605F103</singleDescendant>
            <moduleName>Input7</moduleName>
            <locationX>268</locationX>
            <locationY>99</locationY>
            <moduleId>702E95609FF5429E8E9E3DEED30934AB</moduleId>
            <data>
                <dispo>
                    <id>-17</id>
                    <name>Caller Disconnected</name>
                </dispo>
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
                <grammar xsi:type="boolean" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
                    <mainReturnValue>
                        <name>MEANING</name>
                        <type>STRING</type>
                        <varName>__BUFFER__</varName>
                    </mainReturnValue>
                    <stringProperty>
                        <type>LANGUAGE</type>
                        <value xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">en-US</value>
                        <enabled>true</enabled>
                    </stringProperty>
                    <listProperty>
                        <type>Y</type>
                        <value xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">1</value>
                        <enabled>true</enabled>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">0</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">1</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">2</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">3</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">4</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">5</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">6</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">7</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">8</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">9</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">*</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">#</list>
                    </listProperty>
                    <listProperty>
                        <type>N</type>
                        <value xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">2</value>
                        <enabled>true</enabled>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">0</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">1</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">2</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">3</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">4</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">5</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">6</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">7</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">8</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">9</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">*</list>
                        <list xsi:type="xs:string" xmlns:xs="http://www.w3.org/2001/XMLSchema">#</list>
                    </listProperty>
                    <additionalReturnValues>
                        <name>SWI_literal</name>
                        <type>STRING</type>
                    </additionalReturnValues>
                    <propertiesToVariables/>
                </grammar>
                <prompts>
                    <prompt>
                        <filePrompt>
                            <promptData>
                                <promptSelected>true</promptSelected>
                                <prompt>
                                    <id>-9</id>
                                    <name>EIGHT</name>
                                </prompt>
                                <isRecordedMessage>false</isRecordedMessage>
                            </promptData>
                        </filePrompt>
                        <interruptible>false</interruptible>
                        <canChangeInterruptableOption>true</canChangeInterruptableOption>
                        <ttsEnumed>false</ttsEnumed>
                        <exitModuleOnException>false</exitModuleOnException>
                    </prompt>
                    <count>1</count>
                </prompts>
                <recoEvents>
                    <event>NO_MATCH</event>
                    <count>1</count>
                    <compoundPrompt>
                        <multiLanguagesPromptItem>
                            <prompt>9BED287531A2493591A3DE63BF138F0C</prompt>
                        </multiLanguagesPromptItem>
                        <interruptible>false</interruptible>
                        <canChangeInterruptableOption>true</canChangeInterruptableOption>
                        <ttsEnumed>false</ttsEnumed>
                        <exitModuleOnException>false</exitModuleOnException>
                    </compoundPrompt>
                    <action>CONTINUE</action>
                </recoEvents>
                <recoEvents>
                    <event>NO_INPUT</event>
                    <count>1</count>
                    <compoundPrompt>
                        <multiLanguagesPromptItem>
                            <prompt>114228A09322479E93E8211044F2B775</prompt>
                        </multiLanguagesPromptItem>
                        <interruptible>false</interruptible>
                        <canChangeInterruptableOption>true</canChangeInterruptableOption>
                        <ttsEnumed>false</ttsEnumed>
                        <exitModuleOnException>false</exitModuleOnException>
                    </compoundPrompt>
                    <action>REPROMPT</action>
                </recoEvents>
                <confirmData>
                    <confirmRequired>REQUIRED</confirmRequired>
                    <requiredConfidence>75</requiredConfidence>
                    <maxAttemptsToConfirm>3</maxAttemptsToConfirm>
                    <noInputTimeout>3</noInputTimeout>
                    <prompt>
                        <multiLanguagesPromptItem>
                            <prompt>755F78B037834CD09B29C42D1A63A1B6</prompt>
                        </multiLanguagesPromptItem>
                        <interruptible>true</interruptible>
                        <canChangeInterruptableOption>true</canChangeInterruptableOption>
                        <ttsEnumed>false</ttsEnumed>
                        <exitModuleOnException>false</exitModuleOnException>
                    </prompt>
                    <recoEvents>
                        <event>NO_MATCH</event>
                        <count>1</count>
                        <compoundPrompt>
                            <multiLanguagesPromptItem>
                                <prompt>9BED287531A2493591A3DE63BF138F0C</prompt>
                            </multiLanguagesPromptItem>
                            <interruptible>false</interruptible>
                            <canChangeInterruptableOption>true</canChangeInterruptableOption>
                            <ttsEnumed>false</ttsEnumed>
                            <exitModuleOnException>false</exitModuleOnException>
                        </compoundPrompt>
                        <action>REPROMPT</action>
                    </recoEvents>
                    <recoEvents>
                        <event>NO_INPUT</event>
                        <count>1</count>
                        <compoundPrompt>
                            <multiLanguagesPromptItem>
                                <prompt>114228A09322479E93E8211044F2B775</prompt>
                            </multiLanguagesPromptItem>
                            <interruptible>false</interruptible>
                            <canChangeInterruptableOption>true</canChangeInterruptableOption>
                            <ttsEnumed>false</ttsEnumed>
                            <exitModuleOnException>false</exitModuleOnException>
                        </compoundPrompt>
                        <action>REPROMPT</action>
                    </recoEvents>
                </confirmData>
                <confidenceTreshold>5</confidenceTreshold>
                <maxTimeToEnter>15</maxTimeToEnter>
                <noInputTimeout>5</noInputTimeout>
                <sensitivity>50</sensitivity>
                <incompleteTimeout>2</incompleteTimeout>
                <completeTimeout>0</completeTimeout>
                <swirecNbestListLength>2</swirecNbestListLength>
                <maxAttempts>3</maxAttempts>
                <collapsible>false</collapsible>
                <recognizeConfigParameters>
                    <recognizeConfigParameter>NO_INPUT_TIMEOUT</recognizeConfigParameter>
                    <recognizeConfigParameter>CONFIDENCE_TRESHOLD</recognizeConfigParameter>
                    <recognizeConfigParameter>MAX_TIME_TO_ENTER</recognizeConfigParameter>
                </recognizeConfigParameters>
                <recordUserInput>false</recordUserInput>
                <dtmfHelpButton>null</dtmfHelpButton>
            </data>
        </input>
    </modules>
    <modulesOnHangup>
        <startOnHangup>
            <singleDescendant>B955689FF7E54E4A8CD84E30C38DD82F</singleDescendant>
            <moduleName>StartOnHangup1</moduleName>
            <locationX>34</locationX>
            <locationY>62</locationY>
            <moduleId>633140A87B554504A6D5DC276C3CF50F</moduleId>
        </startOnHangup>
        <hangup>
            <ascendants>633140A87B554504A6D5DC276C3CF50F</ascendants>
            <moduleName>Hangup1</moduleName>
            <locationX>134</locationX>
            <locationY>62</locationY>
            <moduleId>B955689FF7E54E4A8CD84E30C38DD82F</moduleId>
            <data>
                <dispo>
                    <id>-17</id>
                    <name>Caller Disconnected</name>
                </dispo>
                <returnToCallingModule>true</returnToCallingModule>
                <errCode>
                    <isVarSelected>false</isVarSelected>
                    <integerValue>
                        <value>0</value>
                    </integerValue>
                </errCode>
                <overwriteDisposition>false</overwriteDisposition>
            </data>
        </hangup>
    </modulesOnHangup>
    <userVariables/>
    <multiLanguagesPrompts>
        <entry>
            <key>114228A09322479E93E8211044F2B775</key>
            <value>
                <promptId>114228A09322479E93E8211044F2B775</promptId>
                <name>NoInputPrompt</name>
                <description>Default prompt for NoInput event</description>
                <type>AUDIO</type>
                <prompts>
                    <entry key="en-US">
                        <ttsPrompt>
                            <xml>H4sIAAAAAAAAAIWRQYvCMBCF7/6KkHsdvS1LWnFhBc+7eh/toMV0Ip2p2H9vrVBNu7I5Je8b3ntM
3OJaenOhSorAqZ1PZ9YQ70Ne8CG1m99V8mGNKHKOPjCltiGxi2zi5Ex4+vZUEms2Me1xqFoVu1pJ
HkIneuTDsgVPqZMZS8ra6M87d9A944nebYu+pi8UMpf7LbXEyebHwksIxCkOhlVcoVS+1hJslhLV
HwcLDMjApNeVrvqn1/+esfc7uAt5k63NHpmDmiNhZZpQTx10YFwI3jZyMNwEjFfRD7Uw+uYbQw31
uC4CAAA=</xml>
                            <promptTTSEnumed>false</promptTTSEnumed>
                        </ttsPrompt>
                    </entry>
                </prompts>
                <defaultLanguage>en-US</defaultLanguage>
                <isPersistent>true</isPersistent>
            </value>
        </entry>
        <entry>
            <key>755F78B037834CD09B29C42D1A63A1B6</key>
            <value>
                <promptId>755F78B037834CD09B29C42D1A63A1B6</promptId>
                <name>ConfirmPrompt</name>
                <description>Default prompt for user input confirmation</description>
                <type>AUDIO</type>
                <prompts>
                    <entry key="en-US">
                        <ttsPrompt>
                            <xml>H4sIAAAAAAAAANVTUWvCMBB+91ccebeZexqSVhw4EMTB1A2fSrQ3DWsTyaXO/vulHXNttQwGe1jI
Q/J9ue++3CVidMpSOKIlZXTIBsENA9Rbkyi9C9lq+dC/Y0BO6kSmRmPICiQ2inqCDijfJilmqF3U
Az+EdM6qTe6QPoEKTKXejT3xDVWwlhlGPvWw5AWvts0TZ7VnmeZ4LwnhWK5Chrq/WjBeS8KbWQRv
WxHKYVa3RbIYU8P+ZWLiLaYlcsYdntxVrZ81m9pd5MYkRbQ2uW+NQ4vJ0F+xwi698E4zgreLwK9X
QRylVXKT4q/q0wa/xOZli+N48TKNZ9Pl5Gk8i2PBG2zNWaeF/9i5ABayALb2XweMhYNFImADBuoV
3F4R+Lk11uLWBfDo9mjflX/uVAbNTT3mlv1V48+HPNn42h9LoLssIgQAAA==</xml>
                            <promptTTSEnumed>false</promptTTSEnumed>
                        </ttsPrompt>
                    </entry>
                </prompts>
                <defaultLanguage>en-US</defaultLanguage>
                <isPersistent>true</isPersistent>
            </value>
        </entry>
        <entry>
            <key>9BED287531A2493591A3DE63BF138F0C</key>
            <value>
                <promptId>9BED287531A2493591A3DE63BF138F0C</promptId>
                <name>NoMatchPrompt</name>
                <description>Default prompt for NoMatch event</description>
                <type>AUDIO</type>
                <prompts>
                    <entry key="en-US">
                        <ttsPrompt>
                            <xml>H4sIAAAAAAAAAIWRQYvCMBCF7/6KkLvO7k0krSis4Fnd+2gHCZtOpTMV++83VqimXdmcku+F9x4z
bnkrg7lSLb7izH7OPqwhPlWF53NmD/vNdG6NKHKBoWLKbEtil/nEyYXw5ytQSaz5xMTjULX2x0ZJ
HqCDAfm8isITdZixpDxGL+66g+6Z/ujdvjE0tEYhc73fMks8PewsvIRAmuJgWMV5pfK1lmC7kqT+
OFhgoAxMeq500z+9/vdMvd+Jx6po821cS13TSY3nS6MzBx0e14G3fRwM5wDjQfSfopgs+RfRtLpi
LAIAAA==</xml>
                            <promptTTSEnumed>false</promptTTSEnumed>
                        </ttsPrompt>
                    </entry>
                </prompts>
                <defaultLanguage>en-US</defaultLanguage>
                <isPersistent>true</isPersistent>
            </value>
        </entry>
        <entry>
            <key>FDA5FB8B16E24EE1BAEE0D3A4360161D</key>
            <value>
                <promptId>FDA5FB8B16E24EE1BAEE0D3A4360161D</promptId>
                <name>ConfirmPromptWithoutVSR</name>
                <description>Default prompt for user input confirmation with disabled voice recognition</description>
                <type>AUDIO</type>
                <prompts>
                    <entry key="en-US">
                        <ttsPrompt>
                            <xml>H4sIAAAAAAAAANVT0UrDMBR931eEvK9xPslIOypMGAwVtyk+lbS9bsE0GUla17837bA23YYg+GDJ
Q3PO5ZyTexM6OxQCVaANVzLEk+AKI5CZyrnchnizvhvfYGQskzkTSkKIazB4Fo2o2QN7nwsoQNpo
hNxHmbWap6UFcwRaUDC5jR3xDbWwZAVEznra8JS0W7+iU3tmooRbZgBVzV+IQY43K0x6JsR3oWQY
hXILRT+WYXVsvPinxoYMmIFIh1s42LNaP2v62pfIVOV19KpKNxoLGvKpO2KLnWYhF8NQMmwCOd8F
WjHNWSrgV/0Zgl9i982Ik2T1skiWi/X8KV4mCSUe20t2McJ/nFyAHjUYg/AEI/6G7I4b5FamtIbM
BujB7kB/cHfD98e6a/xX8+2KHOm94E/oteYzCQQAAA==</xml>
                            <promptTTSEnumed>false</promptTTSEnumed>
                        </ttsPrompt>
                    </entry>
                </prompts>
                <defaultLanguage>en-US</defaultLanguage>
                <isPersistent>true</isPersistent>
            </value>
        </entry>
    </multiLanguagesPrompts>
    <multiLanguagesVIVRPrompts/>
    <multiLanguagesMenuChoices/>
    <languages/>
    <defaultLanguage>en-US</defaultLanguage>
    <defaultMethod>GET</defaultMethod>
    <defaultFetchTimeout>5</defaultFetchTimeout>
    <showLabelNames>true</showLabelNames>
    <defaultVivrTimeout>5</defaultVivrTimeout>
    <unicodeEncoding>false</unicodeEncoding>
    <timeoutInMilliseconds>3600000</timeoutInMilliseconds>
    <version>950014</version>
</ivrScript>
`
