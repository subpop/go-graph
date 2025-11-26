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
