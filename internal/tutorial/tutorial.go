package tutorial

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectToDB() *sql.DB {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("File .env can't loaded", err)
	}
	dsn := os.Getenv("DB_DSN")

	// Open and Close database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
