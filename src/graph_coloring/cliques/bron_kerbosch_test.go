package cliques

import (
	"testing"
	"graph_coloring/testdata"
	"graph_coloring/graph"
)

func TestFindAllMaximalCliques(t *testing.T) {
	tests := []struct {
		graph           *graph.G
		expectedResults []Clique
	}{
		{
			graph:           testdata.Gc_4_1_Graph(),
			expectedResults: []Clique{{0, 1}, {1, 2}, {1, 3}},
		},
		{
			graph:           testdata.Gc_5_0_Graph(),
			expectedResults: []Clique{{0, 1, 2}, {0, 2, 3}, {0, 3, 4}},
		},
		{
			graph: testdata.MustMakeGraph(`4 6
			0 1
			0 2
			0 3
			1 2
			1 3
			2 3`),
			expectedResults: []Clique{{0, 1, 2, 3}},
		},
	}

	for _, test := range tests {
		cliques := BronKerbosch().FindAllMaximalCliques(test.graph)

		verifyCliques(t, cliques.Cliques, test.expectedResults)

		if cliques.NumOfVertices != test.graph.NumOfVertices {
			t.Errorf("Expected number of vertices to be %v but was %v", test.graph.NumOfVertices,
				cliques.NumOfVertices)
		}
	}
}
