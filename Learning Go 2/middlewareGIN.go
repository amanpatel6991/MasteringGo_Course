package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func GetDummyEndpoint(c *gin.Context) {
	//resp := map[string]string{"hello":"world"}
	//c.JSON(200, resp)
	c.Writer.Write([]byte("DONE !"))
}

func GetDummyEndpointNew(c *gin.Context) {
	resp := map[string]string{"how":"are you"}
	c.JSON(200, resp)
}

func DummyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if (c.Request.Method != "GET") {
			fmt.Println("exiting !")
			c.Abort()
			return
		} else {
			fmt.Println("success NEXT !")
			c.Next()
		}

	}
}

func main() {
	api := gin.Default()


	api.Use(DummyMiddleware())
	//api.Group("/api",GetDummyEndpoint , GetDummyEndpointNew)

	//api.Group()
	api.GET("/",GetDummyEndpoint)
	api.GET("/new",GetDummyEndpointNew)
	api.Run(":5000")
}