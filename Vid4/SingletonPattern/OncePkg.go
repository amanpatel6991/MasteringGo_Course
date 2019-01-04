package main

import (
	"sync"
	"fmt"
)

func main() {

	var once sync.Once

	job := func() {
		fmt.Println("executeing job")
	}

	done := make(chan bool)

	for i:=0;i<10;i++ {
		go func() {
			once.Do(job)
			done<-true
		}()
	}

	for j:=0;j<10 ;j++  {
		fmt.Println(<-done)
	}

}
