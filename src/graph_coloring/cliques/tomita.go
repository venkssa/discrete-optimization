package cliques

import "github.com/venkssa/discrete-optimization/src/graph_coloring/graph"

// Similar to BK except choose u P U X highest number of neigh in  P
// v in P \ N(u)
// SUBG = P U X (P == CAND, X == FINI)
// u in SUBG, where maximize | CAND ^ N(u) |
type tomitaAlgo struct{}

func TomitaAlgo() MaximalCliqueFinder {
	return tomitaAlgo{}
}

func (ta tomitaAlgo) FindAllMaximalCliques(graph *graph.G) *Cliques {
	candidates := NewBitSet(graph.NumOfVertices)
	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		candidates.Set(idx)
	}

	return tomitaMaximalClique(
		make(Clique, 0, graph.NumOfVertices),
		candidates,
		NewBitSet(graph.NumOfVertices),
		NewBitSetPool(graph.NumOfVertices),
		newPivotFinder(neighborsBitSet(graph)),
		&Cliques{Cliques: []Clique{}, NumOfVertices: graph.NumOfVertices})
}

func tomitaMaximalClique(
	r Clique,
	candidate *BitSet,
	finished *BitSet,
	pool *BitSetPool,
	pivotFinder *pivotFinder,
	result *Cliques) *Cliques {

	if candidate.IsZero() && finished.IsZero() {
		result.Add(r.Clone())
		return result
	}

	candidateCopy := pool.Borrow()
	finishedCopy := pool.Borrow()

	pivot := pivotFinder.find(candidate, finished)
	candidateMinusPivotNeighbor := pool.Borrow()
	Minus(candidateMinusPivotNeighbor, candidate, pivotFinder.neighbors[pivot])

	candidateMinusPivotNeighbor.LoopOverSetIndices(func(vIdx uint32) {
		neighbors := pivotFinder.neighbors[vIdx]
		Intersection(candidateCopy, neighbors, candidate)
		Intersection(finishedCopy, neighbors, finished)

		tomitaMaximalClique(append(r, vIdx), candidateCopy, finishedCopy, pool, pivotFinder, result)
		candidate.UnSet(vIdx)
		finished.Set(vIdx)
	})

	pool.Return(candidateCopy)
	pool.Return(finishedCopy)
	pool.Return(candidateMinusPivotNeighbor)
	return result
}

type pivotFinder struct {
	neighbors                   []*BitSet
	subg                        *BitSet
	candidateIntersectNeighbors *BitSet
}

func newPivotFinder(neighbors []*BitSet) *pivotFinder {
	return &pivotFinder{
		neighbors: neighbors,
		subg:      NewBitSet(uint32(len(neighbors))),
		candidateIntersectNeighbors: NewBitSet(uint32(len(neighbors))),
	}
}

func (pf *pivotFinder) find(candidate *BitSet, finished *BitSet) uint32 {
	Union(pf.subg, candidate, finished)

	var maxVertexIdx uint32
	var maxCount uint32

	pf.subg.LoopOverSetIndices(func(vIdx uint32) {
		Intersection(pf.candidateIntersectNeighbors, candidate, pf.neighbors[vIdx])
		count := pf.candidateIntersectNeighbors.NumOfBitsSet()
		if maxCount <= count {
			maxCount = count
			maxVertexIdx = vIdx
		}
	})

	return maxVertexIdx
}
