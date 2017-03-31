package graph_coloring

import (
	"github.com/venkssa/discrete-optimization/src/graph_coloring/cliques"
	"github.com/venkssa/discrete-optimization/src/graph_coloring/graph"
)

func ColorGraph(graph *graph.G, maxColor color) result {
	constraints := append(buildNotEqualConstraints(graph.NumOfVertices), MaxColor{maxColor: maxColor})
	constraints = append(constraints, BuildAllDifferentConstraint(graph, maxColor)...)

	clique := NewCliques(*cliques.BronKerbosch().FindAllMaximalCliques(graph))

	ta := &result{
		graph:       graph,
		constraints: constraints,
		maxColor:    maxColor,
		searchOrder: OrderVerticesByCliqueLen(clique),
		Stats:       make([][]uint32, graph.NumOfVertices),
		Coloring:    nil}
	for idx := 0; idx < len(ta.Stats); idx++ {
		ta.Stats[idx] = make([]uint32, maxColor+1)
	}

	ta.tryAll(NewDomainStore(graph.NumOfVertices), 0)
	return *ta
}

type result struct {
	Stats    [][]uint32
	Coloring []color

	graph       *graph.G
	constraints []Constraint
	maxColor    color
	searchOrder []uint32
}

func (res *result) tryAll(domain *DomainStore, vertexIdx uint32) {
	vertex := res.searchOrder[vertexIdx]
	for currentColor := color(1); currentColor <= res.maxColor; currentColor++ {
		res.Stats[vertex][currentColor]++
		cd := MakeACopy(domain)
		cd.Set(vertex, currentColor)
		if propogate(res.graph, cd, res.constraints) {
			if cd.IsAllVertexColored() {
				res.Coloring = cd.vertexColors
				return
			}
			res.tryAll(cd, vertexIdx+1)
			if res.Coloring != nil {
				return
			}
		}
	}
	return
}

func propogate(graph *graph.G, domainStore *DomainStore, constraints []Constraint) bool {
	for keepPropogating := true; keepPropogating; {
		for _, constraint := range constraints {
			if !constraint.IsFeasible(graph, domainStore) {
				return false
			}
		}

		pruneCount := 0
		for _, constraint := range constraints {
			if constraint.Prune(graph, domainStore) {
				pruneCount++
			}
		}
		keepPropogating = pruneCount > 0
	}
	return true
}

func buildNotEqualConstraints(numOfVertices uint32) []Constraint {
	notEqual := make([]Constraint, numOfVertices)
	for idx := uint32(0); idx < numOfVertices; idx++ {
		notEqual[idx] = NotEqual{idx}
	}
	return notEqual
}
