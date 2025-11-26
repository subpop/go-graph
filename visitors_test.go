package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDepthFirstSearch(t *testing.T) {
	tests := []struct {
		graph     Graph[string]
		input     string
		want      []string
		wantError error
	}{
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
			want:  []string{"a", "b", "c"},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     "x",
			want:      []string{},
			wantError: &MissingVertexErr[string]{"x"},
		},
	}

	for i, test := range tests {
		got, err := test.graph.DepthFirstSearch(test.input, Outbound)

		if test.wantError != nil {
			if !cmp.Equal(err, test.wantError, cmp.AllowUnexported(MissingVertexErr[string]{})) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, test.want) {
				t.Errorf("%v: %v != %v", i, got, test.want)
			}
		}
	}
}

func TestDepthFirstVisit(t *testing.T) {
	tests := []struct {
		graph     Graph[string]
		input     string
		want      []string
		wantError error
	}{
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
			want:  []string{"a", "b", "c"},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     "x",
			want:      []string{},
			wantError: &MissingVertexErr[string]{"x"},
		},
	}

	for i, test := range tests {
		got := make([]string, 0)
		err := test.graph.DepthFirstVisit(test.input, Outbound, func(v string) (stop bool) {
			got = append(got, v)
			return
		})

		if test.wantError != nil {
			if !cmp.Equal(err, test.wantError, cmp.AllowUnexported(MissingVertexErr[string]{})) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, test.want) {
				t.Errorf("%v: %v != %v", i, got, test.want)
			}
		}
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		graph     Graph[string]
		input     string
		want      []string
		wantError error
	}{
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
			want:  []string{"a", "b", "c"},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     "x",
			want:      []string{},
			wantError: &MissingVertexErr[string]{"x"},
		},
	}

	for i, test := range tests {
		got, err := test.graph.BreadthFirstSearch(test.input, Outbound)

		if test.wantError != nil {
			if !cmp.Equal(err, test.wantError, cmp.AllowUnexported(MissingVertexErr[string]{})) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, test.want) {
				t.Errorf("%v: %v != %v", i, got, test.want)
			}
		}
	}
}

func TestBreadthFirstVisit(t *testing.T) {
	tests := []struct {
		graph     Graph[string]
		input     string
		want      []string
		wantError error
	}{
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
			want:  []string{"a", "b", "c"},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     "x",
			want:      []string{},
			wantError: &MissingVertexErr[string]{"x"},
		},
	}

	for i, test := range tests {
		got := make([]string, 0)
		err := test.graph.BreadthFirstVisit(test.input, Outbound, func(v string) (stop bool) {
			got = append(got, v)
			return
		})

		if test.wantError != nil {
			if !cmp.Equal(err, test.wantError, cmp.AllowUnexported(MissingVertexErr[string]{})) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, test.want) {
				t.Errorf("%v: %v != %v", i, got, test.want)
			}
		}
	}
}
