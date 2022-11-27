package controller

import (
	"net/http"
	"packform/full-stack-test-backend/config"
	"packform/full-stack-test-backend/models"

	"github.com/gin-gonic/gin"
)

func DeliveryController(c *gin.Context) {
	deliveries := []models.Deliveries{}
	config.DB.Find(&deliveries)
	c.JSON(http.StatusOK, gin.H{"status": "data fetched sucessfully", "data": &deliveries})
}
