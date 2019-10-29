package main

import "fmt"

func main() {

	arr := []int{0,1,0,2,1,0,1,3,2,1,2,1}

	vol := trap(arr)

	fmt.Println(vol)
}

func trap(height []int) int {

	maxHeightLeftMap := make(map[int]int , 0)      //index , height
	maxHeightRightMap := make(map[int]int , 0)      //index , height

	maxVol :=0

	start := 0
	end := len(height) - 1

	for start < len(height) && end >= 0 {
		//if val2,ok:=maxHeightLeftMap[start-2]; ok {
		//	maxHeightLeftMap[start] = MaxVal(val2 , height[start-1])
		//} else
		if val1,ok:=maxHeightLeftMap[start-1]; ok {
			maxHeightLeftMap[start] = MaxVal(val1 , height[start-1])
		} else {
			maxHeightLeftMap[start] = 0
		}

		//if val2,ok:=maxHeightRightMap[end+2]; ok {
		//	maxHeightRightMap[end] = MaxVal(val2 , height[end-1])
		//} else
		if val1,ok:=maxHeightRightMap[end+1]; ok {
			maxHeightRightMap[end] = MaxVal(val1 , height[end+1])
		} else {
			maxHeightRightMap[end] = 0
		}

		start++
		end--
	}


	fmt.Println("adasd")
	for i:=0;i<len(height);i++ {
		fmt.Print(maxHeightLeftMap[i] , " ")
	}
	fmt.Println()
	for j:=0;j<len(height);j++ {
		fmt.Print(maxHeightRightMap[j] , " ")
	}

	for k:=0;k<len(height);k++ {
		tempVol := MinVal(maxHeightLeftMap[k] , maxHeightRightMap[k]) - height[k]

		if tempVol >=1 {
			maxVol += tempVol
		}

	}

	fmt.Println()

	return maxVol

}

func MaxVal(a , b int) int {
	if a > b {
		return a
	}
	return b
}

func MinVal(a , b int) int {
	if a > b {
		return b
	}
	return a
}