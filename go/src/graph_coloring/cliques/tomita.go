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

	return tomitaMaximalClique(
		make(Clique, 0, graph.NumOfVertices),
		candidates,
		NewBitSet(graph.NumOfVertices),
		neighborsBitSet(graph),
		&Cliques{Cliques: []Clique{}, NumOfVertices: graph.NumOfVertices})
}

func tomitaMaximalClique(
	r Clique,
	candidate *BitSet,
	finished *BitSet,
	allNeighbors []*BitSet,
	result *Cliques) *Cliques {

	if candidate.IsZero() && finished.IsZero() {
		result.Add(r.Clone())
		return result
	}

	pivot := findPivot(candidate, finished, allNeighbors)

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

		tomitaMaximalClique(append(r, v), candidateCopy, finishedCopy, allNeighbors, result)

		candidate.UnSet(v)
		finished.Set(v)
	}
	return result
}

func findPivot(candidate *BitSet, finished *BitSet, neighbors []*BitSet) uint32 {
	subg := candidate.Union(finished)

	intersection := NewBitSet(subg.Len())

	var maxVertexIdx uint32
	var maxCount uint32

	for idx := uint32(0); idx < subg.Len(); idx++ {
		if !subg.IsSet(idx) {
			continue
		}

		Intersection(intersection, candidate, neighbors[idx])

		count := candidate.NumOfBitsSet()

		if maxCount < count {
			maxCount = count
			maxVertexIdx = idx
		}

	}

	return maxVertexIdx
}
