package knapsack

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestNewKnapsackWithValidInput(t *testing.T) {
	input := `4 11
		15 8
		8 4
		10 5
		4 3`
	knapsack, err := NewKnapsack(ioutil.NopCloser(strings.NewReader(input)))
	if err != nil {
		t.Fatal(err)
	}

	if knapsack.Capacity != 11 {
		t.Errorf("Expected capacity to be 11 but was %d", knapsack.Capacity)
	}

	expectedItems := []Item{
		{
			Idx:    1,
			Value:  8,
			Weight: 4,
		},
		{
			Idx:    2,
			Value:  10,
			Weight: 5,
		},
		{
			Idx:    0,
			Value:  15,
			Weight: 8,
		},
		{
			Idx:    3,
			Value:  4,
			Weight: 3,
		},
	}

	if len(expectedItems) != len(knapsack.Items) {
		t.Errorf("Expected %d items but was %d", len(expectedItems), len(knapsack.Items))
	}

	for idx, expectedItem := range expectedItems {
		if idx >= len(knapsack.Items) {
			t.Errorf("Expected %v but was absent", expectedItem)
			continue
		}
		actualItem := knapsack.Items[idx]
		if expectedItem.Idx != actualItem.Idx {
			t.Errorf("Expected Idx to be %d but was %d", expectedItem.Idx, actualItem.Idx)
		}
		if expectedItem.Weight != actualItem.Weight {
			t.Errorf("Expected weight to be %d but was %d", expectedItem.Weight, actualItem.Weight)
		}
		if expectedItem.Value != actualItem.Value {
			t.Errorf("Expected value to be %d but was %d", expectedItem.Value, actualItem.Value)
		}
	}

	for idx := len(expectedItems); idx < len(knapsack.Items); idx++ {
		t.Errorf("Unexpected item %v", knapsack.Items[idx])
	}
}

func TestNewKnapsackWithInvalidInputs(t *testing.T) {
	tests := []struct {
		input string
		err   error
	}{
		{
			input: ` `,
			err:   fmt.Errorf("Line should contain 2 numbers; but was instead []"),
		},
		{
			input: `4 11
			1 2`,
			err: fmt.Errorf("Expected 4 items but only got 1"),
		},
		{
			input: `4 11
			2
			`,
			err: fmt.Errorf("Line should contain 2 numbers; but was instead [2]"),
		},
		{
			input: `crap 11`,
			err: &strconv.NumError{
				Func: "ParseUint",
				Num:  "crap",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			input: `4 crap`,
			err: &strconv.NumError{
				Func: "ParseUint",
				Num:  "crap",
				Err:  strconv.ErrSyntax,
			},
		},
	}

	for _, test := range tests {
		inputReader := strings.NewReader(test.input)
		_, err := NewKnapsack(ioutil.NopCloser(inputReader))

		if err == nil {
			t.Errorf("Expected an error %v but was nil", test.err)
			continue
		}
		if err.Error() != test.err.Error() {
			t.Errorf("Expected error '%v' but was '%v'", test.err, err)
		}
	}
}
