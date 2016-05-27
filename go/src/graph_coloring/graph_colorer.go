package graph_coloring

func ColorGraph(graph *Graph, maxColor color) []color {
	domainStore := NewDomainStore(graph.NumOfVertices)

	for idx := range domainStore.vertexColors {
		domainStore.Set(uint32(idx), color(idx+1))
	}

	return domainStore.vertexColors
}
