package knapsack

import (
	"reflect"
	"testing"
)

func TestNode_IsLeft(t *testing.T) {
	tests := []struct {
		node   Node
		isLeft bool
	}{
		{
			node:   Node{Idx: 0, selections: []selection{SELECTED}},
			isLeft: true,
		},
		{
			node:   Node{Idx: 0, selections: []selection{SKIPPED}},
			isLeft: false,
		},
	}

	for _, test := range tests {
		if test.node.isLeft() != test.isLeft {
			t.Errorf("Expected isLeft = %v but was %v (%#v)", test.isLeft, test.node.isLeft(), test.node)
		}
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
	node := Node{Idx: 0, selections: selections[0:1], usedCapacity: 5, estimate: Estimate(knapsack, selections)}

	nextNodes := node.NextNodes(knapsack)

	if len(nextNodes) != 2 {
		t.Fatalf("Expected 2 next node but was %d", len(nextNodes))
	}

	if nextNodes[0].isLeft() != false {
		t.Error("Expected next node to be a right node but was left.")
	}
	if nextNodes[0].Idx != node.Idx {
		t.Errorf("Expected idx to be %d but was %d", node.Idx, nextNodes[0].Idx)
	}
}

func TestNode_NextNodes_ReturnsTheCorrectChildNode(t *testing.T) {
	tests := []struct {
		capacity     uint32
		expectedNode Node
	}{
		{
			capacity: 7,
			expectedNode: Node{
				Idx:          1,
				selections:   []selection{SELECTED, SELECTED},
				usedCapacity: 7,
				estimate:     7.0,
			},
		},
		{
			capacity: 6,
			expectedNode: Node{
				Idx:          1,
				selections:   []selection{SELECTED, SKIPPED},
				usedCapacity: 2,
				estimate:     5.0,
			},
		},
	}

	for _, test := range tests {
		knapsack := &Knapsack{
			Capacity: test.capacity,
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
		node := Node{Idx: 0, selections: selections[0:1], usedCapacity: 2,
			estimate: Estimate(knapsack, selections)}

		nextNodes := node.NextNodes(knapsack)

		if len(nextNodes) != 2 {
			t.Fatalf("Expected 2 next node but was %d", len(nextNodes))
		}

		verifyNode(t, nextNodes[1], &test.expectedNode)
	}
}

func TestRootNode(t *testing.T) {
	tests := []struct {
		knapsack     *Knapsack
		expectedNode Node
	}{
		{
			knapsack: ks_4_0_Knapsack(),
			expectedNode: Node{
				Idx:          0,
				selections:   []selection{SELECTED, SKIPPED, SKIPPED, SKIPPED}[0:1],
				usedCapacity: 4,
				estimate:     21.75,
			},
		},
		{
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

	for _, test := range tests {
		verifyNode(t, rootNode(test.knapsack), &test.expectedNode)
	}
}

func TestComputeOptimumKnapsack(t *testing.T) {
	tests := []struct {
		knapsack         *Knapsack
		expectedOptValue float64
	}{
		{
			knapsack:         ks_4_0_Knapsack(),
			expectedOptValue: 19,
		},
		{
			knapsack:         ks_19_0_Knapsack(),
			expectedOptValue: 12248,
		},
		{
			knapsack:         ks_30_0_Knapsack(),
			expectedOptValue: 99798,
		},
		{
			knapsack:         ks_40_0_Knapsack(),
			expectedOptValue: 99924,
		},
		{
			knapsack:         ks_50_0_Knapsack(),
			expectedOptValue: 142156,
		},
		{
			knapsack:         ks_200_0_Knapsack(),
			expectedOptValue: 100236,
		},
		{
			knapsack:         ks_400_0_Knapsack(),
			expectedOptValue: 3967180,
		},
		{
			knapsack:         ks_1000_0_Knapsack(),
			expectedOptValue: 109899,
		},
		{
			knapsack:         ks_10000_0_Knapsack(),
			expectedOptValue: 1099893,
		},
	}

	for _, test := range tests {
		node := ComputeOptimumKnapsack(*test.knapsack)
		if node.estimate != test.expectedOptValue {
			t.Errorf("Expected optimum value to be %f but was %f", test.expectedOptValue, node.estimate)
		}

		t.Log(node)
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

	if actualNode.estimate != expectedNode.estimate {
		t.Errorf("Expected estimate of node to be %f but was %f", expectedNode.estimate, actualNode.estimate)
	}
}
