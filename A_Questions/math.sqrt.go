package main

import (
	"fmt"
	"math"
)

const EPSILON = 0.0001

func main() {

	sqrt := findSqrt(8)
	fmt.Printf("%0.2f",sqrt)
	fmt.Println()
	fmt.Printf("%0.2f",math.Pow(sqrt,2))
}

func findSqrt(num float64) float64 {
	var low float64= 0;
	var high float64 = num;
	var mid float64 = 0;

	for high - low > EPSILON {
		mid = low + (high - low) / 2; // finding mid value
		if mid * mid > num {
			high = mid;
		} else if mid * mid < num {
			low = mid;
		} else {
			break
		}
	}
	return mid;
}
