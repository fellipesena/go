package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/products", getProducts)

	router.Run("localhost:3000")
}

func getProducts(c *gin.Context){
	product := new (Product)

	product.Id = 1
	product.Name = "nome produto"
	product.Value = 3.20
	product.CreateDate = time.Now()

	c.IndentedJSON(200, product)
}

type Product struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Value float64 `json:"value"`
	CreateDate time.Time `json:"createDate"`
}