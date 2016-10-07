package cliques

import (
	"testing"
	"graph_coloring/testdata"
	"fmt"
	"graph_coloring/graph"
)

func TestCompareRunTimesFor_DifferentFindAllMaximalCliques(t *testing.T) {
	graphNames := []testdata.GraphName{
		testdata.GC_70_1,
		testdata.GC_70_3,
		testdata.GC_70_5,
		testdata.GC_70_7,
		//testdata.GC_70_9, // Too slow for BK
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

	tomitaCliqueFinder := TomitaAlgo()
	bkCliqueFinder := BronKerbosch()
	parallelBkCliqueFinder := ParallelBKAlgo()
	parallelTomitaCliqueFinder := ParallelTomita()

	for _, graphName := range graphNames {
		g := testdata.Graph(graphName)
		t.Run(fmt.Sprint("ParallelTomita_", graphName),
			testFindAllMaximalCliques(g, parallelTomitaCliqueFinder))
		t.Run(fmt.Sprint("ParallelBK_", graphName), testFindAllMaximalCliques(g, parallelBkCliqueFinder))
		t.Run(fmt.Sprint("Tomita_", graphName), testFindAllMaximalCliques(g, tomitaCliqueFinder))
		t.Run(fmt.Sprint("BK_", graphName), testFindAllMaximalCliques(g, bkCliqueFinder))
	}
}

func testFindAllMaximalCliques(g *graph.G, algo MaximalCliqueFinder) func(*testing.T) {
	return func(t *testing.T) {
		c := algo.FindAllMaximalCliques(g)
		t.Logf("Number of cliques %d", len(c.Cliques))
	}
}

func BenchmarkTomita_FindAllMaximalCliques(b *testing.B) {
	benchmarkFindAllMaximalCliques(b, TomitaAlgo())
}

func BenchmarkFindAllMaximalCliques(b *testing.B) {
	benchmarkFindAllMaximalCliques(b, BronKerbosch())
}

func benchmarkFindAllMaximalCliques(b *testing.B, algo MaximalCliqueFinder) {
	g := testdata.Gc_70_7_Graph()
	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		algo.FindAllMaximalCliques(g)
	}
}
