package main

import (
	"fmt"
)

func main() {

	var c chan int = make(chan int)
	var c1 chan int = make(chan int)

	go pinger(c)
	go ponger(c1)
	//go printer(c , c1)

	//var hold string
	//fmt.Scan(&hold)

	for r := range c{
		fmt.Println(r)
	}
	for r := range c1{
		fmt.Println(r)
	}


}


func pinger(c chan<- int) {                      //c can only be sent to
	for i:=0;i<10;i++{
		c <- i
	}
	close(c)
}

func ponger(c1 chan<- int) {                      //c can only be sent to
	for i:=0;i<10;i++{
		c1 <- i
	}
	close(c1)
}

func printer(c <-chan int , c1 <-chan int) {                        //can only receive from c
	//for {
	//	select {                                          //select works as switch case but only for channels
	//	case <-c :
	//		fmt.Println(<-c)
	//	case <-c1 :
	//		fmt.Println(<-c1)
	//	}
	//}

	for {
			fmt.Println(<-c)
			fmt.Println(<-c1)
	}
}