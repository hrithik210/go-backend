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
