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

	mcf := &maximalCliqueFinder{
		neighbors:     neighbors,
		candidates:    []*BitSet{candidates},
		finishedSlice: []*BitSet{NewBitSet(graph.NumOfVertices)},
		current:       1,
	}
	return tomitaMaximalClique(
		make(Clique, 0, graph.NumOfVertices),
		mcf,
		newPivotFinder(neighbors),
		&Cliques{Cliques: []Clique{}, NumOfVertices: graph.NumOfVertices})
}

func tomitaMaximalClique(r Clique, mcf *maximalCliqueFinder, pivotFinder *pivotFinder, result *Cliques) *Cliques {
	if mcf.resultFound() {
		result.Add(r.Clone())
		return result
	}

	candidate := mcf.currentCandidate()
	finished := mcf.currentFinished()
	pivot := pivotFinder.find(candidate, finished)

	candidateMinusPivotNeighbor := mcf.candidateMinusNeighborsOfPivot(pivot)

	for v := uint32(0); v < candidateMinusPivotNeighbor.Len(); v++ {
		if !candidateMinusPivotNeighbor.IsSet(v) {
			continue
		}

		mcf.increment(v)

		tomitaMaximalClique(append(r, v), mcf, pivotFinder, result)

		candidate.UnSet(v)
		finished.Set(v)
		mcf.decrement()
	}
	return result
}

type maximalCliqueFinder struct {
	neighbors []*BitSet

	candidates               []*BitSet
	finishedSlice            []*BitSet // rename
	candidatesMinusNeighbors []*BitSet
	current                  uint32
}

func (mcf *maximalCliqueFinder) resultFound() bool {
	return mcf.candidates[mcf.current-1].IsZero() && mcf.finishedSlice[mcf.current-1].IsZero()
}

func (mcf *maximalCliqueFinder) candidateMinusNeighborsOfPivot(pivot uint32) *BitSet {
	if mcf.current > uint32(len(mcf.candidatesMinusNeighbors)) {
		mcf.candidatesMinusNeighbors = append(mcf.candidatesMinusNeighbors, NewBitSet(uint32(len(mcf.neighbors))))
	}
	Minus(mcf.candidatesMinusNeighbors[mcf.current-1], mcf.currentCandidate(), mcf.neighbors[pivot])
	return mcf.candidatesMinusNeighbors[mcf.current-1]
}

func (mcf *maximalCliqueFinder) increment(vertexIdx uint32) {
	if mcf.current+1 > uint32(len(mcf.candidates)) {
		mcf.candidates = append(mcf.candidates, NewBitSet(uint32(len(mcf.neighbors))))
		mcf.finishedSlice = append(mcf.finishedSlice, NewBitSet(uint32(len(mcf.neighbors))))
	}

	prevCandidate := mcf.currentCandidate()
	prevFinished := mcf.currentFinished()

	mcf.current++
	Intersection(mcf.currentCandidate(), mcf.neighbors[vertexIdx], prevCandidate)
	Intersection(mcf.currentFinished(), mcf.neighbors[vertexIdx], prevFinished)
}

func (mcf *maximalCliqueFinder) decrement() {
	mcf.current--
}

func (mcf *maximalCliqueFinder) currentCandidate() *BitSet {
	return mcf.candidates[mcf.current-1]
}

func (mcf *maximalCliqueFinder) currentFinished() *BitSet {
	return mcf.finishedSlice[mcf.current-1]
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
