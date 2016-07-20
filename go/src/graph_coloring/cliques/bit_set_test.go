package cliques

import "testing"

func TestBitSet_Set(t *testing.T) {
	bs := newBitSet(64)

	bs.Set(0)
	bs.Set(1)
	bs.Set(63)
	bs.UnSet(1)
	bs.UnSet(0)

	if bs.blocks[0] != 1<<(bitsPerWord-1) {
		t.Error(bs)
	}

	if bs.IsSet(63) != true {
		t.Error(bs)
	}

	if bs.IsSet(0) == true {
		t.Error(bs)
	}
}

func BenchmarkBitSet(b *testing.B) {
	neighbors := newBitSet(50)
	neighbors.Set(1)
	neighbors.Set(3)
	neighbors.Set(5)
	neighbors.Set(49)

	p := &bitSet{blocks: []block{maxBlock}}

	result := newBitSet(50)
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		and(result, neighbors, p)
	}
}
