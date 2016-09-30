package graph_coloring

import (
	"reflect"
	"testing"
	"graph_coloring/cliques"
)

func TestCliques_VertexCountPerCliqueLen(t *testing.T) {
	tests := []struct {
		cliques  *Cliques
		expected []map[uint32]uint32
	}{
		{
			cliques: NewCliques(
				cliques.Cliques{[]cliques.Clique{{0, 1, 2}, {0, 2, 3}, {3, 4}}, 5},
			),
			expected: []map[uint32]uint32{
				{3: 2},
				{3: 1},
				{3: 2},
				{2: 1, 3: 1},
				{2: 1},
			},
		},
		{
			cliques: NewCliques(
				cliques.Cliques{[]cliques.Clique{}, 3},
			),
			expected: []map[uint32]uint32{{}, {}, {}},
		},
	}

	for _, test := range tests {
		for vertexIdx := uint32(0); vertexIdx < uint32(len(test.expected)); vertexIdx++ {
			for cliqueLen, occurrences := range test.expected[vertexIdx] {
				expectedOccurrences := test.cliques.Occurrences(vertexIdx, cliqueLen)
				if expectedOccurrences != occurrences {
					t.Errorf("Expected %d occurrences for vertex %d, cliqueLen %d, but was %d",
						expectedOccurrences, vertexIdx, cliqueLen, occurrences)
				}
			}

		}
	}
}

func TestOrderVertexForSearch(t *testing.T) {
	cliques := NewCliques(cliques.Cliques{[]cliques.Clique{{0, 1, 2}, {0, 2, 3}, {3, 4}}, 5})

	actualOrder := OrderVerticesByCliques(cliques)

	expectedOrder := []uint32{2, 0, 3, 1, 4}

	if !reflect.DeepEqual(actualOrder, expectedOrder) {
		t.Errorf("Expected %v but was %v", expectedOrder, actualOrder)
	}
}
