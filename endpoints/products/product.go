package endpoints

import (
	"first/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	
	fn := func(c *gin.Context){
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil { 
			log.Printf("Error when try decode: %v", err)
		}

		c.JSON(200, product)
	}

	return gin.HandlerFunc(fn)
}