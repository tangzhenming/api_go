package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tang-projects/api_go/internal/router"
)

var serveCmd *cobra.Command

func init() {
	var port string

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Run the router and server",
		Long:  `Run the router with the specified database and port.`,
		Run: func(cmd *cobra.Command, _ []string) {
			router.Run(port)
		},
	}

	serveCmd.Flags().StringVarP(&port, "port", "p", "80", "To specify the port you're going to run the router and server.")
}
