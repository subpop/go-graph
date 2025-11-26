package graph

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		want        bool
	}{
		{
			description: "directed graph with cycle",
			input: Graph[int]{
				isDirected: true,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 0,
						},
						Implicit: edgeMap[int]{
							3: 0,
						},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 0,
						},
						Implicit: edgeMap[int]{
							1: 0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0,
						},
						Implicit: edgeMap[int]{
							2: 0,
						},
					},
				},
			},
			want: true,
		},
		{
			description: "directed acyclic graph (DAG)",
			input: Graph[int]{
				isDirected: true,
				vertices: set[int]{
					0: true,
					1: true,
					2: true,
					3: true,
				},
				adjacencyMap: adjacencyMap[int]{
					0: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							1: 0,
						},
					},
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							0: 0,
						},
						Implicit: edgeMap[int]{
							2: 0,
						},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0,
						},
						Implicit: edgeMap[int]{
							3: 0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			want: false,
		},
		{
			description: "undirected graph with cycle",
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
							2: 0,
							3: 0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0,
							3: 0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0,
							2: 0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			want: true,
		},
		{
			description: "undirected graph without cycle (tree)",
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
							2: 0,
							3: 0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0,
							4: 0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			want: false,
		},
		{
			description: "empty graph",
			input: Graph[int]{
				isDirected:   true,
				vertices:     set[int]{},
				adjacencyMap: adjacencyMap[int]{},
			},
			want: false,
		},
		{
			description: "single vertex graph",
			input: Graph[int]{
				isDirected: true,
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
			want: false,
		},
		{
			description: "single vertex with self-loop",
			input: Graph[int]{
				isDirected: true,
				vertices: set[int]{
					1: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0,
						},
						Implicit: edgeMap[int]{
							1: 0,
						},
					},
				},
			},
			want: true,
		},
		{
			description: "disconnected graph with cycle in one component",
			input: Graph[int]{
				isDirected: true,
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
							2: 0,
						},
						Implicit: edgeMap[int]{
							3: 0,
						},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 0,
						},
						Implicit: edgeMap[int]{
							1: 0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0,
						},
						Implicit: edgeMap[int]{
							2: 0,
						},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							5: 0,
						},
						Implicit: edgeMap[int]{},
					},
					5: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							4: 0,
						},
					},
				},
			},
			want: true,
		},
		{
			description: "disconnected graph without cycles",
			input: Graph[int]{
				isDirected: true,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							1: 0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							3: 0,
						},
					},
				},
			},
			want: false,
		},
		{
			description: "directed graph with back edge creating cycle",
			input: Graph[int]{
				isDirected: true,
				vertices: set[int]{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 0,
						},
						Implicit: edgeMap[int]{
							1: 0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 0,
						},
						Implicit: edgeMap[int]{
							2: 0,
						},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 0,
						},
						Implicit: edgeMap[int]{
							3: 0,
						},
					},
				},
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := test.input.HasCycle()
			if got != test.want {
				t.Errorf("HasCycle() = %v, want %v", got, test.want)
			}
		})
	}
}
