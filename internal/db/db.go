package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tang-projects/api_go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("File .env can't loaded", err)
	}
	dsn := os.Getenv("DB_DSN")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.Migrator().DropTable(&models.User{})
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
