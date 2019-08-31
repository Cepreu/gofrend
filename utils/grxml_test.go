package utils

import (
	"strings"
	"testing"
)

func TestGramGeneratorSimple(t *testing.T) {
	testxml := `
<?xml version="1.0"?>
<grammar version="1.0" 
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
         xsi:schemaLocation="http://www.w3.org/2001/06/grammar 
                             http://www.w3.org/TR/speech-grammar/grammar.xsd"
         xmlns="http://www.w3.org/2001/06/grammar">
	<rule id="town"><item repeat="2-"><item repeat="1-3">Townsville</item><item repeat="1-">Beantown</item><item repeat="0-2">Livermore</item></item><item repeat="3-">London</item></rule>
</grammar>
`

	err := grammarParser(strings.NewReader(testxml))
	if err != nil {
		t.Errorf("\nGRXML parser error: %v", err)
		return
	}

	err = pns["town"].randomRun(40)
	t.Errorf("\ngot: %v Result net: %v", err, pns)
}
