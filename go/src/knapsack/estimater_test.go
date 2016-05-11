package knapsack

import "testing"

func TestEstimate(t *testing.T) {
	tests := []struct {
		selections       []selection
		expectedEstimate float64
	}{
		{
			selections:       []selection{},
			expectedEstimate: 21.75,
		},
		{
			selections:       []selection{SELECTED},
			expectedEstimate: 21.75,
		},
		{
			selections:       []selection{SKIPPED, SKIPPED, SELECTED},
			expectedEstimate: 19.0,
		},
		{
			selections:       []selection{SKIPPED, SKIPPED, SKIPPED, SELECTED},
			expectedEstimate: 4.0,
		},
		{
			selections:       []selection{SKIPPED, SKIPPED, SKIPPED, SKIPPED},
			expectedEstimate: 0.0,
		},
	}

	for _, test := range tests {
		estimate := Estimate(ks_4_0_Knapsack(), test.selections)

		if estimate != test.expectedEstimate {
			t.Errorf("Expected %f but was %f", test.expectedEstimate, estimate)
		}
	}
}

func BenchmarkEstimate_ks_4_0(b *testing.B) {
	runBenchmarkEstimate(b, 0, ks_4_0_Knapsack())
}

func BenchmarkEstimate_ks_4_0_with_first_2_skipped(b *testing.B) {
	runBenchmarkEstimate(b, 2, ks_4_0_Knapsack())
}

func BenchmarkEstimate_ks_19_0(b *testing.B) {
	runBenchmarkEstimate(b, 0, ks_19_0_Knapsack())
}

func BenchmarkEstimate_ks_19_0_with_first_9_skipped(b *testing.B) {
	runBenchmarkEstimate(b, 9, ks_19_0_Knapsack())
}

func BenchmarkEstimate_ks_50_0(b *testing.B) {
	runBenchmarkEstimate(b, 0, ks_50_0_Knapsack())
}

func BenchmarkEstimate_ks_50_0_with_first_25_skipped(b *testing.B) {
	runBenchmarkEstimate(b, 25, ks_50_0_Knapsack())
}

func BenchmarkEstimate_ks_50_0_with_first_45_skipped(b *testing.B) {
	runBenchmarkEstimate(b, 45, ks_50_0_Knapsack())
}

func runBenchmarkEstimate(b *testing.B, selectionsLen int, knapasack *Knapsack) {
	selections := make([]selection, selectionsLen)
	for idx := 0; idx < b.N; idx++ {
		Estimate(knapasack, selections)
	}
}
