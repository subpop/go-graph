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

func TestHasVertex(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       string
		want        bool
	}{
		{
			description: "vertex exists",
			graph: Graph[string]{
				vertices: set[string]{"a": true, "b": true},
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
			want:  true,
		},
		{
			description: "vertex does not exist",
			graph: Graph[string]{
				vertices: set[string]{"a": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "b",
			want:  false,
		},
		{
			description: "empty graph",
			graph: Graph[string]{
				vertices:     set[string]{},
				adjacencyMap: adjacencyMap[string]{},
			},
			input: "a",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := test.graph.HasVertex(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestHasEdge(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		from, to    string
		want        bool
	}{
		{
			description: "edge exists",
			graph: Graph[string]{
				vertices: set[string]{"a": true, "b": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			from: "a",
			to:   "b",
			want: true,
		},
		{
			description: "edge does not exist",
			graph: Graph[string]{
				vertices: set[string]{"a": true, "b": true},
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
			from: "a",
			to:   "b",
			want: false,
		},
		{
			description: "vertex does not exist",
			graph: Graph[string]{
				vertices: set[string]{"a": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			from: "a",
			to:   "c",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := test.graph.HasEdge(test.from, test.to)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestGetEdgeWeight(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		from, to    string
		want        float64
		wantError   error
	}{
		{
			description: "edge exists",
			graph: Graph[string]{
				vertices: set[string]{"a": true, "b": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 5.5},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			from: "a",
			to:   "b",
			want: 5.5,
		},
		{
			description: "from vertex missing",
			graph: Graph[string]{
				vertices: set[string]{"b": true},
				adjacencyMap: adjacencyMap[string]{
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			from:      "a",
			to:        "b",
			wantError: &MissingVertexErr[string]{"a"},
		},
		{
			description: "to vertex missing",
			graph: Graph[string]{
				vertices: set[string]{"a": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			from:      "a",
			to:        "b",
			wantError: &MissingVertexErr[string]{"b"},
		},
		{
			description: "edge does not exist",
			graph: Graph[string]{
				vertices: set[string]{"a": true, "b": true},
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
			from:      "a",
			to:        "b",
			wantError: &MissingEdgeErr[string]{"a", "b"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.graph.GetEdgeWeight(test.from, test.to)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("error: got %v, want %v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if got != test.want {
					t.Errorf("got %v, want %v", got, test.want)
				}
			}
		})
	}
}

func TestGetAllVertices(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		want        []string
	}{
		{
			description: "multiple vertices",
			graph: Graph[string]{
				vertices: set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			want: []string{"a", "b", "c"},
		},
		{
			description: "empty graph",
			graph: Graph[string]{
				vertices:     set[string]{},
				adjacencyMap: adjacencyMap[string]{},
			},
			want: []string{},
		},
		{
			description: "single vertex",
			graph: Graph[string]{
				vertices: set[string]{"a": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			want: []string{"a"},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := test.graph.GetAllVertices()
			if !cmp.Equal(got, test.want, cmpopts.SortSlices(func(x, y string) bool {
				return x < y
			})) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestGetAllEdges(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		want        []struct {
			From, To string
			Weight   float64
		}
	}{
		{
			description: "undirected graph with edges",
			graph: Graph[string]{
				isDirected: false,
				vertices:   set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0, "c": 2.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"a": 1.0},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"a": 2.0},
						Implicit: edgeMap[string]{},
					},
				},
			},
			want: []struct {
				From, To string
				Weight   float64
			}{
				{"a", "b", 1.0},
				{"a", "c", 2.0},
			},
		},
		{
			description: "directed graph with edges",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"c": 2.0},
						Implicit: edgeMap[string]{"a": 1.0},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{"b": 2.0},
					},
				},
			},
			want: []struct {
				From, To string
				Weight   float64
			}{
				{"a", "b", 1.0},
				{"b", "c", 2.0},
			},
		},
		{
			description: "empty graph",
			graph: Graph[string]{
				vertices:     set[string]{},
				adjacencyMap: adjacencyMap[string]{},
			},
			want: []struct {
				From, To string
				Weight   float64
			}{},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := test.graph.GetAllEdges()
			if !cmp.Equal(got, test.want, cmpopts.SortSlices(func(x, y struct {
				From, To string
				Weight   float64
			}) bool {
				if x.From != y.From {
					return x.From < y.From
				}
				return x.To < y.To
			})) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestNumEdges(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		want        int
	}{
		{
			description: "undirected graph with edges",
			graph: Graph[string]{
				isDirected: false,
				vertices:   set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0, "c": 2.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"a": 1.0},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"a": 2.0},
						Implicit: edgeMap[string]{},
					},
				},
			},
			want: 2,
		},
		{
			description: "directed graph with edges",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"c": 2.0},
						Implicit: edgeMap[string]{"a": 1.0},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{"b": 2.0},
					},
				},
			},
			want: 2,
		},
		{
			description: "empty graph",
			graph: Graph[string]{
				vertices:     set[string]{},
				adjacencyMap: adjacencyMap[string]{},
			},
			want: 0,
		},
		{
			description: "vertices without edges",
			graph: Graph[string]{
				vertices: set[string]{"a": true, "b": true},
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
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := test.graph.NumEdges()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestDegree(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       string
		want        int
		wantError   error
	}{
		{
			description: "vertex with multiple edges",
			graph: Graph[string]{
				isDirected: false,
				vertices:   set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0, "c": 2.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"a": 1.0},
						Implicit: edgeMap[string]{},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"a": 2.0},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
			want:  2,
		},
		{
			description: "isolated vertex",
			graph: Graph[string]{
				isDirected: false,
				vertices:   set[string]{"a": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
			want:  0,
		},
		{
			description: "missing vertex",
			graph: Graph[string]{
				isDirected: false,
				vertices:   set[string]{"a": true},
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
		{
			description: "directed graph",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input:     "a",
			wantError: DirectedGraphErr{},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.graph.Degree(test.input)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("error: got %v, want %v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if got != test.want {
					t.Errorf("got %v, want %v", got, test.want)
				}
			}
		})
	}
}

func TestInDegree(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       string
		want        int
		wantError   error
	}{
		{
			description: "vertex with incoming edges",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"c": 2.0},
						Implicit: edgeMap[string]{"a": 1.0, "c": 3.0},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 3.0},
						Implicit: edgeMap[string]{"b": 2.0},
					},
				},
			},
			input: "b",
			want:  2,
		},
		{
			description: "vertex with no incoming edges",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true, "b": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{"a": 1.0},
					},
				},
			},
			input: "a",
			want:  0,
		},
		{
			description: "missing vertex",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true},
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
		{
			description: "undirected graph",
			graph: Graph[string]{
				isDirected: false,
				vertices:   set[string]{"a": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.graph.InDegree(test.input)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("error: got %v, want %v", err, test.wantError)
				}
			} else if test.description == "undirected graph" {
				// For undirected graph, we expect UndirectedGraphErr
				if err == nil {
					t.Error("expected error for undirected graph, got nil")
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if got != test.want {
					t.Errorf("got %v, want %v", got, test.want)
				}
			}
		})
	}
}

func TestOutDegree(t *testing.T) {
	tests := []struct {
		description string
		graph       Graph[string]
		input       string
		want        int
		wantError   error
	}{
		{
			description: "vertex with outgoing edges",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true, "b": true, "c": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0, "c": 2.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{"a": 1.0},
					},
					"c": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{"a": 2.0},
					},
				},
			},
			input: "a",
			want:  2,
		},
		{
			description: "vertex with no outgoing edges",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true, "b": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{"b": 1.0},
						Implicit: edgeMap[string]{},
					},
					"b": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{"a": 1.0},
					},
				},
			},
			input: "b",
			want:  0,
		},
		{
			description: "missing vertex",
			graph: Graph[string]{
				isDirected: true,
				vertices:   set[string]{"a": true},
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
		{
			description: "undirected graph",
			graph: Graph[string]{
				isDirected: false,
				vertices:   set[string]{"a": true},
				adjacencyMap: adjacencyMap[string]{
					"a": struct{ Explicit, Implicit edgeMap[string] }{
						Explicit: edgeMap[string]{},
						Implicit: edgeMap[string]{},
					},
				},
			},
			input: "a",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.graph.OutDegree(test.input)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("error: got %v, want %v", err, test.wantError)
				}
			} else if test.description == "undirected graph" {
				// For undirected graph, we expect UndirectedGraphErr
				if err == nil {
					t.Error("expected error for undirected graph, got nil")
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if got != test.want {
					t.Errorf("got %v, want %v", got, test.want)
				}
			}
		})
	}
}
