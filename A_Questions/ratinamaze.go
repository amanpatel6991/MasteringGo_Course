package main

import (
	"fmt"
)

var path [][]int
var sol [][]int

func main() {

	path = [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1},
	}

	sol = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	//fmt.Println(path)

	pathExists := checkPath(path, 0, 0)
	fmt.Println(pathExists)
	for _, v := range sol {
		fmt.Println(v)
	}

}

func checkPath(path [][]int, x, y int) bool {

	if x == 3 && y == 3 {
		sol[x][y] = 1
		return true
	}

	if isSafe(path, x, y) {
		fmt.Println(x,y)
		sol[x][y] = 1

		if checkPath(path, x + 1, y) {
			return true
		}
		if checkPath(path, x, y + 1) {
			return true
		}

		//backtracking
		sol[x][y] = 0
		return false
	}

	return false

}

func isSafe(path [][]int, x, y int) bool {
	return (x >= 0&&x <= 3) && (y >= 0&&y <= 3) && path[x][y] == 1
}