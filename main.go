package main

import (
	"log"

	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/router"
)

func main() {
	DB, err := db.DBConnection()
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	router.Run(DB)
}
