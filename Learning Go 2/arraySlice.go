package main

import (
	"fmt"
	"reflect"
)

func main() {
	//a := []int{}

	a := make([]int , 0)
a =append(a,2)
	var b [2]int

	fmt.Println( a ,reflect.TypeOf(a).Kind())
	fmt.Println( b ,reflect.TypeOf(b).Kind())
}
