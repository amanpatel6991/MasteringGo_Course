package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
)

type Person struct {
	Id int
	Name string
}

func main() {

	api:=gin.Default()

	api.GET("/" , getTest)
	api.Run(":8000")

}

func getTest(c *gin.Context) {

	var p Person
	_ = json.NewDecoder(c.Request.Body).Decode(&p)

	fmt.Fprintln(c.Writer,p)

}
