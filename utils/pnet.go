package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type PPosition struct {
	id     string
	weight int
}
type PArc struct {
	p     *PPosition
	arity int
}
type PTransition struct {
	id        string
	inbounds  []PArc
	outbounds []PArc
	f         func()
}

type PNet struct {
	positions   []*PPosition
	transitions []*PTransition
}

func (pn *PNet) run(limit int) (err error) {
	var fired = true

	for n := 0; fired && n < limit; n++ {
		fired = false
		for _, t := range pn.transitions {
			fired = t.fire() || fired
		}
		//		fmt.Println(n, fired, pn.positions)
	}
	return
}
func (pn *PNet) randomRun(limit int) (err error) {
	for n := 0; n < limit; n++ {
		var canFire = []*PTransition{}
		for _, tr := range pn.transitions {
			if tr.canFire() {
				canFire = append(canFire, tr)
			}
		}
		switch len(canFire) {
		case 0:
			break
		case 1:
			canFire[0].fire()
		default:
			rand.Seed(time.Now().UnixNano())
			canFire[rand.Intn(len(canFire))].fire()
		}
	}
	return
}

func (pn *PNet) String() string {
	return fmt.Sprintf("{\n\tpositions: %s,\n\ttransitions: %s\n}",
		pn.positions, pn.transitions)
}

func (p *PPosition) String() string {
	return fmt.Sprintf("\n\tPPosition: {id: %s, weight: %d}", p.id, p.weight)
}

func (tr *PTransition) canFire() bool {
	for _, ti := range tr.inbounds {
		if ti.arity > ti.p.weight {
			return false
		}
	}
	return true
}

func (tr *PTransition) fire() bool {
	if !tr.canFire() {
		return false
	}
	for _, ti := range tr.inbounds {
		ti.p.weight -= ti.arity
	}
	for _, to := range tr.outbounds {
		to.p.weight += to.arity
	}
	if tr.f != nil {
		tr.f()
	}
	return true
}

func (tr *PTransition) String() string {
	var (
		ips string
		ops string
	)
	for _, ip := range tr.inbounds {
		ips += fmt.Sprintf("\t%s[%d],", ip.p, ip.arity)
	}
	for _, op := range tr.outbounds {
		ops += fmt.Sprintf("\t%s[%d],", op.p, op.arity)
	}
	return fmt.Sprintf("PTransition: {\nid: %s,\ninbounds: %s,\noutbounds: %s\n}", tr.id, ips, ops)
}
