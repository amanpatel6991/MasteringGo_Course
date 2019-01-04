package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	Field1 int           `json:"field_1"`
	Field2 string        `json:"field_2"`
	Field3 float64       `json:"field_3"`
}

func main() {

	mys := myStruct{1, "data", 5.4}

	inspectStruct(&mys)

}

func inspectStruct(mys interface{}) {
	mysVal := reflect.ValueOf(mys)
	fmt.Println(mys , mysVal , mysVal.Kind() , mysVal.Type())

	mysVal = mysVal.Elem()
	fmt.Println(mysVal.CanSet())

	mysVal.Field(0).SetInt(2)

	fmt.Println(mys , mysVal ,mysVal.Type() , mysVal.Kind() , reflect.TypeOf(mysVal))
}