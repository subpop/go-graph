package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		description string
		input       Graph[int]
		want        []int
		wantError   error
	}{
		{
			description: "linear dependency",
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
			want: []int{3, 2, 1, 0},
		},
		{
			description: "undirected graph",
			input: Graph[int]{
				isDirected:   false,
				vertices:     set[int]{},
				adjacencyMap: adjacencyMap[int]{},
			},
			want: nil,
			wantError: &UndirectedGraphErr[int]{
				g: &Graph[int]{
					isDirected:   false,
					vertices:     set[int]{},
					adjacencyMap: adjacencyMap[int]{},
				},
			},
		},
		{
			description: "cycle detected",
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
			want: nil,
			wantError: &CycleDetectedErr[int]{
				g: &Graph[int]{
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
				if !cmp.Equal(got, test.want, cmp.AllowUnexported(Graph[int]{})) {
					t.Errorf("%+v != %+v", got, test.want)
				}
			}
		})
	}
}
