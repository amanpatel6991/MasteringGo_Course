package main

import (
	"fmt"
	"reflect"
)

type Printer interface {
	Print(s string)
}

type pStruct struct {
	s string
}

func(p *pStruct) Print(s string) {
	p.s = s
	fmt.Println(s)
}

func main() {

	//var p Printer
	//p = &pStruct{}
	//p.Print("data")
	////fmt.Print(p)
	//
	//switch p.(type) {
	//case *pStruct:
	//	fmt.Print("implements")
	//}


	p := &pStruct{}

	inspectType(p)

}

func inspectType(p interface{}) {
	val := reflect.ValueOf(p)
	typ := val.Type()
	fmt.Println(p , val , typ)
	fmt.Println(reflect.TypeOf(p) , val.Elem())

	myinterface := reflect.TypeOf((*Printer)(nil)).Elem()

	fmt.Println(myinterface , typ.Implements(myinterface))        //check if a type typ implements the interface

	if typ.Implements(myinterface) {
		printFunc := val.MethodByName("Print")
		args := []reflect.Value{reflect.ValueOf("print hello !")}  //because args should br []reflect.Value{}
 		printFunc.Call(args)
	}

}