package graph

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		input     Graph
		want      []interface{}
		wantError error
	}{
		{
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

	for i, test := range tests {
		got, err := test.input.TopologicalSort()

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatalf("%v: %v", i, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%v: %+v != %+v", i, got, test.want)
			}
		}
	}
}
