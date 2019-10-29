package main

import "fmt"

var board [][]int
var N = 4

func main() {

	board = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	isPossible := solveNQueens(board, 0)

	fmt.Println(isPossible)
	for _, v := range board {
		fmt.Println(v)
	}

}

func solveNQueens(board [][]int, col int) bool {

	if col >= N {
		return true             //all queens placed
	}

	for i := 0; i < N; i++ {
		if isCellSafe(board, i, col) {
			board[i][col] = 1

			if solveNQueens(board, i + 1) {
				return true
			}

			board[i][col] = 0
		}
	}
	return false
}

func isCellSafe(board [][]int, row, col int) bool {

	//check left cells of this row
	for i := 0; i < col; i++ {
		if board[row][col] == 1 {
			return false
		}
	}

	//check top left diagonal
	j := col
	i := row
	for i >= 0 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	//check bottom left diagonal
	j = col
	i = row
	for i < N && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i++
		j--
	}

	return true

}