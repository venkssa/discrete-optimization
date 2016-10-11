package cliques

import (
	"testing"
	"graph_coloring/testdata"
	"graph_coloring/graph"
)

func TestParallelTomita_FindAllMaximalCliques(t *testing.T) {
	g := testdata.Graph(testdata.GC_50_1)
	expectedCliques :=  BronKerbosch().FindAllMaximalCliques(g).Cliques
	parallelBKCliques := ParallelTomita().FindAllMaximalCliques(g).Cliques

	verifyCliques(t, parallelBKCliques, expectedCliques)
}

func BenchmarkParallelTomita(b *testing.B) {
	parallelTomitaCliqueFinder := ParallelTomita()
	for i := 0; i < b.N; i++ {
		for _, testGraph := range testGraphs() {
			cliques := parallelTomitaCliqueFinder.FindAllMaximalCliques(testGraph)
			b.Log(len(cliques.Cliques))
		}
	}
}

func testGraphs() []*graph.G {
	gs := []*graph.G{}
	graphNames := []testdata.GraphName{
		testdata.GC_70_1,
		testdata.GC_70_3,
		testdata.GC_70_5,
		testdata.GC_70_7,
		testdata.GC_70_9,
		testdata.GC_100_1,
		testdata.GC_100_3,
		testdata.GC_100_5,
		testdata.GC_100_7,
		//testdata.GC_100_9, // Too slow for Tomita and BK
		testdata.GC_250_1,
		testdata.GC_250_3,
		testdata.GC_250_5,
		//testdata.GC_250_7, // Too slow and too much memory
		//testdata.GC_250_9, // Did not even try to run
		testdata.GC_500_1,
		testdata.GC_500_3,
		//testdata.GC_500_5, // Too slow and too much memory
		//testdata.GC_500_7, // Did not even try to run
		//testdata.GC_500_9, // Did not even try to run
		testdata.GC_1000_1,
		//testdata.GC_1000_3, // Too slow and too much memory
	}
	for _, graphName := range graphNames {
		gs = append(gs, testdata.Graph(graphName))
	}

	return gs
}

func tt(count uint32, a *[]string) {
	if count == 0 {
		return
	}
	*a = append(*a, "hello")
	tt(count - 1, a)
}

func TestParallelTomita(t *testing.T) {
	a := []string{}
	tt(2, &a)
	t.Log(a)

	t.Log(a[len(a) - 1])
	t.Log(a[0:len(a) - 1])
}

