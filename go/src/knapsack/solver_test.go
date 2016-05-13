package knapsack

import "testing"

func TestNode_IsLeft(t *testing.T) {
	tests := []struct {
		node   Node
		isLeft bool
	}{
		{
			node:   Node{idx: 0, selections: []selection{SELECTED}},
			isLeft: true,
		},
		{
			node:   Node{idx: 0, selections: []selection{SKIPPED}},
			isLeft: false,
		},
	}

	for _, test := range tests {
		if test.node.IsLeft() != test.isLeft {
			t.Errorf("Expected isLeft = %v but was %v (%#v)", test.isLeft, test.node.IsLeft(), test.node)
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
	node := Node{idx: 0, selections: selections[0:1], usedCapacity: 5, estimate: Estimate(knapsack, selections)}

	nextNodes := node.NextNodes(knapsack)

	if len(nextNodes) != 2 {
		t.Fatalf("Expected 2 next node but was %d", len(nextNodes))
	}

	if nextNodes[0].IsLeft() != false {
		t.Error("Expected next node to be a right node but was left.")
	}
	if nextNodes[0].idx != node.idx {
		t.Errorf("Expected idx to be %d but was %d", node.idx, nextNodes[0].idx)
	}
}

func TestNode_NextNodes_ReturnsOneLeftChildNode(t *testing.T) {
	knapsack := &Knapsack{
		Capacity: 7,
		Items: []Item{
			{
				Value: 5,
				Weight: 2,
			},
			{
				Value: 2,
				 Weight: 5,
			},
		},
	}

	selections := []selection{SELECTED, SKIPPED}
	node := Node{idx: 0, selections: selections[0:1], usedCapacity: 2, estimate: Estimate(knapsack, selections)}

	nextNodes := node.NextNodes(knapsack)

	if len(nextNodes) != 2 {
		t.Fatalf("Expected 2 next node but was %d", len(nextNodes))
	}

	if nextNodes[1].IsLeft() == false {
		t.Errorf("Expected %#v to be a left node but was right.", nextNodes[1])
	}

	if nextNodes[1].idx != node.idx+1 {
		t.Errorf("Expected idx to be %d but was %d", node.idx+1, nextNodes[1].idx)
	}
}

func TestNode_NextNodes_ReturnsOneRightChildNode(t *testing.T) {
	knapsack := &Knapsack{
		Capacity: 7,
		Items: []Item{
			{
				Value: 5,
				Weight: 2,
			},
			{
				Value: 2,
				Weight: 6,
			},
		},
	}

	selections := []selection{SELECTED, SKIPPED}
	node := Node{idx: 0, selections: selections[0:1], usedCapacity: 2, estimate: Estimate(knapsack, selections)}

	nextNodes := node.NextNodes(knapsack)

	if len(nextNodes) != 2 {
		t.Fatalf("Expected 2 next node but was %d", len(nextNodes))
	}

	if nextNodes[1].IsLeft() == true {
		t.Errorf("Expected %#v to be a right node but was left.", nextNodes[1])
	}

	if nextNodes[1].idx != node.idx+1 {
		t.Errorf("Expected idx to be %d but was %d", node.idx+1, nextNodes[1].idx)
	}
}

func TestRootNode(t *testing.T) {
	tests := []struct {
		knapsack *Knapsack
		expectedSelectionCap int
		expectedSelections []selection
		expectedUsedCapacity uint32
		expectedEstimate float64
	}{
		{
			knapsack: ks_4_0_Knapsack(),
			expectedSelections: []selection{SELECTED},
			expectedSelectionCap: 4,
			expectedUsedCapacity: 4,
			expectedEstimate: 21.75,
		},
		{
			knapsack: &Knapsack{
				Capacity: 7,
				Items: []Item{
					{
						Value: 10,
						Weight: 10,
					},
					{
						Value: 2,
						Weight: 2,
					},
				},
			},
			expectedSelections: []selection{SKIPPED},
			expectedSelectionCap: 2,
			expectedUsedCapacity: 0,
			expectedEstimate: 2.0,
		},
	}

	for _, test := range tests {
		rootNode := RootNode(test.knapsack)

		if rootNode.idx != 0 {
			t.Errorf("Expected a root node with idx 0 but was %#v", rootNode)
		}

		if len(rootNode.selections) != len(test.expectedSelections) {
			t.Errorf("Expected root node to have %d selection but had %d",
				len(test.expectedSelections), len(rootNode.selections))
		}

		for idx, selection := range test.expectedSelections {
			if selection != rootNode.selections[idx] {
				t.Errorf("Expected root node to have a selection %d at idx %d but was %d",
					selection, idx, rootNode.selections[idx])
			}
		}

		if cap(rootNode.selections) != test.expectedSelectionCap {
			t.Errorf("Expected capacity of root node's selection to be %d but was %d",
				test.expectedSelectionCap, cap(rootNode.selections))
		}

		if rootNode.usedCapacity != test.expectedUsedCapacity {
			t.Errorf("Expected usedCapacity of root node to be %d but was %d",
				test.expectedUsedCapacity, rootNode.usedCapacity)
		}

		if rootNode.estimate != test.expectedEstimate {
			t.Errorf("Expected estimate of root node to be %f but was %f",
				test.expectedEstimate, rootNode.estimate)
		}
	}
}
