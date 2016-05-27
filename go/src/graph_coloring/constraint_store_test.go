package graph_coloring

import "testing"

func TestAllDifferent_IsFeasible(t *testing.T) {
	tests := []struct {
		vertexColors []color
	}{
		{make([]color, 4)},
		{[]color{1, UNSET, 1, 1}},
		{[]color{1, 2, 1, 1}},
	}

	for _, test := range tests {
		graph := gc_4_1_Graph()
		domainStore := &DomainStore{test.vertexColors}

		constraint := AllDifferent{uint32(1)}

		if constraint.IsFeasible(graph, domainStore) != true {
			t.Errorf("Expected All different Constraint %v to be feasible on domain %v",
				constraint, domainStore)
		}
	}
}

func TestAllDifferent_IsNotFeasible(t *testing.T) {
	graph := gc_4_1_Graph()

	domainStore := &DomainStore{[]color{1, 1, 1, 1}}

	constraint := AllDifferent{uint32(1)}

	if constraint.IsFeasible(graph, domainStore) != false {
		t.Errorf("Expected constraint %v to be in-feasible on the domain %v", constraint, domainStore)
	}
}

func TestAllDifferent_Prune(t *testing.T) {
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
		domainStore := &DomainStore{test.vertexColors}

		constraint := AllDifferent{uint32(1)}

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

func TestAllDifferent_NothingToPrune(t *testing.T) {
	graph := gc_4_1_Graph()
	domainStore := &DomainStore{[]color{UNSET, UNSET, 1, 1}}
	constraint := AllDifferent{uint32(1)}

	if constraint.Prune(graph, domainStore) != false {
		t.Errorf("Expected constraint %v not to prune on the domain %v", constraint, domainStore)
	}
}

func BenchmarkAllDifferent_IsFeasible_UnsetVertex(b *testing.B) {
	benchmarkIsFeasible(b,
		gc_4_1_Graph(),
		&DomainStore{[]color{1, UNSET, 1, 1}},
		AllDifferent{uint32(1)})
}

func BenchmarkAllDifferent_IsFeasible_SetVertex(b *testing.B) {
	benchmarkIsFeasible(b,
		gc_4_1_Graph(),
		&DomainStore{[]color{1, 2, 1, 1}},
		AllDifferent{uint32(1)})
}

func BenchmarkAllDifferent_Prune(b *testing.B) {
	graph := gc_4_1_Graph()
	domainStore := &DomainStore{[]color{1, UNSET, 1, 1}}
	constraint := AllDifferent{uint32(1)}

	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		constraint.Prune(graph, domainStore)
	}
}

func benchmarkIsFeasible(b *testing.B, graph *Graph, domainStore *DomainStore, constraint AllDifferent) {
	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		constraint.IsFeasible(graph, domainStore)
	}
}
