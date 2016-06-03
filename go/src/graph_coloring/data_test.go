package graph_coloring

import (
	"io/ioutil"
	"strings"
)

func gc_4_1_Graph() *Graph {
	return mustMakeGraph(`4 3
		     0 1
		     1 2
		     1 3`)
}

func gc_5_0_Graph() *Graph {
	return mustMakeGraph(`5 7
			0 1
			0 2
			0 3
			0 4
			1 2
			2 3
			3 4`)
}

func gc_20_1_Graph() *Graph {
	return mustMakeGraph(`20 23
		0 16
		1 2
		1 6
		1 7
		1 8
		2 11
		2 16
		2 17
		3 14
		3 16
		3 17
		4 7
		4 13
		4 17
		5 6
		5 11
		6 18
		9 12
		10 13
		11 17
		13 15
		15 17
		16 19`)
}

func mustMakeGraph(graphStr string) *Graph {
	graph, err := NewGraph(ioutil.NopCloser(strings.NewReader(graphStr)))
	if err != nil {
		panic(err)
	}
	return graph
}
