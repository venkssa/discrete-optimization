package cliques

import (
	"sort"
	"testing"
	"graph_coloring/test_data"
	"graph_coloring/graph"
)

func Test70Graph(t *testing.T) {
	cliques := FindAllMaximalCliques(test_data.Gc_70_7_Graph())

	stats := map[int]int{}

	for _, clique := range cliques {
		stats[len(clique)] += 1
	}

	t.Log(stats)
}

func TestFindAllMaximalCliques(t *testing.T) {
	tests := []struct {
		graph           *graph.G
		expectedResults [][]uint32
	}{
		{
			graph:           test_data.Gc_4_1_Graph(),
			expectedResults: [][]uint32{{0, 1}, {1, 2}, {1, 3}},
		},
		{
			graph:           test_data.Gc_5_0_Graph(),
			expectedResults: [][]uint32{{0, 1, 2}, {0, 2, 3}, {0, 3, 4}},
		},
		{
			graph: test_data.MustMakeGraph(`4 6
			0 1
			0 2
			0 3
			1 2
			1 3
			2 3`),
			expectedResults: [][]uint32{{0, 1, 2, 3}},
		},
	}

	for _, test := range tests {
		cliques := FindAllMaximalCliques(test.graph)

		verifyCliques(t, cliques, test.expectedResults)
	}
}

func verifyCliques(t *testing.T, actualResults [][]uint32, expectedResults [][]uint32) {
	for idx := 0; idx < len(expectedResults); idx++ {
		matches := false
		for jdx := 0; jdx < len(actualResults); jdx++ {
			if verifyClique(actualResults[jdx], expectedResults[idx]) {
				matches = true
			}
		}
		if !matches {
			t.Errorf("Expeced %v but did not find in %v", expectedResults[idx], actualResults)
		}
	}

	if len(actualResults) > len(expectedResults) {
		t.Errorf("Unexpected elements found %v", actualResults)
	} else if len(expectedResults) > len(actualResults) {
		t.Errorf("Expected elements %v but did not find any", expectedResults[len(actualResults):])
	}
}

func verifyClique(actualResult []uint32, expectedResult []uint32) bool {
	matches := true

	sortedActualResult := make(sortedResult, len(actualResult))

	for idx, value := range actualResult {
		sortedActualResult[idx] = value
	}

	sort.Sort(sortedActualResult)

	for idx := 0; idx < len(sortedActualResult) && idx < len(expectedResult); idx++ {
		if sortedActualResult[idx] != expectedResult[idx] {
			matches = false
		}
	}

	if len(actualResult) > len(expectedResult) {
		matches = false
	} else if len(expectedResult) > len(actualResult) {
		matches = false
	}

	return matches
}

type sortedResult []uint32

func (sr sortedResult) Len() int {
	return len(sr)
}

func (sr sortedResult) Swap(i, j int) {
	sr[j], sr[i] = sr[i], sr[j]
}

func (sr sortedResult) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func BenchmarkFindAllMaximalCliques(b *testing.B) {
	graph := test_data.Gc_70_7_Graph()
	b.ResetTimer()
	b.ReportAllocs()
	for idx := 0; idx < b.N; idx++ {
		FindAllMaximalCliques(graph)
	}
}

func TestBitSet_Set(t *testing.T) {
	bs := newBitSet(64)

	bs.Set(0)
	bs.Set(1)
	bs.Set(63)
	bs.UnSet(1)
	bs.UnSet(0)

	if bs.blocks[0] != 1<<(bitsPerWord-1) {
		t.Error(bs)
	}

	if bs.IsSet(63) != true {
		t.Error(bs)
	}

	if bs.IsSet(0) == true {
		t.Error(bs)
	}
}

func BenchmarkNeighborUsingMap(b *testing.B) {
	neighbors := []uint32{1, 3, 5, 49}
	p := map[uint32]bool{}
	for idx := uint32(0); idx < 50; idx++ {
		p[idx] = true
	}

	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		pCopy := map[uint32]bool{}
		for _, neighbor := range neighbors {
			if _, ok := p[neighbor]; ok {
				pCopy[neighbor] = true
			}
		}
	}
}

func BenchmarkBitSet(b *testing.B) {
	neighbors := newBitSet(50)
	neighbors.Set(1)
	neighbors.Set(3)
	neighbors.Set(5)
	neighbors.Set(49)

	p := &bitSet{blocks: []block{maxBlock}}

	result := newBitSetLike(p)
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		result.And(neighbors, p)
	}
}

func TestEq(t *testing.T) {
	neighbors := []uint32{1, 3, 5}
	p := map[uint32]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true}
	pCopy := map[uint32]bool{}
	for _, neighbor := range neighbors {
		if _, ok := p[neighbor]; ok {
			pCopy[neighbor] = true
		}
	}

	neighborsBS := newBitSet(1)
	neighborsBS.Set(1)
	neighborsBS.Set(3)
	neighborsBS.Set(5)
	pBS := &bitSet{blocks: []block{maxBlock}}

	pBSCopy := newBitSetLike(pBS)

	pBSCopy.And(neighborsBS, pBS)

	for idx := uint32(0); idx < 8; idx++ {
		if pBSCopy.IsSet(idx) != pCopy[idx] {
			t.Errorf("Expected %v for idx %d, but was %v", pCopy[idx], idx, pBSCopy.IsSet(idx))
		}
	}
}
