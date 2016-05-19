package graph_coloring

import (
	"io/ioutil"
	"strings"
)

func gc_4_1_Graph() *Graph {
	graphStr := `4 3
		     0 1
		     1 2
		     1 3`
	graph, err := NewGraph(ioutil.NopCloser(strings.NewReader(graphStr)))
	if err != nil {
		panic(err)
	}
	return graph
}
