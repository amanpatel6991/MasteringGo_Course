package main

import "fmt"

func main() {

	//BECAUSE All types implement empty interface (as empty interface has 0 methods and every one implements atleast 0 methods)
	checkType("text")
	checkType(3)
	checkType(4.5)

}

func checkType(i interface{}) {

	switch dataType := i.(type) {
	case int:
		fmt.Println("type is int :", dataType)
	case string:
		fmt.Println("type is string :", dataType)
	case float64:
		fmt.Println("type is float :", dataType)
	}

}