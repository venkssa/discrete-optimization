package knapsack

func ComputeOptimumKnapsack(knapsack Knapsack) *Node {
	queue := newQueue(rootNode(&knapsack))

	for !queue.IsEmpty() {
		node := queue.Pop()

		rightNode, childNode := node.NextNodes(&knapsack)

		queue.Add(rightNode)
		queue.Add(childNode)

		if len(node.selections) == len(knapsack.Items) {
			queue.UpdateOptNode(node)
		}
	}

	return queue.optNode
}

type Node struct {
	Idx          uint32
	selections   []selection
	usedCapacity uint32
	currentValue uint32
	estimate     float64
}

func rootNode(knapsack *Knapsack) *Node {
	selections := make([]selection, len(knapsack.Items))
	var usedCapacity uint32
	var currentValue uint32

	if knapsack.Items[0].Weight <= knapsack.Capacity {
		selections[0] = SELECTED
		usedCapacity = knapsack.Items[0].Weight
		currentValue = knapsack.Items[0].Value
	}

	return &Node{
		Idx:          0,
		selections:   selections[0:1],
		usedCapacity: usedCapacity,
		currentValue: currentValue,
		estimate:     estimate(knapsack, usedCapacity, currentValue, selections[0:1]),
	}
}

func (n *Node) isLeft() bool {
	return n.selections[n.Idx] == SELECTED
}

func (n *Node) NextNodes(knapsack *Knapsack) (*Node, *Node) {
	var rightNode *Node
	if n.isLeft() {
		rightNode = n.rightNode(knapsack)
	}

	childIdx := n.Idx + 1
	if childIdx < uint32(len(knapsack.Items)) {
		remainingCapacity := knapsack.Capacity - n.usedCapacity
		if remainingCapacity >= knapsack.Items[childIdx].Weight {
			return rightNode, n.leftChildNode(knapsack)
		} else {
			return rightNode, n.rightChildNode(knapsack)
		}
	}

	return rightNode, nil
}

func (n *Node) rightNode(knapsack *Knapsack) *Node {
	selections := make([]selection, len(knapsack.Items))
	numElemCopied := copy(selections, n.selections)
	selections[n.Idx] = SKIPPED
	selections = selections[0:numElemCopied]
	usedCapacity := n.usedCapacity - knapsack.Items[n.Idx].Weight
	currentValue := n.currentValue - knapsack.Items[n.Idx].Value

	return &Node{
		Idx:          n.Idx,
		selections:   selections,
		usedCapacity: usedCapacity,
		currentValue: currentValue,
		estimate:     estimate(knapsack, usedCapacity, currentValue, selections)}
}

func (n *Node) leftChildNode(knapsack *Knapsack) *Node {
	childIdx := n.Idx + 1
	selections := n.selections[0 : childIdx+1]
	selections[childIdx] = SELECTED
	usedCapacity := n.usedCapacity + knapsack.Items[childIdx].Weight
	currentValue := n.currentValue + knapsack.Items[childIdx].Value

	return &Node{
		Idx:          childIdx,
		selections:   selections,
		usedCapacity: usedCapacity,
		currentValue: currentValue,
		estimate:     estimate(knapsack, usedCapacity, currentValue, selections),
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
		currentValue: n.currentValue,
		estimate:     estimate(knapsack, n.usedCapacity, n.currentValue, selections),
	}
}

type selection byte

const (
	SKIPPED selection = iota
	SELECTED
)

func estimate(knapsack *Knapsack, usedCapacity uint32, currentValue uint32, selections []selection) float64 {
	estimate := float64(currentValue)
	capacityLeft := knapsack.Capacity - usedCapacity

	for idx := len(selections); idx < len(knapsack.Items); idx++ {
		item := knapsack.Items[idx]
		if capacityLeft >= item.Weight {
			estimate += float64(item.Value)
			capacityLeft -= item.Weight
		} else {
			return estimate + item.ValuePerUnitWeight * float64(capacityLeft)
		}
	}
	return estimate
}

type queue struct {
	nodes   []*Node
	lastIdx int
	optNode *Node
}

func newQueue(rootNode *Node) *queue {
	return &queue{
		nodes:   []*Node{rootNode},
		lastIdx: 0,
	}
}

func (q *queue) Add(node *Node) {
	if node == nil {
		return
	}

	if q.lastIdx == len(q.nodes)-1 {
		nodes := make([]*Node, len(q.nodes) * 2)
		copy(nodes, q.nodes)
		q.nodes = nodes
	}
	if q.optNode == nil || node.estimate > q.optNode.estimate {
		q.lastIdx++
		q.nodes[q.lastIdx] = node
	}
}

func (q *queue) UpdateOptNode(node *Node) {
	if q.optNode == nil || node.estimate > q.optNode.estimate {
		q.optNode = node
	}
}

func (q *queue) IsEmpty() bool {
	return q.lastIdx < 0
}

func (q *queue) Pop() *Node {
	lastNode := q.nodes[q.lastIdx]
	q.nodes[q.lastIdx] = nil
	q.lastIdx--
	return lastNode
}
