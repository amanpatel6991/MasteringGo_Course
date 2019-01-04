package main

import (
	"math/rand"
	"time"
	"fmt"
	"net"
	"log"
	"bufio"
	"os"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonynomous %d" , rand.Intn(400))
	fmt.Println("Starting Hydra Chat Client.....")
	fmt.Println("Your Name ?")
	fmt.Scanln(&name)

	fmt.Println("Hello " + name + " Connecting to hydra chat system.. ! \n")
	conn , err := net.Dial("tcp" , "127.0.0.1:2300")
	if err!=nil {
		log.Fatal("could not connect to chat system !!!!!! : " , err)
	}

	fmt.Println("Connected to hydra chat system")

	name += ":"
	defer conn.Close()
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err != nil {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, name + msg + "\n")
	}



}
