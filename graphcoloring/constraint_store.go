package graphcoloring

import "github.com/venkssa/discrete-optimization/graphcoloring/graph"

type Constraint interface {
	IsFeasible(graph *graph.G, domainStore *DomainStore) bool
	Prune(graph *graph.G, domainStore *DomainStore) bool
}

type NotEqual struct {
	vertex uint32
}

func (neq NotEqual) IsFeasible(graph *graph.G, domainStore *DomainStore) bool {
	vertexColor := domainStore.Color(neq.vertex)

	if vertexColor == UNSET {
		return true
	}

	for _, neighbor := range graph.Neighbors(neq.vertex) {
		neighborColor := domainStore.Color(neighbor)

		if neighborColor != UNSET && neighborColor == vertexColor {
			return false
		}
	}

	return true
}

func (neq NotEqual) Prune(graph *graph.G, domainStore *DomainStore) bool {
	colorPalette := make([]bool, graph.NumOfVertices)

	if domainStore.IsSet(neq.vertex) {
		return false
	}

	for _, neighbor := range graph.Neighbors(neq.vertex) {
		neighborColor := domainStore.Color(neighbor)
		if neighborColor == UNSET {
			return false
		}
		colorPalette[neighborColor-1] = true
	}

	for idx, isColored := range colorPalette {
		if !isColored {
			domainStore.Set(neq.vertex, color(idx+1))
			return true
		}
	}

	return false
}

type MaxColor struct {
	maxColor color
}

func (mc MaxColor) IsFeasible(graph *graph.G, domainStore *DomainStore) bool {
	for _, color := range domainStore.vertexColors {
		if color > mc.maxColor {
			return false
		}
	}
	return true
}

func (mc MaxColor) Prune(graph *graph.G, domainStore *DomainStore) bool {
	return false
}
