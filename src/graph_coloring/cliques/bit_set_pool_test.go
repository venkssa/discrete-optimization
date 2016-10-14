package cliques

import "testing"

func TestBitSetPool_Borrow_FromEmptyPool(t *testing.T) {
	pool := NewBitSetPool(2)

	bs := pool.Borrow()
	if bs.Len() != 2 {
		t.Errorf("Expected bit set of len %v but was %v", 2, bs.Len())
	}

	if len(pool.available) != 0 {
		t.Errorf("Expected the pool to be empty but was %v", pool.available)
	}
}

func TestBitSetPool_Borrow_ReturnToPool(t *testing.T) {
	pool := NewBitSetPool(2)

	bs := pool.Borrow()

	if len(pool.available) != 0 {
		t.Errorf("Expected the pool to be empty but was %v", pool.available)
	}

	pool.Return(bs)
	if len(pool.available) != 1 {
		t.Errorf("Expected the pool to have 1 item but was %v", pool.available)
	}
}

func TestBitSetPool_Borrow_ReturnToPoolAndBorrowAgainShouldMakePoolEmpty(t *testing.T) {
	pool := NewBitSetPool(2)

	bs := pool.Borrow()
	if len(pool.available) != 0 {
		t.Errorf("Expected the pool to be empty but was %v", pool.available)
	}

	pool.Return(bs)
	if len(pool.available) != 1 {
		t.Errorf("Expected the pool to have 1 item but was %v", pool.available)
	}

	pool.Borrow()
	if len(pool.available) != 0 {
		t.Errorf("Expected the pool to be empty but was %v", pool.available)
	}
}

func BenchmarkBitSetPool(b *testing.B) {
	pool := NewBitSetPool(250)
	for i := 0; i < b.N; i++ {
		for j := 0; j <= 10; j++ {
			bs := pool.Borrow()
			pool.Return(bs)
		}
	}
}
