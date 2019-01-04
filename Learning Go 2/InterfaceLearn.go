package main

import (
	"fmt"
	"strconv"
)

//A type implements a interface by implementing its methods e.g: (type Temp implements interface MyStringer by implementing its method String())

type MyStringer interface {
	String1() string
}

type Temp int

type Point struct {
	x, y int
}

func (t Temp) String1() string {                       //Explictly called by using x.String1()
	return strconv.Itoa(int(t)) + " °C"
}

func (p *Point) String() string {                      //Gets automatically implemented due to the popular Stringer interface
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func main() {

	//var x MyStringer

	x := Temp(24)
	//fmt.Println(x) // "24 °C"
	fmt.Println(x.String1()) // "24 °C"

	y := &Point{1, 2}
	fmt.Println(y) // "(1,2)"
	//fmt.Println(x.String()) // "(1,2)"

}
