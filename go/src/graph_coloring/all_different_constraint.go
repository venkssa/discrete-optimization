package graph_coloring

type AllDifferentConstraint struct {
	vertices [3]uint32
	maxColor color
	adg      *allDifferentGraph
}

func (adc *AllDifferentConstraint) IsFeasible(graph *Graph, domainStore *DomainStore) bool {
	if adc.adg == nil {
		adc.adg = &allDifferentGraph{
			map[int32][]color{},
			make([]int32, adc.maxColor+2),
			NewVisitKeeper(int(graph.NumOfVertices)),
		}
	}
	for idx := range adc.adg.colorToVertexEdge {
		adc.adg.colorToVertexEdge[idx] = -1
	}

	for _, vertex := range adc.vertices {
		if color := domainStore.Color(vertex); color != UNSET {
			adc.adg.colorToVertexEdge[int(color)] = int32(vertex)
		}

		colorPalette := make([]bool, int(adc.maxColor)+2)
		for _, neighbor := range graph.Neighbors(vertex) {
			neighborColor := domainStore.Color(neighbor)
			if neighborColor == UNSET {
				continue
			}
			colorPalette[neighborColor] = true
		}
		if color := domainStore.Color(vertex); color != UNSET {
			colorPalette[int(color)] = true
		}
		possibleColors := []color{}
		for idx, isColored := range colorPalette[1 : len(colorPalette)-1] {
			if !isColored {
				possibleColors = append(possibleColors, color(idx+1))
			}
		}
		adc.adg.vertexToPossibleColors[int32(vertex)] = possibleColors
	}

	if adc.adg.MaximumMatching() == int32(len(adc.vertices)) {
		return true
	}
	return false
}

func (adc *AllDifferentConstraint) Prune(graph *Graph, domainStore *DomainStore) bool {
	return false
}

func BuildAllDifferentConstraint(graph *Graph, maxColor color) []Constraint {
	res := find3VerticesCompleteGraph(graph)

	constraints := make([]Constraint, len(res))

	for idx := 0; idx < len(res); idx++ {
		constraints[idx] = &AllDifferentConstraint{res[idx], maxColor, nil}
	}

	return constraints
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

type allDifferentGraph struct {
	vertexToPossibleColors map[int32][]color
	colorToVertexEdge      []int32
	vk                     *visitKeeper
}

func (adg *allDifferentGraph) IsFreeVertex(freeVeretex int32) bool {
	return adg.VertexColor(freeVeretex) == 0
}

func (adg *allDifferentGraph) VertexColor(v int32) int32 {
	for idx, vertex := range adg.colorToVertexEdge {
		if vertex == v {
			return int32(idx)
		}
	}
	return 0
}

func (adg *allDifferentGraph) MaximumMatching() int32 {
	for vertex := range adg.vertexToPossibleColors {
		if adg.IsFreeVertex(vertex) {
			adg.vk.Reset()
			path := adg.findAlternatingPath(uint32(vertex))
			if len(path) == 0 {
				break
			}
			for i := 0; i < len(path); i += 2 {
				possibleColors := make([]color, 0, len(adg.vertexToPossibleColors[path[i]])-1)
				for _, possibleColor := range adg.vertexToPossibleColors[path[i]] {
					if possibleColor == color(path[i+1]) {
						continue
					}
					possibleColors = append(possibleColors, color(path[i+1]))
				}
				adg.colorToVertexEdge[path[i+1]] = path[i]
			}

		}
	}

	var maxMatching int32

	for vertex := range adg.vertexToPossibleColors {
		if !adg.IsFreeVertex(vertex) {
			maxMatching++
		}
	}

	return maxMatching
}

type visitKeeper struct {
	visitedVertices map[int32]bool
	visitedBy       map[int32]int32
}

func NewVisitKeeper(numVertices int) *visitKeeper {
	return &visitKeeper{map[int32]bool{}, map[int32]int32{}}
}

func (vk *visitKeeper) Reset() {
	vk.visitedVertices = map[int32]bool{}
	vk.visitedBy = map[int32]int32{}
}

func (adg *allDifferentGraph) findAlternatingPath(freeVertex uint32) []int32 {
	stack := []int32{int32(freeVertex)}
	vk := adg.vk

	vk.visitedBy[int32(freeVertex)] = int32(freeVertex)

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1 : len(stack)-1]

		vk.visitedVertices[vertex] = true
		for _, color := range adg.vertexToPossibleColors[vertex] {
			connectedVertex := adg.colorToVertexEdge[int32(color)]
			if connectedVertex == -1 {
				return append(buildAlternatingPath(adg, vk.visitedBy, vertex), int32(color))
			}
			if vk.visitedVertices[connectedVertex] == false {
				stack = append(stack, connectedVertex)
				vk.visitedBy[connectedVertex] = vertex
			}
		}
	}
	return []int32{}
}

func buildAlternatingPath(adg *allDifferentGraph, visitedBy map[int32]int32, lastVisited int32) []int32 {
	path := make([]int32, 0, 2)
	visitNext := lastVisited
	loop := true

	for loop {
		loop = visitedBy[visitNext] != visitNext
		if loop {
			path = append([]int32{adg.VertexColor(visitNext), visitNext}, path...)
		} else {
			path = append([]int32{visitNext}, path...)
		}
		visitNext = visitedBy[visitNext]
	}

	return path
}
