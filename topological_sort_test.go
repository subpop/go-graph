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
				vertices: set{
					0: true,
					1: true,
					2: true,
					3: true,
					4: true,
					5: true,
				},
				adjacencyMap: adjacencyMap{
					0: edgeMap{},
					1: edgeMap{},
					2: edgeMap{
						3: 0,
					},
					3: edgeMap{
						1: 0,
					},
					4: edgeMap{
						0: 0,
						1: 0,
					},
					5: edgeMap{
						2: 0,
						0: 0,
					},
				},
			},
			want: []interface{}{5, 4, 2, 3, 1, 0},
		},
	}

	for _, test := range tests {
		got, err := test.input.TopologicalSort()

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v != %v", err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%v != %v", got, test.want)
			}
		}
	}
}
