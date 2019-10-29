package main

import (
	"fmt"
	"reflect"
	"math"
)

func main() {

	arr := [8]float64{-2, -3, 4, -1, -2, 1, 5, -3}

	fmt.Println(reflect.TypeOf(arr), reflect.ValueOf(arr), reflect.TypeOf(arr).Kind())

	currMax := arr[0]
	maxSoFar := arr[0]

	for i := 1; i < len(arr); i++ {
		currMax = math.Max(arr[i], currMax + arr[i])
		maxSoFar = math.Max(maxSoFar, currMax)
		fmt.Println(i, currMax, maxSoFar)
	}

	fmt.Println(maxSoFar)

}
