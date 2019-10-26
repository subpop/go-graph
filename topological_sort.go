package graph

import (
	"fmt"

	"github.com/subpop/go-adt"
)


// An UndirectedGraphErr describes a graph that is undirected.
type UndirectedGraphErr struct {
	g *Graph
}

func (e *UndirectedGraphErr) Error() string {
	return fmt.Sprintf("err: graph is undirected: %v", e.g)
}

// TopologicalSort performs a variation on a depth-first search to order a
// directed acyclic graph's vertices in such a way that for every vertex, all
// adjacent vertices appear before it in the list. If graph is undirected, an
// error is returned.
func (g *Graph) TopologicalSort() ([]interface{}, error) {
	if !g.isDirected {
		return nil, &UndirectedGraphErr{g: g}
	}

	var stack adt.Stack

	visited := make(map[interface{}]bool)

	for v := range g.vertices {
		if _, ok := visited[v]; !ok {
			g.topologicalSort(v, visited, &stack)
		}
	}

	sorted := make([]interface{}, 0)
	for {
		f := stack.Pop()
		if f == nil {
			break
		}
		sorted = append(sorted, f)
	}
	return sorted, nil
}

func (g *Graph) topologicalSort(v interface{}, visited map[interface{}]bool, stack *adt.Stack) {
	visited[v] = true

	for n := range g.adjacencyMap[v] {
		if _, ok := visited[n]; !ok {
			g.topologicalSort(n, visited, stack)
		}
	}
	stack.Push(v)
}
