package main

import (
	"fmt"
)

type Set map[string]struct{}                         //set is an arraylist which has uniques items

func main() {

	set := make(Set)
	//fmt.Println(reflect.TypeOf(set).Kind())

	set["item1"] = struct{}{}
	set["item2"] = struct{}{}

	for v := range set {
		fmt.Println(v)
	}

}
