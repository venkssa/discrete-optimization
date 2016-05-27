package graph_coloring

import (
	"common"
	"io"
)

type Graph struct {
	NumOfVertices  uint32
	VertextToEdges [][]uint32
}

func (g *Graph) Neighbors(vertex uint32) []uint32 {
	return g.VertextToEdges[int(vertex)]
}

func NewGraph(rc io.ReadCloser) (*Graph, error) {
	graph := Graph{}

	err := common.Parse(rc, func(line common.LineNum, d1 uint32, d2 uint32) {
		if line == 1 {
			graph.NumOfVertices = d1
			graph.VertextToEdges = make([][]uint32, d1)
		} else {
			graph.VertextToEdges[d1] = append(graph.VertextToEdges[d1], uint32(d2))
			graph.VertextToEdges[d2] = append(graph.VertextToEdges[d2], uint32(d1))
		}
	})

	return &graph, err
}
