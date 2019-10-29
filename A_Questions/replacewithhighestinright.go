package main

import "fmt"

func main() {

	arr := []int{16,17,4,3,5,2}

	replace(arr)

	fmt.Print(arr)

}

func replace(arr []int) {

	len1 := len(arr)
	highestFromRight := arr[len1-1]
        //arr[len1-1] = -1
	for i:=len1-2;i>=0;i--{
		if arr[i] < highestFromRight {
			arr[i] = highestFromRight
		}else {
			highestFromRight = arr[i]
		}
	}

}