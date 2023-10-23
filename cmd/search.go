package cmd

import (
	"fmt"
	"log"

	"github.com/joelseq/surreal-search/internal/searcher"
	"github.com/spf13/cobra"
	"github.com/typesense/typesense-go/typesense"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the index for a page in the docs",
	Long:  "Search the index for a page in the docs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Searched for %s\n", args[0])
		client := typesense.NewClient(
			typesense.WithServer("http://localhost:8108"),
			typesense.WithAPIKey("xyz"),
		)
		searcher := searcher.NewSearcher(client)
		results, err := searcher.Search(args[0])

		if err != nil {
			log.Fatalf("%v\n", err)
		}

		for _, result := range results {
			fmt.Printf("Title: %+v\n", result.Page.Title)
			fmt.Printf("URL: %+v\n", result.Page.Url)
			fmt.Printf("Highlight: %v\n", result.Highlight)
		}
	},
}
