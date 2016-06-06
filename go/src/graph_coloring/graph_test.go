package graph_coloring

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	graph := gc_4_1_Graph()

	if graph.NumOfVertices != 4 {
		t.Errorf("Expected 4 vertices but got %d", graph.NumOfVertices)
	}

	var numOfEdges uint32
	for _, edge := range graph.VertexToEdges {
		numOfEdges += uint32(len(edge))
	}

	if (numOfEdges / 2) != 3 {
		t.Errorf("Expecetd 3 edges but got %d", numOfEdges)
	}

	expectedEdges := [][]uint32{
		{
			1,
		},
		{
			0,
			2,
			3,
		},
		{
			1,
		},
		{
			1,
		},
	}

	if len(expectedEdges) != len(graph.VertexToEdges) {
		t.Errorf("Expected %d vertex but got %d", len(expectedEdges), len(graph.VertexToEdges))
	}
}
