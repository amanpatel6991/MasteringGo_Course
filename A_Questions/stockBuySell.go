package main

import "fmt"

func main() {

	const days int = 10
	//prices := [days]int{100, 180, 260, 310, 40, 535, 695}
	prices := [days]int{23, 13, 25, 29, 33, 19, 34, 45, 65, 67}
	//prices := [days]int{23, 13, 6, 4, 2, 19, 34, 45, 65, 67}
	//prices := [days]int{23, 13, 6,5,4,3,2,1}

	buyIndex := -1
	sellIndex := len(prices) - 1

	flag := false

	for i := len(prices) - 1; i > 0; i-- {
		if prices[i - 1] > prices[i] {
			buyIndex = i
			if buyIndex != sellIndex {
				fmt.Println("buy on :", buyIndex, "sell on : ", sellIndex)
				flag = true
			}
			sellIndex = i - 1
		}
	}

	if sellIndex != 0 {
		fmt.Println("buy on :", 0, "sell on : ", sellIndex)
		flag = true
	}

	if !flag {
		fmt.Println("No Profit !")
	}

}
