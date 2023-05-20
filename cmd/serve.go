package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/router"
)

var port string
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the router and server",
	Long:  `Run the router with the specified database and port.`,
	Run: func(cmd *cobra.Command, args []string) {
		router.Run(DB, port)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "To specify the port you're going to run the router and server.")
}
