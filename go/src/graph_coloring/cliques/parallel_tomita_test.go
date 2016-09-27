package cliques

import (
	"testing"
	"graph_coloring/testdata"
)

func TestParallelTomita_FindAllMaximalCliques(t *testing.T) {
	g := testdata.Graph(testdata.GC_20_3)
	expectedCliques :=  BronKerbosch().FindAllMaximalCliques(g).Cliques
	parallelBKCliques := ParallelTomita().FindAllMaximalCliques(g).Cliques

	verifyCliques(t, parallelBKCliques, expectedCliques)
}
