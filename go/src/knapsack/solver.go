package knapsack

type node struct {
	idx        uint32
	selections []bool
	estimate   float64
}

func (n *node) IsLeft() bool {
	return n.selections[n.idx]
}

func (n *node) NextNodes(knapsack *Knapsack) []*node {
	nodes := []*node{}
	if n.IsLeft() {
		selectionsForRight := append([]bool{}, n.selections...)
		selectionsForRight[n.idx] = false
		rightNode := &node{
			idx:        n.idx,
			selections: selectionsForRight,
			estimate:   knapsack.computeEstimate(selectionsForRight)}
		nodes = append(nodes, rightNode)
	}

	return nodes
}
