package graph

import (
	"github.com/google/uuid"
)

type Vertex struct {
	id    uuid.UUID
	value interface{}
}

func NewVertex(value interface{}) Vertex {
	return Vertex{
		id:    uuid.New(),
		value: value,
	}
}

func (v Vertex) ID() string {
	return v.id.String()
}
