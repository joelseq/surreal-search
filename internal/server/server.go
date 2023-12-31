package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joelseq/surreal-search/api/types"
	"github.com/joelseq/surreal-search/internal/searcher"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/typesense/typesense-go/typesense"
)

type Server struct {
	Port uint
}

func NewServer(port uint) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Serve() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Surreal Search API")
	})
	e.GET("/search", searchHandler)
	portString := fmt.Sprintf(":%d", s.Port)
	e.Logger.Fatal(e.Start(portString))
}

func searchHandler(c echo.Context) error {
	search := c.QueryParam("q")

	client := typesense.NewClient(
		typesense.WithServer(os.Getenv("TYPESENSE_API_ENDPOINT")),
		typesense.WithAPIKey(os.Getenv("TYPESENSE_API_KEY")),
	)
	searcher := searcher.NewSearcher(client)
	results, err := searcher.Search(search)

	if err != nil {
		return err
	}

	output := make([]types.SearchResult, len(results))

	for i, result := range results {
		output[i] = *result
	}

	return c.JSON(http.StatusOK, output)
}
