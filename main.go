package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"ecommerce-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Product{})

	routes.ProductRoutes(r)

	r.Run(":3000")
}
