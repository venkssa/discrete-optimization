package cliques

import (
	"testing"
	"graph_coloring/test_data"
)

func TestT(t *testing.T) {
	graph := test_data.Gc_70_7_Graph()
	TomitaAlgo{}.FindAllMaximalCliques(graph)
}

func BenchmarkTomita_FindAllMaximalCliques(b *testing.B) {
	benchmarkFindAllMaximalCliques(b, TomitaAlgo{})
}

func BenchmarkFindAllMaximalCliques(b *testing.B) {
	benchmarkFindAllMaximalCliques(b, BronKerbosch())
}

func benchmarkFindAllMaximalCliques(b *testing.B, algo MaximalCliqueFinder) {
	graph := test_data.Gc_70_7_Graph()
	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		algo.FindAllMaximalCliques(graph)
	}
}
