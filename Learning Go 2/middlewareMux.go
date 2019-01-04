//package main
//
//import (
//	"log"
//	"net/http"
//	"github.com/gorilla/mux"
//)
//
//func middlewareOne(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("Executing middlewareOne")
//		next.ServeHTTP(w, r)
//		log.Println("Executing middlewareOne again")
//	})
//}
//
//func middlewareTwo(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("Executing middlewareTwo")
//		if r.URL.Path != "/" {
//			return
//		}
//		next.ServeHTTP(w, r)
//		log.Println("Executing middlewareTwo again")
//	})
//}
//
//func final(w http.ResponseWriter, r *http.Request) {
//	log.Println("Executing finalHandler")
//	w.Write([]byte("OK"))
//}
//
//func main() {
//	finalHandler := http.HandlerFunc(final)                         // The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers
//
//	router:=mux.NewRouter()
//
//	router.Handle("/", middlewareTwo(finalHandler))
//	http.ListenAndServe(":3000", router)
//}

//////////////////////////////////////////////////////////////////// OR //////////////////////////////////////////////////////////////////////


package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func middlewareTwo(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
		next(w, r)
		log.Println("Executing middlewareTwo again")
	}
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {
	router:=mux.NewRouter()

	router.HandleFunc("/", middlewareTwo(final))
	http.ListenAndServe(":3000", router)
}