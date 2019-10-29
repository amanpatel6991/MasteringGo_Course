package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := []string{"A", "B", "C"}
	str2 := []string{"D", "E", "F"}
	str3 := []string{"G", "H", "I"}
	str4 := []string{"J", "K", "L"}

	keypad := [][]string{str1, str2, str3, str4}
	input := []int{2, 3, 4}

	findCombinations(keypad, input, 3)

	//strings.Index()

}

func findCombinations(keypad [][]string, input []int, size int) {

	if size == 0 {
		//fmt.Println()
	}

	for i := 0; i < 3; i++ {
		//findCombinations(keypad)
	}
}