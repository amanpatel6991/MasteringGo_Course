package main

import (
	"time"
	"math/rand"
	"fmt"
)

type MyError struct {
	ServerName string
	Code       int
	Msg        string
}

func(e MyError) Error() string{
	return "record not found on "+ e.ServerName+" , error msg : " + e.Msg
}

var sampleDb = map[string]int{
	"Aman" : 1,
	"Banga" : 2,
}

func main() {
	name := "Ajman"                     //name := "dsada"

	if res, err := SServer(name, "Server1"); err != nil {
		fmt.Println("error occured :", err)
		if v,ok:=err.(MyError);ok {
			fmt.Println("server : " , v.ServerName)
		}
	} else {
		fmt.Println("Member found with level : ", res)
	}

}

func SServer(name string, server string) (int, error) {

	fmt.Println("searching in :", server)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	if v, ok := sampleDb[name]; !ok {
		//return -1, errNotFound
		return -1, MyError{server, -1, "record not found !"}
	} else {
		return v, nil
	}

}