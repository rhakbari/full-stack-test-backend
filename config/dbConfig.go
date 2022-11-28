package config

import (
	"os"
	"packform/full-stack-test-backend/models"
	"packform/full-stack-test-backend/query"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	DB_HOST := os.Getenv("S3_BUCKET")
	DB_USERNAME := os.Getenv("S3_BUCKET")
	DB_PASSWORD := os.Getenv("S3_BUCKET")
	DATABASE := os.Getenv("S3_BUCKET")
	DBPORT := os.Getenv("SECRET_KEY")

	db, err := gorm.Open(postgres.Open("host=" + DB_HOST + "user=" + DB_USERNAME + "password=" + DB_PASSWORD + "dbname=" + DATABASE + "port=" + DBPORT))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Orders{}, &models.Customers{}, &models.OrderItems{}, &models.Deliveries{}, &models.CustomerCompanies{})
	DB = db

	query.FillDB(DB)
}
