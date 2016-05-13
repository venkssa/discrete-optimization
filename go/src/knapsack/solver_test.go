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
