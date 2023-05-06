package main

import (
	"github.com/tang-projects/api_go/internal/database"
	"github.com/tang-projects/api_go/internal/router"
)

func main() {
	database.ConnectPostgreSQL()
	database.CreateTables()
	database.ClosePostgreSQL()
	router.RunServe()
}
