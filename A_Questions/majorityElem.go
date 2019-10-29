package main

import "fmt"

var arr []int

func main() {
	arr = []int{5, 6, 5, 5, 3, 0, 2}

	findMajority(arr, 0, len(arr) - 1)

}

func findMajority(arr []int, low, high int) {
	maj_index := 0
	count := 1

	for i := low+1; i <= high; i++ {
		 if arr[maj_index] == arr[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			maj_index = i
			count = 1
		}
	}

	fmt.Println(maj_index , count)

}