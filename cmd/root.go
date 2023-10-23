package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sursearch",
	Short: "Search the SurrealDB docs",
	Long:  `An unofficial search API for the SurrealDB documentation.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
