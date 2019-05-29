package utils

import "fmt"

// A Program to detect cycle in a graph

type Graph [][]int // Pointer to an array containing adjacency lists

func newGraph(V int) Graph {
	return make([][]int, V)
}

func (g Graph) addEdge(v int, w int) error {
	if v < len(g) {
		g[v] = append(g[v], w) // Add w to v’s list.
		return nil
	}
	return fmt.Errorf("incorrect vertix index: %d", v)
}

// used by isCyclic()
// This function is a variation of DFSUytil() in https://www.geeksforgeeks.org/archives/18212
func (g Graph) isCyclicUtil(v int, visited []bool, recStack []bool) bool {
	if visited[v] == false {
		// Mark the current node as visited and part of recursion stack
		visited[v] = true
		recStack[v] = true

		// Recur for all the vertices adjacent to this vertex
		for _, i := range g[v] {
			if !visited[i] && g.isCyclicUtil(i, visited, recStack) {
				return true
			} else if recStack[i] {
				return true
			}
		}
	}
	recStack[v] = false // remove the vertex from recursion stack
	return false
}

// Returns true if the graph contains a cycle, else false.
// This function is a variation of DFS() in https://www.geeksforgeeks.org/archives/18212
// Also, see https://www.geeksforgeeks.org/detect-cycle-in-a-graph/
func (g Graph) isCyclic() bool {
	// Mark all the vertices as not visited and not part of recursion
	// stack
	visited := make([]bool, len(g))
	recStack := make([]bool, len(g))

	// Call the recursive helper function to detect cycle in different
	// DFS trees
	for i := range g {
		if g.isCyclicUtil(i, visited, recStack) {
			return true
		}
	}
	return false
}
