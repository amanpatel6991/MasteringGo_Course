package main

import (
	"time"
	"fmt"
)
//DOUBT......
func main() {

	//go myTimer()
	go myTicker()


	time.Sleep(time.Second * 15)
}

func myTimer() {

	i := 0
	for {
		fmt.Println("doing task!")
		time.Sleep(time.Second*2)

		timer := time.NewTimer(time.Second)
		<-timer.C
		i++
		println("i :", i)
	}
}

func myTicker() {

	j := 0
	for {

		fmt.Println("doing task!")
		time.Sleep(time.Second*2)

		ticker := time.NewTicker(time.Second)
		<-ticker.C
		j++
		println("i :", j)
	}
}