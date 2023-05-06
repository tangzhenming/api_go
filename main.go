package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/tang-projects/api_go/internal/router"
)

const (
	host     = "db-for-api-go"
	port     = 5432
	user     = "api-go"
	password = "123456"
	dbname   = "api-go_dev"
)

func main() {
	// Set up the database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

	router.RunServe()

	// // Set up the Gin router
	// r := gin.Default()

	// // Add a route that queries the database and returns the result
	// r.GET("/ping", func(c *gin.Context) {
	// 	var result string
	// 	err := db.QueryRow("SELECT 'pong'").Scan(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"message": result,
	// 	})
	// })

	// // Start the Gin server
	// r.Run()
}
