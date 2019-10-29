package main

import "fmt"

func main() {
	//str := "abcabcbb"
	//str := "bbbbbbb"
	//str := "pwwkew"
	str := "a"

	ans := lengthOfLongestSubstring(str)

	fmt.Println(ans)

}

func lengthOfLongestSubstring(s string) int {

	// if s == "" || s == " " || len(s) == 1 {
	//     return len(s)
	// }

	mapStr := make(map[byte]int,0)

	maxCount := 0
	i := 0

	for j:=0;j<len(s);j++ {
		if _,ok := mapStr[s[j]]; ok {
			i = Max(mapStr[s[j]],i)
		}
		maxCount = Max(maxCount , j-i+1)
		mapStr[s[j]] = j + 1
	}

	return maxCount
}

func Max(a , b int) int {
	if a>b {
		return a
	}
	return b
}