package graph_coloring

import (
	"testing"
	"graph_coloring/graph"
	"graph_coloring/test_data"
)

func TestNotEqual_IsFeasible(t *testing.T) {
	tests := []struct {
		vertexColors []color
	}{
		{make([]color, 4)},
		{[]color{1, UNSET, 1, 1}},
		{[]color{1, 2, 1, 1}},
	}

	for _, test := range tests {
		graph := test_data.Gc_4_1_Graph()
		domainStore := &DomainStore{test.vertexColors, 0}

		constraint := NotEqual{uint32(1)}

		if constraint.IsFeasible(graph, domainStore) != true {
			t.Errorf("Expected All different Constraint %v to be feasible on domain %v",
				constraint, domainStore)
		}
	}
}

func TestNotEqual_IsNotFeasible(t *testing.T) {
	graph := test_data.Gc_4_1_Graph()

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
		graph := test_data.Gc_4_1_Graph()
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
	graph := test_data.Gc_4_1_Graph()
	domainStore := &DomainStore{[]color{UNSET, UNSET, 1, 1}, 0}
	constraint := NotEqual{uint32(1)}

	if constraint.Prune(graph, domainStore) != false {
		t.Errorf("Expected constraint %v not to prune on the domain %v", constraint, domainStore)
	}
}

func BenchmarkNotEqual_IsFeasible_UnsetVertex(b *testing.B) {
	benchmarkIsFeasible(b,
		test_data.Gc_4_1_Graph(),
		&DomainStore{[]color{1, UNSET, 1, 1}, 0},
		NotEqual{uint32(1)})
}

func BenchmarkNotEqual_IsFeasible_SetVertex(b *testing.B) {
	benchmarkIsFeasible(b,
		test_data.Gc_4_1_Graph(),
		&DomainStore{[]color{1, 2, 1, 1}, 0},
		NotEqual{uint32(1)})
}

func BenchmarkNotEqual_Prune(b *testing.B) {
	graph := test_data.Gc_4_1_Graph()
	domainStore := &DomainStore{[]color{1, UNSET, 1, 1}, 0}
	constraint := NotEqual{uint32(1)}

	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		constraint.Prune(graph, domainStore)
	}
}

func benchmarkIsFeasible(b *testing.B, graph *graph.G, domainStore *DomainStore, constraint NotEqual) {
	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		constraint.IsFeasible(graph, domainStore)
	}
}
