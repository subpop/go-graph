package graph

import (
	"github.com/subpop/go-adt"
)

// TopologicalSort performs a variation on a depth-first search to order a
// directed acyclic graph's vertices in such a way that for every vertex, all
// adjacent vertices appear before it in the list.
func (g *Graph) TopologicalSort() ([]interface{}, error) {
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
