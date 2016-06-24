package graph_coloring

import (
	"sort"
	"testing"
)

func TestFindAllMaximalCliques(t *testing.T) {
	tests := []struct {
		graph           *Graph
		expectedResults [][]uint32
	}{
		{
			graph:           gc_4_1_Graph(),
			expectedResults: [][]uint32{{0, 1}, {1, 2}, {1, 3}},
		},
		{
			graph:           gc_5_0_Graph(),
			expectedResults: [][]uint32{{0, 1, 2}, {0, 2, 3}, {0, 3, 4}},
		},
		{
			graph: mustMakeGraph(`4 6
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
