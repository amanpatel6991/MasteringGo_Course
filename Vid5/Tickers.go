package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(time.Duration(1) * time.Second)

	done := make(chan bool)
	go tickCounter(done, ticker)

	time.Sleep(time.Second * 5)
	ticker.Stop()
	done <- true
	time.Sleep(time.Second * 4)
	fmt.Println("exit ticker !")
}

func tickCounter(done chan bool, ticker *time.Ticker) {

	i := 0
	// NewTicker returns a new Ticker containing a channel that will send the
	// time with a period specified by the duration argument.
	Loop:
	for {
		select {
		case v := <-ticker.C:
			i++
			fmt.Println("count :", i, "at :", v)
		case <-done:
			fmt.Println("ticker stopped , closing ticker channel")
			break Loop
		}
	}
	fmt.Println("Exiting !")
}