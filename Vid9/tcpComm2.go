package main

import (
	"flag"
	"strings"
	"net"
	"fmt"
	"bufio"
	"os"
	"io"
	"time"
)

func main() {

	choice := flag.String("type" , "" , "Server S or Client C")
	flag.Parse()

	switch strings.ToUpper(*choice) {
	case "S":
		startServer()
	case "C":
		startClient()
	}
}

func startClient() {
	conn , err := net.Dial("tcp" , ":8000")
	if err!=nil{
		fmt.Println("Err :" , err)
	}

	defer conn.Close()

	fmt.Println("What Msg do you want to send ?")
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("Writing... " , scanner.Text())
		conn.Write(append(scanner.Bytes() , '\r'))
		buffer := make([]byte , 1024)
		_ , err:=conn.Read(buffer)
		if err != nil && err != io.EOF{
			fmt.Println(err)
		}else if err == io.EOF {
			fmt.Println("conn is closed")
		}

		fmt.Println(string(buffer))
		fmt.Println("What Msg do you want to send ?")
	}
}

func startServer() {
	l , err := net.Listen("tcp" , ":8000")
	if err != nil {
		fmt.Println("err :" , err)
	}

	defer l.Close()

	for {
		conn , err := l.Accept()
		if err != nil {
			fmt.Println("err2 :" , err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	//Use this instead of conn.Read and conn.Write for more functionalities
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		conn.SetDeadline(time.Now().Add(time.Second*20))
		line , err := reader.ReadString('\r')
		if err != nil && err != io.EOF{
			fmt.Println(err)
			return
		}else if err == io.EOF {
			fmt.Println("conn is closed")
			return
		}

		fmt.Printf("Received %s from address %s \n" , line[:len(line)-1] , conn.RemoteAddr())
		writer.WriteString("Msg Received...")                //same as conn.Write([]byte("msg..")
		writer.Flush()                           //to flush the data to underlying writer (here the conn interface)
	}

}