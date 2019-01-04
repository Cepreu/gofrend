package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type incomingCallModule struct {
	generalInfo
}

func (*incomingCallModule) normalize(*IVRScript) error {
	return nil
}

func newIncomingCallModule(decoder *xml.Decoder) Module {
	var pICM = new(incomingCallModule)

F:
	for {
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
				break F /// <----------------------------------- Return should be HERE!
			}
		}
	}
	return pICM
}
