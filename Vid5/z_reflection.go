package main

import (
	"reflect"
	"fmt"
)

func main() {
	// TypeOf returns the reflection Type that represents the dynamic type of i.
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i.

	v := 4.5

	analyzeValue(v)

}

func analyzeValue(v interface{}) {
	val := reflect.ValueOf(v)                    //returns a reflect value from interface value ,on which methods like, Type , Kind can work
	fmt.Println(v ,val , val.Type())
	fmt.Println(val.Kind() == reflect.Float64)        //val.Kind returns original underlying type
	fmt.Println("float value :" , val.Float())

	interfaceVal := val.Interface()                 //again returns the interface value from reflect value
	fmt.Println(interfaceVal)

	switch t:=interfaceVal.(type) {
	case float64:
		fmt.Println("yes" , t)
	case float32:
		fmt.Println("no" , t)
	}

}