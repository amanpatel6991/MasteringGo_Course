package main

import (
	"fmt"
)

func main() {

	arr := [4]int{1,3,2,5}

	arr2 := []int{1,4}
	//zt := reflect.TypeOf(arr).Kind()

	//fmt.Printf("%T: %s\n", zt, zt)
	arr2 = append(arr2 , 33)

	testFunc(arr , arr2)
	fmt.Println(arr,arr2)

}

func testFunc(arr [4]int , arr2 []int) {
	fmt.Println(arr,arr2)

}


