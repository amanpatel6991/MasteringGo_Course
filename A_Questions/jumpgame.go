package main

import (
	"math"
	"fmt"
)

func main() {

	//nums := []int{2,3,1,1,4}
	nums := []int{3,2,1,0,4}
	jumps := make([]int,len(nums))

	jumps[0]=0
	lastIndexBool := false

	for i:=1 ; i< len(nums) ; i++ {
		jumps[i] = math.MaxInt64
		lastIndexBool = false
		for j:=0 ; j<i;j++ {
			if i<=j+nums[j] && jumps[j] != math.MaxInt64 {
				jumps[i] = MinVal(jumps[i],jumps[j]+1)
				lastIndexBool = true
			}
		}
	}

	fmt.Println(jumps , lastIndexBool)

}

func MinVal(a,b int) int {
	if a<b{
		return a
	}
	return b
}
