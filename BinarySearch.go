package main

import "fmt"

func main() {

	arr := []int{1,2,3,4,5,6,7,8,9,10}

	searchByRec(arr , 0 , len(arr)-1 , 1)

}

func searchByRec(arr []int , low int , high int , val int) {

	middle := (low+high)/2
	if (val < arr[middle]) {
		searchByRec(arr , low , middle , val)
	}else if (val > arr[middle]) {
		searchByRec(arr , middle , high , val)
	}else {
		fmt.Println("found at :" , middle , " " , arr[middle])
	}

}