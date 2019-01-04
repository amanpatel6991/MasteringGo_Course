package main

import (
	"github.com/zcustomPkg/calc"
	"fmt"
)

func main() {
	var n1 int =10
	var n2 int =20

	sum := calc.Add(n1,n2)

	fmt.Println("sum is :" , sum)
}
