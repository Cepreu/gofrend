package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

type xmlHangupModule struct {
	s *ivr.HangupModule
}

func newHangupModule(decoder *xml.Decoder) normalizer {
	var (
		inModule = true
		pHM      = new(ivr.HangupModule)
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
				parse(&pHM.ErrCode, decoder)
			} else if v.Name.Local == "errDescription" {
				parse(&pHM.ErrDescr, decoder)
			} else if v.Name.Local == "overwriteDisposition" {
				innerText, err := decoder.Token()
				if err == nil {
					pHM.OverwriteDisp = string(innerText.(xml.CharData)) == "true"
				}
			} else {
				parseGeneralInfo(pHM, decoder, &v)
			}
		case xml.EndElement:
			if v.Name.Local == cHangup {
				inModule = false /// <--- Return should be HERE!
			}
		}
	}
	return xmlHangupModule{pHM}
}

func (module xmlHangupModule) normalize(s *ivr.IVRScript) error {
	s.Modules[module.s.ID] = module.s
	return nil
}
