package graph_coloring

import "testing"

func TestColorGraph(t *testing.T) {
	graph := gc_4_1_Graph()
	maxColor := color(4)
	coloring := ColorGraph(graph, maxColor)

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
}
