package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

/*
        (From <N-1> or Begin)  |
	                           V
	                         /^^^\ statePosN
	                        (  0  )------.
                             \___/<-.    |         /^^^\ maxRepeatPosN
	                           |    |    |    ,---(  M  )
	                           |    |    V    V    \___/
	                           |    |   ========
	                           |    |    |    |  t1_N
		                       |    `----'    V
                               |            /~~~\
                               |      ,-m--( 0/1 )  minRepeatPosN
						       |      |     \___/
                               V      V
                             ===========
	                   t2_N    |
	                           V
	                         /^^^\ statePosN+1
	                        (  0  )
              	             \___/

*/
func (pn *PNet) itemGen(initPos *PPosition, ma, mp, M, n int, content string) *PPosition {
	funcTxt := fmt.Sprintf("T1_%d, %s", n, content)
	maxRepeatPos := &PPosition{fmt.Sprintf("maxrepeat%d", n), M}
	minRepeatPos := &PPosition{fmt.Sprintf("minrepeat%d", n), mp}
	statePos := &PPosition{fmt.Sprintf("state%d", n+1), 0}
	pn.positions = append(pn.positions, statePos, maxRepeatPos, minRepeatPos)
	t1 := PTransition{
		id: fmt.Sprintf("t1_%d", n),
		inbounds: []PArc{
			PArc{initPos, 1},
			PArc{maxRepeatPos, 1},
		},
		outbounds: []PArc{
			PArc{initPos, 1},
			PArc{minRepeatPos, 1},
		},
		f: func() { fmt.Println(funcTxt) },
	}
	t2 := PTransition{
		id: fmt.Sprintf("t2_%d", n),
		inbounds: []PArc{
			PArc{initPos, 1},
			PArc{minRepeatPos, ma},
		},
		outbounds: []PArc{
			PArc{statePos, 1}, //
		},
	}
	pn.transitions = append(pn.transitions, &t1, &t2)
	return statePos
}

func mM(repeater string, limit4unlimited int) (ma, mp, M int) {
	var m int
	if repeater == "" {
		m = 1
		M = 1
	} else {
		va := strings.Split(repeater, "-")
		if va[0] == "" {
			m = 1
		} else {
			m, _ = strconv.Atoi(va[0])
		}
		if len(va) < 2 {
			M = m
		} else if va[1] == "" {
			M = m + limit4unlimited
		} else {
			M, _ = strconv.Atoi(va[1])
		}
	}
	mp = 0
	ma = m
	if m == 0 {
		mp = 1
		ma = 1
	}
	return ma, mp, M
}

func grammarParser(src io.Reader) (pns []*PNet, err error) {
	var (
		decoder = xml.NewDecoder(src)
	)
	decoder.CharsetReader = charset.NewReaderLabel

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			// io.EOF is a successful end
			break
		}
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "rule" {
				if pn, err := ruleParser(decoder); err == nil {
					pns = append(pns, pn)
				}
			}
		}
	}
	return pns, nil
}

func ruleParser(decoder *xml.Decoder) (*PNet, error) {
	pn := &PNet{}
	initPos := &PPosition{"state0", 1}
	pn.positions = append(pn.positions, initPos)
	inRule := true
	n := 0

	for inRule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in itemParser with '%s'\n", err)
			return nil, err
		}
		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "item" {
				initPos, err = itemParser(decoder, v.Attr, pn, initPos, n+1)
			}
		case xml.EndElement:
			if v.Name.Local == "rule" {
				inRule = false
			}
		}
	}
	return pn, nil
}

func itemParser(decoder *xml.Decoder, attrs []xml.Attr, pn *PNet, initPos *PPosition, n int) (*PPosition, error) {
	var (
		inItem    = true
		ma, mp, M int
	)

	for _, attr := range attrs {
		if attr.Name.Local == "repeat" {
			ma, mp, M = mM(attr.Value, 10)
			break
		}
	}

	for inItem {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in itemParser with '%s'\n", err)
			return nil, err
		}
		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "item" {
				initPos, err = itemParser(decoder, v.Attr, pn, initPos, n+1)
			} else if v.Name.Local == "one-of" {
				oneOfParser(decoder, pn)
			}
		case xml.CharData:
			initPos = pn.itemGen(initPos, ma, mp, M, n, string(v))
		case xml.EndElement:
			if v.Name.Local == "item" {
				inItem = false
			}
		}
	}
	return initPos, nil
}

func oneOfParser(decoder *xml.Decoder, pn *PNet) error {
	inOneOf := true

	for inOneOf {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in itemParser with '%s'\n", err)
			return err
		}

		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "item" {
				//				itemParser(decoder, pn)
			}
		case xml.CharData:
			fmt.Println("ferfe")
		case xml.EndElement:
			if v.Name.Local == "one-of" {
				inOneOf = false
			}
		}
	}
	return nil
}
