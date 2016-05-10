package knapsack

import (
	"io/ioutil"
	"strings"
)

func ks_4_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailure(`4 11
		8 4
		10 5
		15 8
		4 3`)
}

func ks_19_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailure(`19 31181
		1945 4990
		321 1142
		2945 7390
		4136 10372
		1107 3114
		1022 2744
		1101 3102
		2890 7280
		962 2624
		1060 3020
		805 2310
		689 2078
		1513 3926
		3878 9656
		13504 32708
		1865 4830
		667 2034
		1833 4766
		16553 40006`)
}

func ks_50_0_Knapsack() *Knapsack {
return newKnapsackOrPanicOnFailure(`50 341045
		1906 4912
		41516 99732
		23527 56554
		559 1818
		45136 108372
		2625 6750
		492 1484
		1086 3072
		5516 13532
		4875 12050
		7570 18440
		4436 10972
		620 1940
		50897 122094
		2129 5558
		4265 10630
		706 2112
		2721 6942
		16494 39888
		29688 71276
		3383 8466
		2181 5662
		96601 231302
		1795 4690
		7512 18324
		1242 3384
		2889 7278
		2133 5566
		103 706
		4446 10992
		11326 27552
		3024 7548
		217 934
		13269 32038
		281 1062
		77174 184848
		952 2604
		15572 37644
		566 1832
		4103 10306
		313 1126
		14393 34886
		1313 3526
		348 1196
		419 1338
		246 992
		445 1390
		23552 56804
		23552 56804
		67 634`)
}

func newKnapsackOrPanicOnFailure(str string) *Knapsack {
	knapsack, err := NewKnapsack(ioutil.NopCloser(strings.NewReader(str)))
	if err != nil {
		panic(err)
	}
	return &knapsack
}
