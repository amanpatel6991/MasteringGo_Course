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

	inspectStruct(mys)

}

func inspectStruct(mys interface{}) {
	mysVal := reflect.ValueOf(mys)
	mysType := reflect.TypeOf(mys)
	fmt.Println(mys, mysVal, mysType , mysVal.Type() , mysVal.Kind()==reflect.Struct)

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	for i := 0; i < mysType.NumField(); i++ {
		fieldType := mysType.Field(i)                  //type StructField -> has some metadata of struct fields
		fieldVal := mysVal.Field(i)
		fmt.Println(fieldType.Name , fieldType.Tag.Get("json"))
		fmt.Println(fieldVal)
		fmt.Println(fieldVal.Interface())
	}
}