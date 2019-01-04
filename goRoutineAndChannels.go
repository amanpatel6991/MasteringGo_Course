package main

import "fmt"

func main() {

	ch := make(chan bool)
	go WaitAndSay(ch, "world")
	fmt.Println("Hello ")

	ch <- true

	<- ch
}

func WaitAndSay(ch chan bool, s string) {

	if val := <-ch; val {
		fmt.Println(s)
	}
	ch <- true
}