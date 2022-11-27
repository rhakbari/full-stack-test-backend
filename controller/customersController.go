package controller

import (
	"net/http"
	"packform/full-stack-test-backend/config"
	"packform/full-stack-test-backend/models"

	"github.com/gin-gonic/gin"
)

func CustomersController(c *gin.Context) {
	customers := []models.Customers{}
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"status": "data fetched sucessfully", "data": &customers})
}
