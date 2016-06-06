package graph_coloring

type Constraint interface {
	IsFeasible(graph *Graph, domainStore *DomainStore) bool
	Prune(graph *Graph, domainStore *DomainStore) bool
}

type NotEqual struct {
	vertex uint32
}

func (neq NotEqual) IsFeasible(graph *Graph, domainStore *DomainStore) bool {
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

func (neq NotEqual) Prune(graph *Graph, domainStore *DomainStore) bool {
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

func (mc MaxColor) IsFeasible(graph *Graph, domainStore *DomainStore) bool {
	for _, color := range domainStore.vertexColors {
		if color > mc.maxColor {
			return false
		}
	}
	return true
}

func (mc MaxColor) Prune(graph *Graph, domainStore *DomainStore) bool {
	return false
}

func find3VerticesCompleteGraph(graph *Graph) [][3]uint32 {
	res := [][3]uint32{}
	for vertexIdx, neighbors := range graph.VertexToEdges {
		if len(neighbors) < 2 {
			continue
		}

		for _, outerNeighborIdx := range neighbors {
			for _, innerNeighborIdx := range neighbors {
				if outerNeighborIdx < uint32(vertexIdx) || innerNeighborIdx < outerNeighborIdx {
					continue
				}
				if graph.AreNeighbors(outerNeighborIdx, innerNeighborIdx) {
					res = append(res, [3]uint32{uint32(vertexIdx), outerNeighborIdx, innerNeighborIdx})
				}
			}
		}
	}
	return res
}
