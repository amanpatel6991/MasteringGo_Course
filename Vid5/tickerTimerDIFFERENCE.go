package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTimer(time.Duration(1) * time.Second)

	go tickCounter(ticker)

	time.Sleep(time.Second * 5)
	ticker.Stop()
	time.Sleep(time.Second * 4)
	fmt.Println("exit ticker !")
}

func tickCounter(ticker *time.Timer) {

	i := 0

	for t := range ticker.C {

		i++
		fmt.Println("count :", i, "at :", t)
	}
}