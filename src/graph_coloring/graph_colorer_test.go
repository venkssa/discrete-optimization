package graph_coloring

import (
	"testing"

	"github.com/venkssa/discrete-optimization/src/graph_coloring/testdata"
)

func TestColorGraph(t *testing.T) {
	graph := testdata.Gc_50_3_Graph()
	maxColor := color(6)
	result := ColorGraph(graph, maxColor)

	coloring := result.Coloring

	if uint32(len(coloring)) != graph.NumOfVertices {
		t.Errorf("Expected all nodes to be colored but was not %v", coloring)
	}

	for idx, color := range coloring {
		if color == UNSET {
			t.Errorf("Expected vertex %v to be colored", idx)
		} else if color > maxColor {
			t.Errorf("Expected vertex %v to be colored <= %v but was colored %v", idx, maxColor, color)
		}
	}

	reorders := make([]color, len(coloring))
	for idx, vertexIdx := range result.searchOrder {
		reorders[idx] = coloring[vertexIdx]
	}
	t.Log(reorders)
	t.Log(result.searchOrder)
	for k, v := range result.Stats {
		t.Log(k, v)
	}
}
