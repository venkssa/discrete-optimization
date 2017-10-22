package knapsack

import (
	"os"
)

func ks_4_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_4_0")
}

func ks_19_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_19_0")
}

func ks_30_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_30_0")
}

func ks_40_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_40_0")
}

func ks_50_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_50_0")
}

func ks_200_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_200_0")
}

func ks_400_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_400_0")
}

func ks_1000_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_1000_0")
}

func ks_10000_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_10000_0")
}

func newKnapsackOrPanicOnFailureFile(path string) *Knapsack {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	knapsack, err := NewKnapsack(file)
	if err != nil {
		panic(err)
	}
	return &knapsack
}
