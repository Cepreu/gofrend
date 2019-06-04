package ivrparser

import (
	"encoding/xml"
	"fmt"
)

type hangupModule struct {
	GeneralInfo
	Return2Caller bool

	ErrCode       parametrized
	ErrDescr      parametrized
	OverwriteDisp bool
}

func (*hangupModule) normalize(*IVRScript) error {
	return nil
}

func newHangupModule(decoder *xml.Decoder) Module {
	var (
		inModule = true
		pHM      = new(hangupModule)
	)
	for inModule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "returnToCallingModule" {
				innerText, err := decoder.Token()
				if err == nil {
					pHM.Return2Caller = string(innerText.(xml.CharData)) == "true"
				}
			} else if v.Name.Local == "errCode" {
				pHM.ErrCode.parse(decoder)
			} else if v.Name.Local == "errDescription" {
				pHM.ErrDescr.parse(decoder)
			} else if v.Name.Local == "overwriteDisposition" {
				innerText, err := decoder.Token()
				if err == nil {
					pHM.OverwriteDisp = string(innerText.(xml.CharData)) == "true"
				}
			} else {
				pHM.parseGeneralInfo(decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == cHangup {
				inModule = false /// <--- Return should be HERE!
			}
		}
	}
	return pHM
}
