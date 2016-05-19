package graph_coloring

import (
	"common"
	"io"
)

type Vertex uint32

type Graph struct {
	NumOfVertices  uint32
	VertextToEdges [][]Vertex
}

func NewGraph(rc io.ReadCloser) (*Graph, error) {
	graph := Graph{}

	err := common.Parse(rc, func(line common.LineNum, d1 uint32, d2 uint32) {
		if line == 1 {
			graph.NumOfVertices = d1
			graph.VertextToEdges = make([][]Vertex, d1)
		} else {
			graph.VertextToEdges[d1] = append(graph.VertextToEdges[d1], Vertex(d2))
			graph.VertextToEdges[d2] = append(graph.VertextToEdges[d2], Vertex(d1))
		}
	})

	return &graph, err
}
