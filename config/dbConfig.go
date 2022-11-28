package config

import (
	"os"
	"packform/full-stack-test-backend/models"
	"packform/full-stack-test-backend/query"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	HOST := os.Getenv("DB_HOST")
	USERNAME := os.Getenv("USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DATABASE := os.Getenv("DATABASE")
	DB_PORT := os.Getenv("DB_PORT")

	db, err := gorm.Open(postgres.Open("host=" + HOST + " user=" + USERNAME + " password=" + DB_PASSWORD + " dbname=" + DATABASE + " port=" + DB_PORT))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Orders{}, &models.Customers{}, &models.OrderItems{}, &models.Deliveries{}, &models.CustomerCompanies{})
	DB = db

	query.FillDB(DB)
}
