package cliques

import "sync"

type Worker interface {
	Work() *Cliques
}

func execute(wrks []Worker, poolSize int) *Cliques {
	workerChan := make(chan Worker, len(wrks))
	for _, wrk := range wrks {
		workerChan <- wrk
	}
	close(workerChan)

	wg := &sync.WaitGroup{}
	wg.Add(poolSize)
	resultChan := make(chan *Cliques)

	for idx := 0; idx < poolSize; idx++ {
		go func() {
			for worker := range workerChan {
				resultChan <- worker.Work()
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	cliques := Cliques{Cliques: []Clique{}}

	for cliquesForSubTree := range resultChan {
		cliques.Add(cliquesForSubTree.Cliques...)
		cliques.NumOfVertices = cliquesForSubTree.NumOfVertices
	}
	return &cliques
}
