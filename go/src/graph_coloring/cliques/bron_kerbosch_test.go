package cliques

import (
	"testing"
	"graph_coloring/test_data"
)

func BenchmarkFindAllMaximalCliques(b *testing.B) {
	graph := test_data.Gc_70_7_Graph()
	b.ResetTimer()
	b.ReportAllocs()
	for idx := 0; idx < b.N; idx++ {
		FindAllMaximalCliques(graph)
	}
}
