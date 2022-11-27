package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Id          string    `json:"id" `
	Created_at  time.Time `json:"created_at"`
	Order_name  string    `json:"order_name" ,gorm:"primary_key"`
	Customer_id string    `json:"customer_id", "foreignKey:Customer_id"`
}
