package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestConnectedComponent(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph
		input       interface{}
		want        []interface{}
		wantError   error
	}{
		{
			description: "utility graph",
			graph: Graph{vertices: set{"a": true, "b": true, "c": true, "x": true, "y": true, "z": true}, adjacencyMap: adjacencyMap{"a": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap{}}, "b": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap{}}, "c": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap{}}, "x": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"a": 0, "b": 0, "c": 0}}, "y": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"a": 0, "b": 0, "c": 0}}, "z": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"a": 0, "b": 0, "c": 0}}}},
			input: "a",
			want:  []interface{}{"a", "b", "c", "x", "y", "z"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.graph.ConnectedComponent(test.input)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(got, test.want, cmpopts.SortSlices(func(x, y interface{}) bool {
					return x.(string) < y.(string)
				})) {
					t.Errorf("%#v != %#v", got, test.want)
				}
			}
		})
	}
}

func TestConnectedComponents(t *testing.T) {
	tests := []struct {
		description string
		input       Graph
		want        [][]interface{}
		wantError   error
	}{
		{
			description: "utility graph",
			input: Graph{vertices: set{"a": true, "b": true, "c": true, "x": true, "y": true, "z": true}, adjacencyMap: adjacencyMap{"a": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap{}}, "b": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap{}}, "c": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap{}}, "x": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"a": 0, "b": 0, "c": 0}}, "y": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"a": 0, "b": 0, "c": 0}}, "z": struct {
				Explicit edgeMap
				Implicit edgeMap
			}{Explicit: edgeMap{"a": 0, "b": 0, "c": 0}}}},
			want: [][]interface{}{{"a", "b", "c", "x", "y", "z"}},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.input.ConnectedComponents()

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(got, test.want, cmpopts.SortSlices(func(x, y interface{}) bool {
					return x.(string) > y.(string)
				})) {
					t.Errorf("%#v != %#v", got, test.want)
				}
			}
		})
	}
}
