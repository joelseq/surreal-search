package visitor

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joelseq/surreal-search/api/types"
	"github.com/typesense/typesense-go/typesense"
)

type Visitor struct {
	client *typesense.Client
}

func NewVisitor(client *typesense.Client) *Visitor {
	return &Visitor{
		client: client,
	}
}

func (v *Visitor) Visit(url string, doc *goquery.Document) {
	bc := "Overview"
	doc.Find("main > crumb > crumb-item").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())

		if i == 0 {
			bc = text
		} else {
			bc += " > " + text
		}
	})

	title := strings.TrimSpace(doc.Find("main > layout-text h2").First().Text())

	var content string

	doc.Find("main > layout-text").Each(func(i int, s *goquery.Selection) {
		content += strings.TrimSpace(s.Text())
	})

	page := &types.Page{
		Url:        url,
		Breadcrumb: bc,
		Title:      title,
		Content:    content,
	}

	v.client.Collection("pages").Documents().Create(page)
}
