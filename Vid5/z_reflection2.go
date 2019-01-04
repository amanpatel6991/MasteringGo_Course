package main

import (
	"fmt"
	"reflect"
)

func main() {

	var val interface{} = 5
	val2 := reflect.ValueOf(val)
	fmt.Println(val , val2 , val2.Type())


	val3 := reflect.ValueOf(&val)
	// Elem returns the original value that the interface v contains
	// or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	valElem:=val3.Elem()
	fmt.Println(valElem.CanSet())

}
