package utils

import (
	"fmt"
	"testing"
)

func TestPetri(t *testing.T) {
	positions := []PPosition{{"p1", 1}, {"p2", 0}, {"p3", 2}, {"p4", 1}}
	pn := &PNet{
		inputs:    []PPosition{},
		positions: positions,
		transitions: []PTransition{
			{
				"t1",
				[]PArc{{&positions[0], 1}},
				[]PArc{{&positions[1], 1}, {&positions[2], 1}},
				func() { fmt.Println("T1 Fired!!!", positions) },
			},
			{
				"t2",
				[]PArc{{&positions[1], 1}, {&positions[2], 1}},
				[]PArc{{&positions[3], 1}, {&positions[0], 1}},
				func() { fmt.Println("T2 Fired!!!", positions) },
			},
		},
	}
	err := pn.run(10)
	t.Errorf("\ngot: %v Result net: %v", err, pn)
}
