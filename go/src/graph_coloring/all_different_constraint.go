package graph_coloring

type allDifferentGraph struct {
	vertexToPossibleColors [][]color
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
	for vertex := 0; vertex < len(adg.vertexToPossibleColors); vertex++ {
		if adg.IsFreeVertex(int32(vertex)) {
			path := adg.findAlternatingPath(uint32(vertex))
			adg.vk.Reset()
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

	for vertex := 0; vertex < len(adg.vertexToPossibleColors); vertex++ {
		if !adg.IsFreeVertex(int32(vertex)) {
			maxMatching++
		}
	}

	return maxMatching
}

type visitKeeper struct {
	visitedVertices []bool
	visitedBy       []int32
}

func NewVisitKeeper(numVertices int) *visitKeeper {
	return &visitKeeper{make([]bool, numVertices), make([]int32, numVertices)}
}

func (vk *visitKeeper) Reset() {
	for idx := 0; idx < len(vk.visitedBy); idx++ {
		vk.visitedBy[idx] = 0
		vk.visitedVertices[idx] = false
	}
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

func buildAlternatingPath(adg *allDifferentGraph, visitedBy []int32, lastVisited int32) []int32 {
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
