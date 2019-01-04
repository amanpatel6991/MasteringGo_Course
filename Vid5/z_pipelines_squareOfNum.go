package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen(2, 3, 4, 5, 6, 7, 8, 9)

	out1 := calcSq(in)                //Work divided in 2 goroutines
	out2 := calcSq(in)

	finalOut := mearge(out1, out2)

	for val := range finalOut {
		fmt.Println(val)
	}

}

func gen(input ...int) <-chan int {
	//Can only receive from c channel now

	c := make(chan int)
	go func() {
		for _, v := range input {
			c <- v
		}
		close(c)
	}()
	return c
}

func calcSq(in <-chan int) chan int {
	//Can only receive from in and out channel now

	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func mearge(out ...chan int) <-chan int {

	var wg sync.WaitGroup
	finalOut := make(chan int)

	wg.Add(len(out))
	output := func(c <-chan int) {
		for v := range c {
			finalOut <- v
		}
		wg.Done()
	}

	for _, c := range out {
		go output(c)

	}

	go func() {
		wg.Wait()
		close(finalOut)
	}()

	return finalOut

}
