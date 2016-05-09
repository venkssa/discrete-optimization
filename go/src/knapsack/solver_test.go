package knapsack

import "testing"

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

	knapsack := &Knapsack{
		Capacity: 11,
		Items: Items([]Item{
			{
				Idx:    1,
				Value:  8,
				Weight: 4,
			},
			{
				Idx:    2,
				Value:  10,
				Weight: 5,
			},
			{
				Idx:    0,
				Value:  15,
				Weight: 8,
			},
			{
				Idx:    3,
				Value:  4,
				Weight: 3,
			},
		}),
	}

	for _, test := range tests {
		estimate := computeEstimate(knapsack, test.selections)

		if estimate != test.expectedEstimate {
			t.Errorf("Expected %f but was %f", test.expectedEstimate, estimate)
		}
	}
}

func BenchmarkComputeEstimate(b *testing.B) {
	selections := []bool{false, false, false, true}

	knapsack := &Knapsack{
		Capacity: 11,
		Items: Items([]Item{
			{
				Idx:    1,
				Value:  8,
				Weight: 4,
			},
			{
				Idx:    2,
				Value:  10,
				Weight: 5,
			},
			{
				Idx:    0,
				Value:  15,
				Weight: 8,
			},
			{
				Idx:    3,
				Value:  4,
				Weight: 3,
			},
		}),
	}
	for idx := 0; idx < b.N; idx++ {
		computeEstimate(knapsack, selections)
	}
}
