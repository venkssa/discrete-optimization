package cliques

type BitSetPool struct {
	bitSetSize uint32
	available []*BitSet
	borrowCount uint32
	returnCount uint32
}

func NewBitSetPool(bitSetSize uint32) *BitSetPool {
	return &BitSetPool{bitSetSize: bitSetSize}
}

// Returns a bitset from the pool or creates a new one if the pool is empty.
// This method is not safe for concurrent access.
func (p *BitSetPool) Borrow() *BitSet {
	p.borrowCount++
	if len(p.available) == 0 {
		return NewBitSet(p.bitSetSize)
	}

	lastIdx := len(p.available) - 1
	bs := p.available[lastIdx]
	p.available = p.available[0:lastIdx]
	return bs
}

// Returns the bitset to the pool.
// This method is not safe for concurrent access.
func (p *BitSetPool) Return(bs *BitSet) {
	p.returnCount++
	p.available = append(p.available, bs)
}

func (p *BitSetPool) Stats() map[string]uint32 {
	return map[string]uint32 {
		"BorrowCount": p.borrowCount,
		"ReturnCount": p.returnCount,
		"AvailablePoolSize": uint32(len(p.available)),
	}
}
