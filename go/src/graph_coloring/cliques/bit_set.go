package cliques

import "math"

const bitsPerWord = 64

const maxBlock = math.MaxUint64

type block uint64

type bitSet struct {
	blocks []block
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

func (bs *bitSet) IsZero() bool {
	for _, block := range bs.blocks {
		if block != 0 {
			return false
		}
	}
	return true
}

func and(result *bitSet, first *bitSet, second *bitSet) {
	for idx := 0; idx < len(result.blocks); idx++ {
		result.blocks[idx] = first.blocks[idx] & second.blocks[idx]
	}
}
