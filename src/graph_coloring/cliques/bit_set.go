package cliques

const bitsPerWord = 64

type block uint64

// I have no clue what this code does and how it computes number of bits set.
func (blk block) BitCount() uint32 {
	blk = blk - ((blk >> 1) & 0x5555555555555555)
	blk = (blk & 0x3333333333333333) + ((blk >> 2) & 0x3333333333333333)
	blk = (blk + (blk >> 4)) & 0x0f0f0f0f0f0f0f0f
	blk = blk + (blk >> 8)
	blk = blk + (blk >> 16)
	blk = blk + (blk >> 32)
	return uint32(blk) & 0x7f
}

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

func (bs *BitSet) Len() uint32 {
	return bs.numOfElements
}

func (bs *BitSet) NumOfBitsSet() uint32 {
	var count uint32
	for _, blk := range bs.blocks {
		if blk != 0 {
			count += blk.BitCount()
		}
	}
	return count
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

func (bs *BitSet) Minus(second *BitSet) *BitSet {
	result := NewBitSet(bs.Len())
	Minus(result, bs, second)
	return result
}

func (bs *BitSet) Union(second *BitSet) *BitSet {
	result := NewBitSet(bs.Len())
	Union(result, bs, second)
	return result
}

func (bs *BitSet) Intersection(second *BitSet) *BitSet {
	result := NewBitSet(bs.Len())
	Intersection(result, bs, second)
	return result
}

func (bs *BitSet) Not() *BitSet {
	result := NewBitSet(bs.Len())
	Not(result, bs)
	return result
}

func (bs *BitSet) LoopOverSetIndices(fn func(setIdx uint32)) {
	lastIdxInBlock := uint32(bitsPerWord)
	for blockIdx, block := range bs.blocks {
		if block != 0 {
			if blockIdx == len(bs.blocks)-1 {
				lastIdxInBlock = bs.numOfElements % bitsPerWord
			}

			for idx := uint32(0); idx < lastIdxInBlock; idx++ {
				if (block & (1 << idx)) == 0 {
					continue
				}

				vIdx := idx + uint32(blockIdx*bitsPerWord)
				fn(vIdx)
			}
		}
	}
}

func Intersection(result *BitSet, first *BitSet, second *BitSet) {
	for idx := 0; idx < len(result.blocks); idx++ {
		result.blocks[idx] = first.blocks[idx] & second.blocks[idx]
	}
}

func Minus(result *BitSet, first *BitSet, second *BitSet) {
	for idx := 0; idx < len(result.blocks); idx++ {
		result.blocks[idx] = first.blocks[idx] & (first.blocks[idx] ^ second.blocks[idx])
	}
}

func Union(result *BitSet, first *BitSet, second *BitSet) {
	for idx := 0; idx < len(result.blocks); idx++ {
		result.blocks[idx] = first.blocks[idx] | second.blocks[idx]
	}
}

func Not(result *BitSet, first *BitSet) {
	for idx := 0; idx < len(result.blocks); idx++ {
		result.blocks[idx] = ^first.blocks[idx]
	}
}
