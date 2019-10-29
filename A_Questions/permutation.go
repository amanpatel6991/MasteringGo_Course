package main

import (
	"fmt"
)

func main() {

	str := "ABC"

	permutation(str, 0, len(str) - 1)

}

func permutation(str string, l int, h int) {

	if l == h {
		fmt.Println(str)
		return
	}

	for i := l; i <= h; i++ {
		str = swap(str, l, i)
		permutation(str, l + 1, h)
		str = swap(str, l, i)                 //backtracking ??
	}

}

func swap(str string, l int, h int) string{
	str1 := []rune(str)
	temp := str1[l]
	str1[l] = str1[h]
	str1[h] = temp

	return string(str1)

}