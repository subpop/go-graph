package graph

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		source      int
		want        map[int]PathResult[int]
		wantError   error
	}{
		{
			description: "simple directed graph",
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
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 2.0,
							4: 5.0,
						},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 1.0,
						},
						Implicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
						},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							2: 5.0,
							3: 1.0,
						},
					},
				},
			},
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
				2: {Distance: 1.0, Path: []int{1, 2}},
				3: {Distance: 3.0, Path: []int{1, 2, 3}},
				4: {Distance: 4.0, Path: []int{1, 2, 3, 4}},
			},
		},
		{
			description: "undirected graph",
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
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							3: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
				2: {Distance: 1.0, Path: []int{1, 2}},
				3: {Distance: 3.0, Path: []int{1, 2, 3}},
			},
		},
		{
			description: "disconnected vertex",
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
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{},
					},
				},
			},
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
				2: {Distance: 1.0, Path: []int{1, 2}},
				3: {Distance: math.Inf(1), Path: nil},
			},
		},
		{
			description: "missing source vertex",
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
			source:    99,
			want:      nil,
			wantError: &MissingVertexErr[int]{v: 99},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.Dijkstra(test.source)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatalf("%v", err)
				}
				if !cmp.Equal(got, test.want) {
					t.Errorf("%+v != %+v", got, test.want)
				}
			}
		})
	}
}

func TestBellmanFord(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		source      int
		want        map[int]PathResult[int]
		wantError   error
	}{
		{
			description: "graph with negative weights",
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
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: -2.0,
							4: 5.0,
						},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 1.0,
						},
						Implicit: edgeMap[int]{
							1: 4.0,
							2: -2.0,
						},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							2: 5.0,
							3: 1.0,
						},
					},
				},
			},
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
				2: {Distance: 1.0, Path: []int{1, 2}},
				3: {Distance: -1.0, Path: []int{1, 2, 3}},
				4: {Distance: 0.0, Path: []int{1, 2, 3, 4}},
			},
		},
		{
			description: "negative cycle detection",
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
							2: 1.0,
						},
						Implicit: edgeMap[int]{
							3: -3.0,
						},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 1.0,
						},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: -3.0,
						},
						Implicit: edgeMap[int]{
							2: 1.0,
						},
					},
				},
			},
			source:    1,
			want:      nil,
			wantError: &NegativeCycleErr[int]{},
		},
		{
			description: "simple positive weights",
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
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 2.0,
						},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
						},
					},
				},
			},
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
				2: {Distance: 1.0, Path: []int{1, 2}},
				3: {Distance: 3.0, Path: []int{1, 2, 3}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.BellmanFord(test.source)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatalf("%v", err)
				}
				if !cmp.Equal(got, test.want) {
					t.Errorf("%+v != %+v", got, test.want)
				}
			}
		})
	}
}

func TestFloydWarshall(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		want        map[int]map[int]PathResult[int]
		wantError   error
	}{
		{
			description: "simple directed graph",
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
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 2.0,
						},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
						},
					},
				},
			},
			want: map[int]map[int]PathResult[int]{
				1: {
					1: {Distance: 0.0, Path: []int{1}},
					2: {Distance: 1.0, Path: []int{1, 2}},
					3: {Distance: 3.0, Path: []int{1, 2, 3}},
				},
				2: {
					1: {Distance: math.Inf(1), Path: nil},
					2: {Distance: 0.0, Path: []int{2}},
					3: {Distance: 2.0, Path: []int{2, 3}},
				},
				3: {
					1: {Distance: math.Inf(1), Path: nil},
					2: {Distance: math.Inf(1), Path: nil},
					3: {Distance: 0.0, Path: []int{3}},
				},
			},
		},
		{
			description: "undirected graph",
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
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
							3: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			want: map[int]map[int]PathResult[int]{
				1: {
					1: {Distance: 0.0, Path: []int{1}},
					2: {Distance: 1.0, Path: []int{1, 2}},
					3: {Distance: 3.0, Path: []int{1, 2, 3}},
				},
				2: {
					1: {Distance: 1.0, Path: []int{2, 1}},
					2: {Distance: 0.0, Path: []int{2}},
					3: {Distance: 2.0, Path: []int{2, 3}},
				},
				3: {
					1: {Distance: 3.0, Path: []int{3, 2, 1}},
					2: {Distance: 2.0, Path: []int{3, 2}},
					3: {Distance: 0.0, Path: []int{3}},
				},
			},
		},
		{
			description: "graph with negative weights",
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
							2: 1.0,
						},
						Implicit: edgeMap[int]{
							3: 1.0,
						},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: -2.0,
						},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 1.0,
						},
						Implicit: edgeMap[int]{
							2: -2.0,
						},
					},
				},
			},
			want: map[int]map[int]PathResult[int]{
				1: {
					1: {Distance: 0.0, Path: []int{1}},
					2: {Distance: 1.0, Path: []int{1, 2}},
					3: {Distance: -1.0, Path: []int{1, 2, 3}},
				},
				2: {
					1: {Distance: -1.0, Path: []int{2, 3, 1}},
					2: {Distance: 0.0, Path: []int{2}},
					3: {Distance: -2.0, Path: []int{2, 3}},
				},
				3: {
					1: {Distance: 1.0, Path: []int{3, 1}},
					2: {Distance: 2.0, Path: []int{3, 1, 2}},
					3: {Distance: 0.0, Path: []int{3}},
				},
			},
		},
		{
			description: "negative cycle detection",
			input: Graph[int]{
				isDirected: true,
				vertices: set[int]{
					1: true,
					2: true,
				},
				adjacencyMap: adjacencyMap[int]{
					1: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: -1.0,
						},
						Implicit: edgeMap[int]{
							2: -1.0,
						},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: -1.0,
						},
						Implicit: edgeMap[int]{
							1: -1.0,
						},
					},
				},
			},
			want:      nil,
			wantError: &NegativeCycleErr[int]{},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.FloydWarshall()

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatalf("%v", err)
				}
				if !cmp.Equal(got, test.want) {
					t.Errorf("%+v != %+v", got, test.want)
				}
			}
		})
	}
}

func TestAStar(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		source      int
		target      int
		heuristic   func(int) float64
		want        PathResult[int]
		wantError   error
	}{
		{
			description: "simple path with zero heuristic",
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
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 2.0,
							4: 5.0,
						},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 1.0,
						},
						Implicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
						},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							2: 5.0,
							3: 1.0,
						},
					},
				},
			},
			source: 1,
			target: 4,
			heuristic: func(v int) float64 {
				return 0.0 // Zero heuristic reduces A* to Dijkstra
			},
			want: PathResult[int]{
				Distance: 4.0,
				Path:     []int{1, 2, 3, 4},
			},
		},
		{
			description: "path with admissible heuristic",
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
							2: 1.0,
							3: 4.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 2.0,
							4: 5.0,
						},
						Implicit: edgeMap[int]{
							1: 1.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 1.0,
						},
						Implicit: edgeMap[int]{
							1: 4.0,
							2: 2.0,
						},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							2: 5.0,
							3: 1.0,
						},
					},
				},
			},
			source: 1,
			target: 4,
			heuristic: func(v int) float64 {
				// Manhattan distance assuming vertices represent grid positions
				distances := map[int]float64{
					1: 3.0,
					2: 2.0,
					3: 1.0,
					4: 0.0,
				}
				return distances[v]
			},
			want: PathResult[int]{
				Distance: 4.0,
				Path:     []int{1, 2, 3, 4},
			},
		},
		{
			description: "no path exists",
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
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{},
					},
				},
			},
			source: 1,
			target: 3,
			heuristic: func(v int) float64 {
				return 0.0
			},
			want: PathResult[int]{
				Distance: math.Inf(1),
				Path:     nil,
			},
		},
		{
			description: "missing source vertex",
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
			source: 99,
			target: 1,
			heuristic: func(v int) float64 {
				return 0.0
			},
			want:      PathResult[int]{},
			wantError: &MissingVertexErr[int]{v: 99},
		},
		{
			description: "missing target vertex",
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
			source: 1,
			target: 99,
			heuristic: func(v int) float64 {
				return 0.0
			},
			want:      PathResult[int]{},
			wantError: &MissingVertexErr[int]{v: 99},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.AStar(test.source, test.target, test.heuristic)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatalf("%v", err)
				}
				if !cmp.Equal(got, test.want) {
					t.Errorf("%+v != %+v", got, test.want)
				}
			}
		})
	}
}

func TestBFSShortestPath(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		source      int
		want        map[int]PathResult[int]
		wantError   error
	}{
		{
			description: "simple directed graph",
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
							2: 0.0,
							3: 0.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 0.0,
						},
						Implicit: edgeMap[int]{
							1: 0.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							4: 0.0,
						},
						Implicit: edgeMap[int]{
							1: 0.0,
						},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							2: 0.0,
							3: 0.0,
						},
					},
				},
			},
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
				2: {Distance: 1.0, Path: []int{1, 2}},
				3: {Distance: 1.0, Path: []int{1, 3}},
				4: {Distance: 2.0, Path: []int{1, 2, 4}},
			},
		},
		{
			description: "undirected graph",
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
							2: 0.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							1: 0.0,
							3: 0.0,
						},
						Implicit: edgeMap[int]{},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							2: 0.0,
							4: 0.0,
						},
						Implicit: edgeMap[int]{},
					},
					4: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{
							3: 0.0,
						},
						Implicit: edgeMap[int]{},
					},
				},
			},
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
				2: {Distance: 1.0, Path: []int{1, 2}},
				3: {Distance: 2.0, Path: []int{1, 2, 3}},
				4: {Distance: 3.0, Path: []int{1, 2, 3, 4}},
			},
		},
		{
			description: "disconnected vertex",
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
							2: 0.0,
						},
						Implicit: edgeMap[int]{},
					},
					2: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{
							1: 0.0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap[int] }{
						Explicit: edgeMap[int]{},
						Implicit: edgeMap[int]{},
					},
				},
			},
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
				2: {Distance: 1.0, Path: []int{1, 2}},
				3: {Distance: math.Inf(1), Path: nil},
			},
		},
		{
			description: "single vertex",
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
			source: 1,
			want: map[int]PathResult[int]{
				1: {Distance: 0.0, Path: []int{1}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.BFSShortestPath(test.source)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatalf("%v", err)
				}
				// For BFS, check distances and path lengths (not exact paths due to map iteration order)
				for vertex, wantResult := range test.want {
					gotResult, ok := got[vertex]
					if !ok {
						t.Errorf("missing vertex %v in result", vertex)
						continue
					}
					if gotResult.Distance != wantResult.Distance {
						t.Errorf("vertex %v: distance %f != %f", vertex, gotResult.Distance, wantResult.Distance)
					}
					if len(gotResult.Path) != len(wantResult.Path) {
						t.Errorf("vertex %v: path length %d != %d", vertex, len(gotResult.Path), len(wantResult.Path))
					}
					// Verify path starts with source and ends with vertex
					if len(gotResult.Path) > 0 {
						if gotResult.Path[0] != test.source {
							t.Errorf("vertex %v: path doesn't start with source %v: %v", vertex, test.source, gotResult.Path)
						}
						if gotResult.Path[len(gotResult.Path)-1] != vertex {
							t.Errorf("vertex %v: path doesn't end with vertex: %v", vertex, gotResult.Path)
						}
					}
				}
			}
		})
	}
}

func TestDijkstraWithStrings(t *testing.T) {
	// Test with string vertices to verify generics work properly
	g := Graph[string]{
		isDirected: true,
		vertices: set[string]{
			"A": true,
			"B": true,
			"C": true,
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
					"C": 2.0,
				},
				Implicit: edgeMap[string]{
					"A": 1.0,
				},
			},
			"C": struct{ Explicit, Implicit edgeMap[string] }{
				Explicit: edgeMap[string]{},
				Implicit: edgeMap[string]{
					"A": 4.0,
					"B": 2.0,
				},
			},
		},
	}

	got, err := g.Dijkstra("A")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := map[string]PathResult[string]{
		"A": {Distance: 0.0, Path: []string{"A"}},
		"B": {Distance: 1.0, Path: []string{"A", "B"}},
		"C": {Distance: 3.0, Path: []string{"A", "B", "C"}},
	}

	if !cmp.Equal(got, want) {
		t.Errorf("%+v != %+v", got, want)
	}
}
