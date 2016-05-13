package knapsack

type Node struct {
	idx          uint32
	selections   []selection
	usedCapacity uint32
	estimate     float64
}

func (n *Node) IsLeft() bool {
	return n.selections[n.idx] == SELECTED
}

func (n *Node) NextNodes(knapsack *Knapsack) []*Node {
	nodes := []*Node{}
	if n.IsLeft() {
		nodes = append(nodes, n.rightNode(knapsack))
	}

	childIdx := n.idx + 1
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
	selections := append([]selection{}, n.selections...)
	selections[n.idx] = SKIPPED
	return &Node{
		idx:          n.idx,
		selections:   selections,
		usedCapacity: n.usedCapacity - knapsack.Items[n.idx].Weight,
		estimate:     Estimate(knapsack, selections)}
}

func (n *Node) leftChildNode(knapsack *Knapsack) *Node {
	childIdx := n.idx + 1
	selections := n.selections[0 : childIdx+1]
	selections[childIdx] = SELECTED
	return &Node{
		idx:          childIdx,
		selections:   selections,
		usedCapacity: n.usedCapacity + knapsack.Items[childIdx].Weight,
		estimate:     Estimate(knapsack, selections),
	}
}

func (n *Node) rightChildNode(knapsack *Knapsack) *Node {
	childIdx := n.idx + 1
	selections := n.selections[0 : childIdx+1]
	selections[childIdx] = SKIPPED

	return &Node{
		idx:          childIdx,
		selections:   selections,
		usedCapacity: n.usedCapacity,
		estimate:     Estimate(knapsack, selections),
	}
}
