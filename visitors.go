package graph

// DepthFirstSearch performs a depth-first traversal of the graph, starting with
// vertex v. It returns a slice of vertices visited during the traversal in
// lexicographic order.
func (g *Graph) DepthFirstSearch(v interface{}) ([]interface{}, error) {
	if _, ok := g.vertices[v]; !ok {
		return nil, &MissingVertexErr{v}
	}
	visited := make(map[interface{}]bool)
	result := make([]interface{}, 0)
	g.visit(v, visited, func(v interface{}) (stop bool) {
		result = append(result, v)
		return
	})
	return result, nil
}

// DepthFirstVisit performs a depth-first traversal of the graph, starting with
// vertex v. The visitorFunc is invoked each time a vertex is visited.
func (g *Graph) DepthFirstVisit(v interface{}, visitorFunc func(v interface{}) (stop bool)) error {
	if _, ok := g.vertices[v]; !ok {
		return &MissingVertexErr{v}
	}
	visited := make(map[interface{}]bool)
	g.visit(v, visited, visitorFunc)
	return nil
}

func (g *Graph) visit(v interface{}, visited map[interface{}]bool, visitorFunc func(v interface{}) (stop bool)) {
	if stop := visitorFunc(v); stop {
		return
	}
	visited[v] = true
	for n := range g.adjacencyMap[v] {
		if _, ok := visited[n]; !ok {
			g.visit(n, visited, visitorFunc)
		}
	}
}

// BreadthFirstSearch performs a breadth-first traversal of the graph, starting
// with vertex v. It returns a slice of vertices visited during the traversal.
func (g *Graph) BreadthFirstSearch(v interface{}) ([]interface{}, error) {
	if _, ok := g.vertices[v]; !ok {
		return nil, &MissingVertexErr{v}
	}
	result := make([]interface{}, 0)
	g.BreadthFirstVisit(v, func(v interface{}) (stop bool) {
		result = append(result, v)
		return
	})
	return result, nil
}

// BreadthFirstVisit performs a breadth-first traversal of the graph, starting
// with vertex v. The visitorFunc is invoked each time a vertex is visited.
func (g *Graph) BreadthFirstVisit(v interface{}, visitorFunc func(v interface{}) (stop bool)) error {
	if _, ok := g.vertices[v]; !ok {
		return &MissingVertexErr{v}
	}
	var q queue
	visited := make(map[interface{}]bool)
	visited[v] = true
	q.enqueue(v)
	for v := q.dequeue(); v != nil; v = q.dequeue() {
		if stop := visitorFunc(v); stop {
			return nil
		}
		for n := range g.adjacencyMap[v] {
			if _, ok := visited[n]; !ok {
				visited[n] = true
				q.enqueue(n)
			}
		}
	}

	return nil
}
