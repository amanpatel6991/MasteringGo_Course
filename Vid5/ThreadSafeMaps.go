package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

type threadSafeCounter struct {
	m map[int]int
	sync.RWMutex
}

func main() {

	mc := threadSafeCounter{m: make(map[int]int)}
	go runWriters(&mc, 10)
	go runReaders(&mc, 10)
	go runReaders(&mc, 10)

	time.Sleep(time.Second * 10)

}

func runWriters(mc *threadSafeCounter, n int) {
	for i := 0; i < n; i++ {
		// If the lock is already locked for reading or writing,
		// Lock blocks until the lock is available.                  (similar to Mutex wala Lock()) ??
		mc.Lock()
		mc.m[i] = i * 10
		mc.Unlock()
		time.Sleep(time.Second)
	}
}

func runReaders(mc *threadSafeCounter, n int) {
	for {
		// If the lock is already locked for reading or writing,
		// Lock blocks until the lock is available.
		mc.RLock()
		v := mc.m[rand.Intn(n)]     //0 to n-1
		mc.RUnlock()
		fmt.Println(v)
		time.Sleep(time.Second)
	}
}
