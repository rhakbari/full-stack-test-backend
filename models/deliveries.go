package models

import "gorm.io/gorm"

type Deliveries struct {
	gorm.Model
	Id                 string  `json:"id" ,gorm: "primary_key" "unique"`
	Order_item_id      *string `json:"order_item_id" ,gorm:"foreignKey:order_item_id"`
	Delivered_quantity int32   `json:"delivered_quantity" `
}
