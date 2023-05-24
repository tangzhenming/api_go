// $ go run main.go drop

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
)

var dropCmd *cobra.Command

func init() {
	var (
		err       error
		modelName string
	)

	dropCmd = &cobra.Command{
		Use:   "drop",
		Short: "Drop  table",
		Long:  `Drop  table from the database.`,
		Run: func(cmd *cobra.Command, args []string) {
			switch modelName {
			case "User":
				err = db.PG.Migrator().DropTable(&models.User{})
			default:
				log.Fatalf("Unknown model: %s", modelName)
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Successfully drop %s \n", modelName)
		},
	}

	dropCmd.Flags().StringVarP(&modelName, "drop", "d", "", "The name of the model to drop")
}
