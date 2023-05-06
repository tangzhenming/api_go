package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnect() {
	user := "api-go"
	password := "123456"
	host := "api-go-mysql"
	port := "3306"
	dbname := "api-go-mysql-dev"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Successfully connected to the MySQL database!!! ===")
}

func MySQLClose() {
	DB.Close()
	fmt.Println("=== Successfully closed the MySQL database!!! ===")
}

func MySQLCreateTables() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Successfully created MySQL table: users!!! ===")
}
