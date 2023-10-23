package crawler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Crawler struct {
	baseUrl    string
	initialUrl string
	maxDepth   uint8
	numWorkers uint
	visitor    Visitor
}

type Visitor interface {
	Visit(url string, doc *goquery.Document)
}

type visitNode struct {
	Url   string
	Depth uint8
}

type inputMessage struct {
	node visitNode
}

type outputMessage struct {
	node  visitNode
	queue []visitNode
}

func NewCrawler(baseUrl, initialUrl string, maxDepth uint8, numWorkers uint, visitor Visitor) *Crawler {
	return &Crawler{
		baseUrl:    baseUrl,
		initialUrl: initialUrl,
		maxDepth:   maxDepth,
		numWorkers: numWorkers,
		visitor:    visitor,
	}
}

func (c *Crawler) Crawl() error {
	// queue := []visitNode{{Url: c.initialUrl, Depth: 0}}
	queued := map[string]bool{c.initialUrl: true}
	visited := map[string]bool{}
	context := context.Background()
	inputChan := make(chan inputMessage)
	outputChan := make(chan outputMessage, 5)
	var wg sync.WaitGroup

	go func() {
		inputChan <- inputMessage{node: visitNode{Url: c.initialUrl, Depth: 0}}
		queued[c.initialUrl] = true
	}()

	for i := uint(0); i < c.numWorkers; i++ {
		wg.Add(1)
		go c.worker(context, &wg, inputChan, outputChan)
	}

	for output := range outputChan {
		visited[output.node.Url] = true
		for _, queueNode := range output.queue {
			// Only queue up strings that start with `/docs`, are within the max depth,
			// and have not already been queued.
			if strings.HasPrefix(queueNode.Url, "/docs") && queueNode.Depth < c.maxDepth && !queued[queueNode.Url] {
				// queue = append(queue, node)
				queued[queueNode.Url] = true
				go func(qN visitNode) {
					inputChan <- inputMessage{node: qN}
				}(queueNode)
			}
		}
		fmt.Printf("Len visited: %d, Len queued: %d\n", len(visited), len(queued))

		if len(visited) == len(queued) {
			close(inputChan)
			close(outputChan)
		}
	}

	wg.Wait()

	fmt.Printf("Visited %d links\n", len(visited))

	return nil
}

func (c *Crawler) worker(context context.Context, wg *sync.WaitGroup, inputChan chan inputMessage, outputChan chan outputMessage) {
	for message := range inputChan {
		fmt.Printf("Received message for %s\n", message.node.Url)
		links, err := c.visit(context, message.node)

		var output []visitNode

		if err != nil {
			fmt.Printf("%v\n", err)

			outputChan <- outputMessage{node: message.node, queue: output}
		}

		for _, link := range links {
			output = append(output, visitNode{
				Url:   link,
				Depth: message.node.Depth + 1,
			})
		}

		outputChan <- outputMessage{node: message.node, queue: output}
	}
	wg.Done()
}

func (c *Crawler) visit(context context.Context, node visitNode) ([]string, error) {
	fmt.Printf("Visiting: %s\n", node.Url)
	res, err := http.Get(c.baseUrl + node.Url)

	if err != nil {
		fmt.Printf("Received error: %v\n", err)
		return nil, err
	}

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var links []string

	doc.Find("body > content a").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Attr("href")

		if exists {
			links = append(links, url)
		}
	})

	c.visitor.Visit(c.baseUrl+node.Url, doc)

	return links, nil
}
