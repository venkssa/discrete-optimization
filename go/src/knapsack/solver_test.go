package knapsack

import (
	"testing"
)

func TestComputeEstimate(t *testing.T) {
	tests := []struct {
		selections       []bool
		expectedEstimate float64
	}{
		{
			selections:       []bool{false, false, false, false},
			expectedEstimate: 21.75,
		},
		{
			selections:       []bool{false, false, true, false},
			expectedEstimate: 21.0,
		},
		{
			selections:       []bool{false, false, false, true},
			expectedEstimate: 20.0,
		},
	}

	for _, test := range tests {
		estimate := computeEstimate(ks_4_0_Knapsack(), test.selections)

		if estimate != test.expectedEstimate {
			t.Errorf("Expected %f but was %f", test.expectedEstimate, estimate)
		}
	}
}

func BenchmarkComputeEstimate_ks_4_0(b *testing.B) {
	runBenchmarkComputeEstimate(b, ks_4_0_Knapsack())
}

func BenchmarkComputeEstimate_ks_19_0(b *testing.B) {
	runBenchmarkComputeEstimate(b, ks_19_0_Knapsack())
}

func BenchmarkComputeEstimate_ks_50_0(b *testing.B) {
	runBenchmarkComputeEstimate(b, ks_50_0_Knapsack())
}

func runBenchmarkComputeEstimate(b *testing.B, knapasack *Knapsack) {
	selections := make([]bool, len(knapasack.Items))

	for idx := 0; idx < b.N; idx++ {
		computeEstimate(knapasack, selections)
	}
}
