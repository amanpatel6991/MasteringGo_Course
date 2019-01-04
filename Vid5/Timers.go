package main

import (
	"time"
	"fmt"
)

func main() {

	dchannel := make(chan int)
	stopChannel := make(chan bool)

	go slowCounter(1, dchannel, stopChannel)

	time.Sleep(time.Second * 5)

	dchannel <- 2
	time.Sleep(time.Second * 10)

	stopChannel <- true
	time.Sleep(time.Second * 1)
}

func slowCounter(n int, dch chan int, sch chan bool) {
	i := 0;

	timeInSec := time.Duration(n) * time.Second           //duration in sec
	Loop:
	for {
		//NewTimer creates a new Timer that will send
		//the current time on its channel after at least duration d.
		//timer := time.NewTimer(timeInSec)        //sends current time on channel after timeInSec duration has elapsed
		//<- timer.C

		select {
		case <-time.After(timeInSec):
			i++
			fmt.Println(i)
		case n = <-dch:
			fmt.Println("duration changed to :", n)
			timeInSec = time.Duration(n) * time.Second
		case <-sch:
			fmt.Println("timer stopped !")
			break Loop
		}

	}
}