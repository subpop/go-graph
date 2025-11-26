package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNeighborhood(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       struct {
			v     string
			order uint
			min   uint
		}
		want      []string
		wantError error
	}{
		{
			description: "utility graph",
			graph:       UtilityGraph(),
			input: struct {
				v     string
				order uint
				min   uint
			}{
				v:     "a",
				order: 1,
				min:   1,
			},
			want: []string{"x", "y", "z"},
		},
		{
			description: "disconnected",
			graph: Graph[string]{
				vertices: set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct {
						Explicit edgeMap[string]
						Implicit edgeMap[string]
					}{},
					"b": struct {
						Explicit edgeMap[string]
						Implicit edgeMap[string]
					}{},
					"c": struct {
						Explicit edgeMap[string]
						Implicit edgeMap[string]
					}{},
				},
			},
			input: struct {
				v     string
				order uint
				min   uint
			}{
				v:     "a",
				order: 1,
				min:   1,
			},
			want: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.graph.Neighborhood(test.input.v, test.input.order, test.input.min, NoDirection)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(got, test.want, cmpopts.SortSlices(func(x, y string) bool { return x > y })) {
					t.Errorf("%v", cmp.Diff(got, test.want))
				}
			}
		})
	}
}
