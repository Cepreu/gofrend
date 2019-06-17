package xmlparser

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
	"golang.org/x/net/html/charset"
)

func TestMenu(t *testing.T) {
	var xmlData = `<menu>
	<ascendants>35EBEC8BD6294BA2B6CF4C96D54BE72D</ascendants>
	<exceptionalDescendant>C4BA62EF0F3F406A819ADE2CABD1669C</exceptionalDescendant>
	<moduleName>Copy of Menu5</moduleName>
	<locationX>294</locationX>
	<locationY>91</locationY>
	<moduleId>0B19BD7CEE3E4B85BA6C6631F1CCA222</moduleId>
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
		<branches>
			<entry>
				<key>No Match</key>
				<value>
					<name>No Match</name>
					<desc>C4BA62EF0F3F406A819ADE2CABD1669C</desc>
				</value>
			</entry>
			<entry>
				<key>apples</key>
				<value>
					<name>apples</name>
					<desc>C4BA62EF0F3F406A819ADE2CABD1669C</desc>
				</value>
			</entry>
			<entry>
				<key>plums</key>
				<value>
					<name>plums</name>
					<desc>C4BA62EF0F3F406A819ADE2CABD1669C</desc>
				</value>
			</entry>
			<entry>
				<key>appricots</key>
				<value>
					<name>appricots</name>
					<desc>C4BA62EF0F3F406A819ADE2CABD1669C</desc>
				</value>
			</entry>
			<entry>
				<key>tutti frutti</key>
				<value>
					<name>tutti frutti</name>
					<desc>C4BA62EF0F3F406A819ADE2CABD1669C</desc>
				</value>
			</entry>
		</branches>
		<useSpeechRecognition>true</useSpeechRecognition>
		<useDTMF>true</useDTMF>
		<recordUserInput>false</recordUserInput>
		<maxAttempts>3</maxAttempts>
		<confidenceTreshold>60</confidenceTreshold>
		<saveInput>
			<name>__BUFFER__</name>
		</saveInput>
		<recoEvents>
			<event>NO_MATCH</event>
			<count>1</count>
			<compoundPrompt>
				<TtsPrompt>
					<xml>H4sIAAAAAAAAAIWRQYvCMBCF7/6KkLs7602WtKKg4Fnd+2gHCZtOpTMV++83dqFrWsWcku+F9x4z
bnErg7lSLb7izM4+Pq0hPlWF53NmD/vNdG6NKHKBoWLKbEtiF/nEyYXwZx2oJNZ8YuJxqFr7Y6Mk
f6CDAfm8jMI/6jBjSXmM/rrrDrpn+qN3+8bQ0AqFzPV+yyzx9LCz8BACaYqDYRXnlcrHWoLtUpL6
42CBgTIw6bnSTZ96vfdMvV+Jx6po821cS13TSY3nS6MOOjpuAy/rOBiOAcZz6D9FMdnxL2oN76or
AgAA</xml>
					<promptTTSEnumed>true</promptTTSEnumed>
				</TtsPrompt>
				<interruptible>false</interruptible>
				<canChangeInterruptableOption>true</canChangeInterruptableOption>
				<ttsEnumed>true</ttsEnumed>
				<exitModuleOnException>false</exitModuleOnException>
			</compoundPrompt>
			<action>CONTINUE</action>
		</recoEvents>
		<recoEvents>
			<event>NO_INPUT</event>
			<count>1</count>
			<compoundPrompt>
				<TtsPrompt>
					<xml>H4sIAAAAAAAAAIWRQYvCMBCF7/6KkLs7602WtKKg4Fnd+2gHCZtOpTMV++83dqFrWsWcku+F9x4z
bnErg7lSLb7izM4+Pq0hPlWF53NmD/vNdG6NKHKBoWLKbEtiF/nEyYXwZx2oJNZ8YuJxqFr7Y6Mk
f6CDAfm8jMI/6jBjSXmM/rrrDrpn+qN3+8bQ0AqFzPV+yyzx9LCz8BACaYqDYRXnlcrHWoLtUpL6
42CBgTIw6bnSTZ96vfdMvV+Jx6po821cS13TSY3nS6MOOjpuAy/rOBiOAcZz6D9FMdnxL2oN76or
AgAA</xml>
					<promptTTSEnumed>true</promptTTSEnumed>
				</TtsPrompt>
				<interruptible>false</interruptible>
				<canChangeInterruptableOption>true</canChangeInterruptableOption>
				<ttsEnumed>true</ttsEnumed>
				<exitModuleOnException>false</exitModuleOnException>
			</compoundPrompt>
			<action>REPROMPT</action>
		</recoEvents>
		<recoEvents>
			<event>NO_INPUT</event>
			<count>2</count>
			<compoundPrompt>
				<multiLanguagesPromptItem>
					<prompt>DAA16853152C4EBFA4DEEDAB00227C2A</prompt>
				</multiLanguagesPromptItem>
				<interruptible>false</interruptible>
				<canChangeInterruptableOption>true</canChangeInterruptableOption>
				<ttsEnumed>false</ttsEnumed>
				<exitModuleOnException>false</exitModuleOnException>
			</compoundPrompt>
			<action>REPROMPT</action>
		</recoEvents>
		<prompts>
			<prompt>
				<TtsPrompt>
					<xml>H4sIAAAAAAAAAIVRQW7CMBC88wrLd3B7qyoniErtuVIL96VZkIWzrrybiPweJ0gBJ0X1aT2znhmN
7fpce9ViZBeo0M+rJ62QfkLl6Fjo7ffH8kUrFqAKfCAsdIes1+XC8i/C6d1jjSTlQqVjQSS6fSPI
V2AAPdBxk4gbNMAENZbJ+rXnrRmu+caotgPf4BswqrafCo203H5pc2dichdrplGsE6zvYzF0G87i
z43ZTJiJyIgLnuVPrf81c+1H5D5UXfnpsS+hr0p1oYnqAG2I6Z06xMbJypphbR7PPMxnzbQXMy9m
XEpk9ukXyjrBtTwCAAA=</xml>
					<promptTTSEnumed>true</promptTTSEnumed>
				</TtsPrompt>
				<interruptible>true</interruptible>
				<canChangeInterruptableOption>true</canChangeInterruptableOption>
				<ttsEnumed>true</ttsEnumed>
				<exitModuleOnException>false</exitModuleOnException>
			</prompt>
			<count>1</count>
		</prompts>
		<prompts>
			<prompt>
				<multiLanguagesPromptItem>
					<prompt>FCD392D339FB4A20B95105FAFF8AC7ED</prompt>
				</multiLanguagesPromptItem>
				<interruptible>false</interruptible>
				<canChangeInterruptableOption>true</canChangeInterruptableOption>
				<ttsEnumed>true</ttsEnumed>
				<exitModuleOnException>false</exitModuleOnException>
			</prompt>
			<count>2</count>
		</prompts>
		<ConfirmData>
			<confirmRequired>REQUIRED</confirmRequired>
			<requiredConfidence>75</requiredConfidence>
			<maxAttemptsToConfirm>3</maxAttemptsToConfirm>
			<noInputTimeout>3</noInputTimeout>
			<prompt>
				<multiLanguagesPromptItem>
					<prompt>87F5501915CA4ACC874D8767B8C4369E</prompt>
				</multiLanguagesPromptItem>
				<interruptible>true</interruptible>
				<canChangeInterruptableOption>true</canChangeInterruptableOption>
				<ttsEnumed>true</ttsEnumed>
				<exitModuleOnException>false</exitModuleOnException>
			</prompt>
			<recoEvents>
				<event>NO_MATCH</event>
				<count>1</count>
				<compoundPrompt>
					<TtsPrompt>
						<xml>H4sIAAAAAAAAAIWRQYvCMBCF7/6KkLs7602WtKKg4Fnd+2gHCZtOpTMV++83dqFrWsWcku+F9x4z
bnErg7lSLb7izM4+Pq0hPlWF53NmD/vNdG6NKHKBoWLKbEtiF/nEyYXwZx2oJNZ8YuJxqFr7Y6Mk
f6CDAfm8jMI/6jBjSXmM/rrrDrpn+qN3+8bQ0AqFzPV+yyzx9LCz8BACaYqDYRXnlcrHWoLtUpL6
42CBgTIw6bnSTZ96vfdMvV+Jx6po821cS13TSY3nS6MOOjpuAy/rOBiOAcZz6D9FMdnxL2oN76or
AgAA</xml>
						<promptTTSEnumed>true</promptTTSEnumed>
					</TtsPrompt>
					<interruptible>false</interruptible>
					<canChangeInterruptableOption>true</canChangeInterruptableOption>
					<ttsEnumed>true</ttsEnumed>
					<exitModuleOnException>false</exitModuleOnException>
				</compoundPrompt>
				<action>REPROMPT</action>
			</recoEvents>
			<recoEvents>
				<event>NO_INPUT</event>
				<count>1</count>
				<compoundPrompt>
					<TtsPrompt>
						<xml>H4sIAAAAAAAAAIWRT4vCMBDF736KkHsdvcmSVlxQ8Oyf+2gHLZtOpDMV++23dqFr2pXNKfm94b3H
xC0fpTd3qqQInNr5dGYN8TnkBV9Se9hvkoU1osg5+sCU2obELrOJkxvh19pTSazZxLTHoWpVnGol
+QEd9MiXVSv8og4zlpS10R9P3UH3jCd6tyP6mj5RyNyft9QSJ4edhZcQiFMcDKu4Qql8rSXYrCSq
Pw4WGCgDk54rPfRPr/89Y+934inkTbY1Z2QOaq6ElWlC7aDj4z7wtpCD4SJgvIl+qBWjX/4GAU/u
hS0CAAA=</xml>
						<promptTTSEnumed>true</promptTTSEnumed>
					</TtsPrompt>
					<interruptible>false</interruptible>
					<canChangeInterruptableOption>true</canChangeInterruptableOption>
					<ttsEnumed>true</ttsEnumed>
					<exitModuleOnException>false</exitModuleOnException>
				</compoundPrompt>
				<action>REPROMPT</action>
			</recoEvents>
			<recoEvents>
				<event>HELP</event>
				<count>1</count>
				<compoundPrompt>
					<interruptible>false</interruptible>
					<canChangeInterruptableOption>true</canChangeInterruptableOption>
					<ttsEnumed>true</ttsEnumed>
					<exitModuleOnException>false</exitModuleOnException>
				</compoundPrompt>
				<action>REPROMPT</action>
			</recoEvents>
		</ConfirmData>
		<items>
			<choice>
				<type>VALUE</type>
				<value>apple</value>
				<showInVivr>true</showInVivr>
			</choice>
			<match>APPR</match>
			<thumbnail>
				<type>VALUE</type>
				<value></value>
				<showInVivr>true</showInVivr>
			</thumbnail>
			<dtmf>DTMF_AUTO</dtmf>
			<ActionType>BRANCH</ActionType>
			<actionName>apples</actionName>
		</items>
		<items>
			<choice>
				<type>VARIABLE</type>
				<value></value>
				<varName>__BUFFER__</varName>
				<showInVivr>true</showInVivr>
			</choice>
			<match>APPR</match>
			<thumbnail>
				<type>VALUE</type>
				<value></value>
				<showInVivr>true</showInVivr>
			</thumbnail>
			<dtmf>DTMF_AUTO</dtmf>
			<ActionType>BRANCH</ActionType>
			<actionName>appricots</actionName>
		</items>
		<items>
			<choice>
				<type>MODULE</type>
				<value></value>
				<showInVivr>true</showInVivr>
				<module>35EBEC8BD6294BA2B6CF4C96D54BE72D</module>
				<moduleField>__BUFFER__</moduleField>
			</choice>
			<match>APPR</match>
			<thumbnail>
				<type>VALUE</type>
				<value></value>
				<showInVivr>true</showInVivr>
			</thumbnail>
			<dtmf>DTMF_AUTO</dtmf>
			<ActionType>BRANCH</ActionType>
			<actionName>plums</actionName>
		</items>
		<items>
			<choice>
				<type>ML_ITEM</type>
				<value></value>
				<showInVivr>true</showInVivr>
				<mlItem>FFBEE18B77834E10B034889F5457DEB4</mlItem>
			</choice>
			<match>APPR</match>
			<thumbnail>
				<type>VALUE</type>
				<value></value>
				<showInVivr>true</showInVivr>
			</thumbnail>
			<dtmf>DTMF_AUTO</dtmf>
			<ActionType>BRANCH</ActionType>
			<actionName>tutti frutti</actionName>
		</items>
		<items>
			<choice>
				<type>ML_ITEM</type>
				<value></value>
				<showInVivr>true</showInVivr>
				<mlItem>7D0A23161B1B47CEA640C61D490D6FEF</mlItem>
			</choice>
			<match>APPR</match>
			<thumbnail>
				<type>VALUE</type>
				<value></value>
				<showInVivr>true</showInVivr>
			</thumbnail>
			<dtmf>DTMF_AUTO</dtmf>
			<ActionType>BRANCH</ActionType>
			<actionName>appricots</actionName>
		</items>
		<items>
			<choice>
				<type>ML_ITEM</type>
				<value></value>
				<showInVivr>true</showInVivr>
				<mlItem>3EF9450C730E462FB97E9ADA7A0E37AE</mlItem>
			</choice>
			<match>APPR</match>
			<thumbnail>
				<type>VALUE</type>
				<value></value>
				<showInVivr>true</showInVivr>
			</thumbnail>
			<dtmf>DTMF_AUTO</dtmf>
			<ActionType>BRANCH</ActionType>
			<actionName>plums</actionName>
		</items>
		<maxTimeToEnter>5</maxTimeToEnter>
		<noInputTimeout>5</noInputTimeout>
		<speechCompleteTimeout>1</speechCompleteTimeout>
		<collapsible>false</collapsible>
	</data>
</menu>
`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	prompts := make(ivr.ScriptPrompts)
	res := newMenuModule(decoder, prompts)
	if res == nil {
		t.Fatalf("Menu module wasn't parsed...")
	}

	var mmm = (res.(xmlMenuModule)).m

	expected := &ivr.MenuModule{
		VoicePromptIDs: ivr.ModulePrompts{
			ivr.AttemptPrompts{
				LangPrArr: []ivr.LanguagePrompts{
					{
						PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_V__T_21"},
						Language: "Default",
					}, {
						PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_V__T_21"},
						Language: "en-US",
					}, {
						PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_V__T_21"},
						Language: "uk",
					},
				},
				Count: 1,
			},
			ivr.AttemptPrompts{
				LangPrArr: []ivr.LanguagePrompts{
					{
						PrArr:    []ivr.PromptID{"FCD392D339FB4A20B95105FAFF8AC7ED_en-US_T_36"},
						Language: "Default",
					}, {
						PrArr:    []ivr.PromptID{"FCD392D339FB4A20B95105FAFF8AC7ED_en-US_T_36"},
						Language: "en-US",
					}, {
						PrArr:    []ivr.PromptID{},
						Language: "uk",
					},
				},
				Count: 2,
			},
		},
		Branches: []*ivr.OutputBranch{
			{"No Match", "C4BA62EF0F3F406A819ADE2CABD1669C", nil},
			{"apples", "C4BA62EF0F3F406A819ADE2CABD1669C", nil},
			{"plums", "C4BA62EF0F3F406A819ADE2CABD1669C", nil},
			{"tutti frutti", "C4BA62EF0F3F406A819ADE2CABD1669C", nil},
		},
		Items: []*ivr.MenuItem{
			{
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_24"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_24"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_24"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "apples"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_25"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_25"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_25"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "appricots"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_26"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_26"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_26"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "plums"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"FFBEE18B77834E10B034889F5457DEB4_en-US_A"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"FFBEE18B77834E10B034889F5457DEB4_en-US_A"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"FFBEE18B77834E10B034889F5457DEB4_uk_A"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "tutti frutti"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"7D0A23161B1B47CEA640C61D490D6FEF_en-US_A"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"7D0A23161B1B47CEA640C61D490D6FEF_en-US_A"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"7D0A23161B1B47CEA640C61D490D6FEF_uk_A"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "appricots"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"3EF9450C730E462FB97E9ADA7A0E37AE_en-US_A"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"3EF9450C730E462FB97E9ADA7A0E37AE_en-US_A"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"3EF9450C730E462FB97E9ADA7A0E37AE_uk_A"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "plums"},
			},
		},
		UseASR:  true,
		UseDTMF: true,
		Events: []*ivr.RecoEvent{
			{
				Event:  "NO_MATCH",
				Action: "CONTINUE",
				CountAndPrompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222__RE__T_19"},
							Language: "Default",
						},
					},
					Count: 1,
				},
			}, {
				Event:  "NO_INPUT",
				Action: "REPROMPT",
				CountAndPrompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222__RE__T_20"},
							Language: "Default",
						},
					},
					Count: 1,
				},
			}, {
				Event:  "NO_INPUT",
				Action: "REPROMPT",
				CountAndPrompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"DAA16853152C4EBFA4DEEDAB00227C2A"},
							Language: "Default",
						},
					},
					Count: 2,
				},
			},
		},
		ConfData: &ivr.ConfirmData{
			ConfirmRequired:      "REQUIRED",
			RequiredConfidence:   75,
			MaxAttemptsToConfirm: 3,
			NoInputTimeout:       3,
			VoicePromptIDs: ivr.ModulePrompts{
				ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"87F5501915CA4ACC874D8767B8C4369E"},
							Language: "Default",
						},
					},
					Count: 1,
				},
			},
			Events: []*ivr.RecoEvent{
				{
					Event:  "NO_MATCH",
					Action: "REPROMPT",
					CountAndPrompt: ivr.AttemptPrompts{
						LangPrArr: []ivr.LanguagePrompts{
							{
								PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_CD___RE__T_22"},
								Language: "Default",
							},
						},
						Count: 1,
					},
				}, {
					Event:  "NO_INPUT",
					Action: "REPROMPT",
					CountAndPrompt: ivr.AttemptPrompts{
						LangPrArr: []ivr.LanguagePrompts{
							{
								PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_CD___RE__T_23"},
								Language: "Default",
							},
						},
						Count: 1,
					},
				}, {
					Event:  "HELP",
					Action: "REPROMPT",
					CountAndPrompt: ivr.AttemptPrompts{
						LangPrArr: []ivr.LanguagePrompts{
							{
								PrArr:    []ivr.PromptID{},
								Language: "Default",
							},
						},
						Count: 1,
					},
				},
			},
		},
		RecoParams: struct {
			SpeechCompleteTimeout int
			MaxTimeToEnter        int
			NoInputTimeout        int
		}{1, 5, 5},
	}
	expected.SetGeneralInfo("Copy of Menu5", "0B19BD7CEE3E4B85BA6C6631F1CCA222",
		[]ivr.ModuleID{"35EBEC8BD6294BA2B6CF4C96D54BE72D"}, "", "C4BA62EF0F3F406A819ADE2CABD1669C",
		"Caller Disconnected", "false")

	if expected.GetID() != mmm.GetID() {
		t.Errorf("\nMenu module, moduleID: \n%v \nwas expected, in reality: \n%v",
			expected.GetID(), mmm.GetID())
	}
	if expected.GetDescendant() != mmm.GetDescendant() {
		t.Errorf("\nMenu module, DescendantID: \n%v \nwas expected, in reality: \n%v", expected.GetDescendant(), mmm.GetDescendant())
	}

}
