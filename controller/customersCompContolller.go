package controller

import (
	"net/http"
	"packform/full-stack-test-backend/config"
	"packform/full-stack-test-backend/models"

	"github.com/gin-gonic/gin"
)

func CustomerCompController(c *gin.Context) {
	customerCompanies := []models.CustomerCompanies{}
	config.DB.Find(&customerCompanies)
	c.JSON(http.StatusOK, gin.H{"status": "data fetched sucessfully", "data": &customerCompanies})
}
