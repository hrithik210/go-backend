package routes

import (
	"ecommerce-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.POST("/", controllers.CreateProducts)
		products.GET("/", controllers.GetAllProducts)
		products.GET("/:id", controllers.GetProductById)
		products.PUT("/:id", controllers.UpdateProduct)
		products.DELETE("/:id", controllers.DeleteProduct)
	}
}
