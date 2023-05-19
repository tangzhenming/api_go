package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("File .env can't loaded", err)
	}
	dsn := os.Getenv("DB_DSN")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	return DB
}
