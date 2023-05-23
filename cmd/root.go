// $ go run . --help

package cmd

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/db"
	"gorm.io/gorm"
)

var modelName string
var err error
var DB *gorm.DB
var RedisClient *redis.Client
var rootCmd = &cobra.Command{Use: "app"}

func init() {
	DB = db.DBConnection()
	RedisClient = db.NewRedis()
}

func Execute() {
	rootCmd.AddCommand(serveCmd)

	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(dropCmd)

	if err = rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
