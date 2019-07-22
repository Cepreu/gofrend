package xmlparser

import (
	"encoding/xml"
	"fmt"

	"github.com/Cepreu/gofrend/ivr"
)

type xmlSkillTransferModule struct {
	m *ivr.SkillTransferModule
}

func newSkillTransferModule(decoder *xml.Decoder, script *ivr.IVRScript) normalizer {
	var (
		inside = true
		pSTM   = new(ivr.SkillTransferModule)
	)

	for inside {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in SkillTransfer with '%s'", err)
			return nil
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "data" {
				if err := decoder.DecodeElement(&pSTM.Data, &v); err != nil {
					fmt.Printf("decoder.DecodeElement() failed in SkillTransfer with '%s'", err)
				}

			} else {
				parseGeneralInfo(pSTM, decoder, &v)
			}

		case xml.EndElement:
			if v.Name.Local == cSkillTransfer {
				inside = false
			}
		}
	}
	return xmlSkillTransferModule{pSTM}
}

func (module xmlSkillTransferModule) normalize(s *ivr.IVRScript) error {
	s.Modules[module.m.ID] = module.m
	return nil
}
