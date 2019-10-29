package main

import "fmt"

func main() {

	for i := 0; i < 8; i++ {
		//fmt.Print(fibonacci(i), " ")
		fmt.Print(fibonacciDP(i), " ")
	}

}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}

func fibonacciDP(n int) int {                    //bottom-up
	storage := make(map[int]int)

	storage[0] = 0
	storage[1] = 1

	for i := 2; i <= n; i++ {
		storage[i] = storage[i-1] + storage[i-2]
	}

	return storage[n]
}

var storage = make(map[int]int)

func fibonacciDP2(n int) int {                       //top-down

	if storage[n] == 0 {
		if n<=1{
			storage[n] = n
		}else {
			storage[n] = fibonacci(n - 1) + fibonacci(n - 2)
		}
	}

	return storage[n]
}