package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

var pns = make(map[string]*PNet)

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
func (pn *PNet) itemGen(initPos *PPosition, ma, mp, M, n int, fu func()) *PPosition {
	maxRepeatPos := &PPosition{id: fmt.Sprintf("maxrepeat%d", n), initWeight: M}
	minRepeatPos := &PPosition{id: fmt.Sprintf("minrepeat%d", n), initWeight: mp}
	statePos := &PPosition{id: fmt.Sprintf("state%d", n+1), initWeight: 0}
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
		f: fu,
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

func grammarParser(src io.Reader) (err error) {
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
				var ruleID string
				for _, attr := range v.Attr {
					if attr.Name.Local == "id" {
						ruleID = attr.Value
						break
					}
				}
				if ruleID == "" {
					return fmt.Errorf("Incorrect grammar: Rule with no ID")
				}
				if pn, err := ruleParser(decoder); err == nil {
					pns[ruleID] = pn
				}
			}
		}
	}
	return err
}

func ruleParser(decoder *xml.Decoder) (*PNet, error) {
	var (
		pn           = &PNet{}
		initPos      = &PPosition{id: "state0", initWeight: 1}
		inRule       = true
		ma, mp, M, n int
	)
	pn.positions = append(pn.positions, initPos)

	for inRule {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in ruleParser  with '%s'\n", err)
			return nil, err
		}
		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "item" {
				for _, attr := range v.Attr {
					if attr.Name.Local == "repeat" {
						ma, mp, M = mM(attr.Value, 10)
						break
					}
				}
				if f, err := itemParser(decoder); err == nil {
					initPos = pn.itemGen(initPos, ma, mp, M, n, f)
					n++
				} else {
					return nil, err
				}
			}
		case xml.EndElement:
			if v.Name.Local == "rule" {
				inRule = false
			}
		}
	}
	return pn, nil
}

func itemParser(decoder *xml.Decoder) (func(), error) {
	var (
		firstToken   = true
		inItem       = true
		ma, mp, M, n int
		f            func()
		pn           = &PNet{}
		initPos      = &PPosition{id: "state0", initWeight: 1}
	)
	pn.positions = append(pn.positions, initPos)

	for inItem {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed in itemParser with '%s'\n", err)
			return nil, err
		}
		if firstToken {
			fmt.Println(">>>>>>>>>>", reflect.TypeOf(t).String())
			if reflect.TypeOf(t).String() == "xml.CharData" {
				funcTxt := string(t.(xml.CharData))
				if Strempty(funcTxt) {
					continue
				}
				f = func() { fmt.Println(funcTxt) }
				decoder.Token() // skip </item> before return
				return f, nil
			}
			firstToken = false
		}
		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == "item" {
				for _, attr := range v.Attr {
					if attr.Name.Local == "repeat" {
						ma, mp, M = mM(attr.Value, 10)
						fmt.Println("ma=", ma, "mp=", mp, "M=", M)
						break
					}
				}
				if f, err := itemParser(decoder); err == nil {
					initPos = pn.itemGen(initPos, ma, mp, M, n, f)
					n++
				} else {
					return nil, err
				}
				inItem = true
			} else if v.Name.Local == "one-of" {
				oneOfParser(decoder, pn)
			}
		case xml.CharData:
			if !inItem {
				funcTxt := string(v)
				f = func() { fmt.Println(funcTxt) }
			}
		case xml.EndElement:
			if v.Name.Local == "item" {
				inItem = false
			}
		}
	}
	f = func() { pn.randomRun(20) }

	return f, nil
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
