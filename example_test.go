package graph_test

import (
	"fmt"
	"log"
	"strconv"

	"github.com/subpop/go-graph"
)

type Package struct {
	Name     string
	Version  string
	Revision int
}

func (p Package) String() string {
	return p.Name + "-" + p.Version + "-" + strconv.FormatInt(int64(p.Revision), 10)
}

func ExampleGraph() {
	foo := Package{
		Name:     "foo",
		Version:  "1.4.2",
		Revision: 1,
	}

	libfoo := Package{
		Name:     "libfoo",
		Version:  "1.5",
		Revision: 2,
	}

	deps := graph.NewGraph[Package](true)
	if err := deps.AddEdge(foo, libfoo, 0); err != nil {
		log.Fatal(err)
	}

	fmt.Println(deps)
	//Output:
	// { (foo-1.4.2-1, libfoo-1.5-2) }
}

func ExampleGraph_Dijkstra() {
	g := graph.NewGraph[string](true)
	g.AddEdge("A", "B", 4.0)
	g.AddEdge("A", "C", 2.0)
	g.AddEdge("B", "C", 1.0)
	g.AddEdge("B", "D", 5.0)
	g.AddEdge("C", "D", 8.0)
	g.AddEdge("C", "E", 10.0)
	g.AddEdge("D", "E", 2.0)

	results, err := g.Dijkstra("A")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Shortest paths from A:\n")
	fmt.Printf("To D: distance=%.0f, path=%v\n", results["D"].Distance, results["D"].Path)
	fmt.Printf("To E: distance=%.0f, path=%v\n", results["E"].Distance, results["E"].Path)

	// Output:
	// Shortest paths from A:
	// To D: distance=9, path=[A B D]
	// To E: distance=11, path=[A B D E]
}
