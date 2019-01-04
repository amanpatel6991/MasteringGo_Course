package main

import (
	"time"
	"math/rand"
	"fmt"
	"github.com/pkg/errors"
)

var errNotFound = errors.New("no record found !!")

var sampleDb = map[string]int{
	"Aman" : 1,
	"Banga" : 2,
}

func main() {
	name := "Aman"                     //name := "dsada"

	if res , err :=SServer(name, "Server1"); err!=nil {
		fmt.Println("error occured :" , err)
	} else {
		fmt.Println("Member found with level : " , res)
	}


}

func SServer(name string, server string) (int, error) {

	fmt.Println("searching in :", server)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	if v, ok := sampleDb[name]; !ok {
		//return -1, errNotFound
		return -1, fmt.Errorf("record wih name %s not found on server %s" , name , server)
	} else {
		return v, nil
	}


}