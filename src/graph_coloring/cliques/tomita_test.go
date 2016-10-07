package cliques

import (
	"graph_coloring/graph"
	"graph_coloring/testdata"
	"testing"
)

func TestTomitaAlgo_FindAllMaximalCliques(t *testing.T) {
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

	tomita := TomitaAlgo()

	for _, test := range tests {
		cliques := tomita.FindAllMaximalCliques(test.graph)
		verifyCliques(t, cliques.Cliques, test.expectedResults)
	}
}

func TestFindPivot(t *testing.T) {
	tests := []struct {
		candidate      *BitSet
		finished       *BitSet
		expectedMaxIdx uint32
	}{
		{
			candidate:      stringToBitSet("1111"),
			finished:       stringToBitSet("0000"),
			expectedMaxIdx: uint32(0),
		},
		{
			candidate:      stringToBitSet("0000"),
			finished:       stringToBitSet("1111"),
			expectedMaxIdx: uint32(3),
		},
		{
			candidate:      stringToBitSet("0010"),
			finished:       stringToBitSet("0100"),
			expectedMaxIdx: uint32(1),
		},
	}

	neigbhors := []*BitSet{
		stringToBitSet("0111"),
		stringToBitSet("1010"),
		stringToBitSet("1100"),
		stringToBitSet("1000"),
	}

	pf := newPivotFinder(neigbhors)

	for _, test := range tests {
		actualMaxIdx := pf.find(test.candidate, test.finished)

		if actualMaxIdx != test.expectedMaxIdx {
			t.Errorf("Expected %v as pivot for candidate %v, finished %v, neighbors %v but was %v",
				test.expectedMaxIdx, test.candidate, test.finished, neigbhors, actualMaxIdx)
		}
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
