package knapsack

import (
	"io/ioutil"
	"strings"
	"os"
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

func ks_30_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailureFile("data/ks_30_0")
}

func ks_40_0_Knapsack() *Knapsack {
	return newKnapsackOrPanicOnFailure(`40 100000
		90001 90000
		89751 89750
		10002 10001
		89501 89500
		10254 10252
		89251 89250
		10506 10503
		89001 89000
		10758 10754
		88751 88750
		11010 11005
		88501 88500
		11262 11256
		88251 88250
		11514 11507
		88001 88000
		11766 11758
		87751 87750
		12018 12009
		87501 87500
		12270 12260
		87251 87250
		12522 12511
		87001 87000
		12774 12762
		86751 86750
		13026 13013
		86501 86500
		13278 13264
		86251 86250
		13530 13515
		86001 86000
		13782 13766
		85751 85750
		14034 14017
		85501 85500
		14286 14268
		85251 85250
		14538 14519
		86131 86130`)
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

func newKnapsackOrPanicOnFailure(str string) *Knapsack {
	knapsack, err := NewKnapsack(ioutil.NopCloser(strings.NewReader(str)))
	if err != nil {
		panic(err)
	}
	return &knapsack
}
