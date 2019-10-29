package main

import (
	"fmt"
)

var intermediateResulSet [16]int

func main() {

	var rows int = 5
	var limit int

	for it := 1; it <= rows; it++ {
		limit = limit + it
	}

	fiboNum := 1

	Loop:
	for i := 1; i < limit; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(fibo(fiboNum), " ")
			fiboNum++
			if fiboNum > limit {
				break Loop
			}
		}
		fmt.Println()
	}

}

func fibo(i int) int {                                //top-down

	if intermediateResulSet[i] == 0 {
		if i <= 1 {
			intermediateResulSet[i] = i
			//intermediateResulSet[i] = i
		} else {
			intermediateResulSet[i] = fibo(i - 1) + fibo(i - 2)
		}

	}
	return intermediateResulSet[i]
}