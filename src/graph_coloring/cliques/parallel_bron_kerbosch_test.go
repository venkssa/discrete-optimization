package cliques

import (
	"testing"
	"graph_coloring/testdata"
)

func TestWorkQueue_GetWork_WhenNoWorkShouldReturnAnError(t *testing.T) {
	q := NewWorkQueue()
	work, err := q.GetWork()

	if err == nil {
		t.Errorf("Expected an error but got %v", work)
	}
	t.Log(work, err)
}

func TestWorkQueue_GetWork(t *testing.T) {
	q := NewWorkQueue()
	expectedWork := Work{VIdx: 1}

	q.AddToQueue(expectedWork)

	actualWork, err := q.GetWork()

	if err != nil {
		t.Fatalf("Expected work but got %v", err)
	}

	if actualWork.VIdx != expectedWork.VIdx {
		t.Fatalf("Expected %v but got %v", expectedWork, actualWork)
	}
}

func TestWorkQueue_GetWork_SecondTimeWithNoWorkShouldReturnAnError(t *testing.T) {
	q := NewWorkQueue()

	q.AddToQueue(Work{VIdx: 1})

	_, err := q.GetWork()
	if err != nil {
		t.Fatalf("Expected work but got %v", err)
	}

	_, err = q.GetWork()
	if err == nil {
		t.Fatal("Expected an error but did not get one")
	}
}

func TestParallelBKAlgo_FindAllMaximalCliques(t *testing.T) {
	g := testdata.Graph(testdata.GC_20_3)
	expectedCliques :=  BronKerbosch().FindAllMaximalCliques(g).Cliques
	parallelBKCliques := ParallelBKAlgo().FindAllMaximalCliques(g).Cliques

	verifyCliques(t, parallelBKCliques, expectedCliques)
}
