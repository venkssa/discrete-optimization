package cliques

func bronKerboschMaximalClique(r []uint32,
			p *bitSet,
			x *bitSet,
			vertexToEdgeBitSet []*bitSet,
			result [][]uint32) [][]uint32 {
	if p.IsZero() && x.IsZero() {
		return append(result, append([]uint32{}, r...))
	}

	numOfVertices := uint32(len(vertexToEdgeBitSet))
	pCopy := newBitSet(numOfVertices)
	xCopy := newBitSet(numOfVertices)

	for v := uint32(0); v < numOfVertices; v++ {
		if !p.IsSet(v) {
			continue
		}

		neighbors := vertexToEdgeBitSet[v]
		and(pCopy, neighbors, p)
		and(xCopy, neighbors, x)

		result = bronKerboschMaximalClique(append(r, v), pCopy, xCopy, vertexToEdgeBitSet, result)

		p.UnSet(v)
		x.Set(v)
	}
	return result
}
