package cliques

import (
	"testing"

	"github.com/venkssa/discrete-optimization/src/graph_coloring/testdata"
)

func TestParallelBKAlgo_FindAllMaximalCliques(t *testing.T) {
	g := testdata.Graph(testdata.GC_20_3)
	expectedCliques := BronKerbosch().FindAllMaximalCliques(g).Cliques
	parallelBKCliques := ParallelBKAlgo().FindAllMaximalCliques(g).Cliques

	verifyCliques(t, parallelBKCliques, expectedCliques)
}
