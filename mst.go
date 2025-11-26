package graph

import (
	"sort"

	"github.com/subpop/go-adt"
)

// edge represents a weighted edge between two vertices
type edge[V comparable] struct {
	from   V
	to     V
	weight float64
}

// Kruskal computes the minimum spanning tree (or forest for disconnected graphs)
// using Kruskal's algorithm. This algorithm only works on undirected graphs.
// If the graph is directed, it returns DirectedGraphErr. The returned graph
// contains only the edges that form the minimum spanning tree/forest.
func (g *Graph[V]) Kruskal() (Graph[V], error) {
	if g.isDirected {
		return Graph[V]{}, DirectedGraphErr{}
	}

	// Create result graph
	mst := NewGraph[V](false)

	// Add all vertices to the result graph
	for v := range g.vertices {
		if err := mst.AddVertex(v); err != nil {
			return Graph[V]{}, err
		}
	}

	// If graph has 0 or 1 vertices, return it as is
	if g.NumVertex() <= 1 {
		return mst, nil
	}

	// Extract all edges (only once per edge pair in undirected graph)
	var edges []edge[V]
	seen := make(map[V]map[V]bool)
	for u := range g.vertices {
		if seen[u] == nil {
			seen[u] = make(map[V]bool)
		}
		for v, weight := range g.adjacencyMap[u].Explicit {
			if !seen[u][v] {
				edges = append(edges, edge[V]{from: u, to: v, weight: weight})
				if seen[v] == nil {
					seen[v] = make(map[V]bool)
				}
				seen[u][v] = true
				seen[v][u] = true
			}
		}
	}

	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	// Create a disjoint set for each vertex
	sets := make(map[V]*adt.DisjointSet[V])
	for v := range g.vertices {
		sets[v] = adt.NewDisjointSet(v)
	}

	// Process edges in order of weight
	for _, e := range edges {
		// If vertices are in different sets, add edge to MST
		if sets[e.from].Find() != sets[e.to].Find() {
			adt.Union(sets[e.from], sets[e.to])
			// AddEdge will add the edge in both directions for undirected graph
			if err := mst.AddEdge(e.from, e.to, e.weight); err != nil {
				return Graph[V]{}, err
			}
		}
	}

	return mst, nil
}

// Prim computes the minimum spanning tree (or forest for disconnected graphs)
// using Prim's algorithm. This algorithm only works on undirected graphs.
// If the graph is directed, it returns DirectedGraphErr. The returned graph
// contains only the edges that form the minimum spanning tree/forest.
func (g *Graph[V]) Prim() (Graph[V], error) {
	if g.isDirected {
		return Graph[V]{}, DirectedGraphErr{}
	}

	// Create result graph
	mst := NewGraph[V](false)

	// Add all vertices to the result graph
	for v := range g.vertices {
		if err := mst.AddVertex(v); err != nil {
			return Graph[V]{}, err
		}
	}

	// If graph has 0 or 1 vertices, return it as is
	if g.NumVertex() <= 1 {
		return mst, nil
	}

	visited := make(map[V]bool)

	// Process each connected component
	for start := range g.vertices {
		if visited[start] {
			continue
		}

		// Priority queue stores edges: (weight, from, to)
		pq := adt.NewPriorityQueue[edge[V]](0)

		// Start with the initial vertex
		visited[start] = true

		// Add all edges from the start vertex
		for neighbor, weight := range g.adjacencyMap[start].Explicit {
			pq.Push(edge[V]{from: start, to: neighbor, weight: weight}, weight)
		}

		// Process edges in order of weight
		for pq.Len() > 0 {
			e := pq.Pop()
			if e == nil {
				break
			}

			// Skip if the destination vertex is already in the MST
			if visited[e.to] {
				continue
			}

			// Add vertex to MST
			visited[e.to] = true

			// Add edge to MST (AddEdge will add in both directions for undirected graph)
			if err := mst.AddEdge(e.from, e.to, e.weight); err != nil {
				return Graph[V]{}, err
			}

			// Add all edges from the newly added vertex
			for neighbor, weight := range g.adjacencyMap[e.to].Explicit {
				if !visited[neighbor] {
					pq.Push(edge[V]{from: e.to, to: neighbor, weight: weight}, weight)
				}
			}
		}
	}

	return mst, nil
}
