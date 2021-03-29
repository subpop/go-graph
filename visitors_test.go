package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDepthFirstSearch(t *testing.T) {
	tests := []struct {
		graph     Graph
		input     interface{}
		want      []interface{}
		wantError error
	}{
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input: "a",
			want:  []interface{}{"a", "b", "c"},
		},
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input:     "x",
			want:      []interface{}{},
			wantError: &MissingVertexErr{"x"},
		},
	}

	for i, test := range tests {
		got, err := test.graph.DepthFirstSearch(test.input, Outbound)

		if test.wantError != nil {
			if !cmp.Equal(err, test.wantError, cmp.AllowUnexported(MissingVertexErr{})) {
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
		graph     Graph
		input     interface{}
		want      []interface{}
		wantError error
	}{
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input: "a",
			want:  []interface{}{"a", "b", "c"},
		},
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input:     "x",
			want:      []interface{}{},
			wantError: &MissingVertexErr{"x"},
		},
	}

	for i, test := range tests {
		got := make([]interface{}, 0)
		err := test.graph.DepthFirstVisit(test.input, Outbound, func(v interface{}) (stop bool) {
			got = append(got, v)
			return
		})

		if test.wantError != nil {
			if !cmp.Equal(err, test.wantError, cmp.AllowUnexported(MissingVertexErr{})) {
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
		graph     Graph
		input     interface{}
		want      []interface{}
		wantError error
	}{
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input: "a",
			want:  []interface{}{"a", "b", "c"},
		},
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input:     "x",
			want:      []interface{}{},
			wantError: &MissingVertexErr{"x"},
		},
	}

	for i, test := range tests {
		got, err := test.graph.BreadthFirstSearch(test.input, Outbound)

		if test.wantError != nil {
			if !cmp.Equal(err, test.wantError, cmp.AllowUnexported(MissingVertexErr{})) {
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
		graph     Graph
		input     interface{}
		want      []interface{}
		wantError error
	}{
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input: "a",
			want:  []interface{}{"a", "b", "c"},
		},
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input:     "x",
			want:      []interface{}{},
			wantError: &MissingVertexErr{"x"},
		},
	}

	for i, test := range tests {
		got := make([]interface{}, 0)
		err := test.graph.BreadthFirstVisit(test.input, Outbound, func(v interface{}) (stop bool) {
			got = append(got, v)
			return
		})

		if test.wantError != nil {
			if !cmp.Equal(err, test.wantError, cmp.AllowUnexported(MissingVertexErr{})) {
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
