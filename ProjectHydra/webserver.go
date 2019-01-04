package main

import (
	"fmt"
	"net/http"
)

func main() {

	//api := gin.Default()
	//api.GET("/" , index)
	//api.Run(":8000")

	http.HandleFunc("/" , index)

	fmt.Println("Listening at : 8000......")
	http.ListenAndServe(":8000" , nil)

}


//func index(context *gin.Context) {
//	fmt.Fprintln(context.Writer , "Welcome to Hydra !")
//}

func index(w http.ResponseWriter , r *http.Request) {
	fmt.Fprintln(w , "Welcome to Hydra !")
}