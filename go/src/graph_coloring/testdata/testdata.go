package testdata

import (
	"fmt"
	"graph_coloring/graph"
	"io/ioutil"
	"os"
	"strings"
)

func Gc_4_1_Graph() *graph.G {
	return MustMakeGraph(`4 3
		     0 1
		     1 2
		     1 3`)
}

func Gc_5_0_Graph() *graph.G {
	return MustMakeGraph(`5 7
			0 1
			0 2
			0 3
			0 4
			1 2
			2 3
			3 4`)
}

func Gc_20_1_Graph() *graph.G {
	return MustMakeGraph(`20 23
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

func Gc_50_3_Graph() *graph.G {
	return mustMakeGraphFile("gc_50_3")
}

func Gc_70_7_Graph() *graph.G {
	return mustMakeGraphFile("gc_70_7")
}

func Gc_100_5_Graph() *graph.G {
	return mustMakeGraphFile("gc_100_5")
}

func Gc_1000_5_Graph() *graph.G {
	return mustMakeGraphFile("gc_1000_5")
}

func mustMakeGraphFile(path string) *graph.G {
	file, err := os.Open(fmt.Sprintf("%s/src/graph_coloring/testdata/%s", os.Getenv("GOPATH"), path))
	if err != nil {
		panic(err)
	}
	g, err := graph.NewGraph(file)
	if err != nil {
		panic(err)
	}
	return g
}

func MustMakeGraph(graphStr string) *graph.G {
	g, err := graph.NewGraph(ioutil.NopCloser(strings.NewReader(graphStr)))
	if err != nil {
		panic(err)
	}
	return g
}
