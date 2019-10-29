package main

import "fmt"

func main() {

	map1 := make(map[int]int)

	changeMap(map1)

	for i,v:=range map1 {
		fmt.Println(i , v)
	}

}

func changeMap(map1 map[int]int)  {

	map1[1] = 2
	map1[2] = 2

}
