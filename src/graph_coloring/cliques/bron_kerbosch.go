package cliques

import "graph_coloring/graph"

type bronKerboschAlgo struct{}

func BronKerbosch() MaximalCliqueFinder {
	return bronKerboschAlgo{}
}

func (bk bronKerboschAlgo) FindAllMaximalCliques(graph *graph.G) *Cliques {
	p := NewBitSet(graph.NumOfVertices)
	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		p.Set(idx)
	}

	return bronKerboschMaximalClique(make(Clique, 0, graph.NumOfVertices),
		p,
		NewBitSet(graph.NumOfVertices),
		neighborsBitSet(graph),
		&Cliques{Cliques: []Clique{}, NumOfVertices: graph.NumOfVertices})
}

func bronKerboschMaximalClique(r Clique,
	p *BitSet,
	x *BitSet,
	vertexToEdgeBitSet []*BitSet,
	result *Cliques) *Cliques {

	if p.IsZero() && x.IsZero() {
		result.Add(r.Clone())
		return result
	}

	numOfVertices := uint32(len(vertexToEdgeBitSet))
	pCopy := NewBitSet(numOfVertices)
	xCopy := NewBitSet(numOfVertices)

	p.LoopOverSetIndices(func (vIdx uint32) {
		neighbors := vertexToEdgeBitSet[vIdx]
		Intersection(pCopy, neighbors, p)
		Intersection(xCopy, neighbors, x)

		bronKerboschMaximalClique(append(r, vIdx), pCopy, xCopy, vertexToEdgeBitSet, result)

		p.UnSet(vIdx)
		x.Set(vIdx)
	})
	return result
}
