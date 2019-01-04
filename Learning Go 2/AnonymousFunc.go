package main

import "fmt"

var DoWork func() = func() {
	fmt.Println("working on func 1")
}

func main() {

	DoWork()

	DoWork = func() {
		fmt.Println("working on func 2")
	}

	DoWork()

}
