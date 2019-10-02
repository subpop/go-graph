package graph

import "github.com/google/uuid"

// A MissingVertexErr describes a vertex that was expected to exist in a Graph.
type MissingVertexErr struct {
	v Vertex
}

func (e *MissingVertexErr) Error() string {
	return "err: missing vertex: " + e.v.id.String()
}

// A MissingEdgeErr describes an edge (a pair of ordered vertices) that does
// not exist in a Graph.
type MissingEdgeErr struct {
	from Vertex
	to   Vertex
}

func (e *MissingEdgeErr) Error() string {
	return "err: missing edge (" + e.from.id.String() + " -> " + e.to.id.String() + ")"
}

// A DuplicateVertexErr describes a vertex that already exists in a Graph.
type DuplicateVertexErr struct {
	v Vertex
}

func (e *DuplicateVertexErr) Error() string {
	return "err: duplicate vertex: " + e.v.id.String()
}

// A DuplicateEdgeErr describes an edge (a pair of ordered vertices) that
// already exist in a Graph.
type DuplicateEdgeErr struct {
	a, b Vertex
}

func (e *DuplicateEdgeErr) Error() string {
	return "err: duplicate edge (" + e.a.id.String() + " -> " + e.b.id.String() + ")"
}

// A Graph is an unordered set of nodes (represented by the Vertex type) along
// with a set of ordered-pair relationships between nodes.
type Graph struct {
	isDirected   bool
	vertices     map[uuid.UUID]Vertex
	adjacencyMap map[uuid.UUID]map[uuid.UUID]*Vertex
}

// NewGraph creates a new Graph, enforcing directed edges if isDirected is true.
func NewGraph(isDirected bool) Graph {
	return Graph{
		isDirected:   isDirected,
		vertices:     make(map[uuid.UUID]Vertex),
		adjacencyMap: make(map[uuid.UUID]map[uuid.UUID]*Vertex),
	}
}

// AddVertex adds v to g. If the Graph already has a Vertex v, it returns
// DuplicateVertexErr.
func (g *Graph) AddVertex(v Vertex) error {
	if _, ok := g.vertices[v.id]; ok {
		return &DuplicateVertexErr{v}
	}

	g.vertices[v.id] = v
	g.adjacencyMap[v.id] = make(map[uuid.UUID]*Vertex)

	return nil
}

// AddEdge creates an edge from a to b. If a or b are not already in the Graph,
// it returns MissingVertexErr. If the Graph is an undirected graph, the inverse
// edge from b to a is also added. If the edge relationship already exists, a
// DuplicateEdgeErr is returned.
func (g *Graph) AddEdge(a, b *Vertex) error {
	if _, ok := g.vertices[a.id]; !ok {
		return &MissingVertexErr{*a}
	}

	if _, ok := g.vertices[b.id]; !ok {
		return &MissingVertexErr{*b}
	}

	if err := g.addEdge(a, b); err != nil {
		return err
	}

	if !g.isDirected {
		if err := g.addEdge(b, a); err != nil {
			return err
		}
	}

	return nil
}

func (g *Graph) addEdge(a, b *Vertex) error {
	neighbors := g.adjacencyMap[a.id]
	if _, ok := neighbors[b.id]; ok {
		return &DuplicateEdgeErr{*a, *b}
	}
	neighbors[b.id] = b
	g.adjacencyMap[a.id] = neighbors

	return nil
}

// RemoveEdge removes an edge from a to b. If a or be are not in the Graph,
// it returns MissingVertexErr. If the Graph is an undirected graph, the inverse
// edge from b to a is also removed. If the edge does not exist, it returns
// MissingEdgeErr.
func (g *Graph) RemoveEdge(a, b *Vertex) error {
	if _, ok := g.vertices[a.id]; !ok {
		return &MissingVertexErr{*a}
	}

	if _, ok := g.vertices[b.id]; !ok {
		return &MissingVertexErr{*b}
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

func (g *Graph) removeEdge(a, b *Vertex) error {
	neighbors, ok := g.adjacencyMap[a.id]
	if !ok {
		return &MissingEdgeErr{*a, *b}
	}
	delete(neighbors, b.id)
	g.adjacencyMap[a.id] = neighbors

	return nil
}

// NumVertex returns the number of vertices in the graph.
func (g Graph) NumVertex() int {
	return len(g.vertices)
}

// Neighbors returns a slice of Vertices adjacent to v.
func (g Graph) Neighbors(v *Vertex) ([]*Vertex, error) {
	if _, ok := g.vertices[v.id]; !ok {
		return nil, &MissingVertexErr{*v}
	}

	neighbors := g.adjacencyMap[v.id]

	vertices := make([]*Vertex, 0, len(neighbors))
	for _, vertex := range neighbors {
		vertices = append(vertices, vertex)
	}

	return vertices, nil
}
