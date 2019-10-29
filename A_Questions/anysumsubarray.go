package main

import (
	"fmt"
)

func main() {

	//arr := []int{1, 4, -2, -2, 5, -4, 3}
	//arr := []int{6, 3, -1, -3, 4, -2, 2, 4, 6, -12, -7}
	arr := []int{3, 4, -7, 1, 3, 3, 1, -4}               //take sum = 7
	randomSum := 7
	largestIndex := 0
	smallestIndex := 0
	//arr := []int{4, 2, 0, 1, 6}
	//arr := []int{-3, 2, 3, 1, 6}

	map1 := make(map[int]struct{})
	sum := 0
	existFlag := false
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		//fmt.Println(sum)
		if _, ok := map1[sum - randomSum]; ok {
			//fmt.Println("exists")
			existFlag = true
			tempSum := 0
			for j := i; j >= 0; j-- {
				tempSum = tempSum + arr[j]
				if tempSum == randomSum {
					if j - i > largestIndex - smallestIndex {
						largestIndex = i
						smallestIndex = j
					}
					fmt.Println("test  ")
					printAnySumSubArray(arr, j, i)
				}
			}
		}
		if sum == randomSum {
			if !existFlag {
				printAnySumSubArray(arr, 0, i)             //to avoid duplicate printing
			}
			existFlag = true
			if i - 0 > largestIndex {
				largestIndex = i
				smallestIndex = 0
			}
		}

		map1[sum] = struct{}{}
		//fmt.Println(existFlag)
	}

	if existFlag {
		fmt.Println("exists")
	} else {
		fmt.Println("does not exist")
	}

	fmt.Print("largest sum subarray is : ", )
	for k := smallestIndex; k <= largestIndex; k++ {
		fmt.Print(arr[k] , " ")
	}
}

func printAnySumSubArray(arr []int, start, end int) {
	fmt.Print("sub array is : ")
	for i := start; i <= end; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}