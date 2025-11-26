
[![PkgGoDev](https://pkg.go.dev/badge/github.com/subpop/go-graph)](https://pkg.go.dev/github.com/subpop/go-graph)
[![Go Report Card](https://goreportcard.com/badge/github.com/subpop/go-graph)](https://goreportcard.com/report/github.com/subpop/go-graph)

Another graph package in Go.

## Features

### Graph Construction
- Generic type support for vertex types
- Directed and undirected graphs
- Weighted edges
- Add/remove vertices and edges

### Shortest Path Algorithms
- Dijkstra (non-negative weights)
- Bellman-Ford (handles negative weights, detects negative cycles)
- Floyd-Warshall (all-pairs shortest paths)
- A* (heuristic-based pathfinding)
- BFS shortest path (unweighted graphs)

### Minimum Spanning Trees
- Kruskal's algorithm
- Prim's algorithm

### Graph Algorithms
- Topological sort (for DAGs)
- Connected components detection
- Neighborhood queries with distance filters

### Graph Traversal
- Depth-first search (DFS)
- Breadth-first search (BFS)
- Visitor pattern support for custom traversal logic
