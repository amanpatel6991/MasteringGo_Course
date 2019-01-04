package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func main() {

	//f, err := os.Create("test1.txt")
	//f, err := os.Open("test1.txt")                       //Opens file for read only
	f, err := os.OpenFile("test1.txt", os.O_APPEND | os.O_RDWR, 0666)                       //Opens file for read only
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	//err = os.Rename("test2New.txt", "test2.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//copyFile("test1.txt", "test2.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}


	//READING A FILE
	//data := make([]byte ,100)
	//f.Read(data)
	//fmt.Println(string(data))
	            //(OR)
	//bytes , _ := ioutil.ReadAll(f)
	//fmt.Println(string(bytes))
	            //(OR)
	//bytes , _ := ioutil.ReadFile("test1.txt")
	//fmt.Println(string(bytes))
	            //(OR)
	//If file is large and requires line by line reading , use BUFIO package
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	//WRITING INTO A FILE
	writerBuffer := bufio.NewWriter(f)
	writerBuffer.WriteString(fmt.Sprintln("written from bufio writer !!"))
	writerBuffer.Flush()
}

func copyFile(name1, name2 string) {
	file1, _ := os.OpenFile(name1 , os.O_APPEND , 0666)
	defer file1.Close()

	file2, _ := os.Open(name2)
	defer file2.Close()

	n, err := io.Copy(file1, file2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("copied : " , n)

	file1.Sync()                                  //Flush content for permanent storage


}