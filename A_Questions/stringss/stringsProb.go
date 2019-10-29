package main

import (
	"fmt"
)

func main() {

	str := "zeekforgeeks"

	strRune := []rune(str)

	for i:=1;i<len(strRune)-1;i++ {
		if strRune[i] == strRune[i+1] {
			if str[i-1] == 97 {
				strRune[i] = 98
			}else {
				strRune[i] = 97
			}
		}
	}

	fmt.Println(string(strRune))

}
