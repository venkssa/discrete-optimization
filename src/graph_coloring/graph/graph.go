package graph

import (
	"io"

	"github.com/venkssa/discrete-optimization/src/common"
)

type G struct {
	NumOfVertices uint32
	VertexToEdges [][]uint32
}

func (g *G) AreNeighbors(vertex1 uint32, vertex2 uint32) bool {
	for _, neighborIdx := range g.Neighbors(vertex1) {
		if neighborIdx == vertex2 {
			return true
		}
	}
	return false
}

func (g *G) Neighbors(vertex uint32) []uint32 {
	return g.VertexToEdges[int(vertex)]
}

func NewGraph(rc io.ReadCloser) (*G, error) {
	graph := G{}

	err := common.Parse(rc, func(line common.LineNum, d1 uint32, d2 uint32) {
		if line == 1 {
			graph.NumOfVertices = d1
			graph.VertexToEdges = make([][]uint32, d1)
		} else {
			graph.VertexToEdges[d1] = append(graph.VertexToEdges[d1], uint32(d2))
			graph.VertexToEdges[d2] = append(graph.VertexToEdges[d2], uint32(d1))
		}
	})

	return &graph, err
}
