package graph

type MissingVertexErr struct {
	v Vertex
}

func (e *MissingVertexErr) Error() string {
	return "err: missing vertex: " + e.v.ID()
}

type DuplicateVertexErr struct {
	v Vertex
}

func (e *DuplicateVertexErr) Error() string {
	return "err: duplicate vertex: " + e.v.ID()
}

type DuplicateEdgeErr struct {
	a, b Vertex
}

func (e *DuplicateEdgeErr) Error() string {
	return "err: duplicate edge (" + e.a.ID() + " -> " + e.b.ID() + ")"
}

type Graph struct {
	isDirected   bool
	vertices     map[string]Vertex
	adjacencyMap map[string]map[string]*Vertex
}

func NewGraph(isDirected bool) Graph {
	return Graph{
		isDirected:   isDirected,
		vertices:     make(map[string]Vertex),
		adjacencyMap: make(map[string]map[string]*Vertex),
	}
}

func (g *Graph) AddVertex(v Vertex) error {
	if _, ok := g.vertices[v.ID()]; ok {
		return &DuplicateVertexErr{v}
	}

	g.vertices[v.ID()] = v
	g.adjacencyMap[v.ID()] = make(map[string]*Vertex)

	return nil
}

func (g *Graph) AddEdge(a, b *Vertex) error {
	if _, ok := g.vertices[a.ID()]; !ok {
		return &MissingVertexErr{*a}
	}

	if _, ok := g.vertices[b.ID()]; !ok {
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
	neighbors := g.adjacencyMap[a.ID()]
	if _, ok := neighbors[b.ID()]; ok {
		return &DuplicateEdgeErr{*a, *b}
	}
	neighbors[b.ID()] = b
	g.adjacencyMap[a.ID()] = neighbors

	return nil
}

func (g Graph) NumVertex() int {
	return len(g.vertices)
}

// Neighbors returns a slice of Vertices adjacent to v.
func (g Graph) Neighbors(v *Vertex) ([]*Vertex, error) {
	if _, ok := g.vertices[v.ID()]; !ok {
		return nil, &MissingVertexErr{*v}
	}

	neighbors := g.adjacencyMap[v.ID()]

	vertices := make([]*Vertex, 0, len(neighbors))
	for _, vertex := range neighbors {
		vertices = append(vertices, vertex)
	}

	return vertices, nil
}
