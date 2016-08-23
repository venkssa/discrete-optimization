package cliques

import (
	"testing"
)

func TestBitCount(t *testing.T) {
	tests := []uint32{1, 4, 10, 20, 50, 64}

	for _, test := range tests {

		bs := NewBitSet(64)
		for idx := uint32(0); idx < test; idx++ {
			bs.Set(idx)
		}

		res := bs.blocks[0].BitCount()

		if res != test {
			t.Errorf("Expected %d but ws %d", test, res)
		}
	}
}

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

func TestBitSet_NumOfBitsSet(t *testing.T) {
	bs := bitsetForTest(3, 0, 2)

	if numbersOfBitsSet := bs.NumOfBitsSet(); numbersOfBitsSet != 2 {
		t.Errorf("Expected number of bits to be set = 2 but was %v", numbersOfBitsSet)
	}
}

func TestBitSet_IsZero(t *testing.T) {
	tests := []struct {
		input    *BitSet
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

func TestBitSet_Intersection(t *testing.T) {
	first := bitsetForTest(3, 0, 2)
	second := bitsetForTest(3, 0)

	resultIntersectionFn := NewBitSet(3)
	Intersection(resultIntersectionFn, first, second)

	if resultIntersectionFn.String() != "100" {
		t.Errorf("Expected %v AND %v to be 100 but was %v", first, second, resultIntersectionFn)
	}

	resultIntersectionMethod := first.Intersection(second)

	if resultIntersectionMethod.String() != "100" {
		t.Errorf("Expected %v AND %v to be 100 but was %v", first, second, resultIntersectionMethod)
	}
}

func TestBitSet_Minus(t *testing.T) {
	tests := []struct {
		first    *BitSet
		second   *BitSet
		expected *BitSet
	}{
		{
			first:    bitsetForTest(3, 0, 1, 2),
			second:   bitsetForTest(3, 1),
			expected: bitsetForTest(3, 0, 2),
		},
		{
			first:    bitsetForTest(4, 2, 3),
			second:   bitsetForTest(4, 1),
			expected: bitsetForTest(4, 2, 3),
		},
	}

	for _, test := range tests {
		resultMinusFn := NewBitSet(test.first.numOfElements)

		Minus(resultMinusFn, test.first, test.second)

		if resultMinusFn.String() != test.expected.String() {
			t.Errorf("Expected %v - %v to be %v but was %v",
				test.first, test.second, test.expected, resultMinusFn)
		}

		resultMinusMethod := test.first.Minus(test.second)

		if resultMinusMethod.String() != test.expected.String() {
			t.Errorf("Expected %v - %v to be %v but was %v",
				test.first, test.second, test.expected, resultMinusMethod)
		}

	}
}

func TestBitSet_Union(t *testing.T) {
	first := bitsetForTest(3, 1)
	second := bitsetForTest(3, 0, 2)

	resultUnionFn := NewBitSet(3)

	Union(resultUnionFn, first, second)

	if resultUnionFn.String() != "111" {
		t.Errorf("Expected %v U %v to be 111 but was %v", first, second, resultUnionFn)
	}

	resultUnionMethod := first.Union(second)

	if resultUnionMethod.String() != "111" {
		t.Errorf("Expected %v U %v to be 111 but was %v", first, second, resultUnionMethod)
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
		Intersection(result, neighbors, p)
		v := uint32(idx % 500)
		p.UnSet(v)
		p.Set(v)
	}
}

func BenchmarkBitSet_(b *testing.B) {
	neighbors := NewBitSet(1000)
	p := NewBitSet(1000)

	b.ResetTimer()
	b.ReportAllocs()

	for idx := 0; idx < b.N; idx++ {
		neighbors.Intersection(p)
		v := uint32(idx % 500)
		p.UnSet(v)
		p.Set(v)
	}
}
