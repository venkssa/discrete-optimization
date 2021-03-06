package graphcoloring

import "testing"

func TestAllDifferentGraph_IsFreeVertex(t *testing.T) {
	adg := &allDifferentGraph{
		vertexToPossibleColors: map[int32][]color{0: {1}},
		colorToVertexEdge:      []int32{-1, -1},
	}

	if adg.IsFreeVertex(0) != true {
		t.Errorf("Expected vertex 0 to be free but was not")
	}
}

func TestAllDifferentGraph_IsNotFreeVertex(t *testing.T) {
	adg := &allDifferentGraph{
		vertexToPossibleColors: map[int32][]color{0: {1}},
		colorToVertexEdge:      []int32{-1, 0},
	}

	if adg.IsFreeVertex(0) != false {
		t.Errorf("Expected vertex 0 to be set but was free")
	}
}

func TestAlternatingPath_FromFreeVertex(t *testing.T) {
	tests := []struct {
		graph        *allDifferentGraph
		freeVertex   uint32
		expectedPath []int32
	}{
		{
			oneHopGraph(),
			1,
			[]int32{1, 3, 4, 4},
		},
		{
			twoHopGraph(),
			0,
			[]int32{0, 2, 3, 4, 4, 5},
		},
	}

	for _, test := range tests {

		actualPath := test.graph.findAlternatingPath(test.freeVertex)

		if len(test.expectedPath) != len(actualPath) {
			t.Errorf("Expected to have a valid alternating path but did not")
		}

		for idx, expectedNode := range test.expectedPath {
			if idx >= len(actualPath) {
				t.Errorf("Expected %v but did not get any", expectedNode)
				continue
			}
			if actualNode := actualPath[idx]; expectedNode != actualNode {
				t.Errorf("Expected %v node but got %v", expectedNode, actualNode)
			}
		}
	}
}

func TestNoAlternatingPath(t *testing.T) {
	dg := allDifferentGraph{
		vertexToPossibleColors: map[int32][]color{0: {1, 2}},
		colorToVertexEdge:      []int32{-1, 1, 2},
		vk:                     NewVisitKeeper(3),
	}

	if len(dg.findAlternatingPath(0)) != 0 {
		t.Errorf("Expected to no valid alternating path but found one.")
	}
}

func TestMaximumMatching(t *testing.T) {
	dg := sample()

	if maxMatching := dg.MaximumMatching(); maxMatching != int32(len(dg.vertexToPossibleColors)) {
		t.Error(dg, maxMatching)
	}
	t.Log(dg.MaximumMatching())
}

func BenchmarkAllDifferentGraph_MaximumMatching(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		oneHopGraph().MaximumMatching()
	}
}

func oneHopGraph() *allDifferentGraph {
	return &allDifferentGraph{
		map[int32][]color{
			0: {1, 2},
			1: {2, 3},
			2: {3},
			3: {4},
			4: {4, 5, 6},
			5: {7},
		},
		[]int32{
			-1,
			2,
			3,
			4,
			-1,
			-1,
			5,
			-1,
		},
		NewVisitKeeper(6),
	}
}

func sample() *allDifferentGraph {
	return &allDifferentGraph{
		map[int32][]color{
			2:  {1, 2, 3},
			11: {1, 2, 3},
			17: {1, 2, 3},
		},
		[]int32{
			-1,
			-1,
			-1,
			-1,
		},
		NewVisitKeeper(3),
	}
}

func twoHopGraph() *allDifferentGraph {
	return &allDifferentGraph{
		map[int32][]color{
			0: {1, 2},
			1: {2},
			2: {3},
			3: {4},
			4: {3, 5, 6},
			5: {7},
		},
		[]int32{
			-1,
			2,
			3,
			1,
			4,
			-1,
			5,
			-1,
		},
		NewVisitKeeper(6),
	}
}
