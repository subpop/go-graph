package graph

// ConnectedComponent performs a depth-first traversal of g to return the
// connected component containing v. In a directed graph, the component returned
// is the strong component.
func (g *Graph[V]) ConnectedComponent(v V) ([]V, error) {
	connectedComponent := make([]V, 0)
	visited := make(set[V])

	var d Direction
	if g.isDirected {
		d = Outbound
	}

	err := g.DepthFirstVisit(v, d, func(x V) (stop bool) {
		visited[x] = true
		connectedComponent = append(connectedComponent, x)
		return
	})
	if err != nil {
		return nil, err
	}

	return connectedComponent, nil
}

// ConnectedComponents performs a depth-first traversal of each vertex in g to
// return the set of connected components in g. In a directed graph, the
// components returned are the strongly connected components.
func (g *Graph[V]) ConnectedComponents() ([][]V, error) {
	components := make([][]V, 0)
	visited := make(set[V])

	var d Direction
	if g.isDirected {
		d = Outbound
	}

	for v := range g.vertices {
		if !visited[v] {
			connectedComponent := make([]V, 0)
			err := g.DepthFirstVisit(v, d, func(x V) (stop bool) {
				visited[x] = true
				connectedComponent = append(connectedComponent, x)
				return
			})
			if err != nil {
				return nil, err
			}
			components = append(components, connectedComponent)
		}
	}

	return components, nil
}
