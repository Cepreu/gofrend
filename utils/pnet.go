package utils

import "fmt"

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
	inputs      []PPosition
	positions   []PPosition
	transitions []PTransition
}

func (p *PNet) run(limit int) (err error) {
	var (
		fired = true
		n     int
	)
	for fired && n < limit {
		fired = false
		for _, t := range p.transitions {
			fired = fired || t.fire(p)
		}
		n++
		fmt.Println(n, fired, p.positions)
	}
	return
}

func (tr *PTransition) fire(p *PNet) bool {
	for _, ti := range tr.inbounds {
		if ti.arity > ti.p.weight {
			return false
		}
	}
	for _, ti := range tr.inbounds {
		ti.p.weight -= ti.arity
	}
	for _, to := range tr.outbounds {
		to.p.weight += to.arity
	}
	tr.f()
	return true
}
