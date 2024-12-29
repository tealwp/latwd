package main

import (
	"errors"
	"fmt"
)

type Link struct {
	URL      string
	Depth    int
	DeadLink DeadLink
	Children []*Link
	Parent   *Link
	Client   *Client
}

type App struct {
	client     *Client
	links      []Link
	maxDepth   int
	maxBreadth int
}

// input params should go in here and put onto the various fields
func NewApp(maxDepth, maxBreadth int) *App {
	return &App{
		client:     NewClient(),
		links:      []Link{},
		maxDepth:   maxDepth,
		maxBreadth: maxBreadth,
	}
}

func (a *App) StartTrolling(url string) {
	// recursively call Troll on the url, then its children, and so on and so forth
	link := Link{URL: url, Depth: 1, Client: a.client}
	link.Troll(url, 1, a.maxDepth, a.maxBreadth)
	a.links = append(a.links, link)
}

func (l *Link) Troll(url string, currentDepth, maxDepth, maxBreadth int) {
	if currentDepth >= maxDepth {
		return
	}
	// call link
	body, err := l.Client.Get(url)
	// if deadlink, add to the link and return for next
	var deadLink DeadLink
	if err != nil && errors.As(err, &deadLink) {
		l.DeadLink = deadLink
		return
	}
	// else just print and return
	if err != nil {
		fmt.Printf("failed call for %s: %v\n", url, err)
		return
	}
	childURLs, err := ParseHTML(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	filteredLinks := FilterLinks(childURLs, url)
	if len(filteredLinks) > maxBreadth {
		filteredLinks = filteredLinks[:maxBreadth]
	}
	for _, link := range childURLs {
		newL := &Link{URL: link, Depth: currentDepth + 1, Client: l.Client, Parent: l}
		l.Children = append(l.Children, newL)
		newL.Troll(link, currentDepth+1, maxDepth, maxBreadth)
	}
}
