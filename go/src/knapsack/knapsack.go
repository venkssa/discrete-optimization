package knapsack

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	Idx    uint32
	Weight uint32
	Value  uint32
}

func (item Item) ValuePerUnitWeight() float64 {
	return float64(item.Value) / float64(item.Weight)
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].ValuePerUnitWeight() > items[j].ValuePerUnitWeight()
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

type Knapsack struct {
	Capacity uint32
	Items    Items
}

func (k *Knapsack) computeEstimate(selections []bool) float64 {
	estimate := 0.0
	capacityLeft := k.Capacity

	for idx, selection := range selections {
		if selection {
			item := k.Items[idx]
			estimate += float64(item.Value)
			capacityLeft -= item.Weight
		}
	}

	for idx, selection := range selections {
		if selection {
			continue
		}

		fmt.Println(selection)
		item := k.Items[idx]
		if capacityLeft >= item.Weight {
			estimate += float64(item.Value)
			capacityLeft -= item.Weight
		} else {
			estimate += item.ValuePerUnitWeight() * float64(capacityLeft)
			break
		}
	}
	return estimate
}

func NewKnapsack(reader io.ReadCloser) (Knapsack, error) {
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	if !scanner.Scan() {
		return Knapsack{}, scanner.Err()
	}

	numberOfItems, capacity, err := splitAndConvertToInt(scanner.Text())
	if err != nil {
		return Knapsack{}, err
	}

	knapsack := Knapsack{Capacity: capacity, Items: Items([]Item{})}
	for idx := uint32(0); idx < numberOfItems; idx++ {
		if !scanner.Scan() {
			break
		}
		value, weight, err := splitAndConvertToInt(scanner.Text())
		if err != nil {
			return knapsack, err
		}
		knapsack.Items = append(knapsack.Items, Item{Idx: idx, Value: value, Weight: weight})
	}

	if uint32(len(knapsack.Items)) != numberOfItems {
		return knapsack, fmt.Errorf("Expected %d items but only got %d", numberOfItems, len(knapsack.Items))
	}
	sort.Sort(knapsack.Items)

	return knapsack, scanner.Err()
}

func splitAndConvertToInt(s string) (uint32, uint32, error) {
	splitStrings := strings.Split(strings.TrimSpace(s), " ")
	if len(splitStrings) != 2 {
		return 0, 0, fmt.Errorf("Line should contain 2 numbers; but was instead %s", splitStrings)
	}
	first, err := strconv.ParseUint(splitStrings[0], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	second, err := strconv.ParseUint(splitStrings[1], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return uint32(first), uint32(second), nil
}
