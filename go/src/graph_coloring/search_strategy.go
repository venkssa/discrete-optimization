package graph_coloring

import (
	"math"
	"sort"
)

type Cliques struct {
	Cliques                [][]uint32
	NumOfVertices          uint32
	MinCliqueLen           uint32
	MaxCliqueLen           uint32

	vertexFreqPerCliqueLen []map[uint32]uint32
}

func NewCliques(cliques [][]uint32, numOfVertices uint32) *Cliques {
	minCliqueLen := uint32(math.MaxUint32)
	maxCliqueLen := uint32(0)

	for _, clique := range cliques {
		length := uint32(len(clique))
		if minCliqueLen > length {
			minCliqueLen = length
		}
		if maxCliqueLen < length {
			maxCliqueLen = length
		}
	}

	return &Cliques{
		Cliques:       cliques,
		NumOfVertices: numOfVertices,
		MinCliqueLen:  minCliqueLen,
		MaxCliqueLen:  maxCliqueLen,

		vertexFreqPerCliqueLen: computeOccurrencesForAllVertices(numOfVertices, cliques),
	}
}

func (c *Cliques) Occurrences(vertexIdx uint32, cliqueLen uint32) uint32 {
	return c.vertexFreqPerCliqueLen[vertexIdx][cliqueLen]
}

func computeOccurrencesForAllVertices(numOfVertices uint32, cliques [][]uint32) []map[uint32]uint32 {
	occurrences := make([]map[uint32]uint32, numOfVertices)

	for vertexIdx := uint32(0); vertexIdx < numOfVertices; vertexIdx++ {
		occurrences[vertexIdx] = map[uint32]uint32{}
	}

	for _, clique := range cliques {
		length := uint32(len(clique))
		for _, vertex := range clique {
			occurrences[vertex][length]++
		}
	}

	return occurrences
}

type orderedVertices struct {
	vertices []uint32
	*Cliques
}

func (ov *orderedVertices) Len() int {
	return len(ov.vertices)
}

func (ov *orderedVertices) Swap(i, j int) {
	ov.vertices[j], ov.vertices[i] = ov.vertices[i], ov.vertices[j]
}

func (ov *orderedVertices) Less(i, j int) bool {
	for currentLen := ov.MaxCliqueLen; currentLen >= ov.MinCliqueLen; currentLen-- {
		if ov.Occurrences(ov.vertices[i], currentLen) < ov.Occurrences(ov.vertices[j], currentLen) {
			return false
		}
	}

	return true
}

func OrderVerticesByCliques(c *Cliques) []uint32 {
	vertices := make([]uint32, c.NumOfVertices)
	for idx := range vertices {
		vertices[idx] = uint32(idx)
	}

	orderedVertices := &orderedVertices{vertices, c}

	sort.Sort(orderedVertices)

	return orderedVertices.vertices
}
