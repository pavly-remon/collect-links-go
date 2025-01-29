package ScrapLinks

import (
	"errors"
	"net/http"

	html "golang.org/x/net/html"
)

func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}

func ScrapLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("couldn't fetch data")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, errors.New("couldn't parse data")
	}

	links := Visit(nil, doc)
	return links, nil

}
