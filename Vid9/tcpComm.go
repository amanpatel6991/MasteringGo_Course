package main

import (
	"flag"
	"strings"
	"net"
	"fmt"
	"bufio"
	"os"
	"io"
	"log"
	"time"
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
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What msg you want to send ?")
	for scanner.Scan() {
		fmt.Println("Writing ", scanner.Text())
		conn.Write(append(scanner.Bytes(), '\r'))                   //sending user typed msg to server

		buffer := make([]byte, 1024)

		_, err := conn.Read(buffer)                    //read from conn into the buffer
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			log.Println("conn is closed !!")
			return nil
		}
		fmt.Println(string(buffer))
		fmt.Println("What msg you want to send ?")
	}
	return scanner.Err()
}

func runServer(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Println("Listening...")
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		conn.SetDeadline(time.Now().Add(time.Second*50))
		line,err:=reader.ReadString('\r')        // ReadString reads until the first occurrence of delim in the input
		if err != nil && err != io .EOF {
			log.Fatal(err)
			return
		} else if err == io.EOF {
			log.Println("conn closed !!")
			return
		}
		fmt.Printf("Received %s from address %s \n" , line[:len(line)-1] , conn.RemoteAddr())
		writer.WriteString("Msg Received...")        //sending to client
		writer.Flush()
	}

}