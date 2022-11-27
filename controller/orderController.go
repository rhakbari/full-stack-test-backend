package controller

import (
	"net/http"
	"packform/full-stack-test-backend/config"
	"packform/full-stack-test-backend/models"

	"github.com/gin-gonic/gin"
)

func OrderController(c *gin.Context) {
	orders := []models.Orders{}
	config.DB.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"status": "data fetched sucessfully", "data": &orders})
}
