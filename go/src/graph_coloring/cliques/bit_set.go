package cliques

import (
	"math"
)

const bitsPerWord = 64

type block uint64

type BitSet struct {
	blocks        []block
	numOfElements uint32
}

func NewBitSet(numOfElements uint32) *BitSet {
	size := numOfElements / bitsPerWord
	if numOfElements%bitsPerWord != 0 {
		size++
	}
	return &BitSet{blocks: make([]block, size), numOfElements: numOfElements}
}

func (bs *BitSet) Set(idx uint32) {
	bs.blocks[idx/bitsPerWord] |= 1 << (idx % bitsPerWord)
}

func (bs *BitSet) UnSet(idx uint32) {
	bs.blocks[idx/bitsPerWord] &= ^(1 << (idx % bitsPerWord))
}

func (bs *BitSet) IsSet(idx uint32) bool {
	return (bs.blocks[idx/bitsPerWord] & (1 << (idx % bitsPerWord))) != 0
}

func (bs *BitSet) IsZero() bool {
	for _, block := range bs.blocks {
		if block != 0 {
			return false
		}
	}
	return true
}

func (bs *BitSet) String() string {
	bitSetAsRune := make([]rune, bs.numOfElements)
	for idx := uint32(0); idx < bs.numOfElements; idx++ {
		if bs.IsSet(idx) {
			bitSetAsRune[idx] = '1'
		} else {
			bitSetAsRune[idx] = '0'
		}
	}
	return string(bitSetAsRune)
}

func And(result *BitSet, first *BitSet, second *BitSet) {
	for idx := 0; idx < len(result.blocks); idx++ {
		result.blocks[idx] = first.blocks[idx] & second.blocks[idx]
	}
}
