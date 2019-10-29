package main

import "fmt"

func main() {

	//s := "barfoothefoobarman"
	words := []string{"foo","bar"}

	//res := make([]int , 0)
	subStrSlice := make([]string,0)

	subStrSlice = permutation(words , 0 , len(words)-1 , subStrSlice)

	fmt.Println(subStrSlice)

	//return res

}

func permutation(words []string , l , h int, subStrSlice []string) []string {

	//fmt.Println("start :",l , h)
	if l == h {
		subStr := getSubStr(words)
		//fmt.Println(subStr)
		subStrSlice = append(subStrSlice , subStr)
		//fmt.Println(words)
		return subStrSlice
	}

	for i := l; i<=h; i++ {
		words = swap(words , l , i)
		//fmt.Println(words , l , i , h)
		subStrSlice = permutation(words , l+1 , h , subStrSlice)
		words = swap(words , l , i)
	}

	return subStrSlice

}

func getSubStr(words []string) string {

	str := ""
	for _,v:=range words {
		str += v
	}

	return str
}

func swap(str1 []string, l int, h int) []string{
	// str1 := []rune(str)
	temp := str1[l]
	str1[l] = str1[h]
	str1[h] = temp

	return str1

}