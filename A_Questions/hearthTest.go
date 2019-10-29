package main

import "fmt"

func main() {
	var n int
	arr := make([]int , 0)

	fmt.Scanln(&n)
	for i:=0;i<n;i++{
		var temp int
		fmt.Scanln(&temp)
		arr = append(arr , temp)
	}

	for i:=n-1;i>=0;i--{
		fmt.Println(arr[i])
	}

}