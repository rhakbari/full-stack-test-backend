package controller

import (
	"net/http"
	"packform/full-stack-test-backend/config"
	"packform/full-stack-test-backend/models"

	"github.com/gin-gonic/gin"
)

func OrderItemsController(c *gin.Context) {
	orderItems := []models.OrderItems{}
	config.DB.Find(&orderItems)
	c.JSON(http.StatusOK, gin.H{"status": "data fetched sucessfully", "data": &orderItems})
}
