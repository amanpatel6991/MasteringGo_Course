package main

import "fmt"

func main() {

	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("mul res : " , mul(5 , add))

}

func mul(n int, mulFunc func(a,b int) int) int{
	return n * mulFunc(2,3)
}

//func add(a, b int) int {
//	return a + b
//}