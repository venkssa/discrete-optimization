package cliques

import (
	"testing"
)

func TestBitSet_Set(t *testing.T) {
	idxsToSet := []uint32{0, 63}
	bs := bitsetForTest(64, idxsToSet...)

	for _, expectedIdxToBeSet := range idxsToSet {
		if bs.IsSet(expectedIdxToBeSet) != true {
			t.Errorf("Expected %v to be set but was not. (%v)", expectedIdxToBeSet, bs)
		}
	}
}

func TestBitSet_UnSet(t *testing.T) {
	bs := bitsetForTest(2, 1)
	bs.UnSet(1)

	if bs.IsSet(1) != false {
		t.Errorf("Expected idx 1 to be unset but was not. (%v)", bs)
	}
}

func TestBitSet_IsZero(t *testing.T) {
	tests := []struct{
		input *BitSet
		expected bool
	}{
		{
			bitsetForTest(2),
			true,
		},
		{
			bitsetForTest(2, 0, 1),
			false,
		},
	}

	for _, test := range tests {
		if test.input.IsZero() != test.expected {
			t.Errorf("Expected IsZero = %v but was %v (%v)", test.expected, test.input.IsZero(),
				test.input)
		}
	}

}

func TestBitSet_String(t *testing.T) {
	bs := bitsetForTest(3, 0, 2)

	if bs.String() != "101" {
		t.Errorf("Expected bitset as string to be 101 but was %v", bs)
	}
}

func TestBitSet_And(t *testing.T) {
	first := bitsetForTest(3, 0, 2)
	second := bitsetForTest(3, 0)

	result := NewBitSet(3)
	And(result, first, second)

	if result.String() != "100" {
		t.Errorf("Expected %v AND %v to be 100 but was %v", first, second, result)
	}
}

func bitsetForTest(numOfElements uint32, idxsToSet ...uint32) *BitSet {
	bs := NewBitSet(numOfElements)

	for _, idx := range idxsToSet {
		bs.Set(idx)
	}

	return bs
}

func BenchmarkBitSet(b *testing.B) {
	neighbors := NewBitSet(1000)
	p := NewBitSet(1000)

	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		result := NewBitSet(1000)
		And(result, neighbors, p)
		v := uint32(idx % 500)
		p.UnSet(v)
		p.Set(v)
	}
}
