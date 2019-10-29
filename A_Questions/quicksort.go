package main

import (
	"fmt"
)

var arr []int

func main() {

	//arr = []int{2, 5, 3, 7, 5, 4, 9, 0, 3}
	//arr = []int{1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0}
	arr = []int{-7, 4, -3, 2, 2, -8, -2, 3, 3, 7, -2, 3, -2}

	quickSort(arr, 0, len(arr) - 1)

	fmt.Println(arr)

}

func quickSort(arr []int, low, high int) {

	if low < high {
		partitionIndex := partition(arr, low, high)

		quickSort(arr, low, partitionIndex - 1)
		quickSort(arr, partitionIndex + 1, high)
	}

}

func partition(arr []int, low, high int) int {

	pivot := arr[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}

	swap(arr, i + 1, high)

	return i + 1

}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}