// $ go run . --help

package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/db"
)

var rootCmd = &cobra.Command{Use: "api-go"}

func init() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("File .env can't loaded", err)
	}

	// 初始化数据库
	db.ConnectPG()
	db.NewRedis()
}

func Execute() {
	// 添加命令
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(dropCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
