package cliques

import (
	"graph_coloring/graph"
	"math"
)

func FindAllMaximalCliques(graph *graph.G) [][]uint32 {
	p := newBitSet(graph.NumOfVertices)
	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		p.Set(idx)
	}
	return bkAlgo(make([]uint32, 0, graph.NumOfVertices), p, newBitSet(graph.NumOfVertices),
		neighborsBitSet(graph), [][]uint32{})
}

func bkAlgo(r []uint32, p *bitSet, x *bitSet, vertexToEdgeBitSet []*bitSet, result [][]uint32) [][]uint32 {
	if p.IsZero() && x.IsZero() {
		return append(result, append([]uint32{}, r...))
	}

	pCopy := newBitSet(uint32(len(vertexToEdgeBitSet)))
	xCopy := newBitSet(uint32(len(vertexToEdgeBitSet)))

	for v := uint32(0); v < uint32(len(vertexToEdgeBitSet)); v++ {
		if !p.IsSet(v) {
			continue
		}

		neighbors := vertexToEdgeBitSet[v]
		pCopy.And(neighbors, p)
		xCopy.And(neighbors, x)

		result = bkAlgo(append(r, v), pCopy, xCopy, vertexToEdgeBitSet, result)

		p.UnSet(v)
		x.Set(v)
	}
	return result
}

func neighborsBitSet(graph *graph.G) []*bitSet {
	vertexToEdgeBitSet := make([]*bitSet, graph.NumOfVertices)

	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		bs := newBitSet(graph.NumOfVertices)

		for _, neighbor := range graph.VertexToEdges[int(idx)] {
			bs.Set(neighbor)
		}
		vertexToEdgeBitSet[idx] = bs
	}
	return vertexToEdgeBitSet
}

const bitsPerWord = 64

const maxBlock = math.MaxUint64

type block uint64

type bitSet struct {
	blocks []block
}

func newBitSetLike(bs *bitSet) *bitSet {
	return &bitSet{blocks: make([]block, len(bs.blocks))}
}

func newBitSet(numOfElements uint32) *bitSet {
	size := numOfElements / bitsPerWord
	if numOfElements%bitsPerWord != 0 {
		size++
	}
	return &bitSet{blocks: make([]block, size)}
}

func (bs *bitSet) Set(idx uint32) {
	bs.blocks[idx/bitsPerWord] |= 1 << (idx % bitsPerWord)
}

func (bs *bitSet) UnSet(idx uint32) {
	bs.blocks[idx/bitsPerWord] &= ^(1 << (idx % bitsPerWord))
}

func (bs *bitSet) IsSet(idx uint32) bool {
	return (bs.blocks[idx/bitsPerWord] & (1 << (idx % bitsPerWord))) != 0
}

func (bs *bitSet) Clone() *bitSet {
	return &bitSet{blocks: append([]block{}, bs.blocks...)}
}
func (bs *bitSet) IsZero() bool {
	for _, block := range bs.blocks {
		if block != 0 {
			return false
		}
	}
	return true
}

func (bs *bitSet) And(first *bitSet, second *bitSet) {
	for idx := 0; idx < len(bs.blocks); idx++ {
		bs.blocks[idx] = first.blocks[idx] & second.blocks[idx]
	}

}
