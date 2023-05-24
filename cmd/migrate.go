// $ go run . migrate

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
)

var migrateCmd *cobra.Command

func init() {
	var (
		err       error
		modelName string
	)

	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database",
		Long:  `Migrate the database by creating or updating tables.`,
		Run: func(cmd *cobra.Command, args []string) {
			switch modelName {
			case "User":
				err = db.PG.AutoMigrate(&models.User{})
			default:
				log.Fatalf("Unknown model: %s", modelName)
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Successfully migrate %s \n", modelName)
		},
	}

	migrateCmd.Flags().StringVarP(&modelName, "migrate", "m", "", "The name of the model to migrate")
}
