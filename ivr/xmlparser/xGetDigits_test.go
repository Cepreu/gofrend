package xmlparser

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/ivr"
	"golang.org/x/net/html/charset"
)

func TestGetDigits(t *testing.T) {
	var xmlData = `
	<getDigits>
	<ascendants>B612F85EA52D4B2586CE5F57579D6EC7</ascendants>
	<exceptionalDescendant>ED132095BE1E4F47B51DA0BB842C3EEF</exceptionalDescendant>
	<singleDescendant>A96A2609FDDE4C499773122F6C6296A1</singleDescendant>
	<moduleName>GetDigits6</moduleName>
	<locationX>341</locationX>
	<locationY>69</locationY>
	<moduleId>F1E142D8CF27471D8940713A637A1C1D</moduleId>
	<data>
		<prompt>
			<TtsPrompt>
				<xml>H4sIAAAAAAAAAIWRzW7CMBCE73mKlc+ELT1VlRNEpfIClN6XZoWi+k9ZB5G3b5JKASdF9cn+xp4Z
rfX2ag1cuJHau0Jt1k8K2H35qnbnQh0/9vmLAonkKjLecaE6FrUtMy2B6fvdsGUXywz6pSnGpj61
keUXjNCQO+964YZG7Mhy2Ue/DrrG8ZjemNw+ybT8RsJwGXaFYpcfDwrvQjBN0TivouvI9r6WULeT
pP4yWHCmzEwmHvka//T63zP1fiSefNWVe98AhWBYIIfQsAhsVtlAg2ntDT6vNY4PlkXxYVON8wnh
ckTTpV5Mvv8HvZuE1kYCAAA=</xml>
				<promptTTSEnumed>false</promptTTSEnumed>
			</TtsPrompt>
			<multiLanguagesPromptItem>
				<prompt>8917F7BDD985458F9E9D33445F4D941D</prompt>
			</multiLanguagesPromptItem>
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
		<numberOfDigits>4</numberOfDigits>
		<maxTime>20</maxTime>
		<maxSilence>2</maxSilence>
		<terminateDigit>#</terminateDigit>
		<clearDigitBuffer>false</clearDigitBuffer>
		<targetVariableName>sergei._TIME_</targetVariableName>
		<format>HHMMP</format>
		<collapsible>true</collapsible>
	</data>
</getDigits>
`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	prompts := make(ivr.ScriptPrompts)
	res := newGetDigitsModule(decoder, prompts)
	if res == nil {
		t.Errorf("Play module wasn't parsed...")
	}
	//res.normalize()
	var mGD = (res.(xmlGetDigitsModule)).m

	expected := &ivr.GetDigitsModule{
		VoicePromptIDs: ivr.ModulePrompts{
			ivr.AttemptPrompts{
				LangPrArr: []ivr.LanguagePrompts{
					{
						PrArr:    []ivr.PromptID{"8917F7BDD985458F9E9D33445F4D941D"},
						Language: "Default",
					},
				},
				Count: 1,
			},
		},
		VisualPromptIDs:    nil,
		TextPromptIDs:      nil,
		TargetVariableName: "sergei._TIME_",
		InputInfo: struct {
			NumberOfDigits   int
			TerminateDigit   string
			ClearDigitBuffer bool
			MaxTime          int
			MaxSilence       int
			Format           string
		}{4, "#", false, 20, 2, "HHMMP"},
	}

	expected.SetGeneralInfo("GetDigits6", "F1E142D8CF27471D8940713A637A1C1D",
		[]ivr.ModuleID{"B612F85EA52D4B2586CE5F57579D6EC7"}, "A96A2609FDDE4C499773122F6C6296A1", "ED132095BE1E4F47B51DA0BB842C3EEF",
		"Caller Disconnected", "true")

	exp, err1 := json.MarshalIndent(expected, "", "  ")
	setv, err2 := json.MarshalIndent(mGD, "", "  ")

	if err1 != nil || err2 != nil || string(exp) != string(setv) {
		t.Errorf("\nGetDigits module: \n%s \n\nwas expected, in reality: \n\n%s", string(exp), string(setv))
	}
}
