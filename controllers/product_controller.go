package controllers

import (
	"net/http"

	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateProducts(c *gin.Context) {
	var Product models.Product

	if err := c.ShouldBindJSON(&Product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if Product.IsMainCard {
		config.DB.Model(&models.Product{}).Where("is_main_card = ?", true).Update("is_main_card", false)
	}

	if err := config.DB.Create(&Product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Product)
}
func getAllProducts(c *gin.Context) {
	var products []models.Product

	config.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product

	var id string = c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.IsMainCard {
		config.DB.Model(&models.Product{}).Where("is_main_card = ?", true).Update("is_main_card = ?", false)
	}

	config.DB.Model(&product).Updates(&input)
	c.JSON(http.StatusOK, product)
}
