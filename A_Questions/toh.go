package main

import (
	"fmt"
	"math"
)

var c int

func main() {

	n := 3
	toh(n, "FROM", "TO", "AUX")
	fmt.Println("moves :" , math.Pow(2.0 , float64(n))-1)
}

func toh(n int, from, to, aux string) {

	if n == 1 {
		fmt.Println("Move disk 1 from ", from, " to ", to)
		return
	}
        //c++
	toh(n - 1, from, aux, to)
	fmt.Println("Move disk ", n, " from ", from, " to ", to)
	toh(n - 1, aux, to, from)

}