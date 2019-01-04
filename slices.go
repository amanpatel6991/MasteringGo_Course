package main

import "fmt"

func main() {

	var sl []int
	s2:=[]int{3,4,8,9}
	//sl = append(sl , s2...)
	sl = append(sl , s2[:3]...)

	fmt.Println(sl)

	s3:=[]int{3,4,8,9}
	s4 := s3[1:3]
	s4[0] = 3232
	fmt.Println(s3)

}
