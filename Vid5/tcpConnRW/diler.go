package main

import (
	"net"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {

	conn , err := net.Dial("tcp" , ":8000")
	if err!= nil {
		log.Fatal(err)
	}

	defer conn.Close()

	data , err := ioutil.ReadAll(conn)
	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}
