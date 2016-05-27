package graph_coloring

type AllDifferent struct {
	vertex uint32
}

func (ad AllDifferent) IsFeasible(graph *Graph, domainStore *DomainStore) bool {
	vertexColor := domainStore.Color(ad.vertex)

	if vertexColor == UNSET {
		return true
	}

	for _, neighbor := range graph.Neighbors(ad.vertex) {
		neighborColor := domainStore.Color(neighbor)

		if neighborColor != UNSET && neighborColor == vertexColor {
			return false
		}
	}

	return true
}

func (ad AllDifferent) Prune(graph *Graph, domainStore *DomainStore) bool {
	colorPalette := make([]bool, graph.NumOfVertices)

	for _, neighbor := range graph.Neighbors(ad.vertex) {
		neighborColor := domainStore.Color(neighbor)
		if neighborColor == UNSET {
			return false
		}
		colorPalette[neighborColor-1] = true
	}

	for idx, isColored := range colorPalette {
		if !isColored {
			domainStore.Set(ad.vertex, color(idx+1))
		}
	}

	return true
}
