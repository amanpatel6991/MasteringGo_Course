package main

import (
	//"sort"
	"fmt"
)

func main() {
	//nums := []int{3,3}
	nums := []int{2,7,11,15}
	target := 9

	result := twoSum(nums , target)

	fmt.Println(result)
}


//func twoSum(nums []int, target int) []int {
//	result := make([]int,0)
//	dataMap := make(map[int]int)
//
//	for k , v := range nums {
//		dataMap[v] = k
//	}
//
//	sort.Ints(nums)
//
//
//	a := 0
//	b := len(nums) - 1
//	fmt.Println("a,b :" , a,b)
//	fmt.Println("nums :" , nums)
//	for a < b {
//		fmt.Println(a,b)
//		fmt.Println(nums[a],nums[b])
//		if nums[a] + nums[b] > target {
//			b--
//		} else if nums[a] + nums[b] < target {
//			a++
//		} else {
//			fmt.Println("apsaps")
//			fmt.Println(a,b)
//
//			result = append(result,dataMap[nums[a]])
//			result = append(result,dataMap[nums[b]])
//			break
//		}
//		fmt.Println(a,b)
//	}
//	return result
//}

func twoSum(nums []int, target int) []int {
	result := make([]int,2)
	resMap := make(map[int]int)

	// Loop:
	// for i,v1 := range nums {
	//     for j,v2 := range nums {
	//         if i == j {
	//             continue
	//         }
	//         if v1+v2 == target {
	//             result = append(result,i)
	//             result = append(result,j)
	//             break Loop
	//         }
	//     }
	// }


	for i,v := range nums {
		if v1 , ok := resMap[target - v]; ok {
			result[0] = v1
			result[1] = i
			break
		} else {
			resMap[v] = i
		}
	}

	return result
}