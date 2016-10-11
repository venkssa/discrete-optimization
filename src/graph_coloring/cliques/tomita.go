package cliques

import (
	"graph_coloring/graph"
)

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
		newBitSetPool(graph.NumOfVertices),
		newPivotFinder(neighborsBitSet(graph)),
		&Cliques{Cliques: []Clique{}, NumOfVertices: graph.NumOfVertices})
}

func tomitaMaximalClique(
	r Clique,
	candidate *BitSet,
	finished *BitSet,
	pool *bitSetPool,
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
	pool.Return(candidateMinusPivotNeighbor)
	pool.Return(candidateCopy)
	pool.Return(finishedCopy)

	return result
}

type bitSetPool struct {
	bitSetSize uint32
	available []*BitSet
}

func newBitSetPool(bitSetSize uint32) *bitSetPool {
	return &bitSetPool{bitSetSize: bitSetSize}
}

func (p *bitSetPool) Borrow() *BitSet {
	if len(p.available) == 0 {
		return NewBitSet(p.bitSetSize)
	}

	lastIdx := len(p.available) - 1
	bs := p.available[lastIdx]
	p.available = p.available[0:lastIdx]
	return bs
}

func (p *bitSetPool) Return(bs *BitSet) {
	p.available = append(p.available, bs)
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
