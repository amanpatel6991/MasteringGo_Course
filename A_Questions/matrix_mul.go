package main

import "fmt"

func main() {

	first := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	}

	m := 3
	//n := 3

	second := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	}

	p := 3
	q := 3

	multiply := [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}

	fmt.Println(len(first), len(second))

	//multiply
	for c := 0; c < m; c++ {
		for d := 0; d < q; d++ {
			sum := 0
			for k := 0; k < p; k++ {
				sum = sum + first[c][k] * second[k][d];
			}

			multiply[c][d] = sum;
			sum = 0;
		}
	}

	for _, v := range multiply {
		fmt.Println(v)
	}

}
