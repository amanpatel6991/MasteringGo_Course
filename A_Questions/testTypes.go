package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type test struct {
	i int
	name string
}

func main() {

	var data test = test{1,"Aman"}
	//a = "asa"                                 //not allowed due to type saftey

	fmt.Println(data, reflect.TypeOf(data).Kind() , reflect.TypeOf(data).Align() , unsafe.Alignof(data))

}
