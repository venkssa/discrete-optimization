package cliques

import (
	"testing"
	"time"
)

type TestWorker struct {
	fn func() Cliques
}

func (tw TestWorker) Work() *Cliques {
	res := tw.fn()
	return &res
}

func sleepingTestWorker(sleepTime time.Duration, result Cliques) Worker {
	return TestWorker{
		fn: func() Cliques {
			time.Sleep(sleepTime)
			return result
		},
	}
}

func TestExecute_ExecutesWorkAndCollectsResult(t *testing.T) {
	sleepTime := 1 * time.Millisecond
	workerResults := []Cliques{
		{[]Clique{{0, 1, 2}}, 5},
		{[]Clique{{0, 2, 3}}, 5},
		{[]Clique{{0, 3, 4}}, 5},
		{[]Clique{{0, 1, 4}}, 5},
	}

	t.Run("WorkerPool_1", testExecuteWorkAndCollectResult(1, sleepTime, workerResults, 4 * time.Millisecond))
	t.Run("WorkerPool_2", testExecuteWorkAndCollectResult(2, sleepTime, workerResults, 2 * time.Millisecond))
	t.Run("WorkerPool_4", testExecuteWorkAndCollectResult(4, sleepTime, workerResults, 1 * time.Millisecond))
}

func testExecuteWorkAndCollectResult(
	poolSize int,
	sleepTime time.Duration,
	workerResults []Cliques,
	expectedTimeTaken time.Duration) func(*testing.T) {
	return func(t *testing.T) {
		workers := []Worker{}
		expected := Cliques{}
		for _, workerResult := range workerResults {
			workers = append(workers, sleepingTestWorker(sleepTime, workerResult))
			expected.Add(workerResult.Cliques...)
			expected.NumOfVertices = workerResult.NumOfVertices
		}

		startTime := time.Now()
		actual := execute(workers, poolSize)
		timeTaken := time.Now().Sub(startTime)

		verifyCliques(t, actual.Cliques, expected.Cliques)

		if expectedTimeTaken > timeTaken {
			t.Errorf("Expected to take %v but took %v", expectedTimeTaken, timeTaken)
		}
	}
}
