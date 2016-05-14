package knapsack

func ComputeOptimumKnapsack(knapsack Knapsack) *Node {
	queue := []*Node{rootNode(&knapsack)}
	var optNode *Node

	for len(queue) != 0 {
		lastIdx := len(queue) - 1
		node := queue[lastIdx]
		nextNodes := node.NextNodes(&knapsack)

		queue = queue[0:lastIdx:lastIdx]

		for _, nextNode := range nextNodes {
			if optNode == nil {
				queue = append(queue, nextNode)
			} else if nextNode.estimate > optNode.estimate {
				queue = append(queue, nextNode)
			}
		}

		if len(node.selections) != len(knapsack.Items) {
			continue
		}

		if optNode == nil || node.estimate > optNode.estimate {
			optNode = node
		}
	}
	return optNode
}

type Node struct {
	Idx          uint32
	selections   []selection
	usedCapacity uint32
	estimate     float64
}

func rootNode(knapsack *Knapsack) *Node {
	selections := make([]selection, len(knapsack.Items))
	usedCapacity := uint32(0)
	if knapsack.Items[0].Weight <= knapsack.Capacity {
		selections[0] = SELECTED
		usedCapacity = knapsack.Items[0].Weight
	}

	return &Node{
		Idx:          0,
		selections:   selections[0:1],
		usedCapacity: usedCapacity,
		estimate:     Estimate(knapsack, selections[0:1]),
	}
}

func (n *Node) isLeft() bool {
	return n.selections[n.Idx] == SELECTED
}

func (n *Node) NextNodes(knapsack *Knapsack) []*Node {
	nodes := []*Node{}
	if n.isLeft() {
		nodes = append(nodes, n.rightNode(knapsack))
	}

	childIdx := n.Idx + 1
	if childIdx < uint32(len(knapsack.Items)) {
		remainingCapacity := knapsack.Capacity - n.usedCapacity
		if remainingCapacity >= knapsack.Items[childIdx].Weight {
			nodes = append(nodes, n.leftChildNode(knapsack))
		} else {
			nodes = append(nodes, n.rightChildNode(knapsack))
		}
	}

	return nodes
}

func (n *Node) rightNode(knapsack *Knapsack) *Node {
	selections := make([]selection, len(knapsack.Items))
	numElemCopied := copy(selections, n.selections)
	selections[n.Idx] = SKIPPED
	selections = selections[0:numElemCopied]
	return &Node{
		Idx:          n.Idx,
		selections:   selections,
		usedCapacity: n.usedCapacity - knapsack.Items[n.Idx].Weight,
		estimate:     Estimate(knapsack, selections)}
}

func (n *Node) leftChildNode(knapsack *Knapsack) *Node {
	childIdx := n.Idx + 1
	selections := n.selections[0 : childIdx+1]
	selections[childIdx] = SELECTED
	return &Node{
		Idx:          childIdx,
		selections:   selections,
		usedCapacity: n.usedCapacity + knapsack.Items[childIdx].Weight,
		estimate:     Estimate(knapsack, selections),
	}
}

func (n *Node) rightChildNode(knapsack *Knapsack) *Node {
	childIdx := n.Idx + 1
	selections := n.selections[0 : childIdx+1]
	selections[childIdx] = SKIPPED

	return &Node{
		Idx:          childIdx,
		selections:   selections,
		usedCapacity: n.usedCapacity,
		estimate:     Estimate(knapsack, selections),
	}
}
