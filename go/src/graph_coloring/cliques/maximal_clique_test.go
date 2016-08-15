package cliques

import (
	"sort"
	"testing"
	"graph_coloring/test_data"
	"reflect"
)

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

func verifyCliques(t *testing.T, actualResults []Clique, expectedResults []Clique) {
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

func verifyClique(actualResult Clique, expectedResult Clique) bool {
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
