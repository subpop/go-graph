package graph

import (
	"github.com/google/uuid"
)

// A Vertex represents a node in a graph.
type Vertex struct {
	id    uuid.UUID
	value interface{}
}

// NewVertex creates a new Vertex.
func NewVertex(value interface{}) Vertex {
	return Vertex{
		id:    uuid.New(),
		value: value,
	}
}

// String implements the Stringer interface.
func (v Vertex) String() string {
	return v.id.String()
}
