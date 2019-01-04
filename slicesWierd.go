package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := [4]int{1 ,2 , 5,3}

	fmt.Println(reflect.TypeOf(a).Kind() , a)

	s := a[:3]
	s = append( s , 7)                //s = append( s , 7 , 8) , write this and see the wierdness printing below :D

	fmt.Println(reflect.TypeOf(s).Kind() , s , a)



}
