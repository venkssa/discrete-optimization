package cliques

import (
	"graph_coloring/graph"
	"runtime"
)

type parallelTomita struct{}

func ParallelTomita() MaximalCliqueFinder {
	return parallelTomita{}
}

func (ta parallelTomita) FindAllMaximalCliques(g *graph.G) *Cliques {
	neighborsBitSet := neighborsBitSet(g)
	candidate := NewBitSet(g.NumOfVertices)
	for idx := uint32(0); idx < g.NumOfVertices; idx++ {
		candidate.Set(idx)
	}
	finished := candidate.Not()
	finder := newPivotFinder(neighborsBitSet)
	pivot := finder.find(candidate, finished)
	candidateMinusPivotNeighbor := candidate.Minus(neighborsBitSet[pivot])

	wrks := []Worker{}

	for vIdx := uint32(0); vIdx < candidateMinusPivotNeighbor.Len(); vIdx++ {
		if !candidateMinusPivotNeighbor.IsSet(vIdx) {
			continue
		}
		neighbor := neighborsBitSet[vIdx]
		wrks = append(wrks, &tomitaWorker{
			vIdx:            vIdx,
			neighborsBitSet: neighborsBitSet,
			candidate:       neighbor.Intersection(candidate),
			finished:        neighbor.Intersection(finished),
		})

		candidate.UnSet(vIdx)
		finished.Set(vIdx)
	}

	return execute(wrks, runtime.NumCPU())
}

type tomitaWorker struct {
	vIdx            uint32
	neighborsBitSet []*BitSet
	candidate       *BitSet
	finished        *BitSet
}

func (wrk *tomitaWorker) Work() *Cliques {
	numOfVertices := uint32(len(wrk.neighborsBitSet))

	return tomitaMaximalClique(
		append(make(Clique, 0, numOfVertices), wrk.vIdx),
		wrk.candidate,
		wrk.finished,
		wrk.neighborsBitSet,
		newPivotFinder(wrk.neighborsBitSet),
		&Cliques{Cliques: []Clique{}, NumOfVertices: numOfVertices})
}
