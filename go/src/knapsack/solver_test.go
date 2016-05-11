package knapsack

import "testing"

func TestNode_IsLeft(t *testing.T) {
	tests := []struct {
		node   node
		isLeft bool
	}{
		{
			node:   node{idx: 0, selections: []bool{true, false, false, false}},
			isLeft: true,
		},
		{
			node:   node{idx: 0, selections: []bool{false, false, false, false}},
			isLeft: false,
		},
	}

	for _, test := range tests {
		if test.node.IsLeft() != test.isLeft {
			t.Errorf("Expected isLeft = %v but was %v (%#v)", test.isLeft, test.node.IsLeft(), test.node)
		}
	}
}

func TestNode_NextNodesForLeftGivesRightNode(t *testing.T) {
	knapsack := newKnapsackOrPanicOnFailure(`2 5
	1 5
	2 5`)

	selections := []bool{true, false}
	node := node{idx: 0, selections: selections, estimate: knapsack.computeEstimate(selections)}

	nextNodes := node.NextNodes(knapsack)

	if len(nextNodes) != 1 {
		t.Errorf("Expected 1 next node but was %d", len(nextNodes))
	}
}
