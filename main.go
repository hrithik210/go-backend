package main

import (
	"ecommerce-backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hey there",
		})
	})
}
