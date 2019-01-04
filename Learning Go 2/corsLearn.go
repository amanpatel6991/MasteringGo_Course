package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rs/cors"
	"fmt"
)

func corsFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello CORS !!")
}

func main() {

	router := mux.NewRouter();
	router.HandleFunc("/testCors", corsFunc)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com", "http://foo.com:8000"},                      //A list of origins a cross-domain request can be executed from.
		AllowCredentials: true,
		AllowedMethods: []string{"POST"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(router)
	http.ListenAndServe(":8000", handler)
}
