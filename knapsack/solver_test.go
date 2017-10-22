package knapsack

import (
	"reflect"
	"testing"
)

func TestNode_IsLeft(t *testing.T) {
	tests := map[string]struct {
		node   Node
		isLeft bool
	}{
		"IsLeft if SELECTED": {
			node:   Node{Idx: 0, selections: []selection{SELECTED}},
			isLeft: true,
		},
		"Is not left if SKIPPED": {
			node:   Node{Idx: 0, selections: []selection{SKIPPED}},
			isLeft: false,
		},
	}

	for testName, testdata := range tests {
		t.Run(testName, func(t *testing.T) {
			if testdata.node.isLeft() != testdata.isLeft {
				t.Errorf("Expected isLeft = %v but was %v (%#v)", testdata.isLeft, testdata.node.isLeft(), testdata.node)
			}
		})
	}
}

func TestNode_NextNodes_OnALeft_ReturnsRightNodeAsTheFirstResult(t *testing.T) {
	knapsack := &Knapsack{
		Capacity: 5,
		Items: []Item{
			{
				Value:  5,
				Weight: 5,
			},
			{
				Value:  2,
				Weight: 5,
			},
		},
	}

	selections := []selection{SELECTED, SKIPPED}
	node := Node{Idx: 0, selections: selections[0:1], usedCapacity: 5, estimate: estimate(knapsack, 5, 0, selections)}

	rightNode, _ := node.NextNodes(knapsack)

	if rightNode.isLeft() != false {
		t.Error("Expected next node to be a right node but was left.")
	}
	if rightNode.Idx != node.Idx {
		t.Errorf("Expected idx to be %d but was %d", node.Idx, rightNode.Idx)
	}
}

func TestNode_NextNodes(t *testing.T) {
	tests := map[string]struct {
		capacity     uint32
		expectedNode Node
	}{
		"Left child is returned": {
			capacity: 7,
			expectedNode: Node{
				Idx:          1,
				selections:   []selection{SELECTED, SELECTED},
				usedCapacity: 7,
				currentValue: 7,
				estimate:     7.0,
			},
		},
		"Right child is returned": {
			capacity: 6,
			expectedNode: Node{
				Idx:          1,
				selections:   []selection{SELECTED, SKIPPED},
				usedCapacity: 2,
				currentValue: 5,
				estimate:     5.0,
			},
		},
	}

	for testName, testdata := range tests {
		t.Run(testName, func(t *testing.T) {
			knapsack := &Knapsack{
				Capacity: testdata.capacity,
				Items: []Item{
					{
						Value:  5,
						Weight: 2,
					},
					{
						Value:  2,
						Weight: 5,
					},
				},
			}
			selections := []selection{SELECTED, SKIPPED}
			node := Node{Idx: 0, selections: selections[0:1], usedCapacity: 2, currentValue: 5,
				estimate: estimate(knapsack, 2, 5, selections)}

			_, childNode := node.NextNodes(knapsack)

			verifyNode(t, childNode, &testdata.expectedNode)
		})
	}
}

func TestRootNode(t *testing.T) {
	tests := map[string]struct {
		knapsack     *Knapsack
		expectedNode Node
	}{
		"ks 4_0 knapsack should return a root node with left selected": {
			knapsack: ks_4_0_Knapsack(),
			expectedNode: Node{
				Idx:          0,
				selections:   []selection{SELECTED, SKIPPED, SKIPPED, SKIPPED}[0:1],
				usedCapacity: 4,
				currentValue: 8,
				estimate:     21.75,
			},
		},
		"knapsack with first left element exceeding the capacity should return root node with left skipped": {
			knapsack: &Knapsack{
				Capacity: 7,
				Items: []Item{
					{
						Value:  10,
						Weight: 10,
					},
					{
						Value:  2,
						Weight: 2,
					},
				},
			},
			expectedNode: Node{
				Idx:          0,
				selections:   []selection{SKIPPED, SKIPPED}[0:1],
				usedCapacity: 0,
				estimate:     2.0,
			},
		},
	}

	for testName, testdata := range tests {
		t.Run(testName, func(t *testing.T) {
			verifyNode(t, rootNode(testdata.knapsack), &testdata.expectedNode)
		})
	}
}

func TestComputeOptimumKnapsack(t *testing.T) {
	tests := map[string]struct {
		knapsack         *Knapsack
		expectedOptValue float64
	}{
		"KS 4.0 knapsack should be 19": {
			knapsack:         ks_4_0_Knapsack(),
			expectedOptValue: 19,
		},
		"KS 19.0 knapsack should be 12248": {
			knapsack:         ks_19_0_Knapsack(),
			expectedOptValue: 12248,
		},
		"KS 30.0 knapsack should be 99798": {
			knapsack:         ks_30_0_Knapsack(),
			expectedOptValue: 99798,
		},
		"KS 40.0 knapsack should be 99924": {
			knapsack:         ks_40_0_Knapsack(),
			expectedOptValue: 99924,
		},
		"KS 50.0 knapsack should be 142156": {
			knapsack:         ks_50_0_Knapsack(),
			expectedOptValue: 142156,
		},
		"KS 200.0 knapsack should be 100236": {
			knapsack:         ks_200_0_Knapsack(),
			expectedOptValue: 100236,
		},
		"KS 400.0 knapsack should be 3967180": {
			knapsack:         ks_400_0_Knapsack(),
			expectedOptValue: 3967180,
		},
		"KS 1000.0 knapsack should be 109899": {
			knapsack:         ks_1000_0_Knapsack(),
			expectedOptValue: 109899,
		},
		"KS 10000.0 knapsack should be 1099893": {
			knapsack:         ks_10000_0_Knapsack(),
			expectedOptValue: 1099893,
		},
	}

	for testName, testdata := range tests {
		t.Run(testName, func(t *testing.T) {
			node := ComputeOptimumKnapsack(*testdata.knapsack)
			if node.estimate != testdata.expectedOptValue {
				t.Errorf("Expected optimum value to be %f but was %f", testdata.expectedOptValue, node.estimate)
			}
		})
	}

}

func verifyNode(t *testing.T, actualNode *Node, expectedNode *Node) {
	if actualNode.Idx != expectedNode.Idx {
		t.Errorf("Expected idx to be %d but was %d", expectedNode.Idx, actualNode.Idx)
	}

	if len(actualNode.selections) != len(expectedNode.selections) {
		t.Errorf("Expected node to have %d selections but had %d",
			len(expectedNode.selections), len(actualNode.selections))
	}

	if !reflect.DeepEqual(actualNode.selections, expectedNode.selections) {
		t.Errorf("Expected %v selections but was %v", expectedNode.selections, actualNode.selections)
	}

	if cap(actualNode.selections) != cap(expectedNode.selections) {
		t.Errorf("Expected capacity of node's selection to be %d but was %d",
			cap(expectedNode.selections), cap(actualNode.selections))
	}

	if actualNode.usedCapacity != expectedNode.usedCapacity {
		t.Errorf("Expected usedCapacity of node to be %d but was %d",
			expectedNode.usedCapacity, actualNode.usedCapacity)
	}

	if actualNode.currentValue != expectedNode.currentValue {
		t.Errorf("Expected currentValue of node to be %d but was %d",
			expectedNode.currentValue, actualNode.currentValue)
	}

	if actualNode.estimate != expectedNode.estimate {
		t.Errorf("Expected estimate of node to be %f but was %f", expectedNode.estimate, actualNode.estimate)
	}
}

func BenchmarkComputeOptimumKnapsack_ks_4_0(b *testing.B) {
	benchmarkComputeOptimumKnapsack(b, *ks_4_0_Knapsack())
}

func BenchmarkComputeOptimumKnapsack_ks_19_0(b *testing.B) {
	benchmarkComputeOptimumKnapsack(b, *ks_19_0_Knapsack())
}

func BenchmarkComputeOptimumKnapsack_ks_30_0(b *testing.B) {
	benchmarkComputeOptimumKnapsack(b, *ks_30_0_Knapsack())
}

func BenchmarkComputeOptimumKnapsack_ks_50_0(b *testing.B) {
	benchmarkComputeOptimumKnapsack(b, *ks_50_0_Knapsack())
}

func BenchmarkComputeOptimumKnapsack_ks_200_0(b *testing.B) {
	benchmarkComputeOptimumKnapsack(b, *ks_200_0_Knapsack())
}

func BenchmarkComputeOptimumKnapsack_ks_1000_0(b *testing.B) {
	benchmarkComputeOptimumKnapsack(b, *ks_1000_0_Knapsack())
}

func BenchmarkComputeOptimumKnapsack_ks_10000_0(b *testing.B) {
	benchmarkComputeOptimumKnapsack(b, *ks_10000_0_Knapsack())
}

func benchmarkComputeOptimumKnapsack(b *testing.B, knapsack Knapsack) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ComputeOptimumKnapsack(knapsack)
	}
}
