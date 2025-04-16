package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"ecommerce-backend/routes"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Product{})

	routes.ProductRoutes(r)

	r.Run(":3000")
	fmt.Printf("🚀 Server is running at port 3000")
}
