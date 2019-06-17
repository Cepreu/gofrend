package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

func newIncomingCallModule(decoder *xml.Decoder) ivr.Module {
	var (
		inModule = true
		pICM     = new(ivr.IncomingCallModule)
	)
	for inModule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			parseGeneralInfo(pICM, decoder, &v)
		case xml.EndElement:
			if v.Name.Local == "incomingCall" || v.Name.Local == "startOnHangup" {
				inModule = false /// <--- Return should be HERE!
			}
		}
	}
	return pICM
}
