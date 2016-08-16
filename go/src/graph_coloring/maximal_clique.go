package graph_coloring

import "graph_coloring/graph"

func FindAllMaximalCliques(graph *graph.G) [][]uint32 {
	p := map[uint32]bool{}
	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		p[idx] = true
	}
	return bkAlgo(make([]uint32, 0, graph.NumOfVertices), p, map[uint32]bool{}, graph, [][]uint32{})
}

func bkAlgo(r []uint32, p map[uint32]bool, x map[uint32]bool, graph *graph.G, result [][]uint32) [][]uint32 {
	if len(p) == 0 && len(x) == 0 {
		return append(result, append([]uint32{}, r...))
	}

	for v := range p {
		pCopy := map[uint32]bool{}
		xCopy := map[uint32]bool{}

		for _, neighbor := range graph.Neighbors(v) {
			if _, ok := x[neighbor]; ok {
				xCopy[neighbor] = true
			}
			if _, ok := p[neighbor]; ok {
				pCopy[neighbor] = true
			}
		}
		result = bkAlgo(append(r, v), pCopy, xCopy, graph, result)

		delete(p, v)
		x[v] = true
	}
	return result
}
