package endpoints

import (
	"first/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Get() gin.HandlerFunc {
	
	fn := func(c *gin.Context){
		var product models.Product
		log.Printf("aqui: %v", c.Params)

		log.Printf("aaaaaa: %v", c.Params.ByName("id"))
		
		xd, ok := c.Params.Get("id")

		log.Printf("aaaaaa: %v %v", xd, ok)

		c.JSON(200, product)
	}

	return gin.HandlerFunc(fn)
}