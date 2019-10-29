package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0, 9}

	j:=0
	for i := 0; i < len(arr); i++ {
		if arr[i] !=0 {
			swap(arr , i,j)
			j++
		}
	}
	//for i := 0; i < len(arr); i++ {
	//	if arr[i] == 0 {
	//		for j := i; j > 0; j-- {
	//			swap(arr, j, j - 1)
	//		}
	//	}
	//}

	fmt.Println(arr)

	//fmt.Println(2+"45")
	//	tr, _ := parser.ParseExpr("(3-1) * 5")
	//	buf :=make([]byte , 20)
	//	types.WriteExpr(bytes.NewBuffer(buf) , tr)
	//fmt.Println(tr , buf)
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}