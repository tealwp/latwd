package main

import (
	"bytes"
	"fmt"
	"net/url"

	"golang.org/x/net/html"
)

func ParseHTML(body []byte) ([]string, error) {
	node, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}
	return GetAllLinks(node), nil
}

func GetAllLinks(node *html.Node) []string {
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)
	return links
}

func FilterLinks(links []string, currentURL string) []string {
	current, err := url.Parse(currentURL)
	if err != nil {
		return nil
	}
	var filteredLinks []string
	for _, link := range links {
		linkURL, err := url.Parse(link)
		if err != nil {
			continue
		}
		if linkURL.Host == current.Host {
			filteredLinks = append(filteredLinks, link)
		}
	}

	return filteredLinks
}
