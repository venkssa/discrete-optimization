package cliques

import (
	"graph_coloring/graph"
)

func FindAllMaximalCliques(graph *graph.G) [][]uint32 {
	p := newBitSet(graph.NumOfVertices)
	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		p.Set(idx)
	}
	return bronKerboschMaximalClique(make([]uint32, 0, graph.NumOfVertices), p, newBitSet(graph.NumOfVertices),
		neighborsBitSet(graph), [][]uint32{})
}

func neighborsBitSet(graph *graph.G) []*bitSet {
	vertexToEdgeBitSet := make([]*bitSet, graph.NumOfVertices)

	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		bs := newBitSet(graph.NumOfVertices)

		for _, neighbor := range graph.VertexToEdges[int(idx)] {
			bs.Set(neighbor)
		}
		vertexToEdgeBitSet[idx] = bs
	}
	return vertexToEdgeBitSet
}
