// $ go run . --help

package cmd

import (
	"database/sql"
	"log"

	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/tutorial"
)

var modelName string
var err error
var DB *sql.DB
var rootCmd = &cobra.Command{Use: "app"}

func init() {
	DB = tutorial.ConnectToDB()
}

func Execute() {
	// rootCmd.AddCommand(serveCmd)
	// rootCmd.AddCommand(migrateCmd)
	// rootCmd.AddCommand(dropCmd)

	if err = rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
