package ivrparser

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Cepreu/gofrend/utils"
	"golang.org/x/net/html/charset"
)

func TestInBodyPrompts(t *testing.T) {
	var xmlData = `
<prompts>
<prompt>
	<ttsPrompt>
		<xml>H4sIAAAAAAAAAIVSwU7DMAy97yusnNkCN4TSVkMCiTMMzl5rrdHSpKqdsf49bSeVtWUiJ+c9571n
KyY7Vw5O1LANPlEPm3sF5PNQWH9I1O7jdf2ogAV9gS54SlRLrLJ0ZbgmPL44qshLuoLuGBRp7D4K
8QUYQIf+sO2IX2iAPVaUdtZPPW/0cJ12jGqf6CI9IxOc+ipR5Ne7d6WvTPTUxeh5FGOFqutYjO2W
J/GXxqxnzExkxIXO8qfW/5pT7VvkPhRt+gY5ehByDtoQQUoCLkMNZYgN30Fo4NLBJBBr6Cqs62C9
9Lk28FXavITvEF0xvHf2SJnRg/RyJH1zJqPnu9TLZY5NHTn5KD/TopiPcAIAAA==</xml>
		<promptTTSEnumed>true</promptTTSEnumed>
	</ttsPrompt>
	<interruptible>false</interruptible>
	<canChangeInterruptableOption>true</canChangeInterruptableOption>
	<ttsEnumed>true</ttsEnumed>
	<exitModuleOnException>false</exitModuleOnException>
</prompt>
<count>1</count>
</prompts>
<prompts>
<prompt>
	<ttsPrompt>
		<xml>H4sIAAAAAAAAAIVSQWrDMBC85xWLzk3U3kqRbRJoIec27XkTb2MRWTLedYh/X9kBN7YbqoukmWFm
Vshkl9LBmWq2wSfqafWogPwh5NYfE7X7eFs+K2BBn6MLnhLVEqssXRiuCE+vjkryki4gLoMitd03
QnwFetChP64j8Qv1sMeS0hj90vFG99exYnD7RNfQBpng3J0SRX65e1f6JkSPU4yeVjFWqLytxdiu
eVR/Hsx6wkxMBlzoIn96/e859r5H7kPepls4oAch56ANDUhBwEWooAhNzQ8QargqSjwRxB2rKlgv
XasVfBUoPbuFPMB3FEePzOjeeT6RvjuS0dOn1PO3HESRHP2THycFBWhvAgAA</xml>
		<promptTTSEnumed>true</promptTTSEnumed>
	</ttsPrompt>
	<interruptible>false</interruptible>
	<canChangeInterruptableOption>true</canChangeInterruptableOption>
	<ttsEnumed>true</ttsEnumed>
	<exitModuleOnException>false</exitModuleOnException>
</prompt>
<count>2</count>
</prompts>
<prompts>
<prompt>
	<ttsPrompt>
		<xml>H4sIAAAAAAAAAIVSQU7DMBC89xUrn2kNN4ScREUCwRnKfZusmqiON8o6pf49TiqFJqHCJ3tmNDO7
ssnOtYUTtVKxS9TD5l4BuZyLyh0Stft8XT8qEI+uQMuOEhVIVJaujDSExxdLNTmfriAeg9631b7z
JBdgAC26wzYSv9AAO6wpjdFPPW/08JwqRrcvtB09oxCc+luiyK13H0pfhehpitHzKqbyVF/XEgxb
mdRfBoueMTOTEfd09n96/e859b5F7rkI6Tvk6MCTtRC4A18SSMkNlNy1cgfcwkUheUlFZwniHZuG
K+f7Zht442+oMURZSba5mHCBITN6CFgOpm9OZvR8o3q50lEUycl3+QGMRxZMdgIAAA==</xml>
		<promptTTSEnumed>true</promptTTSEnumed>
	</ttsPrompt>
	<interruptible>false</interruptible>
	<canChangeInterruptableOption>true</canChangeInterruptableOption>
	<ttsEnumed>true</ttsEnumed>
	<exitModuleOnException>false</exitModuleOnException>
</prompt>
<count>3</count>
</prompts>`

	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel
	prompts := make(ScriptPrompts)
	ap, err := parseVoicePromptS(decoder, prompts, "_V")
	utils.PrettyPrint(ap)
	utils.PrettyPrint(prompts)
	if err != nil {
		t.Errorf("Prompts weren't parsed...%v", err)
	}
}
