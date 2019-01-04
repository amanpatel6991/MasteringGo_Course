package main

import (
	"strings"
	"fmt"
)

func main() {

	//var games int
	aliceStr := "aman a"
	bobStr := "aman pat"


	for _, v := range aliceStr {
		for _, v1 := range bobStr {
			if v == v1 {
				bobStr = strings.Replace(bobStr, string(v1), "", -1)
				aliceStr = strings.Replace(aliceStr, string(v), "", -1)
			}
		}
	}

	fmt.Println(aliceStr == "" )
	fmt.Println(bobStr == "" )

}