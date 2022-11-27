package main

import (
	"packform/full-stack-test-backend/config"
	"packform/full-stack-test-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
	}))

	routes.GetOrder(r)
	routes.GetOrderItems(r)
	routes.GetCustomerCompany(r)
	routes.GetDelivery(r)
	routes.GetCustomers(r)
	config.Connect()
	r.Run(":8000")
}
