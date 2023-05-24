// $ go run . migrate [modelName]

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Long:  `Migrate the database by creating or updating tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		runMigrateCmd(cmd, args)
	},
}

func runMigrateCmd(cmd *cobra.Command, args []string) {
	var err error

	switch args[0] {
	case "User":
		err = db.PG.AutoMigrate(&models.User{})
	default:
		log.Fatalf("Unknown model: %s", args[0])
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully migrate table by model: %s \n", args[0])
}
