package cliques

import (
	"fmt"
	"graph_coloring/graph"
	"runtime"
	"sync"
)

type ParallelBKAlgo struct{}

func (ta ParallelBKAlgo) FindAllMaximalCliques(g *graph.G) *Cliques {
	n := neighborsBitSet(g)
	q := queueRootCandidates(n)
	workers := newWorkers(n, q, uint32(runtime.NumCPU()))

	cliques := Cliques{Cliques: []Clique{}, NumOfVertices: g.NumOfVertices}

	for cliquesForSubTree := range workers.ResultChan {
		cliques.Add(cliquesForSubTree.Cliques...)
	}
	return &cliques
}

type workers struct {
	ResultChan <-chan *Cliques
	q          *WorkQueue
}

func newWorkers(vertexToEdgeBitSet []*BitSet, q *WorkQueue, numOfWorkers uint32) *workers {
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
				result := bronKerboschMaximalClique(
					append(make(Clique, 0, wrk.Candidate.Len()), wrk.VIdx),
					wrk.Candidate,
					wrk.Finished,
					vertexToEdgeBitSet,
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

func queueRootCandidates(vertexToEdgeBitSet []*BitSet) *WorkQueue {
	q := NewWorkQueue()
	numOfVertices := uint32(len(vertexToEdgeBitSet))
	for vIdx := uint32(0); vIdx < numOfVertices; vIdx++ {
		candidate := NewBitSet(numOfVertices)
		for uIdx := vIdx; uIdx < numOfVertices; uIdx++ {
			candidate.Set(uIdx)
		}
		neighbors := vertexToEdgeBitSet[vIdx]
		finished := candidate.Not()
		Intersection(candidate, neighbors, candidate)
		Intersection(finished, neighbors, finished)
		w := Work{
			VIdx:      vIdx,
			Candidate: candidate,
			Finished:  finished,
		}
		q.AddToQueue(w)
	}

	return q
}

type Work struct {
	VIdx      uint32
	Candidate *BitSet
	Finished  *BitSet
}

type getWorkResponse struct {
	w   Work
	err error
}

type WorkQueue struct {
	addToQueueChan chan<- []Work
	getWorkChan    chan<- chan<- getWorkResponse
}

func NewWorkQueue() *WorkQueue {
	addToQueueChan := make(chan []Work)
	getWorkChan := make(chan chan<- getWorkResponse)

	q := WorkQueue{
		addToQueueChan: addToQueueChan,
		getWorkChan:    getWorkChan,
	}

	go func() {
		var workQueue []Work
		for {
			select {
			case rChan := <-getWorkChan:
				if len(workQueue) == 0 {
					rChan <- getWorkResponse{err: noWorkError}
				} else {
					rChan <- getWorkResponse{w: workQueue[0]}
					workQueue = workQueue[1:]
				}
			case works := <-addToQueueChan:
				workQueue = append(workQueue, works...)
			}
		}
	}()

	return &q
}

var noWorkError = fmt.Errorf("No work present")

func (wq *WorkQueue) GetWork() (Work, error) {
	rChan := make(chan getWorkResponse)

	wq.getWorkChan <- rChan
	resp := <-rChan

	return resp.w, resp.err
}

func (wq *WorkQueue) AddToQueue(w ...Work) {
	wq.addToQueueChan <- w
}
