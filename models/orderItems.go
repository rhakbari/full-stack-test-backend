package models

import "gorm.io/gorm"

type OrderItems struct {
	gorm.Model
	Id             string  `json:"id" `
	Order_id       *string `json:"order_id" ,  gorm:"primary_key" `
	Price_per_unit float32 `json:"price_per_unit"`
	Quantity       int32   `json:"quantity" `
	Product        string  `json:"product" `
}
