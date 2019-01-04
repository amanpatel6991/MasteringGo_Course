package main

import (
	"net/http"
	"MasteringGo_Course/Vid4/SingletonPattern/hydraLogger"
	"fmt"
)

func main() {

	logger := hydraLogger.GetInstance()
	logger.Println("starting log service ....")          //this println is of log type

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	logger := hydraLogger.GetInstance()
	fmt.Fprintln(w , "Welcome to Hydra system !")

	logger.Println("Http Request received !")
}