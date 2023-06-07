// $ go run main.go drop [modelName]

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/db"
	"github.com/tang-projects/api_go/internal/models"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop  table",
	Long:  `Drop  table from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		runDropCmd(cmd, args)
	},
}

func runDropCmd(cmd *cobra.Command, args []string) {
	var err error

	switch args[0] {
	case "User":
		err = db.PG.Migrator().DropTable(&models.User{})
	case "Post":
		err = db.PG.Migrator().DropTable(&models.Post{})
	default:
		log.Fatalf("Unknown model: %s", args[0])
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully drop table by model: %s \n", args[0])
}
