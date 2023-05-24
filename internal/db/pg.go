package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PG *gorm.DB

func ConnectPG() {
	dsn := os.Getenv("DB_DSN")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	PG = DB
}
