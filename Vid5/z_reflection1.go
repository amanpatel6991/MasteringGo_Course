package main

import (
	"reflect"
	"fmt"
)

func main() {

	var v float64 = 3.4

	val := reflect.ValueOf(&v)
	fmt.Println(val.CanSet())

	vElem := val.Elem()
	fmt.Println(vElem.CanSet())

	vElem.SetFloat(2.2)
	fmt.Println(vElem , v)

}
