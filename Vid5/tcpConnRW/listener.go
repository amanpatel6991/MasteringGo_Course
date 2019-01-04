package main

import (
	"net"
	"log"
	"io"
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
		io.WriteString(conn , "hello !!!!!")
		io.WriteString(conn , "hello !!!!!")
		//scanner := bufio.NewScanner(bufio.NewReader(conn))
		//for scanner.Scan() {
		//	io.WriteString(conn , scanner.T)
		//}
		conn.Close()
	}
}
