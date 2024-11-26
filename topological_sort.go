package graph

import (
	"fmt"
	"reflect"

	"github.com/subpop/go-adt"
)

// A CycleDetectedErr describes a graph that contains a cycle.
type CycleDetectedErr struct {
	g *Graph
}

func (e *CycleDetectedErr) Error() string {
	return fmt.Sprintf("err: cycle detected in graph: %v", e.g)
}

func (e *CycleDetectedErr) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// An UndirectedGraphErr describes a graph that is undirected.
type UndirectedGraphErr struct {
	g *Graph
}

func (e *UndirectedGraphErr) Error() string {
	return fmt.Sprintf("err: graph is undirected: %v", e.g)
}

func (e *UndirectedGraphErr) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// TopologicalSort performs a variation on a depth-first search to order a
// directed acyclic graph's vertices in such a way that for every vertex, all
// adjacent vertices appear before it in the list. If graph is undirected, an
// error is returned. If a cycle is detected, an error is returned.
func (g *Graph) TopologicalSort() ([]interface{}, error) {
	if !g.isDirected {
		return nil, &UndirectedGraphErr{g: g}
	}

	var stack adt.Stack

	visited := make(map[interface{}]bool)

	for v := range g.vertices {
		if err := g.topologicalSort(v, visited, &stack); err != nil {
			return nil, err
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

func (g *Graph) topologicalSort(v interface{}, visited map[interface{}]bool, stack *adt.Stack) error {
	if visited[v] {
		return &CycleDetectedErr{g: g}
	}
	visited[v] = true

	for n := range g.adjacencyMap[v].Explicit {
		if _, ok := visited[n]; !ok {
			if err := g.topologicalSort(n, visited, stack); err != nil {
				return err
			}
		}
	}
	stack.Push(v)

	return nil
}
