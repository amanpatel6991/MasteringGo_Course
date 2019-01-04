package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(time.Duration(1) * time.Second)

	done := tickCounter(ticker)

	time.Sleep(time.Second * 5)
	ticker.Stop()
	done <- true
	time.Sleep(time.Second * 4)
	fmt.Println("exit ticker !")
}

func tickCounter(ticker *time.Ticker) chan bool{

	done := make(chan bool)
	go func() {
		i := 0
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
	}()
	return done
}