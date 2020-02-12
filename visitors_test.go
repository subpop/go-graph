package graph

import (
	"reflect"
	"testing"
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

	for _, test := range tests {
		got, err := test.graph.DepthFirstSearch(test.input)

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

	for _, test := range tests {
		got := make([]interface{}, 0)
		err := test.graph.DepthFirstVisit(test.input, func(v interface{}) (stop bool) {
			got = append(got, v)
			return
		})

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

	for _, test := range tests {
		got, err := test.graph.BreadthFirstSearch(test.input)

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

	for _, test := range tests {
		got := make([]interface{}, 0)
		err := test.graph.BreadthFirstVisit(test.input, func(v interface{}) (stop bool) {
			got = append(got, v)
			return
		})

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
