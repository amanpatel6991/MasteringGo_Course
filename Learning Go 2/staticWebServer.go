package main

import (
	"net/http"
	"fmt"
	"time"
)

//func main() {
//
//	mux := http.NewServeMux()
//	fs := http.FileServer(http.Dir("./"))
//	mux.Handle("/" , fs)
//	http.ListenAndServe(":8000" , mux)
//}


func main() {

	http.HandleFunc("/index", home1)

	server := &http.Server{
		Addr: ":8000",
		Handler: nil,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}

func home1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "using default servMux with server configs !")
}