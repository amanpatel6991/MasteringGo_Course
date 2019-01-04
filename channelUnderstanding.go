package main

import "fmt"

func main() {

	ch := make(chan int , 2)

	ch <- 2
	ch <- 3

	fmt.Println(<-ch , <-ch)
	ch <- 3
	ch <- 6
	fmt.Println(<-ch , <-ch)

}
