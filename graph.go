package graph

import (
	"fmt"
	"reflect"
)

type set map[interface{}]bool
type edgeMap map[interface{}]float64
type adjacencyMap map[interface{}]edgeMap

// A MissingVertexErr describes a vertex that does not exist in a Graph.
type MissingVertexErr struct {
	v interface{}
}

func (e *MissingVertexErr) Error() string {
	return "err: missing vertex: " + reflect.ValueOf(e.v).String()
}

// A MissingEdgeErr describes an edge (a pair of ordered vertices) that does
// not exist in a Graph.
type MissingEdgeErr struct {
	from, to interface{}
}

func (e *MissingEdgeErr) Error() string {
	return "err: missing edge (" + reflect.ValueOf(e.from).String() + " - " + reflect.ValueOf(e.to).String() + ")"
}

// A DuplicateVertexErr describes a vertex that already exists in a Graph.
type DuplicateVertexErr struct {
	v interface{}
}

func (e *DuplicateVertexErr) Error() string {
	return "err: duplicate vertex: " + reflect.ValueOf(e.v).String()
}

// A DuplicateEdgeErr describes an edge (a pair of ordered vertices) that
// already exist in a Graph.
type DuplicateEdgeErr struct {
	from, to interface{}
}

func (e *DuplicateEdgeErr) Error() string {
	return "err: duplicate edge (" + reflect.ValueOf(e.from).String() + " - " + reflect.ValueOf(e.to).String() + ")"
}

// A Graph is an unordered set of nodes along with a set of weighted ordered-pair
// relationships between nodes.
type Graph struct {
	isDirected   bool
	vertices     set
	adjacencyMap adjacencyMap
}

// NewGraph creates a new Graph, enforcing directed edges if isDirected is true.
func NewGraph(isDirected bool) Graph {
	return Graph{
		isDirected:   isDirected,
		vertices:     make(set),
		adjacencyMap: make(adjacencyMap),
	}
}

func (g Graph) String() string {
	out := "{ "
	for a, e := range g.adjacencyMap {
		for b := range e {
			out += fmt.Sprintf("(%v, %v) ", a, b)
		}
	}
	return out + "}"
}

// AddVertex adds v to g. If the graph already contains vertex v, it returns
// DuplicateVertexErr.
func (g *Graph) AddVertex(v interface{}) error {
	if _, ok := g.vertices[v]; ok {
		return &DuplicateVertexErr{v}
	}

	g.vertices[v] = true
	g.adjacencyMap[v] = make(edgeMap)

	return nil
}

// AddVertices adds vertices v to g. If the graph already contains a vertex, it
// returns DuplicateVertexErr.
func (g *Graph) AddVertices(v ...interface{}) error {
	for _, vertex := range v {
		if err := g.AddVertex(vertex); err != nil {
			return err
		}
	}

	return nil
}

// RemoveVertex removes v from g. If the graph does not contain vertex v, it
// returns MissingVertexErr.
func (g *Graph) RemoveVertex(v interface{}) error {
	if _, ok := g.vertices[v]; !ok {
		return &MissingVertexErr{v}
	}

	for n := range g.adjacencyMap[v] {
		delete(g.adjacencyMap[n], v)
	}

	delete(g.adjacencyMap, v)

	delete(g.vertices, v)

	return nil
}

// AddEdge creates a weighted edge from a to b. It adds a and b to the graph if
// they are not already present. If the graph is an undirected graph, the inverse
// edge from b to a is also added. If the edge relationship already exists, a
// DuplicateEdgeErr is returned.
func (g *Graph) AddEdge(a, b interface{}, weight float64) error {
	if _, ok := g.vertices[a]; !ok {
		if err := g.AddVertex(a); err != nil {
			return err
		}
	}

	if _, ok := g.vertices[b]; !ok {
		if err := g.AddVertex(b); err != nil {
			return err
		}
	}

	if err := g.addEdge(a, b, weight); err != nil {
		return err
	}

	if !g.isDirected {
		if err := g.addEdge(b, a, weight); err != nil {
			return err
		}
	}

	return nil
}

func (g *Graph) addEdge(a, b interface{}, weight float64) error {
	neighbors := g.adjacencyMap[a]
	if _, ok := neighbors[b]; ok {
		return &DuplicateEdgeErr{a, b}
	}
	neighbors[b] = weight
	g.adjacencyMap[a] = neighbors

	return nil
}

// RemoveEdge removes an edge from a to b. If a or b are not in the graph,
// it returns MissingVertexErr. If the graph is an undirected graph, the inverse
// edge from b to a is also removed. If the edge does not exist, it returns
// MissingEdgeErr.
func (g *Graph) RemoveEdge(a, b interface{}) error {
	if _, ok := g.vertices[a]; !ok {
		return &MissingVertexErr{a}
	}

	if _, ok := g.vertices[b]; !ok {
		return &MissingVertexErr{b}
	}

	if err := g.removeEdge(a, b); err != nil {
		return err
	}

	if !g.isDirected {
		if err := g.removeEdge(b, a); err != nil {
			return err
		}
	}

	return nil
}

func (g *Graph) removeEdge(a, b interface{}) error {
	neighbors := g.adjacencyMap[a]
	if _, ok := neighbors[b]; !ok {
		return &MissingEdgeErr{a, b}
	}
	delete(neighbors, b)
	g.adjacencyMap[a] = neighbors

	return nil
}

// NumVertex returns the number of vertices in the graph.
func (g Graph) NumVertex() int {
	return len(g.vertices)
}

// Neighbors returns a slice of vertices adjacent to v. If the graph does not
// contain vertex v, it returns MissingVertexErr.
func (g Graph) Neighbors(v interface{}) ([]interface{}, error) {
	if _, ok := g.vertices[v]; !ok {
		return nil, &MissingVertexErr{v}
	}

	neighbors := g.adjacencyMap[v]

	vertices := make([]interface{}, 0, len(neighbors))
	for vertex := range neighbors {
		vertices = append(vertices, vertex)
	}

	return vertices, nil
}
