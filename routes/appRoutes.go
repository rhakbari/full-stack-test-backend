package routes

import (
	"packform/full-stack-test-backend/controller"

	"github.com/gin-gonic/gin"
)

func GetOrder(r *gin.Engine) {
	r.GET("/orders", controller.OrderController)
}

func GetOrderItems(r *gin.Engine) {
	r.GET("/order_items", controller.OrderItemsController)
}

func GetDelivery(r *gin.Engine) {
	r.GET("/deliveries", controller.DeliveryController)
}

func GetCustomers(r *gin.Engine) {
	r.GET("/customers", controller.CustomersController)
}

func GetCustomerCompany(r *gin.Engine) {
	r.GET("/customer_company", controller.CustomerCompController)
}
