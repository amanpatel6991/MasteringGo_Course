package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	lenArr := len(arr)
	i := 0

	finalStep := 23
	//cannotJumpBy := make([]int , 0)

	jumps := 0
	jumpedBy := make([]int, 0)

	for {
		tempIndex := lenArr - 1 - i
		//if tempIndex < 0 {
		//	break
		//}
		if tempIndex < 0 || finalStep <= 0 {
			break
		}
		if arr[tempIndex] <= finalStep {
			jumpedBy = append(jumpedBy, arr[tempIndex])
			finalStep = finalStep - arr[tempIndex]
			jumps++
		} else {
			i++
		}
		fmt.Println("test :", finalStep, jumps, i)
	}

	fmt.Println(jumps)
	fmt.Println(jumpedBy)

}
