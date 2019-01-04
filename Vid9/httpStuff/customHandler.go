package main

import (
	"net/http"
	"fmt"
)

type Person struct {
	Id   int
	Name string
}

func(p Person) ServeHTTP(w http.ResponseWriter , r *http.Request) {
	r.URL.Query().Set("id2" , "3")
	fmt.Println(r.URL.Path , r.URL.Query().Get("id"))
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w , "using custom handler !")
	case "/index":
		fmt.Fprintln(w , "using custom handler index !")
	}
}

func main() {

	var p Person
	//http.Handle("/"  ,p)
	//http.Handle("/index"  ,p)
	//http.ListenAndServe(":8000", nil)
	http.ListenAndServe(":8000", p)

}
