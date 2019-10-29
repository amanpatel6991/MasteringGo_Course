package main

import "fmt"

func main() {

	arr := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	m := len(arr)                             //col length
	//n := len(arr[0])                         //row len
	fmt.Println(m)

	for i := 0; i < m / 2; i++ {
		for j := i; j < m-i-1; j++ {

			temp := arr[i][j]

			arr[i][j] = arr[j][m-i-1]             //right to top
			arr[j][m-i-1] = arr[m-i-1][m-j-1]             //bottom to right
			arr[m-i-1][m-j-1] = arr[m-j-1][i]             //left to bottom

			arr[m-j-1][i] = temp


		}
	}

	for _,v:=range arr {
		fmt.Println(v)
	}

	//str := "dsd"
	//
	//for _,v:=range str {
	//	fmt.Println(string(v))
	//}

}
