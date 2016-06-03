package graph_coloring

func ColorGraph(graph *Graph, maxColor color) result {
	constraints := append(buildAllDifferentConstraints(graph.NumOfVertices), MaxColor{maxColor})

	ta := &result{
		graph:       graph,
		constraints: constraints,
		maxColor:    maxColor,
		Stats:       make([][]uint32, graph.NumOfVertices),
		Coloring:    nil}
	for idx := 0; idx < len(ta.Stats); idx++ {
		ta.Stats[idx] = make([]uint32, maxColor+1)
	}

	ta.tryAll(NewDomainStore(graph.NumOfVertices), 0)
	return *ta
}

type result struct {
	graph       *Graph
	constraints []Constraint
	maxColor    color

	Stats       [][]uint32
	Coloring    []color
}

func (res *result) tryAll(domain *DomainStore, vertexIdx uint32) {
	for currentColor := color(1); currentColor <= res.maxColor; currentColor++ {
		res.Stats[vertexIdx][currentColor]++
		cd := MakeACopy(domain)
		cd.Set(vertexIdx, currentColor)
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

func propogate(graph *Graph, domainStore *DomainStore, constraints []Constraint) bool {
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

func buildAllDifferentConstraints(numOfVertices uint32) []Constraint {
	allDiff := make([]Constraint, numOfVertices)
	for idx := uint32(0); idx < numOfVertices; idx++ {
		allDiff[idx] = NotEqual{idx}
	}
	return allDiff
}
