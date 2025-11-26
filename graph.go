package graph

import (
	"fmt"
	"reflect"
)

// Direction is the direction of a relationship between two vertices in
// a directed graph. It has no meaning in undirected graphs.
type Direction int

const (
	// NoDirection represents the absence of direction. In a directed graph,
	// this means both inbound and outbound edges.
	NoDirection Direction = iota

	// Outbound represents only edges going "out" from a vertex.
	Outbound

	// Inbound represents only edges going "in" to a vertex.
	Inbound
)

type set[V comparable] map[V]bool
type edgeMap[V comparable] map[V]float64
type adjacencyMap[V comparable] map[V]struct{ Explicit, Implicit edgeMap[V] }

// A MissingVertexErr describes a vertex that does not exist in a Graph.
type MissingVertexErr[V comparable] struct {
	v V
}

func (e *MissingVertexErr[V]) Error() string {
	return "err: missing vertex: " + reflect.ValueOf(e.v).String()
}

func (e *MissingVertexErr[V]) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// A MissingEdgeErr describes an edge (a pair of ordered vertices) that does
// not exist in a Graph.
type MissingEdgeErr[V comparable] struct {
	from, to V
}

func (e *MissingEdgeErr[V]) Error() string {
	return "err: missing edge (" + reflect.ValueOf(e.from).String() + " - " + reflect.ValueOf(e.to).String() + ")"
}

func (e *MissingEdgeErr[V]) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// A DuplicateVertexErr describes a vertex that already exists in a Graph.
type DuplicateVertexErr[V comparable] struct {
	v V
}

func (e *DuplicateVertexErr[V]) Error() string {
	return "err: duplicate vertex: " + reflect.ValueOf(e.v).String()
}

func (e *DuplicateVertexErr[V]) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// A DuplicateEdgeErr describes an edge (a pair of ordered vertices) that
// already exist in a Graph.
type DuplicateEdgeErr[V comparable] struct {
	from, to V
}

func (e *DuplicateEdgeErr[V]) Error() string {
	return "err: duplicate edge (" + reflect.ValueOf(e.from).String() + " - " + reflect.ValueOf(e.to).String() + ")"
}

func (e *DuplicateEdgeErr[V]) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// A Graph is an unordered set of nodes along with a set of weighted ordered-pair
// relationships between nodes.
type Graph[V comparable] struct {
	isDirected   bool
	vertices     set[V]
	adjacencyMap adjacencyMap[V]
}

// NewGraph creates a new Graph, enforcing directed edges if isDirected is true.
func NewGraph[V comparable](isDirected bool) Graph[V] {
	return Graph[V]{
		isDirected:   isDirected,
		vertices:     make(set[V]),
		adjacencyMap: make(adjacencyMap[V]),
	}
}

func (g Graph[V]) String() string {
	out := "{ "
	for a, e := range g.adjacencyMap {
		for b := range e.Explicit {
			out += fmt.Sprintf("(%v, %v) ", a, b)
		}
	}
	return out + "}"
}

// AddVertex adds v to g. If the graph already contains vertex v, it returns
// DuplicateVertexErr.
func (g *Graph[V]) AddVertex(v V) error {
	if _, ok := g.vertices[v]; ok {
		return &DuplicateVertexErr[V]{v}
	}

	g.vertices[v] = true
	g.adjacencyMap[v] = struct{ Explicit, Implicit edgeMap[V] }{
		Explicit: make(edgeMap[V]),
		Implicit: make(edgeMap[V]),
	}

	return nil
}

// AddVertices adds vertices v to g. If the graph already contains a vertex, it
// returns DuplicateVertexErr.
func (g *Graph[V]) AddVertices(v ...V) error {
	for _, vertex := range v {
		if err := g.AddVertex(vertex); err != nil {
			return err
		}
	}

	return nil
}

// RemoveVertex removes v from g. If the graph does not contain vertex v, it
// returns MissingVertexErr.
func (g *Graph[V]) RemoveVertex(v V) error {
	if _, ok := g.vertices[v]; !ok {
		return &MissingVertexErr[V]{v}
	}

	for n := range g.adjacencyMap[v].Explicit {
		delete(g.adjacencyMap[n].Explicit, v)
	}
	for n := range g.adjacencyMap[v].Implicit {
		delete(g.adjacencyMap[n].Implicit, v)
	}

	delete(g.adjacencyMap, v)

	delete(g.vertices, v)

	return nil
}

// AddEdge creates a weighted edge from a to b. It adds a and b to the graph if
// they are not already present. If the graph is an undirected graph, the inverse
// edge from b to a is also added. If the edge relationship already exists, a
// DuplicateEdgeErr is returned.
func (g *Graph[V]) AddEdge(a, b V, weight float64) error {
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

	if err := g.addExplicitEdge(a, b, weight); err != nil {
		return err
	}

	if g.isDirected {
		// In a directed graph, adding an edge from a to b adds an Implicit edge from b to a.
		if err := g.addImplicitEdge(b, a, weight); err != nil {
			return err
		}
	} else {
		// In an undirected graph, adding an edge from a to b adds an Explicit edge from b to a.
		if err := g.addExplicitEdge(b, a, weight); err != nil {
			return err
		}
	}

	return nil
}

func (g *Graph[V]) addExplicitEdge(a, b V, weight float64) error {
	edges := g.adjacencyMap[a]
	if _, ok := edges.Explicit[b]; ok {
		return &DuplicateEdgeErr[V]{a, b}
	}
	edges.Explicit[b] = weight
	g.adjacencyMap[a] = edges

	return nil
}

func (g *Graph[V]) addImplicitEdge(a, b V, weight float64) error {
	edges := g.adjacencyMap[a]
	if _, ok := edges.Implicit[b]; ok {
		return &DuplicateEdgeErr[V]{a, b}
	}
	edges.Implicit[b] = weight
	g.adjacencyMap[a] = edges

	return nil
}

// RemoveEdge removes an edge from a to b. If a or b are not in the graph,
// it returns MissingVertexErr. If the graph is an undirected graph, the inverse
// edge from b to a is also removed. If the edge does not exist, it returns
// MissingEdgeErr.
func (g *Graph[V]) RemoveEdge(a, b V) error {
	if _, ok := g.vertices[a]; !ok {
		return &MissingVertexErr[V]{a}
	}

	if _, ok := g.vertices[b]; !ok {
		return &MissingVertexErr[V]{b}
	}

	if err := g.removeExplicitEdge(a, b); err != nil {
		return err
	}

	if g.isDirected {
		if err := g.removeImplicitEdge(b, a); err != nil {
			return err
		}
	} else {
		if err := g.removeExplicitEdge(b, a); err != nil {
			return err
		}
	}

	return nil
}

func (g *Graph[V]) removeExplicitEdge(a, b V) error {
	edges := g.adjacencyMap[a]
	if _, ok := edges.Explicit[b]; !ok {
		return &MissingEdgeErr[V]{a, b}
	}
	delete(edges.Explicit, b)
	g.adjacencyMap[a] = edges

	return nil
}

func (g *Graph[V]) removeImplicitEdge(a, b V) error {
	edges := g.adjacencyMap[a]
	if _, ok := edges.Implicit[b]; !ok {
		return &MissingEdgeErr[V]{a, b}
	}
	delete(edges.Implicit, b)
	g.adjacencyMap[a] = edges

	return nil
}

// NumVertex returns the number of vertices in the graph.
func (g Graph[V]) NumVertex() int {
	return len(g.vertices)
}

// Neighbors returns a slice of vertices adjacent to v, given direction d. If
// the graph is undirected, d is ignored. If the graph does not contain vertex
// v, it returns MissingVertexErr.
func (g Graph[V]) Neighbors(v V, d Direction) ([]V, error) {
	if _, ok := g.vertices[v]; !ok {
		return nil, &MissingVertexErr[V]{v}
	}

	var neighbors edgeMap[V]
	if g.isDirected {
		// In a directed graph, the neighbors of a vertex v are the set of
		// vertices:
		// - to which v has an explicit edge if direction d is outbound.
		// - to which v has an implicit edge if direction d is inbound.
		// - to which v has an explicit or implicit edge if direction d is
		//   unqualified.
		switch d {
		case Outbound:
			neighbors = g.adjacencyMap[v].Explicit
		case Inbound:
			neighbors = g.adjacencyMap[v].Implicit
		default:
			neighbors = g.adjacencyMap[v].Explicit
			for k, v := range g.adjacencyMap[v].Implicit {
				neighbors[k] = v
			}
		}
	} else {
		// In an undirected graph, the neighbors of a vertex v are the set of
		// vertices to which v has an explicit edge.
		neighbors = g.adjacencyMap[v].Explicit
	}

	vertices := make([]V, 0, len(neighbors))
	for vertex := range neighbors {
		vertices = append(vertices, vertex)
	}

	return vertices, nil
}
