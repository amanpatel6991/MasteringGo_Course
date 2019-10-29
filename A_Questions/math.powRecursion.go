package main

import "fmt"

func main() {

	res := Power(2, 10)
	fmt.Println(res)

}

func Power(n, p int) int {
	if p == 0 {
		return 1
	}

	n = n * Power(n, p - 1)
	return n
}