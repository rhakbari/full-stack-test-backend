package config

import (
	"packform/full-stack-test-backend/models"
	"packform/full-stack-test-backend/query"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=password dbname=postgres port=5432"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Orders{}, &models.Customers{}, &models.OrderItems{}, &models.Deliveries{}, &models.CustomerCompanies{})
	DB = db

	query.FillDB(DB)
}
