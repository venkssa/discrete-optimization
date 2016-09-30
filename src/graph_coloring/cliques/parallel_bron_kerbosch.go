package cliques

import (
	"graph_coloring/graph"
	"runtime"
)

type parallelBKAlgo struct{}

func ParallelBKAlgo() MaximalCliqueFinder {
	return parallelBKAlgo{}
}

func (ta parallelBKAlgo) FindAllMaximalCliques(g *graph.G) *Cliques {
	wrks := []Worker{}
	neighborsBitSet := neighborsBitSet(g)
	for idx := uint32(0); idx < g.NumOfVertices; idx++ {
		wrks = append(wrks, &bkWorker{idx, neighborsBitSet})
	}

	return execute(wrks, runtime.NumCPU())
}

type bkWorker struct {
	vIdx            uint32
	neighborsBitSet []*BitSet
}

func (wrk *bkWorker) Work() *Cliques {
	numOfVertices := uint32(len(wrk.neighborsBitSet))
	candidate := NewBitSet(numOfVertices)
	for uIdx := wrk.vIdx; uIdx < numOfVertices; uIdx++ {
		candidate.Set(uIdx)
	}
	neighbors := wrk.neighborsBitSet[wrk.vIdx]
	finished := candidate.Not()
	Intersection(candidate, neighbors, candidate)
	Intersection(finished, neighbors, finished)

	return bronKerboschMaximalClique(
		append(make(Clique, 0, numOfVertices), wrk.vIdx),
		candidate,
		finished,
		wrk.neighborsBitSet,
		&Cliques{Cliques: []Clique{}, NumOfVertices: numOfVertices})
}
