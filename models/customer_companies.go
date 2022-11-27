package models

import "gorm.io/gorm"

type CustomerCompanies struct {
	gorm.Model
	Company_id   string `json:"company_id", gorm: "primary_key"`
	Company_name string `json:"company_name" `
}
