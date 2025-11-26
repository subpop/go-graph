package graph

// HasCycle returns true if the graph contains a cycle, false otherwise.
// For directed graphs, it uses DFS with a recursion stack to detect back edges.
// For undirected graphs, it uses DFS with parent tracking to detect cycles.
func (g *Graph[V]) HasCycle() bool {
	visited := make(map[V]bool)

	if g.isDirected {
		// For directed graphs, use recursion stack approach
		recStack := make(map[V]bool)
		for v := range g.vertices {
			if !visited[v] {
				if g.hasCycleDirected(v, visited, recStack) {
					return true
				}
			}
		}
	} else {
		// For undirected graphs, use parent tracking approach
		var zeroValue V
		for v := range g.vertices {
			if !visited[v] {
				if g.hasCycleUndirected(v, visited, zeroValue, true) {
					return true
				}
			}
		}
	}

	return false
}

// hasCycleDirected performs DFS on a directed graph to detect cycles.
// It uses a recursion stack to track vertices in the current DFS path.
// A cycle exists if we encounter a vertex already in the recursion stack.
func (g *Graph[V]) hasCycleDirected(v V, visited map[V]bool, recStack map[V]bool) bool {
	visited[v] = true
	recStack[v] = true

	for n := range g.adjacencyMap[v].Explicit {
		if !visited[n] {
			if g.hasCycleDirected(n, visited, recStack) {
				return true
			}
		} else if recStack[n] {
			// Back edge found - cycle detected
			return true
		}
	}

	recStack[v] = false
	return false
}

// hasCycleUndirected performs DFS on an undirected graph to detect cycles.
// It uses parent tracking to avoid false positives from the edge we came from.
// A cycle exists if we visit a vertex that's already visited and isn't the parent.
func (g *Graph[V]) hasCycleUndirected(v V, visited map[V]bool, parent V, isRoot bool) bool {
	visited[v] = true

	for n := range g.adjacencyMap[v].Explicit {
		if !visited[n] {
			if g.hasCycleUndirected(n, visited, v, false) {
				return true
			}
		} else if isRoot || n != parent {
			// We found a visited vertex that's not our parent - cycle detected
			return true
		}
	}

	return false
}
