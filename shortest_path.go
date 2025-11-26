package graph

import (
	"fmt"
	"math"
	"reflect"

	"github.com/subpop/go-adt"
)

// PathResult stores the result of a shortest path computation from a source
// vertex to a destination vertex.
type PathResult[V comparable] struct {
	Distance float64
	Path     []V
}

// A NegativeCycleErr describes a graph that contains a negative cycle.
type NegativeCycleErr[V comparable] struct {
	cycle []V
}

func (e *NegativeCycleErr[V]) Error() string {
	return fmt.Sprintf("err: negative cycle detected: %v", e.cycle)
}

func (e *NegativeCycleErr[V]) Is(target error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(target)
}

// Dijkstra computes the shortest paths from the source vertex to all other
// vertices in the graph using Dijkstra's algorithm. This algorithm only works
// correctly with non-negative edge weights. If the source vertex does not exist
// in the graph, it returns MissingVertexErr. The returned map contains
// PathResult entries for all vertices reachable from the source. Unreachable
// vertices will have a Distance of math.Inf(1).
func (g *Graph[V]) Dijkstra(source V) (map[V]PathResult[V], error) {
	if _, ok := g.vertices[source]; !ok {
		return nil, &MissingVertexErr[V]{v: source}
	}

	dist := make(map[V]float64)
	prev := make(map[V]*V)
	visited := make(map[V]bool)

	// Initialize distances
	for v := range g.vertices {
		dist[v] = math.Inf(1)
	}
	dist[source] = 0

	// Priority queue
	pq := adt.NewPriorityQueue[V](0)
	pq.Push(source, 0)

	for pq.Len() > 0 {
		u := pq.Pop()
		if u == nil {
			break
		}

		if visited[*u] {
			continue
		}
		visited[*u] = true

		// Get neighbors based on graph type
		neighbors := g.adjacencyMap[*u].Explicit
		for v, weight := range neighbors {
			alt := dist[*u] + weight
			if alt < dist[v] {
				dist[v] = alt
				uCopy := *u
				prev[v] = &uCopy
				pq.Push(v, alt)
			}
		}
	}

	// Build results
	results := make(map[V]PathResult[V])
	for v := range g.vertices {
		path := g.reconstructPath(prev, source, v)
		results[v] = PathResult[V]{
			Distance: dist[v],
			Path:     path,
		}
	}

	return results, nil
}

// BellmanFord computes the shortest paths from the source vertex to all other
// vertices using the Bellman-Ford algorithm. This algorithm can handle negative
// edge weights and detects negative cycles. If a negative cycle is detected,
// it returns NegativeCycleErr. If the source vertex does not exist, it returns
// MissingVertexErr.
func (g *Graph[V]) BellmanFord(source V) (map[V]PathResult[V], error) {
	if _, ok := g.vertices[source]; !ok {
		return nil, &MissingVertexErr[V]{v: source}
	}

	dist := make(map[V]float64)
	prev := make(map[V]*V)

	// Initialize distances
	for v := range g.vertices {
		dist[v] = math.Inf(1)
	}
	dist[source] = 0

	// Relax edges |V| - 1 times
	numVertices := len(g.vertices)
	for i := 0; i < numVertices-1; i++ {
		for u := range g.vertices {
			if dist[u] == math.Inf(1) {
				continue
			}
			neighbors := g.adjacencyMap[u].Explicit
			for v, weight := range neighbors {
				alt := dist[u] + weight
				if alt < dist[v] {
					dist[v] = alt
					uCopy := u
					prev[v] = &uCopy
				}
			}
		}
	}

	// Check for negative cycles
	for u := range g.vertices {
		if dist[u] == math.Inf(1) {
			continue
		}
		neighbors := g.adjacencyMap[u].Explicit
		for v, weight := range neighbors {
			if dist[u]+weight < dist[v] {
				// Negative cycle detected, reconstruct it
				cycle := g.findNegativeCycle(prev, v)
				return nil, &NegativeCycleErr[V]{cycle: cycle}
			}
		}
	}

	// Build results
	results := make(map[V]PathResult[V])
	for v := range g.vertices {
		path := g.reconstructPath(prev, source, v)
		results[v] = PathResult[V]{
			Distance: dist[v],
			Path:     path,
		}
	}

	return results, nil
}

// FloydWarshall computes the shortest paths between all pairs of vertices
// using the Floyd-Warshall algorithm. This algorithm can handle negative edge
// weights but will detect negative cycles. If a negative cycle is detected,
// it returns NegativeCycleErr.
func (g *Graph[V]) FloydWarshall() (map[V]map[V]PathResult[V], error) {
	vertices := make([]V, 0, len(g.vertices))
	for v := range g.vertices {
		vertices = append(vertices, v)
	}

	// Initialize distance and next matrices
	dist := make(map[V]map[V]float64)
	next := make(map[V]map[V]*V)

	for _, u := range vertices {
		dist[u] = make(map[V]float64)
		next[u] = make(map[V]*V)
		for _, v := range vertices {
			if u == v {
				dist[u][v] = 0
			} else {
				dist[u][v] = math.Inf(1)
			}
		}
	}

	// Set distances for existing edges
	for u := range g.vertices {
		neighbors := g.adjacencyMap[u].Explicit
		for v, weight := range neighbors {
			dist[u][v] = weight
			vCopy := v
			next[u][v] = &vCopy
		}
	}

	// Floyd-Warshall algorithm
	for _, k := range vertices {
		for _, i := range vertices {
			for _, j := range vertices {
				if dist[i][k] != math.Inf(1) && dist[k][j] != math.Inf(1) {
					alt := dist[i][k] + dist[k][j]
					if alt < dist[i][j] {
						dist[i][j] = alt
						next[i][j] = next[i][k]
					}
				}
			}
		}
	}

	// Check for negative cycles
	for _, v := range vertices {
		if dist[v][v] < 0 {
			return nil, &NegativeCycleErr[V]{cycle: []V{v}}
		}
	}

	// Build results
	results := make(map[V]map[V]PathResult[V])
	for _, u := range vertices {
		results[u] = make(map[V]PathResult[V])
		for _, v := range vertices {
			path := g.reconstructPathFloydWarshall(next, u, v)
			results[u][v] = PathResult[V]{
				Distance: dist[u][v],
				Path:     path,
			}
		}
	}

	return results, nil
}

// AStar computes the shortest path from source to target using the A* algorithm
// with the provided heuristic function. The heuristic function should return an
// estimated cost from any vertex to the target. For the algorithm to be optimal,
// the heuristic must be admissible (never overestimate the actual cost).
// If source or target do not exist in the graph, it returns MissingVertexErr.
func (g *Graph[V]) AStar(source, target V, heuristic func(V) float64) (PathResult[V], error) {
	if _, ok := g.vertices[source]; !ok {
		return PathResult[V]{}, &MissingVertexErr[V]{v: source}
	}
	if _, ok := g.vertices[target]; !ok {
		return PathResult[V]{}, &MissingVertexErr[V]{v: target}
	}

	gScore := make(map[V]float64)
	fScore := make(map[V]float64)
	prev := make(map[V]*V)
	visited := make(map[V]bool)

	// Initialize scores
	for v := range g.vertices {
		gScore[v] = math.Inf(1)
		fScore[v] = math.Inf(1)
	}
	gScore[source] = 0
	fScore[source] = heuristic(source)

	// Priority queue
	pq := adt.NewPriorityQueue[V](0)
	pq.Push(source, fScore[source])

	for pq.Len() > 0 {
		current := pq.Pop()
		if current == nil {
			break
		}

		if *current == target {
			path := g.reconstructPath(prev, source, target)
			return PathResult[V]{
				Distance: gScore[target],
				Path:     path,
			}, nil
		}

		if visited[*current] {
			continue
		}
		visited[*current] = true

		neighbors := g.adjacencyMap[*current].Explicit
		for neighbor, weight := range neighbors {
			if visited[neighbor] {
				continue
			}

			tentativeGScore := gScore[*current] + weight
			if tentativeGScore < gScore[neighbor] {
				currentCopy := *current
				prev[neighbor] = &currentCopy
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = gScore[neighbor] + heuristic(neighbor)
				pq.Push(neighbor, fScore[neighbor])
			}
		}
	}

	// No path found
	return PathResult[V]{
		Distance: math.Inf(1),
		Path:     nil,
	}, nil
}

// BFSShortestPath computes the shortest paths from the source vertex to all
// other vertices in an unweighted graph (or treating all edges as having weight 1)
// using breadth-first search. This is more efficient than Dijkstra for unweighted
// graphs. If the source vertex does not exist, it returns MissingVertexErr.
func (g *Graph[V]) BFSShortestPath(source V) (map[V]PathResult[V], error) {
	if _, ok := g.vertices[source]; !ok {
		return nil, &MissingVertexErr[V]{v: source}
	}

	dist := make(map[V]float64)
	prev := make(map[V]*V)
	visited := make(map[V]bool)

	// Initialize distances
	for v := range g.vertices {
		dist[v] = math.Inf(1)
	}
	dist[source] = 0

	// BFS using queue
	queue := adt.NewQueue[V]()
	queue.Enqueue(source)
	visited[source] = true

	for queue.Len() > 0 {
		u := queue.Dequeue()
		if u == nil {
			break
		}

		neighbors := g.adjacencyMap[*u].Explicit
		for v := range neighbors {
			if !visited[v] {
				visited[v] = true
				dist[v] = dist[*u] + 1
				prev[v] = u
				queue.Enqueue(v)
			}
		}
	}

	// Build results
	results := make(map[V]PathResult[V])
	for v := range g.vertices {
		path := g.reconstructPath(prev, source, v)
		results[v] = PathResult[V]{
			Distance: dist[v],
			Path:     path,
		}
	}

	return results, nil
}

// reconstructPath reconstructs the path from source to target using the
// predecessor map.
func (g *Graph[V]) reconstructPath(prev map[V]*V, source, target V) []V {
	if source == target {
		return []V{source}
	}

	// Check if target is reachable
	if prev[target] == nil && target != source {
		return nil
	}

	path := []V{target}
	current := target
	visited := make(map[V]bool)
	visited[target] = true

	for current != source {
		p := prev[current]
		if p == nil {
			return nil
		}
		current = *p

		if visited[current] {
			// Cycle detected in predecessor chain
			return nil
		}
		visited[current] = true

		path = append([]V{current}, path...)
	}

	return path
}

// reconstructPathFloydWarshall reconstructs the path from source to target
// using the next matrix from Floyd-Warshall algorithm.
func (g *Graph[V]) reconstructPathFloydWarshall(next map[V]map[V]*V, source, target V) []V {
	if source == target {
		return []V{source}
	}

	if next[source][target] == nil {
		return nil
	}

	path := []V{source}
	current := source
	visited := make(map[V]bool)
	visited[source] = true

	for current != target {
		n := next[current][target]
		if n == nil {
			return nil
		}
		current = *n

		if visited[current] && current != target {
			// Cycle detected
			return nil
		}
		visited[current] = true

		path = append(path, current)
	}

	return path
}

// findNegativeCycle attempts to reconstruct a negative cycle starting from
// vertex v using the predecessor map.
func (g *Graph[V]) findNegativeCycle(prev map[V]*V, v V) []V {
	visited := make(map[V]bool)
	current := v

	// Walk backwards until we find a cycle
	for !visited[current] {
		visited[current] = true
		p := prev[current]
		if p == nil {
			break
		}
		current = *p
	}

	// Reconstruct the cycle
	cycle := []V{current}
	start := current
	p := prev[current]
	if p != nil {
		current = *p
		for current != start {
			cycle = append([]V{current}, cycle...)
			p := prev[current]
			if p == nil {
				break
			}
			current = *p
		}
	}

	return cycle
}
