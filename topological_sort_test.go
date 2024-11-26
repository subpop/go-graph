package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		description string
		input       Graph
		want        []interface{}
		wantError   error
	}{
		{
			description: "linear dependency",
			input: Graph{
				isDirected: true,
				vertices: set{
					0: true,
					1: true,
					2: true,
					3: true,
				},
				adjacencyMap: adjacencyMap{
					0: struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{
							1: 0,
						},
					},
					1: struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							0: 0,
						},
						Implicit: edgeMap{
							2: 0,
						},
					},
					2: struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							1: 0,
						},
						Implicit: edgeMap{
							3: 0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							2: 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			want: []interface{}{3, 2, 1, 0},
		},
		{
			description: "undirected graph",
			input: Graph{
				isDirected:   false,
				vertices:     set{},
				adjacencyMap: adjacencyMap{},
			},
			want: nil,
			wantError: &UndirectedGraphErr{
				g: &Graph{
					isDirected:   false,
					vertices:     set{},
					adjacencyMap: adjacencyMap{},
				},
			},
		},
		{
			description: "cycle detected",
			input: Graph{
				isDirected: true,
				vertices: set{
					1: true,
					2: true,
					3: true,
				},
				adjacencyMap: adjacencyMap{
					1: struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							2: 0,
						},
						Implicit: edgeMap{
							3: 0,
						},
					},
					2: struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							3: 0,
						},
						Implicit: edgeMap{
							1: 0,
						},
					},
					3: struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							1: 0,
						},
						Implicit: edgeMap{
							2: 0,
						},
					},
				},
			},
			want: nil,
			wantError: &CycleDetectedErr{
				g: &Graph{
					isDirected: true,
					vertices: set{
						1: true,
						2: true,
						3: true,
					},
					adjacencyMap: adjacencyMap{
						1: struct{ Explicit, Implicit edgeMap }{
							Explicit: edgeMap{
								2: 0,
							},
							Implicit: edgeMap{
								3: 0,
							},
						},
						2: struct{ Explicit, Implicit edgeMap }{
							Explicit: edgeMap{
								3: 0,
							},
							Implicit: edgeMap{
								1: 0,
							},
						},
						3: struct{ Explicit, Implicit edgeMap }{
							Explicit: edgeMap{
								1: 0,
							},
							Implicit: edgeMap{
								2: 0,
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.TopologicalSort()

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatalf("%v", err)
				}
				if !cmp.Equal(got, test.want, cmp.AllowUnexported(Graph{})) {
					t.Errorf("%+v != %+v", got, test.want)
				}
			}
		})
	}
}
