package main

import (
	"os"
	"packform/full-stack-test-backend/config"
	"packform/full-stack-test-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	godotenv.Load()
	port := os.Getenv("BACKEND_PORT")
	routes.GetOrder(r)
	routes.GetOrderItems(r)
	routes.GetCustomerCompany(r)
	routes.GetDelivery(r)
	routes.GetCustomers(r)
	config.Connect()
	r.Run(port)
}
