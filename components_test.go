package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestConnectedComponent(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       string
		want        []string
		wantError   error
	}{
		{
			description: "utility graph",
			graph: Graph[string]{vertices: set[string]{"a": true, "b": true, "c": true, "x": true, "y": true, "z": true}, adjacencyMap: adjacencyMap[string]{"a": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap[string]{}}, "b": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap[string]{}}, "c": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap[string]{}}, "x": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0}}, "y": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0}}, "z": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0}}}},
			input: "a",
			want:  []string{"a", "b", "c", "x", "y", "z"},
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
				if !cmp.Equal(got, test.want, cmpopts.SortSlices(func(x, y string) bool {
					return x < y
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
		input       Graph[string]
		want        [][]string
		wantError   error
	}{
		{
			description: "utility graph",
			input: Graph[string]{vertices: set[string]{"a": true, "b": true, "c": true, "x": true, "y": true, "z": true}, adjacencyMap: adjacencyMap[string]{"a": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap[string]{}}, "b": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap[string]{}}, "c": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0}, Implicit: edgeMap[string]{}}, "x": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0}}, "y": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0}}, "z": struct {
				Explicit edgeMap[string]
				Implicit edgeMap[string]
			}{Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0}}}},
			want: [][]string{{"a", "b", "c", "x", "y", "z"}},
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
				if !cmp.Equal(got, test.want, cmpopts.SortSlices(func(x, y string) bool {
					return x > y
				})) {
					t.Errorf("%#v != %#v", got, test.want)
				}
			}
		})
	}
}
