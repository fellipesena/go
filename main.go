package main

import (
	endpoints "first/endpoints/products"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/products/:id", endpoints.Get())
	router.POST("/products", endpoints.Create())

	router.Run("localhost:3000")
}
