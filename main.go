package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Product{})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hey there"})
	})

	r.Run(":3000")
}
