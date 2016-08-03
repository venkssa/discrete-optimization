package cliques

import (
	"graph_coloring/graph"
)

func FindAllMaximalCliques(graph *graph.G) [][]uint32 {
	p := NewBitSet(graph.NumOfVertices)
	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		p.Set(idx)
	}
	return bronKerboschMaximalClique(make([]uint32, 0, graph.NumOfVertices), p, NewBitSet(graph.NumOfVertices),
		neighborsBitSet(graph), [][]uint32{})
}

func neighborsBitSet(graph *graph.G) []*BitSet {
	neighborsBitSet := make([]*BitSet, graph.NumOfVertices)

	for vertexIdx := uint32(0); vertexIdx < graph.NumOfVertices; vertexIdx++ {
		neighborsBitSet[vertexIdx] = NewBitSet(graph.NumOfVertices)
		for _, neighborIdx := range graph.Neighbors(vertexIdx) {
			neighborsBitSet[vertexIdx].Set(neighborIdx)
		}
	}

	return neighborsBitSet
}
