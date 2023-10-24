package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joelseq/surreal-search/api/types"
	"github.com/joelseq/surreal-search/internal/crawler"
	"github.com/joelseq/surreal-search/internal/visitor"
	"github.com/spf13/cobra"
	"github.com/typesense/typesense-go/typesense"
)

var workers uint
var depth uint8

func init() {
	indexCmd.Flags().UintVarP(&workers, "workers", "w", 50, "Number of workers to use")
	indexCmd.Flags().Uint8VarP(&depth, "depth", "d", 5, "Max depth to crawl for links")
	rootCmd.AddCommand(indexCmd)
}

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Builds a Typesense index of the Surreal docs",
	Long:  `This command will (re-)build the index of the SurrealDB docs site in Typesense`,
	Run: func(cmd *cobra.Command, args []string) {
		client := typesense.NewClient(
			typesense.WithServer(os.Getenv("TYPESENSE_API_ENDPOINT")),
			typesense.WithAPIKey(os.Getenv("TYPESENSE_API_KEY")),
		)
		createSchema(client)
		v := visitor.NewVisitor(client)
		c := crawler.NewCrawler("https://surrealdb.com", "/docs", depth, workers, v)

		if err := c.Crawl(); err != nil {
			log.Fatalf("%v\n", err)
		}
	},
}

func createSchema(c *typesense.Client) {
	collections, err := c.Collections().Retrieve()

	if err != nil {
		log.Fatalln(err)
	}
	if len(collections) > 0 {
		fmt.Println("Deleting existing collection...")

		_, err = c.Collection("pages").Delete()

		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Creating schema...")
	_, err = c.Collections().Create(types.Schema)

	if err != nil {
		log.Fatalln(err)
	}
}
