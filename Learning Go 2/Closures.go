package main

import "fmt"

func intSequence() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main() {

	intSeq := intSequence()

	fmt.Println(intSeq(), intSeq())

	intSeq2 := intSequence()
	fmt.Println(intSeq2(), intSeq2())
}
