package searcher

import (
	"github.com/joelseq/surreal-search/api/types"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

type Searcher struct {
	client *typesense.Client
}

func NewSearcher(client *typesense.Client) *Searcher {
	return &Searcher{
		client: client,
	}
}

func (s *Searcher) Search(query string) ([]*types.SearchResult, error) {
	searchParameters := &api.SearchCollectionParams{
		Q:       query,
		QueryBy: "title, content",
	}

	res, err := s.client.Collection("pages").Documents().Search(searchParameters)

	if err != nil {
		return nil, err
	}

	pages := make([]*types.SearchResult, len(*res.Hits))
	for i, hit := range *res.Hits {
		p := types.GetPageFromMap(hit.Document)
		searchResult := &types.SearchResult{
			Page: p,
		}

		if len(*hit.Highlights) > 0 {
			searchResult.Highlight = *(*hit.Highlights)[0].Snippet
		}

		pages[i] = searchResult
	}

	return pages, nil
}
