package graph_coloring

import "testing"

func TestNotEqual_IsFeasible(t *testing.T) {
	tests := []struct {
		vertexColors []color
	}{
		{make([]color, 4)},
		{[]color{1, UNSET, 1, 1}},
		{[]color{1, 2, 1, 1}},
	}

	for _, test := range tests {
		graph := gc_4_1_Graph()
		domainStore := &DomainStore{test.vertexColors, 0}

		constraint := NotEqual{uint32(1)}

		if constraint.IsFeasible(graph, domainStore) != true {
			t.Errorf("Expected All different Constraint %v to be feasible on domain %v",
				constraint, domainStore)
		}
	}
}

func TestNotEqual_IsNotFeasible(t *testing.T) {
	graph := gc_4_1_Graph()

	domainStore := &DomainStore{[]color{1, 1, 1, 1}, 0}

	constraint := NotEqual{uint32(1)}

	if constraint.IsFeasible(graph, domainStore) != false {
		t.Errorf("Expected constraint %v to be in-feasible on the domain %v", constraint, domainStore)
	}
}

func TestNotEqual_Prune(t *testing.T) {
	tests := []struct {
		vertexColors  []color
		expectedColor color
	}{
		{
			vertexColors:  []color{1, UNSET, 1, 1},
			expectedColor: 2,
		},
		{
			vertexColors:  []color{1, UNSET, 2, 3},
			expectedColor: 4,
		},
	}

	for _, test := range tests {
		graph := gc_4_1_Graph()
		domainStore := &DomainStore{test.vertexColors, 0}

		constraint := NotEqual{uint32(1)}

		if constraint.Prune(graph, domainStore) != true {
			t.Errorf("Expected constraint %v to prune on the domain %v", constraint, domainStore)
		}

		if domainStore.IsSet(constraint.vertex) != true {
			t.Errorf("Expected a color to be assigned for %v on the domain %v", constraint.vertex, domainStore)
		}

		if actualColor := domainStore.Color(constraint.vertex); actualColor != test.expectedColor {
			t.Errorf("Expected a color %v to be assigned for %v but was %v", test.expectedColor,
				constraint.vertex, actualColor)
		}
	}
}

func TestNotEqual_NothingToPrune(t *testing.T) {
	graph := gc_4_1_Graph()
	domainStore := &DomainStore{[]color{UNSET, UNSET, 1, 1}, 0}
	constraint := NotEqual{uint32(1)}

	if constraint.Prune(graph, domainStore) != false {
		t.Errorf("Expected constraint %v not to prune on the domain %v", constraint, domainStore)
	}
}

func BenchmarkNotEqual_IsFeasible_UnsetVertex(b *testing.B) {
	benchmarkIsFeasible(b,
		gc_4_1_Graph(),
		&DomainStore{[]color{1, UNSET, 1, 1}, 0},
		NotEqual{uint32(1)})
}

func BenchmarkNotEqual_IsFeasible_SetVertex(b *testing.B) {
	benchmarkIsFeasible(b,
		gc_4_1_Graph(),
		&DomainStore{[]color{1, 2, 1, 1}, 0},
		NotEqual{uint32(1)})
}

func BenchmarkNotEqual_Prune(b *testing.B) {
	graph := gc_4_1_Graph()
	domainStore := &DomainStore{[]color{1, UNSET, 1, 1}, 0}
	constraint := NotEqual{uint32(1)}

	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		constraint.Prune(graph, domainStore)
	}
}

func benchmarkIsFeasible(b *testing.B, graph *Graph, domainStore *DomainStore, constraint NotEqual) {
	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		constraint.IsFeasible(graph, domainStore)
	}
}

func TestFind3VerticesCompleteGraph(t *testing.T) {
	tests := []struct {
		graph          *Graph
		expectedResult [][3]uint32
	}{
		{
			graph: mustMakeGraph(`3 3
				0 1
				0 2
				1 2`),
			expectedResult: [][3]uint32{{0, 1, 2}},
		},
		{
			graph:          gc_4_1_Graph(),
			expectedResult: [][3]uint32{},
		},
		{
			graph:          gc_5_0_Graph(),
			expectedResult: [][3]uint32{{0, 1, 2}, {0, 2, 3}, {0, 3, 4}},
		},
		{
			graph:          gc_20_1_Graph(),
			expectedResult: [][3]uint32{{2, 11, 17}},
		},
	}

	for _, test := range tests {
		res := find3VerticesCompleteGraph(test.graph)

		if numOfRes, expectedNumOfRes := len(res), len(test.expectedResult); numOfRes != expectedNumOfRes {
			t.Errorf("Expected %d 3-vertices complete graph but got %d", expectedNumOfRes, numOfRes)
		}

		for idx, vertices := range res {
			if idx < len(test.expectedResult) {
				if vertices != test.expectedResult[idx] {
					t.Errorf("Expected %v but got %v", test.expectedResult[idx], vertices)
				}
			} else {
				t.Errorf("Did not expected result %v", vertices)
			}
		}
	}
}
