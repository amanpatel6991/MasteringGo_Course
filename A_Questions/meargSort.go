package main

import (
	"fmt"
)

var arr []int

func main() {

	//arr = []int{2, 5, 3, 7, 5, 4 , 9 , 0 ,3}
	arr = []int{10, 94, 63, 47, 51}
	//arr = []int{1, 20, 6, 4, 5}
	//arr = []int{1, 5, 6, 4, 20}
	//arr = []int{-7, 4, -3, 2, 2 ,-8, -2, 3, 3, 7, -2, 3, -2}

	inv := meargSort(arr, 0, len(arr) - 1)

	fmt.Println(arr)
	fmt.Println(inv)

}

func meargSort(arr []int, low, high int) int {

	inv := 0

	if low < high {
		middle := (low + high) / 2
		inv = inv + meargSort(arr, low, middle)
		inv = inv + meargSort(arr, middle + 1, high)

		inv = inv + mearge(arr, low, middle, high)
	}

	return inv

}

func mearge(arr []int, low, mid, high int) int {

	var invtemp int

	size1 := mid - low + 1
	size2 := high - mid

	left := make([]int, size1)
	right := make([]int, size2)

	for i := 0; i < size1; i++ {
		left[i] = arr[low + i]
	}

	for j := 0; j < size2; j++ {
		right[j] = arr[mid + j + 1]
	}

	i := 0
	j := 0
	k := low

	for i < size1 && j < size2 {
		//if math.Abs(float64(left[i])) <= math.Abs(float64(right[j])) {               //a variation
		//if left[i]%10 <= right[j]%10 {                                                 //another variation
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
			invtemp = invtemp + (mid - i + 1)
		}
		k++
	}

	for i < size1 {
		arr[k] = left[i]
		i++
		k++
	}

	for j < size2 {
		arr[k] = right[j]
		j++
		k++
	}

	return invtemp

}