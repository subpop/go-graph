package graph

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Helper function to calculate total MST weight
func mstWeight[V comparable](g Graph[V]) float64 {
	total := 0.0
	seen := make(map[V]map[V]bool)
	for u := range g.vertices {
		if seen[u] == nil {
			seen[u] = make(map[V]bool)
		}
		for v, weight := range g.adjacencyMap[u].Explicit {
			if !seen[u][v] {
				total += weight
				if seen[v] == nil {
					seen[v] = make(map[V]bool)
				}
				seen[u][v] = true
				seen[v][u] = true
			}
		}
	}
	return total
}

func TestKruskal(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		wantWeight  float64
		wantEdges   int
		wantError   error
	}{
		{
			description: "simple connected undirected graph",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							3: 2.0,
							4: 5.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
							4: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 5.0,
							3: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 4.0, // edges: (1,2)=1, (2,3)=2, (3,4)=1
			wantEdges:  3,
		},
		{
			description: "disconnected graph (spanning forest)",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 3.0, // edges: (1,2)=1, (3,4)=2
			wantEdges:  2,
		},
		{
			description: "single vertex",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 0.0,
			wantEdges:  0,
		},
		{
			description: "empty graph",
			input: Graph[int]{
				isDirected: false,
				vertices:   set[int]{},
				adjacencyMap: adjacencyMap[int]{},
			},
			wantWeight: 0.0,
			wantEdges:  0,
		},
		{
			description: "graph with equal weight edges",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
							3: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							3: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							2: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 2.0,
			wantEdges:  2,
		},
		{
			description: "directed graph error",
			input: Graph[int]{
				isDirected: true,
				vertices: set[int]{
					1: true,
					2: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
				},
			},
			wantError: DirectedGraphErr{},
		},
		{
			description: "complex graph",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
					5: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 2.0,
							3: 3.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 2.0,
							3: 1.0,
							4: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 3.0,
							2: 1.0,
							4: 5.0,
							5: 6.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 4.0,
							3: 5.0,
							5: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					5: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 6.0,
							4: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 8.0, // edges: (2,3)=1, (4,5)=1, (1,2)=2, (2,4)=4
			wantEdges:  4,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.Kruskal()

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("error: %#v != %#v", err, test.wantError)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check number of vertices matches
			if got.NumVertex() != test.input.NumVertex() {
				t.Errorf("vertex count: %d != %d", got.NumVertex(), test.input.NumVertex())
			}

			// Check total weight
			gotWeight := mstWeight(got)
			if math.Abs(gotWeight-test.wantWeight) > 0.0001 {
				t.Errorf("total weight: %f != %f", gotWeight, test.wantWeight)
			}

			// Check number of edges (count once per edge pair in undirected graph)
			edgeCount := 0
			seen := make(map[int]map[int]bool)
			for u := range got.vertices {
				if seen[u] == nil {
					seen[u] = make(map[int]bool)
				}
				for v := range got.adjacencyMap[u].Explicit {
					if !seen[u][v] {
						edgeCount++
						if seen[v] == nil {
							seen[v] = make(map[int]bool)
						}
						seen[u][v] = true
						seen[v][u] = true
					}
				}
			}
			if edgeCount != test.wantEdges {
				t.Errorf("edge count: %d != %d", edgeCount, test.wantEdges)
			}
		})
	}
}

func TestPrim(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		wantWeight  float64
		wantEdges   int
		wantError   error
	}{
		{
			description: "simple connected undirected graph",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							3: 2.0,
							4: 5.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
							4: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 5.0,
							3: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 4.0,
			wantEdges:  3,
		},
		{
			description: "disconnected graph (spanning forest)",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 3.0,
			wantEdges:  2,
		},
		{
			description: "single vertex",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 0.0,
			wantEdges:  0,
		},
		{
			description: "empty graph",
			input: Graph[int]{
				isDirected: false,
				vertices:   set[int]{},
				adjacencyMap: adjacencyMap[int]{},
			},
			wantWeight: 0.0,
			wantEdges:  0,
		},
		{
			description: "graph with equal weight edges",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
							3: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							3: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							2: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 2.0,
			wantEdges:  2,
		},
		{
			description: "directed graph error",
			input: Graph[int]{
				isDirected: true,
				vertices: set[int]{
					1: true,
					2: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
				},
			},
			wantError: DirectedGraphErr{},
		},
		{
			description: "complex graph",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
					5: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 2.0,
							3: 3.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 2.0,
							3: 1.0,
							4: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 3.0,
							2: 1.0,
							4: 5.0,
							5: 6.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 4.0,
							3: 5.0,
							5: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					5: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 6.0,
							4: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			wantWeight: 8.0,
			wantEdges:  4,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.Prim()

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("error: %#v != %#v", err, test.wantError)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check number of vertices matches
			if got.NumVertex() != test.input.NumVertex() {
				t.Errorf("vertex count: %d != %d", got.NumVertex(), test.input.NumVertex())
			}

			// Check total weight
			gotWeight := mstWeight(got)
			if math.Abs(gotWeight-test.wantWeight) > 0.0001 {
				t.Errorf("total weight: %f != %f", gotWeight, test.wantWeight)
			}

			// Check number of edges
			edgeCount := 0
			seen := make(map[int]map[int]bool)
			for u := range got.vertices {
				if seen[u] == nil {
					seen[u] = make(map[int]bool)
				}
				for v := range got.adjacencyMap[u].Explicit {
					if !seen[u][v] {
						edgeCount++
						if seen[v] == nil {
							seen[v] = make(map[int]bool)
						}
						seen[u][v] = true
						seen[v][u] = true
					}
				}
			}
			if edgeCount != test.wantEdges {
				t.Errorf("edge count: %d != %d", edgeCount, test.wantEdges)
			}
		})
	}
}

// TestKruskalVsPrim verifies that both algorithms produce MSTs with the same total weight
func TestKruskalVsPrim(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
	}{
		{
			description: "simple graph",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							3: 2.0,
							4: 5.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
							4: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 5.0,
							3: 1.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
		},
		{
			description: "complex graph",
			input: Graph[int]{
				isDirected: false,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
					5: true,
					6: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 7.0,
							3: 9.0,
							6: 14.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 7.0,
							3: 10.0,
							4: 15.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 9.0,
							2: 10.0,
							4: 11.0,
							6: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 15.0,
							3: 11.0,
							5: 6.0,
						},
						Implicit: edgeMap[int]{},
					},
					5: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 6.0,
							6: 9.0,
						},
						Implicit: edgeMap[int]{},
					},
					6: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 14.0,
							3: 2.0,
							5: 9.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			kruskalMST, err := test.input.Kruskal()
			if err != nil {
				t.Fatalf("Kruskal error: %v", err)
			}

			primMST, err := test.input.Prim()
			if err != nil {
				t.Fatalf("Prim error: %v", err)
			}

			kruskalWeight := mstWeight(kruskalMST)
			primWeight := mstWeight(primMST)

			if math.Abs(kruskalWeight-primWeight) > 0.0001 {
				t.Errorf("MST weights differ: Kruskal=%f, Prim=%f", kruskalWeight, primWeight)
			}
		})
	}
}

// TestMSTWithStrings verifies MST works with string vertices (testing generics)
func TestMSTWithStrings(t *testing.T) {
	g := Graph[string]{
		isDirected: false,
		vertices: set[string]{
			"A": true,
			"B": true,
			"C": true,
			"D": true,
		},
		adjacencyMap: adjacencyMap[string]{
			"A": struct{ Explicit, Implicit edgeMap[string] }{
				Explicit: edgeMap[string]{
					"B": 1.0,
					"C": 4.0,
				},
				Implicit: edgeMap[string]{},
			},
			"B": struct{ Explicit, Implicit edgeMap[string] }{
				Explicit: edgeMap[string]{
					"A": 1.0,
					"C": 2.0,
					"D": 5.0,
				},
				Implicit: edgeMap[string]{},
			},
			"C": struct{ Explicit, Implicit edgeMap[string] }{
				Explicit: edgeMap[string]{
					"A": 4.0,
					"B": 2.0,
					"D": 1.0,
				},
				Implicit: edgeMap[string]{},
			},
			"D": struct{ Explicit, Implicit edgeMap[string] }{
				Explicit: edgeMap[string]{
					"B": 5.0,
					"C": 1.0,
				},
				Implicit: edgeMap[string]{},
			},
		},
	}

	kruskalMST, err := g.Kruskal()
	if err != nil {
		t.Fatalf("Kruskal error: %v", err)
	}

	primMST, err := g.Prim()
	if err != nil {
		t.Fatalf("Prim error: %v", err)
	}

	kruskalWeight := mstWeight(kruskalMST)
	primWeight := mstWeight(primMST)

	expectedWeight := 4.0
	if math.Abs(kruskalWeight-expectedWeight) > 0.0001 {
		t.Errorf("Kruskal weight: %f != %f", kruskalWeight, expectedWeight)
	}
	if math.Abs(primWeight-expectedWeight) > 0.0001 {
		t.Errorf("Prim weight: %f != %f", primWeight, expectedWeight)
	}
}

