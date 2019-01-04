package main

import "fmt"

type Student struct {
	Id int
	Name string
	//getData func() string
}

func main() {
	var s Student


	s.Id=1
	s.Name="Am"
	fmt.Println(s)
}
