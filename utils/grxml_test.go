package utils

import (
	"fmt"
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
	<rule id="town">
		<item>Townsville</item>
		<item>Beantown</item> 
	</rule>
</grammar>
`
	fmt.Println(testxml)
}
