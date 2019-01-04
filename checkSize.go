package main

import (
	"fmt"
	"unsafe"
)

type A struct {                      //thus 24bytes
	A1 bool                       // _
	A2 float64                    // _ _ _ _ _ _ _ _
	A3 int32                      // _ _ _ _
}

type B struct {                     //thus 16 bytes as bool is accomodated with int32 wala line
	B1 float64                    // _ _ _ _ _ _ _ _
	B2 int32                      // _ _ _ _
	B3 bool                       // _ -> this goes in aboe line itself therefore 8*2 = 16 bytes
}

func main() {

	a := A{}
	b := B{}

	//fmt.Println(unsafe.Sizeof(int32(1)))
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))

}
