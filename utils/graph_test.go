package utils

import (
	"testing"
)

func TestGraphCyclic(t *testing.T) {
	// Create a graph given in the above diagram
	g := newGraph(4)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	if !g.isCyclic() {
		t.Errorf("\nTestGraph : should be True but is False")
	}
}
func TestGraphUncycling(t *testing.T) {
	// Create a graph given in the above diagram
	g := newGraph(4)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 3)
	g.addEdge(0, 3)

	if g.isCyclic() {
		t.Errorf("\nTestGraph : should be False but is True")
	}
}
