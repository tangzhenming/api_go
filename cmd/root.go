// $ go run . --help

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/db"
	"gorm.io/gorm"
)

var modelName string
var err error
var DB *gorm.DB
var rootCmd = &cobra.Command{Use: "app"}

func init() {
	DB = db.DBConnection()
}

func Execute() error {
	return rootCmd.Execute()
}
