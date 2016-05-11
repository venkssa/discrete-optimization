package knapsack

type selection byte

const (
	SKIPPED selection = iota
	SELECTED
)

func Estimate(knapsack *Knapsack, selections []selection) float64 {
	estimate := 0.0
	capacityLeft := knapsack.Capacity

	for idx, selection := range selections {
		if selection == SELECTED {
			item := knapsack.Items[idx]
			estimate += float64(item.Value)
			capacityLeft -= item.Weight

		}
	}

	for idx := len(selections); idx < len(knapsack.Items); idx++ {
		item := knapsack.Items[idx]
		if capacityLeft >= item.Weight {
			estimate += float64(item.Value)
			capacityLeft -= item.Weight
		} else {
			estimate += item.ValuePerUnitWeight() * float64(capacityLeft)
			return estimate
		}
	}

	return estimate
}
