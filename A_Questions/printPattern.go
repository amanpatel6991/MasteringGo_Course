package main

import "fmt"

const maxNum = 4

func main() {


	printPattern(1)
	fmt.Println()
	fmt.Println()
	printPattern2(1 , 1)

}

func printPattern(num int) {

	if num > maxNum {
		return
	}

	for i:=1;i<=num;i++{
		fmt.Print(num, " ")
	}
	fmt.Println()
	printPattern(num + 1)
	for i:=1;i<=num;i++{
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func printPattern2(num ,prev int) {

	if num > maxNum {
		return
	}

	i := prev
	k := prev
	for j:=1;j<=num;j++{
		fmt.Print(i, " ")
		i++
	}
	//prev = i
	fmt.Println()
	printPattern2(num + 1 , i)

	for j:=1;j<=num;j++{
		fmt.Print(k, " ")
		k++
	}
	fmt.Println()
}