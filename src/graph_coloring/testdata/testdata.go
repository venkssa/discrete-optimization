package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/venkssa/discrete-optimization/src/graph_coloring/graph"
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

type GraphName string

const (
	GC_1000_1 = GraphName("gc_1000_1")
	GC_1000_3 = GraphName("gc_1000_3")
	GC_1000_5 = GraphName("gc_1000_5")
	GC_1000_7 = GraphName("gc_1000_7")
	GC_1000_9 = GraphName("gc_1000_9")
	GC_100_1  = GraphName("gc_100_1")
	GC_100_3  = GraphName("gc_100_3")
	GC_100_5  = GraphName("gc_100_5")
	GC_100_7  = GraphName("gc_100_7")
	GC_100_9  = GraphName("gc_100_9")
	GC_20_1   = GraphName("gc_20_1")
	GC_20_3   = GraphName("gc_20_3")
	GC_20_5   = GraphName("gc_20_5")
	GC_20_7   = GraphName("gc_20_7")
	GC_20_9   = GraphName("gc_20_9")
	GC_250_1  = GraphName("gc_250_1")
	GC_250_3  = GraphName("gc_250_3")
	GC_250_5  = GraphName("gc_250_5")
	GC_250_7  = GraphName("gc_250_7")
	GC_250_9  = GraphName("gc_250_9")
	GC_4_1    = GraphName("gc_4_1")
	GC_500_1  = GraphName("gc_500_1")
	GC_500_3  = GraphName("gc_500_3")
	GC_500_5  = GraphName("gc_500_5")
	GC_500_7  = GraphName("gc_500_7")
	GC_500_9  = GraphName("gc_500_9")
	GC_50_1   = GraphName("gc_50_1")
	GC_50_3   = GraphName("gc_50_3")
	GC_50_5   = GraphName("gc_50_5")
	GC_50_7   = GraphName("gc_50_7")
	GC_50_9   = GraphName("gc_50_9")
	GC_5_0    = GraphName("gc_5_0")
	GC_70_1   = GraphName("gc_70_1")
	GC_70_3   = GraphName("gc_70_3")
	GC_70_5   = GraphName("gc_70_5")
	GC_70_7   = GraphName("gc_70_7")
	GC_70_9   = GraphName("gc_70_9")
)

func Graph(graphName GraphName) *graph.G {
	return mustMakeGraphFile(string(graphName))
}

func mustMakeGraphFile(path string) *graph.G {
	gopath, ok := os.LookupEnv("GOPATH")
	if !ok {
		gopath = os.Getenv("HOME") + "/go"
	}
	file, err := os.Open(fmt.Sprintf("%s/src/github.com/venkssa/discrete-optimization/src/graph_coloring/testdata/%s", gopath, path))
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
