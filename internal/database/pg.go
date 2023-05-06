package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func PGConnect() {
	host := "api-go-pg"
	port := 5432
	user := "api-go"
	password := "123456"
	dbname := "api-go-dev"

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
	fmt.Println("=== Successfully connected to the PG database!!! ===")
}

func PGClose() {
	DB.Close()
	fmt.Println("=== Successfully closed the PG database!!! ===")
}

func PGCreateTables() {
	// 创建 users ，可以使用 AI 生成 SQL 语句，remember aviod error and make it suitable for PG
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
	fmt.Println("=== Successfully created PG table: users!!! ===")
}
