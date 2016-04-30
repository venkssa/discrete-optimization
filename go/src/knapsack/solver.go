package knapsack

func computeEstimate(knapsack Knapsack, selections ...bool) float64 {
	estimate := 0.0
	capacityLeft := knapsack.Capacity

	for idx, selection := range selections {
		if selection {
			item := knapsack.Items[idx]
			estimate += float64(item.Value)
			capacityLeft -= item.Weight
		}
	}

	for idx, selection := range selections {
		if !selection {
			item := knapsack.Items[idx]
			if capacityLeft >= item.Weight {
				estimate += float64(item.Value)
				capacityLeft -= item.Weight
			} else {
				estimate += item.ValuePerUnitWeight() * float64(capacityLeft)
				break
			}
		}
	}
	return estimate
}
