package utils

import (
	"encoding/xml"
	"fmt"
)

type rule struct {
	XMLName xml.Name `xml:"rule"`
	Items   []string `xml:"item"`
}

func parse(content string) (pn *PNet, err error) {

	var grammarDef rule
	pn = &PNet{positions: []PPosition{{"state0", 1}}}
	lastPosInd := 0

	if err := xml.Unmarshal([]byte(content), &grammarDef); err != nil {
		return nil, err
	}

	initPos:=&PPosition{"init", 1}
	for n, item := range grammarDef.Items {
		input:= &PPosition{string(item), 0}
		repeatPos:=&PPosition{fmt.Sprintf("repeat%d", n), 1}
		statePos:=&PPosition{fmt.Sprintf("state%d", n), 0}

		pn.inputs = append(pn.inputs, input)
		pn.positions = append(pn.positions, repeatPos)
		tr:=PTransition{
			id:fmt.Sprintf("T%d", n),
			inbounds:{
				PArc{&input, 1},
				PArc{&initPos,1},
				PArc{&repeatPos,1},
			},
			outbounds: {
				PArc{statePos, 1},
				PArc{initPos,1},
			} 
			)
		arc1:=PArc{&pn.inputs[inp], 1}
		arc2:=PArc{&pn.positions[inp], 1}

		pn.transitions = append(pn.transitions, tr}

	// inputs:    []PPosition{},
	// 	positions: positions,
	// 	transitions: []PTransition{
	// 		{
	// 			"t1",
	// 			[]PArc{{&positions[0], 1}},
	// 			[]PArc{{&positions[1], 1}, {&positions[2], 1}},
	// 			func() { fmt.Println("T1 Fired!!!", positions) },
	// 		},
	// 		{
	// 			"t2",
	// 			[]PArc{{&positions[1], 1}, {&positions[2], 1}},
	// 			[]PArc{{&positions[3], 1}, {&positions[0], 1}},
	// 			func() { fmt.Println("T2 Fired!!!", positions) },
	// 		},
	// 	},
	// }

	return pn, nil
}
