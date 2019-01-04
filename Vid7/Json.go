package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id   int        `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
}

func main() {
	p := Person{Id : 1, Name:"Aman"}
	map1 := map[int]string{1:"one", 2 : "two"}

	b, _ := json.Marshal(&p)
	bm, _ := json.Marshal(&map1)

	fmt.Println(string(b))
	fmt.Println(string(bm))

	//f , err := os.OpenFile("jsonTest.json" , os.O_APPEND|os.O_RDWR , 0666)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//json.NewEncoder(f).Encode(&p)

	dataByte := []byte(`{"id":1,"name":"Aman"}`)
	var data Person
	json.Unmarshal(dataByte , &data)

	fmt.Println(data)


}
