package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type hangupModule struct {
	generalInfo
	Return2Caller bool

	ErrCode       parametrized
	ErrDescr      parametrized
	OverwriteDisp bool
}

func (s *IVRScript) parseHangup(decoder *xml.Decoder, v *xml.StartElement, inModulesOnHangup bool) error {
	var pHM = new(hangupModule)
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
			if v.Name.Local == "returnToCallingModule" {
				innerText, err := decoder.Token()
				if err == nil {
					pHM.Return2Caller = string(innerText.(xml.CharData)) == "true"
				}
			} else if v.Name.Local == "errCode" {
				pHM.ErrCode.parse(decoder, &v)
			} else if v.Name.Local == "errDescription" {
				pHM.ErrDescr.parse(decoder, &v)
			} else if v.Name.Local == "overwriteDisposition" {
				innerText, err := decoder.Token()
				if err == nil {
					pHM.OverwriteDisp = (string(innerText.(xml.CharData)) == "true")
				}
			} else {
				pHM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == lastElement {
				break F /// <----------------------------------- Return should be HERE!

			}
		}
	}

	if inModulesOnHangup == false {
		s.HangupModules = append(s.HangupModules, pHM)
	} else {
		s.ModulesOnHangup.HangupModules = append(s.ModulesOnHangup.HangupModules, pHM)
	}
	return nil
}
