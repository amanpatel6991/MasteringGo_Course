package main

import (
	"sync"
	"fmt"
)

type safeCounter struct {
	i int
	*sync.Mutex
}

func main() {

	sc := new(safeCounter)             //See when to take * and when not to
	//fmt.Println(*sc)
	*sc = safeCounter{0,&sync.Mutex{}}
	//fmt.Println(sc.Mutex)

	       //(or)
	//sc := &safeCounter{
	//	Mutex : &sync.Mutex{},
	//	i: 0,
	//}

	for i:=0 ; i< 10 ;i++ {
		go sc.Inc()
		go sc.Dec()
	}

	//time.Sleep(time.Second*2)
	fmt.Println(sc.GetValue())

}


func(sc *safeCounter) Inc() {
	// If the lock is already in use, the calling goroutine
	// blocks until the mutex is available.
	sc.Lock()
	sc.i++
	sc.Unlock()
}

func(sc *safeCounter) Dec() {
	sc.Lock()
	sc.i--
	sc.Unlock()
}

func(sc *safeCounter) GetValue() int{
	sc.Lock()
	val := sc.i
	sc.Unlock()

	return val
}