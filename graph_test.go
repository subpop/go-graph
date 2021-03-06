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
			graph: Graph{
				vertices:     set{},
				adjacencyMap: adjacencyMap{},
			},
			input: "a",
			want: Graph{
				vertices: set{
					"a": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
				},
			},
		},
		{
			graph: Graph{
				vertices: set{
					"a": true,
				},
				adjacencyMap: adjacencyMap{},
			},
			input:     "a",
			wantError: &DuplicateVertexErr{"a"},
		},
	}

	for i, test := range tests {
		err := test.graph.AddVertex(test.input)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.graph, test.want) {
				t.Errorf("%v: %+v != %+v", i, test.graph, test.want)
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
				},
			},
			input: "a",
			want: Graph{
				vertices: set{
					"b": true,
				},
				adjacencyMap: adjacencyMap{
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
				},
			},
		},
		{
			graph: Graph{
				vertices: set{
					"b": true,
				},
				adjacencyMap: adjacencyMap{
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
				},
			},
			input:     "a",
			wantError: &MissingVertexErr{"a"},
		},
	}

	for i, test := range tests {
		err := test.graph.RemoveVertex(test.input)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.graph, test.want) {
				t.Errorf("%v: %+v != %+v", i, test.graph, test.want)
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
			graph: Graph{
				vertices:     set{},
				adjacencyMap: adjacencyMap{},
			},
			input: []interface{}{"a", "b"},
			want: Graph{
				vertices: set{
					"a": true,
					"b": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
				},
			},
		},
		{
			graph: Graph{
				vertices: set{
					"a": true,
				},
				adjacencyMap: adjacencyMap{
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
				},
			},
			input:     []interface{}{"a"},
			wantError: &DuplicateVertexErr{"a"},
		},
	}

	for i, test := range tests {
		err := test.graph.AddVertices(test.input...)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.graph, test.want) {
				t.Errorf("%v: %+v != %+v", i, test.graph, test.want)
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
				},
			},
			input: struct{ a, b string }{"a", "b"},
			want: Graph{
				vertices: set{
					"a": true,
					"b": true,
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
							"a": 0,
						},
						Implicit: edgeMap{},
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"a": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input:     struct{ a, b string }{"a", "b"},
			wantError: &DuplicateEdgeErr{"a", "b"},
		},
	}

	for i, test := range tests {
		err := test.graph.AddEdge(test.input.a, test.input.b, 0)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.graph, test.want) {
				t.Errorf("%v: %+v != %+v", i, test.graph, test.want)
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap{},
					},
					"c": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"c": 0,
						},
						Implicit: edgeMap{},
					},
					"c": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap{},
					},
					"c": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap{},
					},
					"c": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap{},
					},
					"c": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
				},
			},
			input:     struct{ a, b string }{"a", "c"},
			wantError: &MissingEdgeErr{"a", "c"},
		},
	}

	for i, test := range tests {
		err := test.graph.RemoveEdge(test.input.a, test.input.b)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.graph, test.want) {
				t.Errorf("%v: %+v != %+v", i, test.graph, test.want)
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
					},
					"b": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap{},
					},
					"c": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{
							"b": 0,
						},
						Implicit: edgeMap{},
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
					"a": struct{ Explicit, Implicit edgeMap }{
						Explicit: edgeMap{},
						Implicit: edgeMap{},
					},
				},
			},
			input:     "b",
			wantError: &MissingVertexErr{"b"},
		},
	}

	for i, test := range tests {
		got, err := test.graph.Neighbors(test.input, NoDirection)

		if test.wantError != nil {
			if !reflect.DeepEqual(err, test.wantError) {
				t.Errorf("%v: %v != %v", i, err, test.wantError)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%v: %+v != %+v", i, got, test.want)
			}
		}
	}
}
