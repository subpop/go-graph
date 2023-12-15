package graph

import (
	"github.com/subpop/go-adt"
)

// DepthFirstSearch performs a depth-first traversal of the graph, starting with
// vertex v. It returns a slice of vertices visited during the traversal in
// lexicographic order.
func (g *Graph) DepthFirstSearch(v interface{}, d Direction) ([]interface{}, error) {
	if _, ok := g.vertices[v]; !ok {
		return nil, &MissingVertexErr{v}
	}
	visited := make(map[interface{}]bool)
	result := make([]interface{}, 0)
	err := g.visit(v, visited, d, func(v interface{}) (stop bool) {
		result = append(result, v)
		return
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DepthFirstVisit performs a depth-first traversal of the graph, starting with
// vertex v. The visitorFunc is invoked each time a vertex is visited.
func (g *Graph) DepthFirstVisit(v interface{}, d Direction, visitorFunc func(v interface{}) (stop bool)) error {
	if _, ok := g.vertices[v]; !ok {
		return &MissingVertexErr{v}
	}
	visited := make(map[interface{}]bool)
	if err := g.visit(v, visited, d, visitorFunc); err != nil {
		return err
	}
	return nil
}

func (g *Graph) visit(v interface{}, visited map[interface{}]bool, d Direction, visitorFunc func(v interface{}) (stop bool)) error {
	if stop := visitorFunc(v); stop {
		return nil
	}
	visited[v] = true
	neighbors, err := g.Neighbors(v, d)
	if err != nil {
		return err
	}

	for _, n := range neighbors {
		if _, ok := visited[n]; !ok {
			if err := g.visit(n, visited, d, visitorFunc); err != nil {
				return err
			}
		}
	}
	return nil
}

// BreadthFirstSearch performs a breadth-first traversal of the graph, starting
// with vertex v. It returns a slice of vertices visited during the traversal.
func (g *Graph) BreadthFirstSearch(v interface{}, d Direction) ([]interface{}, error) {
	if _, ok := g.vertices[v]; !ok {
		return nil, &MissingVertexErr{v}
	}
	result := make([]interface{}, 0)
	err := g.BreadthFirstVisit(v, d, func(v interface{}) (stop bool) {
		result = append(result, v)
		return
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BreadthFirstVisit performs a breadth-first traversal of the graph, starting
// with vertex v. The visitorFunc is invoked each time a vertex is visited.
func (g *Graph) BreadthFirstVisit(v interface{}, d Direction, visitorFunc func(v interface{}) (stop bool)) error {
	if _, ok := g.vertices[v]; !ok {
		return &MissingVertexErr{v}
	}
	var q adt.Queue
	visited := make(map[interface{}]bool)
	visited[v] = true
	q.Enqueue(v)
	for v := q.Dequeue(); v != nil; v = q.Dequeue() {
		if stop := visitorFunc(v); stop {
			return nil
		}
		neighbors, err := g.Neighbors(v, d)
		if err != nil {
			return err
		}
		for _, n := range neighbors {
			if _, ok := visited[n]; !ok {
				visited[n] = true
				q.Enqueue(n)
			}
		}
	}

	return nil
}
