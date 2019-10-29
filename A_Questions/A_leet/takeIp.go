package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {

	fmt.Println("start......")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	fmt.Println(text)

}
