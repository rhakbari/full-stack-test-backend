package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	// Id           string `json:"id" "unique"`
	User_id      string `json:"user_id", gorm: "primary_key" "unique"`
	Login        string `json:"login" `
	Password     string `json:"password" `
	Name         string `json:"name" `
	Company_id   string `json:"company_id", gorm:"foreignKey:Company_id"`
	Credit_cards string `json:"credit_cards" `
}
