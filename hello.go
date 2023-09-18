package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/hello-world", getHelloWorld)

	router.Run("localhost:3000")
}

func getHelloWorld(c *gin.Context){
	c.IndentedJSON(200, "hello world")
}