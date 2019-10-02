package graph

import (
	"reflect"
	"testing"
)

func TestAddVertex(t *testing.T) {
	tests := []struct {
		graph     Graph
		input     string
		want      Graph
		wantError error
	}{
		{
			graph: NewGraph(false),
			input: "a",
			want: Graph{
				vertices: set{
					"a": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{},
				},
			},
		},
		{
			graph: Graph{
				vertices: set{
					"a": true,
				},
				adjacencyMap: make(adjacencyMap),
			},
			input:     "a",
			want:      Graph{},
			wantError: &DuplicateVertexErr{"a"},
		},
	}

	for _, test := range tests {
		err := test.graph.AddVertex(test.input)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v != %v", err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.graph, test.want) {
				t.Errorf("%+v != %+v", test.graph, test.want)
			}
		}
	}
}

func TestRemoveVertex(t *testing.T) {
	tests := []struct {
		graph     Graph
		input     string
		want      Graph
		wantError error
	}{
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{
						"b": 0,
					},
					"b": edgeMap{
						"a": 0,
					},
				},
			},
			input: "a",
			want: Graph{
				vertices: set{
					"b": true,
				},
				adjacencyMap: adjacencyMap{
					"b": edgeMap{},
				},
			},
		},
		{
			graph: Graph{
				vertices: set{
					"b": true,
				},
				adjacencyMap: adjacencyMap{
					"b": edgeMap{},
				},
			},
			input:     "a",
			wantError: &MissingVertexErr{"a"},
		},
	}

	for _, test := range tests {
		err := test.graph.RemoveVertex(test.input)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v != %v", err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.graph, test.want) {
				t.Errorf("%+v != %+v", test.graph, test.want)
			}
		}
	}
}

func TestAddEdge(t *testing.T) {
	tests := []struct {
		input       struct{ a, b string }
		want        Graph
		shouldError bool
		wantError   error
	}{
		{
			input: struct{ a, b string }{"a", "b"},
			want: Graph{
				vertices: set{
					"a": true,
					"b": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{
						"b": 0,
					},
					"b": edgeMap{
						"a": 0,
					},
				},
			},
		},
	}

	for _, test := range tests {
		got := NewGraph(false)
		if err := got.AddVertex(test.input.a); err != nil {
			t.Fatal(err)
		}
		if err := got.AddVertex(test.input.b); err != nil {
			t.Fatal(err)
		}
		err := got.AddEdge(test.input.a, test.input.b, 0)

		if test.shouldError {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v != %v", err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%+v != %+v", got, test.want)
			}
		}
	}
}

func TestRemoveEdge(t *testing.T) {
	tests := []struct {
		graph       Graph
		input       struct{ a, b string }
		want        Graph
		shouldError bool
		wantError   error
	}{
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{
						"b": 0,
					},
					"b": edgeMap{
						"a": 0,
						"c": 0,
					},
					"c": edgeMap{
						"b": 0,
					},
				},
			},
			input: struct{ a, b string }{"a", "b"},
			want: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{},
					"b": edgeMap{
						"c": 0,
					},
					"c": edgeMap{
						"b": 0,
					},
				},
			},
		},
	}

	for _, test := range tests {
		err := test.graph.RemoveEdge(test.input.a, test.input.b)

		if test.shouldError {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v != %v", err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.graph, test.want) {
				t.Errorf("%+v != %+v", test.graph, test.want)
			}
		}
	}
}

func TestNeighbors(t *testing.T) {
	tests := []struct {
		graph       Graph
		input       string
		want        []interface{}
		shouldError bool
		wantError   error
	}{
		{
			graph: Graph{
				vertices: set{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{
						"b": 0,
					},
					"b": edgeMap{
						"a": 0,
						"c": 0,
					},
					"c": edgeMap{
						"b": 0,
					},
				},
			},
			input: "b",
			want:  []interface{}{"a", "c"},
		},
	}

	for _, test := range tests {
		got, err := test.graph.Neighbors(test.input)

		if test.shouldError {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v != %v", err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%+v != %+v", got, test.want)
			}
		}
	}
}
