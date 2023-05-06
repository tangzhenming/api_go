package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "db-for-api-go"
	port     = 5432
	user     = "api-go"
	password = "123456"
	dbname   = "api-go_dev"
)

func ConnectPostgreSQL() {
	// Set up the database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Successfully connected to the database!!! ===")
}

func ClosePostgreSQL() {
	DB.Close()
	fmt.Println("=== Successfully closed the database!!! ===")
}

func CreateTables() {
	// 创建 users
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Successfully created users!!! ===")
}
