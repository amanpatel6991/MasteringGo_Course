package main

import "fmt"

func main() {

	resSlice := cutSlice()

	fmt.Println(resSlice , " old slice :" ,resSlice[:cap(resSlice)])
}

//func cutSlice() []int{
//	sl := []int{1,2,3,4,5,6,7,8}  //causes memory leak as cap is still not shortened
//
//	return sl[1:4]
//}

func cutSlice() []int{
	sl := []int{1,2,3,4,5,6,7,8}  //no mem leak now
	s2:=make([]int , 3)
	copy(s2,sl[1:4])
	return s2
}