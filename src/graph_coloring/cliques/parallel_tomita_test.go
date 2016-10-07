package cliques

import (
	"testing"
	"graph_coloring/testdata"
)

func TestParallelTomita_FindAllMaximalCliques(t *testing.T) {
	g := testdata.Graph(testdata.GC_50_1)
	expectedCliques :=  BronKerbosch().FindAllMaximalCliques(g).Cliques
	parallelBKCliques := ParallelTomita().FindAllMaximalCliques(g).Cliques

	verifyCliques(t, parallelBKCliques, expectedCliques)
}

func BenchmarkParallelTomita(b *testing.B) {
	g := testdata.Graph(testdata.GC_250_5)

	for i := 0; i < b.N; i++ {
		cliques := ParallelTomita().FindAllMaximalCliques(g).Cliques
		b.Log(len(cliques))
	}
}
