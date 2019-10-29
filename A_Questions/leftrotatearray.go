package main

import "fmt"

func main() {

	arr :=[]int{1, 2, 3, 4, 5, 6, 7}

	nTimes := 2
	for i:=0;i<nTimes;i++ {
		leftByOne(arr,len(arr)-1)
	}

	fmt.Println(arr)

}

func leftByOne(arr []int , len1 int) {
	//fmt.Print(len1)
	for i:=0;i<=len1-1;i++{
		swap(arr , i,i+1)
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}