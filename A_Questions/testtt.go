package main

import "fmt"

func main() {
	a:=9
	func() {
	a=3
	}()

	fmt.Println(a)

}
