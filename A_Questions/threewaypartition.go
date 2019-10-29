package main

import "fmt"

var arr []int

func main() {

	arr = []int{0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0}          //0s ,1s ,2s
	threeWayPartition(arr, len(arr) - 1)

	fmt.Println(arr)

}

func threeWayPartition(arr []int, end int) {
	pivot := 1
	start := 0
	mid := 0

	for mid <= end {
		if arr[mid] < pivot {
			//0
			swap(arr, start, mid)
			start++
			mid++
		} else if arr[mid] > pivot {
			//2
			swap(arr, mid, end)
			end--
		} else {
			//1
			mid++
		}
	}

}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}