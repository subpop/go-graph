package graph

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func UtilityGraph() Graph[string] {
	return Graph[string]{
		vertices: set[string]{"a": true, "b": true, "c": true, "x": true, "y": true, "z": true},
		adjacencyMap: adjacencyMap[string]{
			"a": struct {
					Explicit edgeMap[string]
					Implicit edgeMap[string]
			}{
				Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0},
				Implicit: edgeMap[string]{},
			},
			"b": struct {
					Explicit edgeMap[string]
					Implicit edgeMap[string]
			}{
				Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0},
				Implicit: edgeMap[string]{},
			},
			"c": struct {
					Explicit edgeMap[string]
					Implicit edgeMap[string]
			}{
				Explicit: edgeMap[string]{"x": 0, "y": 0, "z": 0},
				Implicit: edgeMap[string]{},
			},
			"x": struct {
					Explicit edgeMap[string]
					Implicit edgeMap[string]
			}{
				Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0},
			},
			"y": struct {
					Explicit edgeMap[string]
					Implicit edgeMap[string]
			}{
				Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0},
			},
			"z": struct {
					Explicit edgeMap[string]
					Implicit edgeMap[string]
			}{
				Explicit: edgeMap[string]{"a": 0, "b": 0, "c": 0},
			},
		},
	}
}

func TestAddVertex(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       string
		want        Graph[string]
		wantError   error
	}{
		{
			graph: Graph[string]{
				vertices:     set[string]{},
				adjacencyMap: adjacencyMap[string]{},
			},
			input: "a",
			want: Graph[string]{
				vertices: set[string]{
					"a": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
				},
				adjacencyMap: adjacencyMap[string]{},
			},
			input:     "a",
			wantError: &DuplicateVertexErr[string]{"a"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.graph.AddVertex(test.input)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(test.graph, test.want, cmp.AllowUnexported(Graph[string]{})) {
					t.Errorf("%+v != %+v", test.graph, test.want)
				}
			}
		})
	}
}

func TestRemoveVertex(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       string
		want        Graph[string]
		wantError   error
	}{
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
			want: Graph[string]{
				vertices: set[string]{
					"b": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"b": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     "a",
			wantError: &MissingVertexErr[string]{"a"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.graph.RemoveVertex(test.input)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%v != %v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(test.graph, test.want, cmp.AllowUnexported(Graph[string]{})) {
					t.Errorf("%+v != %+v", test.graph, test.want)
				}
			}
		})
	}
}

func TestAddVertices(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       []string
		want        Graph[string]
		wantError   error
	}{
		{
			graph: Graph[string]{
				vertices:     set[string]{},
				adjacencyMap: adjacencyMap[string]{},
			},
			input: []string{"a", "b"},
			want: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     []string{"a"},
			wantError: &DuplicateVertexErr[string]{"a"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.graph.AddVertices(test.input...)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%v != %v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(test.graph, test.want, cmp.AllowUnexported(Graph[string]{})) {
					t.Errorf("%+v != %+v", test.graph, test.want)
				}
			}
		})
	}
}

func TestAddEdge(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       struct{ a, b string }
		want        Graph[string]
		wantError   error
	}{
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: struct{ a, b string }{"a", "b"},
			want: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
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
							"a": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
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
							"a": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     struct{ a, b string }{"a", "b"},
			wantError: &DuplicateEdgeErr[string]{"a", "b"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.graph.AddEdge(test.input.a, test.input.b, 0)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%v != %v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(test.graph, test.want, cmp.AllowUnexported(Graph[string]{})) {
					t.Errorf("%+v != %+v", test.graph, test.want)
				}
			}
		})
	}
}

func TestRemoveEdge(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       struct{ a, b string }
		want        Graph[string]
		wantError   error
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
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: struct{ a, b string }{"a", "b"},
			want: Graph[string]{
				vertices: set[string]{
					"a": true,
					"b": true,
					"c": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
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
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     struct{ a, b string }{"d", "b"},
			wantError: &MissingVertexErr[string]{"d"},
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
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     struct{ a, b string }{"b", "d"},
			wantError: &MissingVertexErr[string]{"d"},
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
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     struct{ a, b string }{"a", "c"},
			wantError: &MissingEdgeErr[string]{"a", "c"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.graph.RemoveEdge(test.input.a, test.input.b)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%v != %v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(test.graph, test.want, cmp.AllowUnexported(Graph[string]{})) {
					t.Errorf("%+v != %+v", test.graph, test.want)
				}
			}
		})
	}
}

func TestNeighbors(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       string
		want        []string
		wantError   error
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
							"a": 0,
							"c": 0,
						},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{
							"b": 0,
						},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "b",
			want:  []string{"a", "c"},
		},
		{
			graph: Graph[string]{
				vertices: set[string]{
					"a": true,
				},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     "b",
			wantError: &MissingVertexErr[string]{"b"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.graph.Neighbors(test.input, NoDirection)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%v != %v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(got, test.want, cmp.AllowUnexported(Graph[string]{}), cmpopts.SortSlices(func(x, y string) bool {
					return fmt.Sprintf("%v", x) < fmt.Sprintf("%v", y)
				})) {
					t.Errorf("%+v != %+v", got, test.want)
				}
			}
		})
	}
}
