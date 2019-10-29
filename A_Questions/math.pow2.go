package main

import "fmt"

func main() {

	res := Power(2, 4)
	fmt.Println(res)

}

func Power(n, p int) int {
	if p == 0 {
		return 1
	}

	temp := Power(n, p / 2)

	if p % 2 == 0 {
		return temp * temp
	} else {
		if p > 0 {
			return n * temp * temp
		} else {
			return (temp * temp) / n
		}
	}
	//return n
}