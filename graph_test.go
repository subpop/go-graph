package graph

import (
	"reflect"
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := NewGraph(false)

	for i := 0; i < 3; i++ {
		v := NewVertex(nil)
		if err := g.AddVertex(v); err != nil {
			t.Fatal(err)
		}
	}

	if g.NumVertex() != 3 {
		t.Fatalf("%v != 3", g.NumVertex())
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(false)

	a := NewVertex(nil)
	b := NewVertex(nil)

	if err := g.AddVertex(a); err != nil {
		t.Fatal(err)
	}
	if err := g.AddVertex(b); err != nil {
		t.Fatal(err)
	}
	if err := g.AddEdge(&a, &b); err != nil {
		t.Fatal(err)
	}

	if g.NumVertex() != 2 {
		t.Fatalf("%v != 2", g.NumVertex())
	}
}

func TestRemoveEdge(t *testing.T) {
	g := NewGraph(false)

	a := NewVertex(nil)
	b := NewVertex(nil)

	if err := g.AddVertex(a); err != nil {
		t.Fatal(err)
	}
	if err := g.AddVertex(b); err != nil {
		t.Fatal(err)
	}
	if err := g.AddEdge(&a, &b); err != nil {
		t.Fatal(err)
	}

	var neighbors []*Vertex
	var err error

	neighbors, err = g.Neighbors(&a)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(neighbors, []*Vertex{&b}) {
		t.Fatalf("%v != %v", neighbors, []*Vertex{&b})
	}

	if err := g.RemoveEdge(&a, &b); err != nil {
		t.Fatal(err)
	}

	neighbors, err = g.Neighbors(&a)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(neighbors, []*Vertex{}) {
		t.Fatalf("%v != %v", neighbors, []*Vertex{})
	}
}

func TestNeighbors(t *testing.T) {
	g := NewGraph(false)

	a := NewVertex(nil)
	b := NewVertex(nil)

	if err := g.AddVertex(a); err != nil {
		t.Fatal(err)
	}
	if err := g.AddVertex(b); err != nil {
		t.Fatal(err)
	}
	if err := g.AddEdge(&a, &b); err != nil {
		t.Fatal(err)
	}

	if g.NumVertex() != 2 {
		t.Fatalf("%v != 2", g.NumVertex())
	}

	neighbors, err := g.Neighbors(&a)
	if err != nil {
		t.Fatal(err)
	}

	if len(neighbors) != 1 {
		t.Fatalf("%v != 1", len(neighbors))
	}
	if !reflect.DeepEqual(*neighbors[0], b) {
		t.Fatalf("%+v != %+v", *neighbors[0], b)
	}
}
