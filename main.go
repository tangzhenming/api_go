package main

import (
	"github.com/tang-projects/api_go/internal/database"
	"github.com/tang-projects/api_go/internal/router"
)

func main() {
	// database.PGConnect()
	// database.PGCreateTables()
	// defer database.PGClose() // how to use defer: https://sl.bing.net/h5D0LAVudgW
	database.MySQLConnect()
	database.MySQLCreateTables()
	defer database.MySQLClose()
	router.RunServe()
}
