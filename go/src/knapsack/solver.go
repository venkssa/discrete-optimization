package knapsack

type node struct {
	idx          uint32
	selections   []selection
	usedCapacity uint32
	estimate     float64
}

func (n *node) IsLeft() bool {
	return n.selections[n.idx] == SELECTED
}

func (n *node) NextNodes(knapsack *Knapsack) []*node {
	nodes := []*node{}
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

func (n *node) rightNode(knapsack *Knapsack) *node {
	selections := append([]selection{}, n.selections...)
	selections[n.idx] = SKIPPED
	return &node{
		idx:          n.idx,
		selections:   selections,
		usedCapacity: n.usedCapacity - knapsack.Items[n.idx].Weight,
		estimate:     Estimate(knapsack, selections)}
}

func (n *node) leftChildNode(knapsack *Knapsack) *node {
	childIdx := n.idx + 1
	selections := n.selections[0 : childIdx+1]
	selections[childIdx] = SELECTED
	return &node{
		idx:          childIdx,
		selections:   selections,
		usedCapacity: n.usedCapacity + knapsack.Items[childIdx].Weight,
		estimate:     Estimate(knapsack, selections),
	}
}

func (n *node) rightChildNode(knapsack *Knapsack) *node {
	childIdx := n.idx + 1
	selections := n.selections[0 : childIdx+1]
	selections[childIdx] = SKIPPED

	return &node{
		idx:          childIdx,
		selections:   selections,
		usedCapacity: n.usedCapacity,
		estimate:     Estimate(knapsack, selections),
	}
}
