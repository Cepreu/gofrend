package ivrparser

import (
	"encoding/xml"
	"strings"
	"testing"

	"golang.org/x/net/html/charset"
)

func TestPlay(t *testing.T) {
	var xmlData = `
<play>
<ascendants>B612F85EA52D4B2586CE5F57579D6EC7</ascendants>
<ascendants>F1E142D8CF27471D8940713A637A1C1D</ascendants>
<singleDescendant>A96A2609FDDE4C499773122F6C6296A1</singleDescendant>
<moduleName>Play4</moduleName>
<locationX>419</locationX>
<locationY>155</locationY>
<moduleId>ED132095BE1E4F47B51DA0BB842C3EEF</moduleId>
<data>
	<prompt>
		<ttsPrompt>
			<xml>H4sIAAAAAAAAANVUsW7DIBTc8xUWe0K7VRV2lEqp1KWV3CZrhcOTiwq4AmzFf19MFddgrC5d4gm/
Q/fu3T1Btmcpsg604Y3K0e3mBmWgTg3jqs7R4e1xfYcyY6liVDQKctSDQdtiRcwX0M+9AAnKFqvM
fYRaq3nVWjA/BV8UVNU7B/yWfFlRCYVrfT/gBPvf8MbIdqSihQdqIOuGU45ArQ+vCE+a4LALwbEU
wi3IqSxD+50J5M8bGxwhEclYt3C2Sa6/OUPuJbBqWF/sy/KlJNif5xrwogiC4+FxenpS6TjTxACJ
3sZqULX9mOc8XvEBX+6l8k6S+egvsUtgvJUo4ZEbaFHBfBdGJLbceZs0gHRUc1oJuK59+Y9VCcZ/
HkJ7OpYb0LrR7wzMieAAnFiZ9Gzkd4kF78c3S60nzIcEAAA=</xml>
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
	<numberOfDigits>0</numberOfDigits>
	<terminateDigit>N/A</terminateDigit>
	<clearDigitBuffer>false</clearDigitBuffer>
	<collapsible>false</collapsible>
</data>
</play>
`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	prompts := make(scriptPrompts)
	res := newPlayModule(decoder, prompts)
	if res == nil {
		t.Errorf("Play module wasn't parsed...")
	}
	// more sanity checking...
}
