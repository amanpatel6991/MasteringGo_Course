package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/" , index)
	http.ListenAndServe(":8000" , nil)

}

func index(w http.ResponseWriter , r *http.Request) {
	fmt.Fprintln(w , "Hello !")
}