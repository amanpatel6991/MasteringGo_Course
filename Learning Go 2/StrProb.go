//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//
//	a:="aman"
//	b:="patel"
//
//	b = b+a
//
//	fmt.Println(b , len(b) , strings.Count(b , "s"))
//
//
//}

package main

import (
	"fmt"
	"strings"
)

func main() {

	var numOfStr int
	fmt.Scan(&numOfStr)

	var tempScan string

	var str []string
	//str = append(str , "sdf")
	for i := 0; i < numOfStr; i++ {
		fmt.Scan(&tempScan)
		str = append(str , tempScan)
	}

	fmt.Println(str)

	var strParent string = ""

	var count int
	var countNode int


	for j:=0 ; j<numOfStr ; j++ {
		strParent = strParent + str[j]
	}

	countNode = len(strParent)
	fmt.Println(countNode)

	for k:=0 ; k < len(strParent) ; k++ {

		count = 0
		for l:=0 ; l < len(strParent) ; l++ {

			if str[l] == str[k] {
				count++
			}

		}

		if count > 1 {
			countNode--
			strParent = strings.Replace(strParent , string(str[k]) , "" , -1)
		}

	}

	fmt.Println(countNode)

}