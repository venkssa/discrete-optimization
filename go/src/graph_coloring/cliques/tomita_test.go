package cliques

import (
	"testing"
	"graph_coloring/test_data"
	"graph_coloring/graph"
	"reflect"
)

func TestTomitaAlgo_FindAllMaximalCliques(t *testing.T) {
	tests := []struct {
		graph           *graph.G
		expectedResults []Clique
	}{
		{
			graph:           test_data.Gc_4_1_Graph(),
			expectedResults: []Clique{{0, 1}, {1, 2}, {1, 3}},
		},
		{
			graph:           test_data.Gc_5_0_Graph(),
			expectedResults: []Clique{{0, 1, 2}, {0, 2, 3}, {0, 3, 4}},
		},
		{
			graph: test_data.MustMakeGraph(`4 6
			0 1
			0 2
			0 3
			1 2
			1 3
			2 3`),
			expectedResults: []Clique{{0, 1, 2, 3}},
		},
	}

	tomita := TomitaAlgo{}

	for _, test := range tests {
		cliques := tomita.FindAllMaximalCliques(test.graph)
		if !reflect.DeepEqual(cliques.Cliques, test.expectedResults) {
			t.Errorf("Expected %v but found %v", test.expectedResults, cliques.Cliques)
		}
	}
}

func TestNumberOfNeighbors(t *testing.T) {
	graph := test_data.MustMakeGraph(`4 5
	0 1
	0 2
	0 3
	1 3
	2 3`)

	actualNeighbors := neighborsBitSet(graph)

	expectedNeighbors := []*BitSet{
		stringToBitSet("0111"),
		stringToBitSet("1001"),
		stringToBitSet("1001"),
		stringToBitSet("1110"),
	}

	if !reflect.DeepEqual(actualNeighbors, expectedNeighbors) {
		t.Errorf("Expected %v as neighbors but found %v", expectedNeighbors, actualNeighbors)
	}
}

func stringToBitSet(str string) *BitSet {
	bs := NewBitSet(uint32(len(str)))
	for idx, r := range str {
		if r == '1' {
			bs.Set(uint32(idx))
		}
	}
	return bs
}

func BenchmarkTomita_FindAllMaximalCliques(b *testing.B) {
	graph := test_data.Gc_70_7_Graph()
	b.ResetTimer()
	b.ReportAllocs()

	tomitaAlgo := TomitaAlgo{}
	for idx := 0; idx < b.N; idx++ {
		tomitaAlgo.FindAllMaximalCliques(graph)
	}
}
