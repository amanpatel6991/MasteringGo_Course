package main

import (
	"flag"
	"strings"
	"net"
	"fmt"
	"bufio"
	"os"
	"log"
)

func main() {

	op := flag.String("type", "", "Server (S) or Client (C) ?")
	//address := flag.String("addr", ":8080", "address? host:port ")
	address := ":8000"
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(address)
	case "C":
		runClient(address)
	}

}

func runClient(address string) error {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}

	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What msg you want to send ?")
	for scanner.Scan() {
		fmt.Println("Writing ", scanner.Text())
		conn.Write(scanner.Bytes())

 		buffer := make([]byte, 1024)

		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buffer))
		fmt.Println("What msg you want to send ?")
	}
	return scanner.Err()
}

func runServer(address string) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		return err
	}

	defer pc.Close()

	buffer := make([]byte , 1024)
	fmt.Println("Listening...")
	for {
		_ , addr , _ := pc.ReadFrom(buffer)
		fmt.Printf("Received %s from address %s \n" , string(buffer) , addr)
		_ , err :=pc.WriteTo([]byte("Msg received") , addr)
		if err!=nil{
			log.Fatal("cant write back to conn" , err)
		}
	}
}