package cliques

import (
	"graph_coloring/graph"
	"runtime"
	"sync"
)

type parallelTomita struct{}

func ParallelTomita() MaximalCliqueFinder {
	return parallelTomita{}
}

func (ta parallelTomita) FindAllMaximalCliques(g *graph.G) *Cliques {
	n := neighborsBitSet(g)
	q := queueTomitaRootCandidates(n)
	workers := newTomitaWorkers(n, q, uint32(runtime.NumCPU()))

	cliques := Cliques{Cliques: []Clique{}, NumOfVertices: g.NumOfVertices}

	for cliquesForSubTree := range workers.ResultChan {
		cliques.Add(cliquesForSubTree.Cliques...)
	}
	return &cliques
}

func queueTomitaRootCandidates(allNeighbors []*BitSet) *WorkQueue {
	finder := newPivotFinder(allNeighbors)
	candidate := NewBitSet(uint32(len(allNeighbors)))
	for idx := uint32(0); idx < uint32(len(allNeighbors)); idx++ {
		candidate.Set(idx)
	}

	// TODO: This is fairly similar to the actual recursive tomita algorithm. Refactor to remove this.
	finished := candidate.Not()
	pivot := finder.find(candidate, finished)
	candidateMinusPivotNeighbor := candidate.Minus(allNeighbors[pivot])

	q := NewWorkQueue()

	for v := uint32(0); v < candidateMinusPivotNeighbor.Len(); v++ {
		if !candidateMinusPivotNeighbor.IsSet(v) {
			continue
		}
		neighbor := allNeighbors[v]
		q.AddToQueue(Work {
			VIdx: v,
			Candidate: neighbor.Intersection(candidate),
			Finished: neighbor.Intersection(finished),
		})
		candidate.UnSet(v)
		finished.Set(v)
	}

	return q
}

// TODO This is similar to BK worker in parallel BK. Refactor to remove this.
type tomitaWorkers struct {
	ResultChan <-chan *Cliques
	q          *WorkQueue
}

func newTomitaWorkers(vertexToEdgeBitSet []*BitSet, q *WorkQueue, numOfWorkers uint32) *workers {
	resultChan := make(chan *Cliques)
	w := workers{
		ResultChan: resultChan,
		q:          q,
	}

	wg := &sync.WaitGroup{}
	wg.Add(int(numOfWorkers))

	for idx := uint32(0); idx < numOfWorkers; idx++ {
		go func() {
			for {
				wrk, err := q.GetWork()
				if err != nil {
					wg.Done()
					return
				}
				result := tomitaMaximalClique(
					append(make(Clique, 0, wrk.Candidate.Len()), wrk.VIdx),
					wrk.Candidate,
					wrk.Finished,
					vertexToEdgeBitSet,
					newPivotFinder(vertexToEdgeBitSet),
					&Cliques{Cliques: []Clique{}, NumOfVertices: wrk.Candidate.Len()})
				resultChan <- result
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	return &w
}
