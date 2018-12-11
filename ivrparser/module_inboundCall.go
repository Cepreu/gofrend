package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type incomingCallModule struct {
	generalInfo
}

func (s *IVRScript) newIncomingCallModule(decoder *xml.Decoder, v *xml.StartElement, inModulesOnHangup bool) error {
	var pICM = new(incomingCallModule)
	var lastElement string
	if v != nil {
		lastElement = v.Name.Local
	}

F:
	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return err
		}

		switch v := t.(type) {
		case xml.StartElement:
			pICM.parseGeneralInfo(decoder, &v)
		case xml.EndElement:
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!

			}
		}
	}

	if inModulesOnHangup == false {
		s.IncomingCall = pICM
	} else {
		s.ModulesOnHangup.StartOnHangup = pICM
	}
	return nil
}
