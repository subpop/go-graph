package graph

import (
	"fmt"
	"reflect"

	"github.com/subpop/go-adt"
)

// A CycleDetectedErr describes a graph that contains a cycle.
type CycleDetectedErr[V comparable] struct {
	g *Graph[V]
}

func (e *CycleDetectedErr[V]) Error() string {
	return fmt.Sprintf("err: cycle detected in graph: %v", e.g)
}

func (e *CycleDetectedErr[V]) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// An UndirectedGraphErr describes a graph that is undirected.
type UndirectedGraphErr[V comparable] struct {
	g *Graph[V]
}

func (e *UndirectedGraphErr[V]) Error() string {
	return fmt.Sprintf("err: graph is undirected: %v", e.g)
}

func (e *UndirectedGraphErr[V]) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// TopologicalSort performs a variation on a depth-first search to order a
// directed acyclic graph's vertices in such a way that for every vertex, all
// adjacent vertices appear before it in the list. If graph is undirected, an
// error is returned. If a cycle is detected, an error is returned.
func (g *Graph[V]) TopologicalSort() ([]V, error) {
	if !g.isDirected {
		return nil, &UndirectedGraphErr[V]{g: g}
	}

	var stack adt.Stack[V]

	visited := make(map[V]bool)
	recStack := make(map[V]bool)

	for v := range g.vertices {
		if !visited[v] {
			if err := g.topologicalSort(v, visited, recStack, &stack); err != nil {
				return nil, err
			}
		}
	}

	sorted := make([]V, 0)
	for {
		f := stack.Pop()
		if f == nil {
			break
		}
		sorted = append(sorted, *f)
	}
	return sorted, nil
}

func (g *Graph[V]) topologicalSort(v V, visited map[V]bool, recStack map[V]bool, stack *adt.Stack[V]) error {
	visited[v] = true
	recStack[v] = true

	for n := range g.adjacencyMap[v].Explicit {
		if !visited[n] {
			if err := g.topologicalSort(n, visited, recStack, stack); err != nil {
				return err
			}
		} else if recStack[n] {
			return &CycleDetectedErr[V]{g: g}
		}
	}

	recStack[v] = false
	stack.Push(v)

	return nil
}
