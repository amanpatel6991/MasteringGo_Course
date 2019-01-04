package main

import (
	"time"
	"math/rand"
	"fmt"
)

var errNotFound = "Member not found in db"

var sampleDb = map[string]int{
	"Aman" : 1,
	"Banga" : 2,
}

func main() {

	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)
	ch2 := make(chan int)

	name := "Aman"

	go SServer(name, "Server1", ch)
	go SServer(name, "Server2", ch2)


	select {
	case sc := <-ch:
		fmt.Println("Found on server 1 : ", sc)
	case sc := <-ch2:
		fmt.Println("Found on server 2 : ", sc)
	case <-time.After(time.Second * 8):
		fmt.Println("TimeOut !")
	//default:                                                 //Prevents waiting to receive from channels
	//	fmt.Println("slow !")
	}

}

//func SServer(name string, server string, c chan int) (int, error) {
//
//	fmt.Println("searching in :", server)
//	time.Sleep(time.Duration(rand.Intn(1)) * time.Second)
//
//	if v, ok := sampleDb[name]; !ok {
//		return -1, errors.New(errNotFound)
//	} else {
//		return v, nil
//	}
//
//	//c <- sampleDb[name]
//
//}

func SServer(name string, server string, c chan int) {

	fmt.Println("searching in :", server)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	c <- sampleDb[name]

}