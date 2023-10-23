package types

import (
	"github.com/mitchellh/mapstructure"
	"github.com/typesense/typesense-go/typesense/api"
)

var Schema *api.CollectionSchema = &api.CollectionSchema{
	Name: "pages",
	Fields: []api.Field{
		{
			Name: "url",
			Type: "string",
		},
		{
			Name: "breadcrumb",
			Type: "string",
		},
		{
			Name: "title",
			Type: "string",
		},
		{
			Name: "content",
			Type: "string",
		},
	},
}

type Page struct {
	Url        string `json:"url" mapstructure:"url"`
	Breadcrumb string `json:"breadcrumb" mapstructure:"breadcrumb"`
	Title      string `json:"title" mapstructure:"title"`
	Content    string `json:"content" mapstructure:"content"`
}

type SearchResult struct {
	Page      *Page  `json:"page"`
	Highlight string `json:"highlight"`
}

func GetPageFromMap(m *map[string]interface{}) *Page {
	p := &Page{}
	mapstructure.Decode(m, p)

	return p
}
