package main

import (
	"flag"
	"strings"
)

func main() {
	op := flag.String("type" , "" , "Server(S) or Client(C)")
	//address
	
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(":8080")
	case "C":
		runClient(":8080")
	}
	
}

func runServer(add string) {

}

func runClient(add string) {
	c := A_ProtoParent.N
}
