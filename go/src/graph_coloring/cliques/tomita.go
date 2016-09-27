package cliques

import (
	"graph_coloring/graph"
)

// Similar to BK except choose u P U X highest number of neigh in  P
// v in P \ N(u)
// SUBG = P U X (P == CAND, X == FINI)
// u in SUBG, where maximize | CAND ^ N(u) |
type TomitaAlgo struct{}

func (ta TomitaAlgo) FindAllMaximalCliques(graph *graph.G) *Cliques {
	candidates := NewBitSet(graph.NumOfVertices)
	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		candidates.Set(idx)
	}

	neighbors := neighborsBitSet(graph)

	return tomitaMaximalClique(
		make(Clique, 0, graph.NumOfVertices),
		candidates,
		NewBitSet(graph.NumOfVertices),
		neighbors,
		newPivotFinder(neighbors),
		&Cliques{Cliques: []Clique{}, NumOfVertices: graph.NumOfVertices})
}

func tomitaMaximalClique(
	r Clique,
	candidate *BitSet,
	finished *BitSet,
	allNeighbors []*BitSet,
	pivotFinder *pivotFinder,
	result *Cliques) *Cliques {

	if candidate.IsZero() && finished.IsZero() {
		result.Add(r.Clone())
		return result
	}

	pivot := pivotFinder.find(candidate, finished)

	candidateMinusPivotNeighbor := candidate.Minus(allNeighbors[pivot])

	candidateCopy := NewBitSet(result.NumOfVertices)
	finishedCopy := NewBitSet(result.NumOfVertices)

	for v := uint32(0); v < candidateMinusPivotNeighbor.Len(); v++ {
		if !candidateMinusPivotNeighbor.IsSet(v) {
			continue
		}

		neighbors := allNeighbors[v]
		Intersection(candidateCopy, neighbors, candidate)
		Intersection(finishedCopy, neighbors, finished)

		tomitaMaximalClique(append(r, v), candidateCopy, finishedCopy, allNeighbors, pivotFinder, result)

		candidate.UnSet(v)
		finished.Set(v)
	}
	return result
}

type pivotFinder struct {
	neighbors              []*BitSet
	subg                   *BitSet
	candidateMinusNeighbor *BitSet
}

func newPivotFinder(neighbors []*BitSet) *pivotFinder {
	return &pivotFinder{
		neighbors: neighbors,
		subg:      NewBitSet(uint32(len(neighbors))),
		candidateMinusNeighbor: NewBitSet(uint32(len(neighbors))),
	}
}

func (pf *pivotFinder) find(candidate *BitSet, finished *BitSet) uint32 {
	Union(pf.subg, candidate, finished)

	var maxVertexIdx uint32
	var maxCount uint32

	for idx := uint32(0); idx < pf.subg.Len(); idx++ {
		if !pf.subg.IsSet(idx) {
			continue
		}

		Intersection(pf.candidateMinusNeighbor, candidate, pf.neighbors[idx])

		count := candidate.NumOfBitsSet()

		if maxCount < count {
			maxCount = count
			maxVertexIdx = idx
		}

	}

	return maxVertexIdx
}
