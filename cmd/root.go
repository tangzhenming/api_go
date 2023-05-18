package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "My app",
	Long:  `My app does great things.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello %s\n", name)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "World", "The name to say hello to")
}

func Execute() error {
	return rootCmd.Execute()
}

// Usage
// go run . --help
// go run . --name Ryan
