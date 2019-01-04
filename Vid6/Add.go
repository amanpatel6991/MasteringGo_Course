package main

import (
	"fmt"
)

func main() {
	res := Add(2, 3)
	var f [2]int
	fmt.Println(res , f)

	//n := strings.Count("amaaaa" , "a")
	//fmt.Println(n)
}

func Add(a, b int) int {
	return a + b
}