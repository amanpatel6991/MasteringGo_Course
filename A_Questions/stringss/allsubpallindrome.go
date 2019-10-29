package main

import "fmt"

var str string

func main() {

	str = "ssgoogle"

	low1 := 0
	high1 := 0

	fmt.Println(len(str))
	for i:=0 ; i< len(str)-1; i++ {
		check(str, i, i , &low1 , &high1)               //for odd length
		check(str, i, i+1, &low1 , &high1)             //for even length
	}

	fmt.Println("longest is :" , string([]rune(str)[low1:high1+1]) , " from index " , low1 ," to " , high1)

}

func check(str string , low , high int , low1 , high1 *int)  {
	for low>=0 && high<=len(str) && rune(str[low]) == rune(str[high]){
		fmt.Println("sub palli is :" , string([]rune(str)[low:high+1]))
		if high - low > *high1 - *low1 {
			*high1 = high
			*low1 = low
		}
		low--
		high++
	}
}