package knapsack

import (
	"fmt"
	"io"
	"sort"

	"github.com/venkssa/discrete-optimization/common"
)

type Item struct {
	Idx                uint32
	Weight             uint32
	Value              uint32
	ValuePerUnitWeight float64
}

func NewItem(idx uint32, value uint32, weight uint32) Item {
	return Item{
		Idx:                idx,
		Value:              value,
		Weight:             weight,
		ValuePerUnitWeight: float64(value) / float64(weight),
	}
}

type Items []Item

type Knapsack struct {
	Capacity uint32
	Items    Items
}

func NewKnapsack(rc io.ReadCloser) (Knapsack, error) {
	var numberOfItems uint32
	knapsack := Knapsack{}

	err := common.Parse(rc, func(line common.LineNum, d1 uint32, d2 uint32) {
		if line == 1 {
			numberOfItems = d1
			knapsack.Capacity = d2
		} else {
			knapsack.Items = append(knapsack.Items, NewItem(uint32(line)-2, d1, d2))
		}
	})

	if err != nil {
		return knapsack, err
	}

	if uint32(len(knapsack.Items)) != numberOfItems {
		return knapsack, fmt.Errorf("Expected %d items but only got %d", numberOfItems, len(knapsack.Items))
	}

	sort.Slice(knapsack.Items, func(i, j int) bool {
		return knapsack.Items[i].ValuePerUnitWeight > knapsack.Items[j].ValuePerUnitWeight
	})

	return knapsack, err
}
