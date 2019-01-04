package main

import (
	"strings"
	"fmt"
)

var finalStr string

func main() {

	str := "Hello , am here !"

	wordsArray := strings.Split(str, " ")

	for i, v := range wordsArray {
		reverse(v)
		if i != len(wordsArray) - 1 {
			finalStr = finalStr + " "
		}
	}

	fmt.Println(finalStr)

}

func reverse(word string) {
	wordRune := []rune(word)
	length := len(wordRune) - 1
	for i := 0; i <= length / 2; i++ {
		temp := wordRune[i]
		wordRune[i] = wordRune[length - i]
		wordRune[length - i] = temp
	}

	finalStr = finalStr + string(wordRune)

}
