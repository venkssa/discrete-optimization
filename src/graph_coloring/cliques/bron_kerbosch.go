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
		newBitSetPool(graph.NumOfVertices),
		neighborsBitSet(graph),
		&Cliques{Cliques: []Clique{}, NumOfVertices: graph.NumOfVertices})
}

func bronKerboschMaximalClique(r Clique,
	candidate *BitSet,
	finished *BitSet,
	pool *bitSetPool,
	vertexToEdgeBitSet []*BitSet,
	result *Cliques) *Cliques {

	if candidate.IsZero() && finished.IsZero() {
		result.Add(r.Clone())
		return result
	}

	candidateCopy := pool.Borrow()
	defer pool.Return(candidateCopy)
	finishedCopy := pool.Borrow()
	defer pool.Return(finishedCopy)

	candidate.LoopOverSetIndices(func (vIdx uint32) {
		neighbors := vertexToEdgeBitSet[vIdx]
		Intersection(candidateCopy, neighbors, candidate)
		Intersection(finishedCopy, neighbors, finished)

		bronKerboschMaximalClique(append(r, vIdx), candidateCopy, finishedCopy, pool, vertexToEdgeBitSet, result)

		candidate.UnSet(vIdx)
		finished.Set(vIdx)
	})
	return result
}
