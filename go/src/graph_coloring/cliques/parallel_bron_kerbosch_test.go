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

func TestInitParallelAlgo(t *testing.T) {
	n := neighborsBitSet(testdata.Graph(testdata.GC_4_1))
	q := queueRootCandidates(n)

	expectedWorks := []Work{
		{
			VIdx: 0,
			Candidate: bitsetForTest(4, 0, 1, 2, 3),
			Finished: bitsetForTest(4),
		},
		{
			VIdx: 1,
			Candidate: bitsetForTest(4, 1, 2, 3),
			Finished: bitsetForTest(4, 0),
		},
		{
			VIdx: 2,
			Candidate: bitsetForTest(4, 2, 3),
			Finished: bitsetForTest(4, 0, 1),
		},
		{
			VIdx: 3,
			Candidate: bitsetForTest(4, 3),
			Finished: bitsetForTest(4, 0, 1, 2),
		},
	}

	for _, expectedWork := range expectedWorks {
		work, err := q.GetWork()
		if err != nil {
			t.Fatalf("Expected work but got %v", err)
		}
		if work.Candidate.String() != expectedWork.Candidate.Intersection(n[work.VIdx]).String() {
			t.Errorf("Expected candidate to be %v but was %v", expectedWork.Candidate, work.Candidate)
		}
		if work.Finished.String() != expectedWork.Finished.Intersection(n[work.VIdx]).String() {
			t.Errorf("Expected finished to be %v but was %v", expectedWork.Finished, work.Finished)
		}
	}
}

func TestParallelBKAlgo_FindAllMaximalCliques(t *testing.T) {
	g := testdata.Graph(testdata.GC_20_3)
	expectedCliques :=  BronKerbosch().FindAllMaximalCliques(g).Cliques
	parallelBKCliques := ParallelBKAlgo{}.FindAllMaximalCliques(g).Cliques

	verifyCliques(t, parallelBKCliques, expectedCliques)
}
