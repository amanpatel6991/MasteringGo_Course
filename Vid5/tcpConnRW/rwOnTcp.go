package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main() {

	li , err := net.Listen("tcp" , ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn , err:= li.Accept()
		if err!=nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintln(conn , "i heard you say :" , ln)
	}
	defer conn.Close()

	fmt.Println("code got here !!")

}