package graph

import (
	"fmt"

	"github.com/subpop/go-adt"
)

// Neighborhood returns a slice of vertices adjacent to v, based on the provided
// parameters.
// - order: the number of "hops" from the vertex to include.
// - minimumDistance: the minimum number of "hops" from the vertex before
// collecting.
// If the graph is undirected, d is ignored. If the graph does not contain
// vertex v, it returns MissingVertexErr.
func (g *Graph) Neighborhood(v interface{}, order uint, minimumDistance uint, d Direction) ([]interface{}, error) {
	if order < minimumDistance {
		return nil, InvalidArgumentErr{fmt.Sprintf("%v !< %v", order, minimumDistance), "order must be greater than or equal to minimumDistance"}
	}

	result := make(set)
	distances := make(map[interface{}]uint)
	queue := adt.NewQueue()

	// insert and mark the initial vertex
	_ = queue.Enqueue(v)
	distances[v] = 0

	for {
		if queue.Len() == 0 {
			break
		}

		current := queue.Dequeue()
		distance := distances[current]

		// vertex current is further from the origin than the desired order, so
		// the vertex can be ignored.
		if distance > order {
			continue
		}

		// vertex current is within the minimumDistance threshold, so the vertex
		// is included.
		if distance >= minimumDistance {
			result[current] = true
		}

		// enqueue the neighbors of the current vertex and record their
		// distances.
		neighbors, err := g.Neighbors(current, d)
		if err != nil {
			return nil, err
		}

		for _, n := range neighbors {
			if _, has := distances[n]; !has {
				distances[n] = distance + 1
				queue.Enqueue(n)
			}
		}
	}

	keys := make([]interface{}, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}

	return keys, nil
}
