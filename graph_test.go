package graph

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

func TestAddVertices(t *testing.T) {
	tests := []struct {
		graph     Graph
		input     []interface{}
		want      Graph
		wantError error
	}{
		{
			graph: NewGraph(false),
			input: []interface{}{"a", "b"},
			want: Graph{
				vertices: set{
					"a": true,
					"b": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{},
					"b": edgeMap{},
				},
			},
		},
		{
			graph: Graph{
				vertices: set{
					"a": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{},
				},
			},
			input:     []interface{}{"a"},
			wantError: &DuplicateVertexErr{"a"},
		},
	}

	for _, test := range tests {
		err := test.graph.AddVertices(test.input...)

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
		graph     Graph
		input     struct{ a, b string }
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
					"a": edgeMap{},
					"b": edgeMap{},
				},
			},
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
			input:     struct{ a, b string }{"a", "b"},
			wantError: &DuplicateEdgeErr{"a", "b"},
		},
	}

	for _, test := range tests {
		err := test.graph.AddEdge(test.input.a, test.input.b, 0)

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

func TestRemoveEdge(t *testing.T) {
	tests := []struct {
		graph     Graph
		input     struct{ a, b string }
		want      Graph
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
			input:     struct{ a, b string }{"d", "b"},
			wantError: &MissingVertexErr{"d"},
		},
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
			input:     struct{ a, b string }{"b", "d"},
			wantError: &MissingVertexErr{"d"},
		},
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
			input:     struct{ a, b string }{"a", "c"},
			wantError: &MissingEdgeErr{"a", "c"},
		},
	}

	for _, test := range tests {
		err := test.graph.RemoveEdge(test.input.a, test.input.b)

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

func TestNeighbors(t *testing.T) {
	tests := []struct {
		graph     Graph
		input     string
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
		{
			graph: Graph{
				vertices: set{
					"a": true,
				},
				adjacencyMap: adjacencyMap{
					"a": edgeMap{},
				},
			},
			input:     "b",
			wantError: &MissingVertexErr{"b"},
		},
	}

	for _, test := range tests {
		got, err := test.graph.Neighbors(test.input)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v != %v", err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, test.want, cmp.Options{cmpopts.SortSlices(func(a, b string) bool { return a < b })}) {
				t.Errorf("%+v != %+v", got, test.want)
			}
		}
	}
}
