package main

import "fmt"

func main() {

	slice:=[]int{1,2,3,4,5,6,7,8,9}

	i:=2 //delete value at this index

	copy(slice[i:] , slice[i+1:])      //copt in slice(2 se last tak) mei slice(3 se last tak)
	slice = slice[:len(slice)-1]
	fmt.Println(slice , len(slice) , cap(slice))

}
