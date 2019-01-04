package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var c chan string = make(chan string)

	go func() {
		defer wg.Done()
		c <- "ping"
	}()

	go func() {
		defer wg.Done()
		c <- "pong"
	}()

	go func() {
		//can only receive from c
		for i:=range c{
			fmt.Println(i)
		}
	}()

	//var in string
	//fmt.Scan(&in)

	wg.Wait()

}
