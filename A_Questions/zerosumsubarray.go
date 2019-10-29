package main

import "fmt"

func main() {

	//arr := []int{1, 4, -2, -2, 5, -4, 3}
	//arr := []int{6, 3, -1, -3, 4, -2, 2, 4, 6, -12, -7}
	arr := []int{3, 4, -7, 1, 3, 3, 1, -4}
	//arr := []int{4, 2, 0, 1, 6}
	//arr := []int{-3, 2, 3, 1, 6}

	map1 := make(map[int]struct{})
	sum := 0
	existFlag := false
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		//fmt.Println(sum)
		if _, ok := map1[sum]; ok {
			//fmt.Println("exists")
			existFlag = true
			tempSum := 0
			for j := i; j >= 0; j-- {
				tempSum = tempSum + arr[j]
				if tempSum == 0 {
					printZeroSubArray(arr, j, i)
				}
			}
		}
		if sum == 0 {
			// if sum is 0, we found a subarray starting
			// from index 0 and ending at index i
			existFlag = true
			printZeroSubArray(arr, 0, i)
		}

		map1[sum] = struct{}{}
		//fmt.Println(existFlag)
	}

	if existFlag {
		fmt.Println("exists")
	} else {
		fmt.Println("does not exist")
	}

}

func printZeroSubArray(arr []int, start, end int) {
	fmt.Print("sub array is : ")
	for i := start; i <= end; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}