package link_collector

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	html "golang.org/x/net/html"
)

type LinkCollector struct {
	targetURL string
	domain    string
	links     []string
}

func New(targetURL string) *LinkCollector {
	return &LinkCollector{targetURL: targetURL}
}

func (l *LinkCollector) visit(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				l.links = append(l.links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		l.visit(c)
	}
}

func (l *LinkCollector) CollectLinks() error {
	u, err := url.Parse(l.targetURL)
	if err != nil {
		log.Fatal(err)
	}
	l.domain = u.Hostname()
	resp, err := http.Get(l.targetURL)
	if err != nil {
		return errors.New("couldn't fetch data")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return errors.New("couldn't parse data")
	}

	l.visit(doc)
	l.getFullPath()
	return nil

}

func (l *LinkCollector) getFullPath() {
	for index, _ := range l.links {
		if string(l.links[index][0]) == "/" {
			l.links[index] = l.domain + l.links[index]
		} else if l.links[index][0:2] == ".." {
			l.links[index] = l.domain + l.links[index][2:]
		}
	}
}

func (l *LinkCollector) GetLinks() []string {
	return l.links
}
