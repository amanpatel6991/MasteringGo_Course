package main

import "fmt"

var arr []int

func main() {

	arr = []int{0, 0, 1, 0, 1, 1, 1, 0, 1, 1}

	findThatZero(arr, 0, len(arr) - 1)

}

func findThatZero(arr []int, low, high int) {

	count := 0
	maxCount := 0
	prevZeroIndex := 0
	indexToRelplace := 0

	for i := 0; i <= high; i++ {
		if arr[i] == 1 {
			count++
		} else {
			count = i - prevZeroIndex     //0 1 2 2
			prevZeroIndex = i             //0 1 1 3
		}

		if count > maxCount {
			maxCount = count
			indexToRelplace = prevZeroIndex
		}
	}

	fmt.Println(maxCount , indexToRelplace)

}
