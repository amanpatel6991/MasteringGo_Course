package main

import (
	"math"
	"fmt"
)

var cost [][]float64

func main() {

	cost = [][]float64{
		{1, 2, 3},
		{4, 8, 2},
		{1, 5, 3},
	};

	//mincost := minPath(cost, 2, 2)
	mincost := minPathDP(cost, 2, 2)

	fmt.Println(mincost)

}

func minPath(cost [][]float64, x, y int) float64 {

	if x < 0 || y < 0 {
		return math.MaxFloat64
	} else if x == 0 && y == 0 {
		return cost[x][y]
	}

	fmt.Println(x, y)
	return cost[x][y] + math.Min(minPath(cost, x - 1, y - 1), math.Min(minPath(cost, x - 1, y), minPath(cost, x, y - 1)))

}

func minPathDP(cost [][]float64, x, y int) float64 {      //bottom-up

	var tc [3][3]float64

	tc[0][0] = cost[0][0]

	for i := 1; i <= x; i++ {
		tc[i][0] = tc[i - 1][0] + cost[i][0]
	}
	for j := 1; j <= y; j++ {
		tc[0][j] = tc[0][j - 1] + cost[0][j]
	}

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			tc[i][j] = math.Min(tc[i - 1][j - 1], math.Min(tc[i - 1][j], tc[i][j - 1])) + cost[i][j]
		}
	}
	//fmt.Println(tc)
	return tc[x][y]
}