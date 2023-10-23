package cmd

import (
	"github.com/joelseq/surreal-search/internal/server"
	"github.com/spf13/cobra"
)

var port uint

func init() {
	serveCmd.Flags().UintVarP(&port, "port", "p", 8080, "Port to run server on")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Backend API",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(port)
		s.Serve()
	},
}
