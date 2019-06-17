package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type IncomingCallModule struct {
	GeneralInfo
}

func (*IncomingCallModule) normalize(*IVRScript) error {
	return nil
}

func newIncomingCallModule(decoder *xml.Decoder) Module {
	var (
		inModule = true
		pICM     = new(IncomingCallModule)
	)
	for inModule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			pICM.parseGeneralInfo(decoder, &v)
		case xml.EndElement:
			if v.Name.Local == "incomingCall" || v.Name.Local == "startOnHangup" {
				inModule = false /// <--- Return should be HERE!
			}
		}
	}
	return pICM
}
